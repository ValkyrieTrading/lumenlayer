#!/usr/bin/env bash
# Builds the Windows desktop app (x64 and ARM64) from ../index.html.
# Needs only Go 1.24+ (no other downloads). Output goes to ../dist/.
set -euo pipefail
cd "$(dirname "$0")"
mkdir -p app ../dist
cp ../index.html app/index.html

go run ./tools/mkres -ico icon.ico -arch amd64 -o resource_windows_amd64.syso
go run ./tools/mkres -ico icon.ico -arch arm64 -o resource_windows_arm64.syso

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-H=windowsgui -s -w" -o ../dist/LumenLayerStudio.exe .
GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-H=windowsgui -s -w" -o ../dist/LumenLayerStudio-arm64.exe .
ls -la ../dist
