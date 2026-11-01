// Lumen Layer Studio desktop launcher.
//
// One executable that embeds the whole web app (app/index.html), serves it on a
// fixed local port and opens it in its own app window (Microsoft Edge or Google
// Chrome in "app" mode, falling back to the default browser). The server only
// listens on 127.0.0.1, so nothing is reachable from other computers.
//
// The port is fixed so the browser treats every launch as the same site, which
// keeps saved presets, watermark and AI settings between runs.
package main

import (
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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
)

// heartbeat lets the launcher notice when the last window has closed.
const heartbeat = `<script>(function(){function p(){fetch('/__alive',{cache:'no-store'}).catch(function(){})}p();setInterval(p,5000)})()</script>`

func main() {
	page := injectHeartbeat(pageHTML)
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", preferredPort))
	if err != nil {
		// Already running? Then just open another window on the same instance.
		existing := fmt.Sprintf("http://127.0.0.1:%d/", preferredPort)
		if isOurServer(existing) {
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
	url := "http://" + ln.Addr().String() + "/"

	var lastBeat atomic.Int64
	lastBeat.Store(time.Now().Unix())
	mux := http.NewServeMux()
	mux.HandleFunc("/__alive", func(w http.ResponseWriter, r *http.Request) {
		lastBeat.Store(time.Now().Unix())
		w.Header().Set(markerHeader, "1")
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
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
	go http.Serve(ln, mux)

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

func injectHeartbeat(html string) string {
	if i := strings.LastIndex(html, "</body>"); i >= 0 {
		return html[:i] + heartbeat + html[i:]
	}
	return html + heartbeat
}

func isOurServer(url string) bool {
	c := http.Client{Timeout: 1500 * time.Millisecond}
	res, err := c.Get(url + "__alive")
	if err != nil {
		return false
	}
	res.Body.Close()
	return res.Header.Get(markerHeader) == "1"
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
