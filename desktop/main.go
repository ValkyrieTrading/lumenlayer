// Lumen Layer Studio desktop launcher.
//
// One executable that embeds the whole web app (app/index.html), serves it on a
// fixed local port and opens it in its own app window (Microsoft Edge or Google
// Chrome in "app" mode, falling back to the default browser).
//
// Photos and folders dropped onto the EXE (or a shortcut to it), or opened with
// it from Explorer, are handed to the app. If the app is already running, they
// are sent to the open window instead of starting a second copy.
//
// The server only listens on 127.0.0.1 and rejects requests from websites, so
// nothing is reachable from other computers or other pages.
package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//go:embed app/index.html
var pageHTML string

//go:embed icon.ico
var iconICO []byte

const (
	appName       = "Lumen Layer Studio"
	preferredPort = 47821
	markerHeader  = "X-Lumen-Layer-Studio"
	maxFiles      = 5000
)

// The heartbeat tells the launcher a window is still open, and collects photos
// handed over by Windows (dropped on the EXE, or "Open with").
const heartbeat = `<script>(function(){function p(){fetch('/__alive',{cache:'no-store'}).then(function(r){return r.status===200?r.json():null}).then(function(j){if(j&&j.files&&j.files.length&&window.lumenImportRemote)window.lumenImportRemote(j.files)}).catch(function(){})}p();setInterval(p,2000)})()</script>`

// Photo types the app can open, including camera RAW (read via the embedded full-size preview).
var photoExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".jfif": true, ".png": true, ".webp": true, ".gif": true, ".bmp": true, ".avif": true, ".tif": true, ".tiff": true,
	".heic": true, ".heif": true, // the app explains that these need converting
	".cr2": true, ".cr3": true, ".crw": true, ".nef": true, ".nrw": true, ".arw": true, ".srf": true, ".sr2": true, ".dng": true,
	".raf": true, ".orf": true, ".rw2": true, ".rwl": true, ".pef": true, ".srw": true, ".3fr": true, ".iiq": true, ".erf": true,
	".kdc": true, ".mrw": true, ".mos": true, ".x3f": true,
}

type fileRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// inbox holds photos waiting to be picked up by the window, and the files it may read.
var inbox struct {
	sync.Mutex
	paths   map[string]string
	pending []fileRef
}

func main() {
	inbox.paths = map[string]string{}
	args := os.Args[1:]
	page := injectHeartbeat(pageHTML)

	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", preferredPort))
	if err != nil {
		existing := fmt.Sprintf("http://127.0.0.1:%d/", preferredPort)
		if isOurServer(existing) {
			// Already running: hand photos to the open window, or open another window.
			if len(args) > 0 && sendToRunning(existing, args) {
				return
			}
			if done := openAppWindow(existing); done == nil {
				openDefaultBrowser(existing)
			}
			return
		}
		// Port taken by another program: use any free port (settings will not carry over).
		if ln, err = net.Listen("tcp", "127.0.0.1:0"); err != nil {
			fatal("Lumen Layer Studio could not start its local server:\n\n" + err.Error())
			return
		}
	}
	port := ln.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://127.0.0.1:%d/", port)
	if len(args) > 0 {
		queuePaths(args)
	}

	var lastBeat atomic.Int64
	lastBeat.Store(time.Now().Unix())
	mux := http.NewServeMux()
	mux.HandleFunc("/__alive", func(w http.ResponseWriter, r *http.Request) {
		lastBeat.Store(time.Now().Unix())
		w.Header().Set(markerHeader, "1")
		w.Header().Set("Cache-Control", "no-store")
		inbox.Lock()
		files := inbox.pending
		inbox.pending = nil
		inbox.Unlock()
		if len(files) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"files": files})
	})
	mux.HandleFunc("/__file/", func(w http.ResponseWriter, r *http.Request) {
		inbox.Lock()
		p, ok := inbox.paths[strings.TrimPrefix(r.URL.Path, "/__file/")]
		inbox.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		f, err := os.Open(p)
		if err != nil {
			http.Error(w, "The photo is no longer available", http.StatusNotFound)
			return
		}
		defer f.Close()
		st, _ := f.Stat()
		w.Header().Set("Cache-Control", "private, max-age=3600")
		http.ServeContent(w, r, filepath.Base(p), st.ModTime(), f)
	})
	mux.HandleFunc("/__open", func(w http.ResponseWriter, r *http.Request) {
		// Only another copy of this program may hand over files: it sends a custom
		// header and no Origin, which a web page cannot do without a CORS preflight.
		if r.Method != http.MethodPost || r.Header.Get("Origin") != "" || r.Header.Get(markerHeader) != "open" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		var body struct{ Paths []string }
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4<<20)).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		n := queuePaths(body.Paths)
		w.Header().Set(markerHeader, "1")
		fmt.Fprintf(w, `{"queued":%d}`, n)
	})
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Header().Set("Content-Length", strconv.Itoa(len(iconICO)))
		w.Header().Set(markerHeader, "1")
		w.Write(iconICO)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set(markerHeader, "1")
		w.Write([]byte(page))
	})
	go http.Serve(ln, localOnly(port, mux))

	// A minimised window may ping slowly; the app itself keeps working without the server.
	idle := 120 * time.Second
	if v, err := strconv.Atoi(os.Getenv("LUMEN_IDLE_SECONDS")); err == nil && v > 0 {
		idle = time.Duration(v) * time.Second
	}

	started := time.Now()
	windowClosed := openAppWindow(url)
	if windowClosed == nil && os.Getenv("LUMEN_NO_BROWSER") == "" {
		openDefaultBrowser(url)
	}
	for {
		select {
		case <-windowClosed:
			// A very quick exit means the browser handed the window to a process that
			// was already running; fall back to watching the heartbeat instead.
			if time.Since(started) > 4*time.Second {
				return
			}
			windowClosed = nil
		case <-time.After(3 * time.Second):
			if time.Since(time.Unix(lastBeat.Load(), 0)) > idle {
				return
			}
		}
	}
}

