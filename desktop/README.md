# Desktop app (Windows EXE)

This folder builds `LumenLayerStudio.exe`: a single file that contains the whole app.

When you run it:
1. It serves the app on `127.0.0.1:47821`, which only this computer can reach.
2. It opens the app in its own Microsoft Edge (or Chrome) app window.
3. It quits when that window closes.

The port stays the same between runs, so saved presets, watermark and AI settings are kept.

Photos or folders dropped onto the EXE (or a shortcut to it), or opened with it from Explorer, are passed in as command-line arguments. Folders are scanned up to 8 levels deep, for at most 5,000 photos. If the app is already running, the new copy sends the files to it and exits.

**How the local server stays private:**
- Any request whose `Host` is not `127.0.0.1` or `localhost` is rejected, which blocks DNS-rebinding attacks from websites.
- Files can be handed over only by another copy of the launcher: the request must carry a custom header and no `Origin`, which a web page can't do.
- A handed-over photo is served only under a random, unguessable ID.

## Build

Needs only Go 1.24 or newer. There are no other dependencies or downloads.

```bash
./build.sh
```

This copies `../index.html` into the build, generates the Windows icon and version resources with `tools/mkres` (standard library only), and writes:

- `../dist/LumenLayerStudio.exe` for Windows 10/11 x64
- `../dist/LumenLayerStudio-arm64.exe` for Windows on ARM

| File | Purpose |
|---|---|
| `main.go` | Local server and app-window launcher |
| `platform_windows.go` / `platform_other.go` | Hidden child windows and error message box |
| `tools/mkres/` | Writes the icon and version info into a `.syso` resource for `go build` |
| `icon.ico`, `icon.png` | App icon |
| `build.sh` | Build script |

The EXE is not code-signed. The first time it runs, Windows SmartScreen asks you to confirm: click **More info → Run anyway**.
