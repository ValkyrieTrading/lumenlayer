# Lumen Layer Studio

**A complete photo editor in one file. It runs in your web browser, with nothing to install and no account, and your photos stay on your computer.**

Lumen Layer Studio takes a shoot from import to finished files:

- **Import** many photos at once, including camera RAW files.
- **Sort** them automatically into Keep, Review and Reject.
- **Edit** with develop sliders, camera color profiles, masked layers and HDR.
- **Copy** your look to every keeper, or let your choice of AI suggest the edit.
- **Export** all keepers in one ZIP: full resolution or Instagram sizes, with an optional watermark and no hidden metadata.

![Lumen Layer Studio](docs/screenshot.png)

## Quick start

### Windows app (recommended)
1. Download **[`dist/LumenLayerStudio.exe`](dist/LumenLayerStudio.exe)**. On that page click **Download raw file**. For Windows on ARM laptops, use `dist/LumenLayerStudio-arm64.exe`.
2. Double-click it. The first time, Windows SmartScreen says *"Windows protected your PC"*, because the app is not code-signed: click **More info → Run anyway**.
3. The app opens in its own window. Drag photos **or whole folders** onto it, or click **Import photos**. You can also drop a folder straight onto the EXE (or a desktop shortcut to it).

There is nothing to install. It is one ~7 MB file that opens through Microsoft Edge (built into Windows 10/11) and closes when you close the window. To rebuild it from source, see [`desktop/README.md`](desktop/README.md).

