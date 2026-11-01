# How to use Lumen Layer Studio

## 1. Install and open

1. On the GitHub page of this repository click **Code → Download ZIP** and unzip it anywhere on your PC.
2. Double-click **`index.html`**. It opens in your default browser.
   - If it opens in another app, right-click it → **Open with** → Chrome, Edge, Firefox or Safari.
   - Tip: bookmark the page so you can come back to it quickly.

The editor opens with six sample photos in the **Library** strip at the bottom. Three are good shots and three are deliberately flawed, so you can try everything straight away.

## 2. The screen at a glance

| Area | What it does |
|---|---|
| **Top bar** | Import photos, Undo / Redo, **Enhance**, Before / After, Shortcuts, Export |
| **Tool options bar** | Settings for the current tool (brush size, crop ratio, and so on) |
| **Tool rail** (far left) | Select, Hand, Crop, Spot Heal, Paint, Masking brush |
| **Presets & History** (left) | Looks by category, and every step you have taken |
| **Canvas** (centre) | Your photo. Scroll to zoom, hold Space and drag to pan |
| **Library** (under the canvas) | All imported photos with score and verdict, plus the batch buttons |
| **Histogram & Develop** (right) | Profile and editing sliders |
| **Layers** (bottom right) | Masked adjustments, paint and fill layers. Click the **Layers** heading to fold it away and give the sliders more room |

## 3. Import photos

You can bring photos in any of these ways:

| How | Works in |
|---|---|
| Drag **photos or whole folders** onto the app window (subfolders included) | Windows app and browser |
| Drag photos or folders onto **`LumenLayerStudio.exe`** or a desktop shortcut to it | Windows app |
| Right-click a photo in Explorer → **Open with** → choose `LumenLayerStudio.exe` | Windows app |
| Click **Import photos** / **Add photos**, or press `Ctrl+O` (`⌘O` on Mac) | Windows app and browser |
| Paste an image (`Ctrl+V`) | Windows app and browser |

- If the Windows app is already open, photos dropped on the EXE or its shortcut are added to the open window, not to a second copy.
- Supported: **JPEG, PNG, WebP, TIFF, BMP, GIF, AVIF** and camera **RAW** (CR2, CR3, NEF, ARW, DNG, RAF, ORF, RW2, PEF, SRW and others).
- **RAW+JPEG pairs** (when the camera saves both) are imported once, using the JPEG. It's the same picture that the RAW file stores inside it.
- **HEIC** photos (the iPhone default) can't be opened. Export them as JPEG, or set the iPhone camera to *Settings → Camera → Formats → Most Compatible*.
- Other files in a dropped folder, such as sidecars and videos, are skipped and counted.
- The first photo opens immediately; the rest are scored in the background.

Tip: make a desktop shortcut (right-click `LumenLayerStudio.exe` → **Send to → Desktop (create shortcut)**), then just drop each new shoot's folder onto it.

## 4. Sort good and bad photos (culling)

Every photo gets a **score from 0 to 100** and a verdict, shown by the colored bar on its card:

| Verdict | Score | Meaning |
|---|---|---|
| 🟢 **Keep** | 68 – 100 | Sharp and well exposed |
| 🟠 **Review** | 45 – 67 | Usable but worth a second look |
| 🔴 **Reject** | 0 – 44 | Blurry, badly exposed, very noisy or flat |

- The reason is written on each card. Hover over a card to see the details.
- **Filter** (All / Keep / Review / Reject) and **sort** (Best first, Weakest first, Import order).
- **Override** a verdict: `P` marks Keep, `X` marks Reject, `U` clears your flag. You can also use the ✓ / ✕ buttons on a card.
- **AI review** asks your chosen AI to judge composition and impact as well (see section 8).
- **Remove rejects** clears rejected photos. Click it twice to confirm.
- `←` / `→` move between photos. Each photo remembers its own edits and history.

## 5. Edit a photo

### Quick start: Enhance
Press `E` (or click **Enhance**):
1. The photo is measured, and exposure, white balance, highlights and shadows, contrast, noise reduction and sharpening are set for you. Each change is listed with its reason.
2. Choose a style: **Natural, Vivid, Bright, Moody** or **HDR**. Use **Strength** to make it stronger or gentler.
3. Optional: click **AI edit** to let your chosen AI pick the settings instead.
4. Click **Apply**, or **Cancel** to put everything back. Every change is an ordinary slider you can fine-tune.

### Camera profile
At the top of Develop, choose a **Profile**:
- **Camera:** Standard, Neutral, Portrait, Landscape, Vivid, Flat
- **Film:** Chrome Slide, Saturated Slide, Warm Negative, Pastel Negative, Cinema Teal & Orange
- **Monochrome:** plain, red filter, green filter, soft matte

