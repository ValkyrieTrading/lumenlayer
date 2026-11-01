# Desktop app (Windows EXE)

This folder builds `LumenLayerStudio.exe`: a single file that contains the whole app.

When you run it:
1. It serves the app on `127.0.0.1:47821`, which only this computer can reach.
2. It opens the app in its own Microsoft Edge (or Chrome) app window.
3. It quits when that window closes.

The port stays the same between runs, so saved presets, watermark and AI settings are kept.

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