### iPhone, iPad and Android
Lumen Layer installs as an app from your phone's browser. It gets a home-screen icon, opens full screen and works offline. You need the web address where it's published; see [Publish it for phones](#publish-it-for-phones) below.

- **iPhone / iPad:** open the address in **Safari** → tap **Share** → **Add to Home Screen** → **Add**. You can also tap **Install app** in the app for these steps.
- **Android:** open the address in **Chrome** → tap **Install app** in the app, or **⋮ → Install app**. Once installed, you can share photos from Gallery straight into **Lumen Layer**.

On phones, use two fingers to zoom and pan. Exports open the share sheet, where you tap **Save Image** to put them in Photos.

### Any computer (browser)
1. On this page click **Code → Download ZIP**, then unzip it.
2. Double-click **`index.html`** to open it in Chrome, Edge, Firefox or Safari.
3. Drag your photos onto the window, or click **Import photos**.

Read **[HOW-TO-USE.md](HOW-TO-USE.md)** for the step-by-step guide and all keyboard shortcuts.

## What it can do

### Edit
| | |
|---|---|
| **Camera & film profiles** | Camera Standard, Neutral, Portrait, Landscape, Vivid and Flat; Chrome Slide, Saturated Slide, Warm Negative, Pastel Negative and Cinema film looks; Monochrome with red and green filters. Each has an Amount slider. |
| **Light & color** | Exposure, contrast, highlights, shadows, whites, blacks, temperature, tint, vibrance, saturation and dehaze |
| **Color fixes** | Auto white balance, **Pick neutral** (click something grey or white), and As shot |
| **Detail & HDR** | HDR tone mapping, clarity, sharpening, and luminance and color noise reduction |
| **Tone & color tools** | Point tone curve (RGB, R, G, B), 3-way color grading wheels, 8-color HSL mixer, live histogram |
| **Enhance** | One click sets exposure, color, noise and detail for the photo, in a Natural, Vivid, Bright, Moody or HDR style, with a written reason for every change |
| **Local edits** | One-click **Select Sky** and **Select Subject**, a masking brush, and Smart Object layers with their own sliders |
| **Retouching** | Spot Healing brush, paint brush, blend modes, fill and vignette layers |
| **Crop & rotate** | Free or fixed ratios, every Instagram size, straighten, 90° rotation and flips |

### Presets that work across very different photos
There are **31 built-in presets** in 10 categories: **Fix, Portrait, Landscape, Street, Night, Food, Interiors, B&W, Film,** and your own **Saved** looks.

Most of them are **adaptive**. They measure each photo's exposure, white balance and noise first, then add the look on top, so a single preset suits bright, dark, warm and cool shots alike.

The **Fix** presets rescue common problems:
- indoor orange light
- shade and overcast
- underwater blue
- dark photos
- hazy photos
- grey snow and sand

### Manage a whole shoot
- **Batch import:** drag in photos or whole folders (subfolders included), or drop them onto the Windows EXE. Supports JPEG, PNG, WebP, TIFF and **RAW** (CR2, CR3, NEF, ARW, DNG, RAF, ORF, RW2, PEF, SRW and more). RAW+JPEG pairs are imported once.
- **Automatic culling:** every photo is scored for focus, exposure, contrast and noise, then sorted into Keep, Review or Reject. You can override any verdict.
- **Merge to HDR:** combines 2–7 bracketed exposures, with automatic alignment for hand-held shots.
- **Copy edits to keepers** and **apply a preset to all keepers:** each photo keeps its own undo history.

### AI edit and AI review, with the AI of your choice
- **AI edit** asks an AI to choose the slider settings for a photo. **AI review** asks it to judge which photos are worth keeping.
- **Choose any provider:**
  - Claude (built in when the page is opened inside Claude)
  - OpenAI
  - Anthropic API
  - Google Gemini
  - OpenRouter
  - **Ollama** or **LM Studio**, running **locally on your own PC**: free and private
  - any OpenAI-compatible server
- **What the AI changes:** only slider values. Your photo is processed by the app's standard adjustments, and no pixels are generated.

### Export
- JPEG, PNG or WebP, at **native resolution** (up to 16,384 px) or at any size.
- **Instagram sizes:** 1:1, 4:5, 3:4, 9:16 and 1.91:1. The frame is centred on the subject.
- **All keepers in one ZIP**, with an optional **text or logo watermark**.
- **No metadata in exported files:** no GPS location, no camera data, and no AI or C2PA tags.

## Good to know

- **Privacy:** photos never leave your computer. The only exception is when you use AI edit or AI review with a cloud provider: that provider then receives a 1024 px preview. With Ollama or LM Studio, nothing leaves your PC.
- **Internet:** needed only for fonts, one styling library and cloud AI. Offline, everything else still works.
- **Imported photos last only while the tab is open.** Export your keepers before closing. Presets, watermark and AI settings *are* remembered by your browser.
- **RAW files** open through the full-resolution JPEG that cameras store inside the RAW file. A browser can't process raw sensor data, so the extra highlight and shadow range of a dedicated RAW converter isn't available.
- **Sky / Subject detection and photo scoring** use simple rules that run on your computer, not a trained model. They work well on clear scenes. Use the masking brush (`K`) or your own flags to correct them.
- **AI providers other than the built-in Claude** work in the downloaded app (`index.html`), not inside Claude itself. API keys are stored only in your browser, and only if you tick "Remember".
- **Social media AI labels:** exports carry no AI or provenance tags, and AI edit only sets ordinary adjustments, so there's nothing in the file that marks it as AI-generated. Each platform still applies its own policies and detection.
- **Very large exports** (40 MP and up) can take several seconds and a lot of memory. Merged HDR photos are capped at about 6 MP.

## Publish it for phones

Phones install the app from a web address, so the files need to be online. **GitHub Pages** does this for free; you only need to set it up once:

1. In this repository open **Settings → Pages**. On iPhone, use Safari and tap **aA → Request Desktop Website** to see Settings.
2. Under **Build and deployment**, set **Source** to *Deploy from a branch*, choose **main** and **/ (root)**, and click **Save**.
3. After a minute or two the address appears at the top of that page, for example `https://valkyrietrading.github.io/Nutition-app/`.

GitHub Pages is free for **public** repositories. For a private repository it needs a paid GitHub plan, so either make the repository public (**Settings → General → Danger Zone → Change visibility**) or host the files elsewhere. Netlify and Cloudflare Pages also host it for free: upload the repository folder.

**App stores:** this is an installable web app, not an App Store or Play Store download. To list it in the Google Play Store, paste the web address into [PWABuilder](https://www.pwabuilder.com) to generate an Android package. The Apple App Store needs a Mac and a paid Apple Developer account.

## Files

| File | Purpose |
|---|---|
| `dist/LumenLayerStudio.exe` | Windows app (x64): double-click to run |
| `dist/LumenLayerStudio-arm64.exe` | Windows app for ARM laptops |
| `index.html` | The whole application: open it in a browser, or publish it for phones |
| `manifest.webmanifest`, `sw.js`, `pwa/` | What makes it installable on phones: app name and icons, offline support, Android "Share to" |
| `src/app.html` | App source. Run `python3 build-web.py` to rebuild `index.html` and `sw.js` |
| `desktop/` | Source and build script for the Windows app |
| `HOW-TO-USE.md` | Step-by-step guide, AI setup and keyboard shortcuts |
| `docs/screenshot.png` | Screenshot used above |