// localOnly rejects requests whose Host is not this machine's loopback address,
// which blocks DNS-rebinding attempts from websites.
func localOnly(port int, h http.Handler) http.Handler {
	ok := map[string]bool{
		fmt.Sprintf("127.0.0.1:%d", port): true,
		fmt.Sprintf("localhost:%d", port): true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ok[r.Host] {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// queuePaths expands files and folders (subfolders included) into photos for the window.
func queuePaths(paths []string) int {
	var found []string
	add := func(p string) {
		if len(found) < maxFiles && photoExt[strings.ToLower(filepath.Ext(p))] && !strings.HasPrefix(filepath.Base(p), ".") {
			found = append(found, p)
		}
	}
	for _, p := range paths {
		abs, err := filepath.Abs(p)
		if err != nil {
			continue
		}
		st, err := os.Stat(abs)
		if err != nil {
			continue
		}
		if !st.IsDir() {
			add(abs)
			continue
		}
		root := strings.Count(abs, string(os.PathSeparator))
		filepath.WalkDir(abs, func(q string, d fs.DirEntry, err error) error {
			if err != nil || len(found) >= maxFiles {
				return nil
			}
			if d.IsDir() {
				if q != abs && (strings.HasPrefix(d.Name(), ".") || strings.Count(q, string(os.PathSeparator))-root > 8) {
					return filepath.SkipDir
				}
				return nil
			}
			add(q)
			return nil
		})
	}
	inbox.Lock()
	defer inbox.Unlock()
	for _, p := range found {
		id := randomID()
		inbox.paths[id] = p
		var size int64
		if st, err := os.Stat(p); err == nil {
			size = st.Size()
		}
		inbox.pending = append(inbox.pending, fileRef{ID: id, Name: filepath.Base(p), Size: size})
	}
	return len(found)
}

func randomID() string {
	b := make([]byte, 12)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func sendToRunning(base string, paths []string) bool {
	var abs []string
	for _, p := range paths {
		if a, err := filepath.Abs(p); err == nil {
			abs = append(abs, a)
		}
	}
	body, _ := json.Marshal(map[string][]string{"paths": abs})
	req, _ := http.NewRequest(http.MethodPost, base+"__open", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(markerHeader, "open")
	c := http.Client{Timeout: 10 * time.Second}
	res, err := c.Do(req)
	if err != nil {
		return false
	}
	res.Body.Close()
	return res.StatusCode == http.StatusOK
}

func injectHeartbeat(html string) string {
	if i := strings.LastIndex(html, "</body>"); i >= 0 {
		return html[:i] + heartbeat + html[i:]
	}
	return html + heartbeat
}

func isOurServer(url string) bool {
	c := http.Client{Timeout: 1500 * time.Millisecond}
	res, err := c.Get(url + "favicon.ico")
	if err != nil {
		return false
	}
	res.Body.Close()
	return res.StatusCode == http.StatusOK && res.Header.Get(markerHeader) == "1"
}

// openAppWindow starts Edge or Chrome in app mode with its own profile folder.
// It returns a channel that closes when that browser process exits, or nil if
// no suitable browser was found.
func openAppWindow(url string) <-chan struct{} {
	exe := findAppBrowser()
	if exe == "" {
		return nil
	}
	profile := filepath.Join(dataDir(), "window-profile")
	os.MkdirAll(profile, 0o755)
	cmd := exec.Command(exe,
		"--app="+url,
		"--user-data-dir="+profile,
		"--no-first-run",
		"--no-default-browser-check",
		"--disable-features=Translate",
		"--window-size=1440,920",
	)
	hideConsole(cmd)
	if err := cmd.Start(); err != nil {
		return nil
	}
	done := make(chan struct{})
	go func() { cmd.Wait(); close(done) }()
	return done
}

func findAppBrowser() string {
	var candidates []string
	switch runtime.GOOS {
	case "windows":
		for _, base := range []string{os.Getenv("ProgramFiles(x86)"), os.Getenv("ProgramFiles"), os.Getenv("LOCALAPPDATA")} {
			if base == "" {
				continue
			}
			candidates = append(candidates,
				filepath.Join(base, "Microsoft", "Edge", "Application", "msedge.exe"),
				filepath.Join(base, "Google", "Chrome", "Application", "chrome.exe"),
				filepath.Join(base, "BraveSoftware", "Brave-Browser", "Application", "brave.exe"),
			)
		}
	case "darwin":
		candidates = []string{
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		}
	default:
		if os.Getenv("LUMEN_NO_BROWSER") != "" {
			return ""
		}
		for _, n := range []string{"microsoft-edge", "google-chrome", "chromium", "chromium-browser"} {
			if p, err := exec.LookPath(n); err == nil {
				candidates = append(candidates, p)
			}
		}
	}
	for _, c := range candidates {
		if st, err := os.Stat(c); err == nil && !st.IsDir() {
			return c
		}
	}
	return ""
}

func openDefaultBrowser(url string) {
	if os.Getenv("LUMEN_NO_BROWSER") != "" {
		return
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	hideConsole(cmd)
	cmd.Start()
}

func dataDir() string {
	if d, err := os.UserConfigDir(); err == nil {
		return filepath.Join(d, "LumenLayerStudio")
	}
	return filepath.Join(os.TempDir(), "LumenLayerStudio")
}