Use **Amount** to blend it in (0–200%). Your sliders work on top of the profile.

### Fix the colors
In the **Color** section:
- **Auto white balance** removes color casts automatically.
- **◎ Pick neutral** (`W`): click something that should be grey or white, such as a wall, paper or a cloud.
- **As shot** resets white balance.

### Sliders
- **Light:** Exposure, Contrast, Highlights, Shadows, Whites, Blacks.
- **Color:** Temp, Tint, Vibrance, Saturation. **Effects:** Dehaze.
- **Detail & HDR:** **HDR** evens out very bright and very dark areas. Also Clarity, Sharpening, Noise reduction and Color noise.
- **Tone Curve:** click to add a point, drag to shape it, double-click to remove it. Use the RGB / Red / Green / Blue tabs.
- **Color Grading:** drag inside the Shadows, Midtones or Highlights wheel, or click a swatch.
- **HSL Mixer:** Hue, Saturation and Luminance for 8 color ranges.
- **Double-click any slider name** to reset it. Each section also has **Reset**.

### Local adjustments with masks
1. Click **Select Sky** or **Select Subject** in the Layers panel. The detected area glows briefly.
2. The Develop sliders now change **only that area** (the panel header says *Editing "Sky Adjustment"*). Click **Edit global** to go back.
3. Press `K` for the **Masking brush** to paint your own area. `Alt` erases, and `O` keeps the overlay visible.
4. Each layer has a blend mode, opacity, mask (None / Sky / Subject / Brush) and Invert. Drag layers to reorder them.

