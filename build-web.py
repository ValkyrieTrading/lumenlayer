#!/usr/bin/env python3
"""Builds index.html (the installable web app) and sw.js from src/app.html.

Run after changing src/app.html, then run desktop/build.sh to rebuild the Windows EXEs.
"""
import hashlib, pathlib
root = pathlib.Path(__file__).resolve().parent
app = (root / 'src' / 'app.html').read_text(encoding='utf-8')
head, body = app.split('</style>\n', 1)
pwa_head = '''<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1,viewport-fit=cover">
<meta name="description" content="Photo editor with camera profiles, adaptive presets, AI edit, HDR, RAW import, culling and batch export.">
<meta name="theme-color" content="#121418">
<link rel="manifest" href="manifest.webmanifest">
<link rel="icon" type="image/png" href="pwa/icon-192.png">
<link rel="apple-touch-icon" href="pwa/apple-touch-icon.png">
<meta name="mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
<meta name="apple-mobile-web-app-title" content="Lumen Layer">
<style>:root{padding-top:env(safe-area-inset-top,0px);padding-bottom:env(safe-area-inset-bottom,0px)}</style>
'''
html = pwa_head + head + '</style>\n</head>\n<body>\n' + body + ('' if body.endswith('\n') else '\n') + '</body>\n</html>\n'
(root / 'index.html').write_text(html, encoding='utf-8')
version = hashlib.sha256(html.encode()).hexdigest()[:12]
(root / 'sw.js').write_text((root / 'sw.template.js').read_text(encoding='utf-8').replace('__VERSION__', version), encoding='utf-8')
print(f'index.html ({len(html):,} bytes) and sw.js (cache lumen-{version}) written')