### Retouch, crop and compare
- **Spot Healing** (`B`): paint over dust or blemishes. Press `B` again for the Paint brush, and use `[` / `]` to change brush size.
- **Crop & rotate** (`C`): choose a ratio or an **Instagram** size, drag the frame, and use **Straighten** to level the horizon. Press **Enter** to apply or **Esc** to cancel.
- **Before / After**: `\` toggles the split. Drag the handle to sweep it.
- **Undo / redo**: `Ctrl+Z` / `Ctrl+Shift+Z`. Click any History step to jump to it.

## 6. Presets

- Choose a category: **Fix, Portrait, Landscape, Street, Night, Food, Interiors, B&W, Film** or **Saved**.
- Hover over a preset to preview it, and click to apply it.
- Presets marked **"Adapts"** measure each photo first (exposure, white balance and noise), then add the look. That's why one preset works on bright, dark, warm and cool photos.
- **→ Keepers** (shown when you hover over a card) applies the preset to every Keep photo, adapting to each one.
- **Save your own look:** type a name and click **Save**. Leave **"Adapt exposure, tone & white balance to each photo"** ticked to make it adaptive too.

Useful starting points:

| Situation | Preset |
|---|---|
| Any photo that needs a general clean-up | **Auto Fix** |
| Orange indoor light | **Fix Indoor Light** |
| Grey, cold day | **Fix Shade & Overcast** |
| Blue-green underwater shots | **Fix Underwater** |
| Too dark or noisy | **Rescue Dark Photo** / **Low Light Clean** |
| Snow or beach looks grey | **Bright Snow & Beach** |
| People | **Soft Skin Portrait**, **Clean Headshot**, **Golden Portrait** |
| Scenery | **Crisp Landscape**, **Moody Mountains**, **Dramatic Sky** |
| City at night | **Night City** |
| Rooms, products, events | **Bright Interior**, **Clean Product**, **Airy Wedding** |

## 7. Work on many photos at once

- **Copy edits to keepers** (`Shift+C`): tick the groups to copy (Light, Color, Profile, Detail & HDR, Tone curve, Color grading, HSL, Sky/Subject layers, Brush-masked layers, Crop), choose **All keepers** or **Keepers + Review**, then click **Paste**. Each photo can undo it in its own History. **Copy only** plus `Shift+V` pastes onto one photo at a time.
- **Merge to HDR:** Ctrl/⌘-click 2–7 bracketed exposures in the library (or use the ◎ button on each card), then click **Merge to HDR**. The merged photo appears as a new library photo.

## 8. Set up AI edit and AI review

Click the **⚙** button in the Enhance panel. The AI settings also open the first time you use AI edit or AI review.

| Provider | What you need |
|---|---|
| **Claude (built in)** | Nothing. This works only when the page is opened inside Claude |
| **OpenAI** | An API key from platform.openai.com |
| **Anthropic API** | An API key from console.anthropic.com |
| **Google Gemini** | An API key from aistudio.google.com |
| **OpenRouter** | One key from openrouter.ai gives access to many models |
| **Ollama (local)** | Free and private. See below |
| **LM Studio (local)** | Free and private. Load a vision model, start the local server, turn on **Enable CORS** |
| **Custom** | Any OpenAI-compatible server address (vLLM, LocalAI, Jan, a company gateway) |

The model you choose **must accept images**: a vision model.

**Set up Ollama (runs on your PC)**
1. Install Ollama from ollama.com.
2. Download a vision model: `ollama pull llama3.2-vision` (or `qwen2.5vl`, `gemma3`).
3. Start it so this page may connect:
   - Windows (PowerShell): `$env:OLLAMA_ORIGINS="*"; ollama serve`
   - macOS / Linux: `OLLAMA_ORIGINS="*" ollama serve`
4. In AI settings choose **Ollama**, click **List models**, pick your model, then click **Test connection** and **Save**.

**Good to know**
- AI providers other than the built-in Claude work in the **downloaded app** (`index.html`), not inside Claude.
- API keys are kept only in your browser, and only if **Remember on this device** is ticked.
- The AI receives a 1024 px preview and the photo's measurements. With Ollama or LM Studio, nothing leaves your PC.
- The AI only chooses **slider values**. Your photo is processed by the app's standard adjustments, and nothing is generated.

## 9. Export

Click **Export** (`Ctrl+S`) for the open photo, or **Export keepers** (`Ctrl+Shift+S`) for every Keep photo.

| Option | Choices |
|---|---|
| **What to export** | This photo, or All keepers (one ZIP file) |
| **Format** | JPEG, PNG (lossless) or WebP, with a quality slider |
| **Instagram** | Off, or Square 1:1 · Portrait 4:5 · Portrait 3:4 · Story 9:16 · Landscape 1.91:1. The frame is centred on the subject. Tick *Full resolution* to keep the ratio at native size |
| **Size** | Native (full original resolution), 75%, 50%, 25%, or an exact long edge up to 16,384 px |
| **Watermark** | Off, Text or Logo. Choose a position (corners, centre or tiled), size and opacity |

Exported files contain **no metadata**: no GPS location, no camera data and no AI or C2PA tags. They are saved to your **Downloads** folder.

## 10. Keyboard shortcuts

| Key | Action |
|---|---|
| `V` | Select tool (click the sky or subject to pick its layer) |
| `H` / hold `Space` | Hand tool: pan |
| `C` or `R` | Crop & rotate |
| `B` | Spot Healing (press again for Paint) |
| `K` | Masking brush on / off |
| `E` | Enhance & AI edit panel |
| `W` | Pick neutral point (white balance) |
| `[` / `]` | Smaller / larger brush |
| `Alt` while painting | Erase |
| `O` | Pin the mask overlay |
| `\` | Before / After split |
| `Enter` / `Esc` | Apply / cancel crop (Esc also closes panels) |
| `P` / `X` / `U` | Mark Keep / Mark Reject / clear your flag |
| `←` / `→` | Previous / next photo |
| Ctrl/⌘-click a card | Select photos for Merge to HDR |
| `Shift+C` / `Shift+V` | Copy edits to keepers / paste copied edits here |
| `Ctrl+Z` / `Ctrl+Shift+Z` | Undo / redo |
| `Ctrl+S` / `Ctrl+Shift+S` | Export this photo / export all keepers |
| `Ctrl+O` | Import photos |
| `Ctrl+0`, `Ctrl+=`, `Ctrl+-` | Fit, zoom in, zoom out |

On a Mac use `⌘` instead of `Ctrl`.

## 11. Troubleshooting

| Problem | What to do |
|---|---|
| "Couldn't reach Ollama" | Start it with `OLLAMA_ORIGINS="*"` as shown in section 8, and check that the address is `http://localhost:11434` |
| "Couldn't reach LM Studio" | Start its local server and turn on **Enable CORS** |
| AI says the key was rejected | Check the key and that your account has credit or API access |
| "Could not find model" | Use **List models**, or check the exact model name on the provider's site |
| AI edit replies in a strange way | Pick a stronger **vision** model |
| Other AI providers don't work inside Claude | Expected: open the downloaded `index.html` on your PC instead |
| A RAW file shows "No readable preview" | That file stores only a small preview. Convert it to JPEG or TIFF with your camera's software |
| Export is slow or fails at a huge size | Pick 75% or 50%, or type a smaller long edge |
| My photos were gone after reopening | Photos are kept only while the tab is open. Export your keepers before closing |
| Sky or Subject missed an area | Select that layer, press `K`, and paint the missing part (`Alt` erases) |
