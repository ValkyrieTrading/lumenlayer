# How to use Lumen Layer Studio

## 1. Install and open

1. On the GitHub page of this repository click **Code → Download ZIP** and unzip it anywhere on your PC.
2. Double-click **`index.html`**. It opens in your default browser.
   - If it opens in another app, right-click it → **Open with** → Chrome, Edge, Firefox or Safari.
   - Tip: bookmark the page or pin the tab so you can come back to it quickly.

The editor opens with six sample photos in the **Library** strip at the bottom. Three are good shots and three are deliberately flawed, so you can try the sorting straight away.

## 2. The screen at a glance

| Area | What it does |
|---|---|
| **Top bar** | Import photos, Undo / Redo, Auto Tone, Before / After, Shortcuts, Export |
| **Tool options bar** | Settings for the current tool (brush size, crop ratio, and so on) |
| **Tool rail** (far left) | Select, Hand, Crop, Spot Heal, Paint, Masking brush |
| **Presets & History** (left) | One-click looks, and every step you have taken |
| **Canvas** (centre) | Your photo. Scroll to zoom, hold Space and drag to pan |
| **Library** (under the canvas) | All imported photos with their score and verdict |
| **Histogram & Develop** (right) | The editing sliders |
| **Layers** (bottom right) | Masked adjustments, paint and fill layers |

## 3. Import photos

- Drag one or many files onto the window, click **Import photos** / **Add photos**, press `Ctrl+O` (`⌘O` on Mac), or paste an image.
- Supported: **JPEG, PNG, WebP** and camera **RAW** (CR2, CR3, NEF, ARW, DNG, RAF, ORF, RW2, PEF, SRW and others).
- The first photo opens immediately; the rest are scored in the background ("Scoring…" on each card).

## 4. Sort good and bad photos (culling)

Every photo gets a **score from 0 to 100** and a verdict, shown by the colored bar on its card:

| Verdict | Score | Meaning |
|---|---|---|
| 🟢 **Keep** | 68 – 100 | Sharp and well exposed |
| 🟠 **Review** | 45 – 67 | Usable but worth a second look |
| 🔴 **Reject** | 0 – 44 | Blurry, badly exposed, very noisy or flat |

The reason is written on each card, for example "Soft or shaken" or "Underexposed". Hover over a card to see the details.

- **Filter** with All / Keep / Review / Reject, and **sort** by Best first, Weakest first or Import order.
- **Override** the verdict yourself: `P` marks Keep, `X` marks Reject, `U` clears your flag. You can also use the ✓ / ✕ buttons on a card.
- **Remove rejects** clears rejected photos from the library. Click it twice to confirm.
- `←` / `→` move to the previous or next photo. Each photo remembers its own edits and history.
- **Review with Claude** (only when opened inside Claude) adds a judgement of composition and impact on top of the technical score.

## 5. Edit a photo

### Develop sliders (right panel)

- **Light**: Exposure, Contrast, Highlights, Shadows, Whites, Blacks.
- **Color**: Temp, Tint, Vibrance, Saturation. **Effects**: Dehaze.
- **Tone Curve**: click to add a point, drag to shape it, and double-click a point (or drag it off the graph) to remove it. Use the RGB / Red / Green / Blue tabs for each channel.
- **Color Grading**: drag the dot inside the Shadows, Midtones or Highlights wheel, or click a color swatch. Hold Shift to change strength only.
- **HSL Mixer**: Hue, Saturation and Luminance for 8 color ranges.
- **Double-click any slider name** to reset it. Each section has its own **Reset**.
- **Auto Tone** (top bar) sets a balanced starting point.

### Local adjustments with masks

1. Click **Select Sky** or **Select Subject** in the Layers panel. The detected area glows briefly.
2. A new **Smart Object** layer is selected, and the panel header shows *Editing "Sky Adjustment"*. The Develop sliders now change **only that area**.
3. Click **Edit global** (or Layer 0) to go back to editing the whole photo.
4. To paint your own area, press `K` for the **Masking brush** and paint. `Alt` erases, and `O` keeps the mask overlay visible.
5. Each layer has a blend mode, opacity, mask (None / Sky / Subject / Brush) and Invert. Drag layers to reorder them; the eye icon hides a layer.

### Remove dust and blemishes

Press `B` for **Spot Healing** and paint over the spot. It copies nearby texture and matches the lighting. Press `B` again to switch to the **Paint** brush. Use `[` and `]` to change brush size.

### Crop & rotate

1. Press `C` (or `R`).
2. Pick an **Aspect**: Original, Free, a fixed ratio, or an **Instagram** size (Square 1:1, Portrait 4:5, Portrait 3:4, Story 9:16, Landscape 1.91:1).
3. Drag the corners or edges, drag inside to move the frame, or drag outside it to draw a new one.
4. Use **Straighten** to level the horizon, the rotate buttons for 90° turns, and **Flip H / V** to mirror.
5. Press **Enter** to apply or **Esc** to cancel. Cropping never deletes pixels, so you can come back and change it.

### Presets

- Click a preset card to apply it, or hover over it to preview. Presets set the sliders **and** add masked layers (for example a sunset wash on the sky).
- Save your own look: type a name under the presets and click **Save**. It is remembered by your browser.

### Before / After and undo

- `\` toggles the Before / After split. Drag the round handle to sweep it.
- `Ctrl+Z` undoes and `Ctrl+Shift+Z` redoes (`⌘` on Mac). Click any step in **History** to jump to it.

## 6. Copy your edits to all keepers

1. Edit one photo the way you like it.
2. Click **Copy edits to keepers** in the Library bar (or press `Shift+C`).
3. Tick the groups to copy: Light, White balance & color, Dehaze, Tone curve, Color grading, HSL, Sky/Subject layers, Brush-masked layers, Crop & rotate.
4. Choose **All keepers** or **Keepers + Review**, then click **Paste**.

Each photo gets a "Paste settings from …" step in its own History, so you can undo it per photo. Sky and Subject layers use each photo's own detected sky and subject. **Copy only** stores the settings so you can press `Shift+V` on any photo later.

## 7. Export

Click **Export** (`Ctrl+S`) for the open photo, or **Export keepers** (`Ctrl+Shift+S`) for every Keep photo.

| Option | Choices |
|---|---|
| **What to export** | This photo, or All keepers (one ZIP file) |
| **Format** | JPEG, PNG (lossless) or WebP, with a quality slider |
| **Instagram** | Off (use your crop) or an Instagram size. The frame is placed inside your crop and centred on the subject. Tick *Full resolution* to keep the ratio at native size |
| **Size** | Native (full original resolution), 75%, 50%, 25%, or an exact long edge up to 16,384 px |
| **Watermark** | Off, Text (any wording and color) or Logo (PNG with transparency works best). Choose a position (corners, centre or tiled), size and opacity |

The preview shows the result, including the watermark. When you click Export, the browser saves the file to your **Downloads** folder.

## 8. Keyboard shortcuts

| Key | Action |
|---|---|
| `V` | Select tool (click the sky or subject to pick its layer) |
| `H` / hold `Space` | Hand tool: pan the canvas |
| `C` or `R` | Crop & rotate |
| `B` | Spot Healing (press again for the Paint brush) |
| `K` | Masking brush on / off |
| `[` / `]` | Smaller / larger brush |
| `Alt` while painting | Erase |
| `O` | Pin the mask overlay on or off |
| `\` | Before / After split |
| `Enter` / `Esc` | Apply / cancel crop |
| `P` / `X` / `U` | Mark Keep / Mark Reject / clear your flag |
| `←` / `→` | Previous / next photo |
| `Shift+C` / `Shift+V` | Copy edits to keepers / paste copied edits here |
| `Ctrl+Z` / `Ctrl+Shift+Z` | Undo / redo |
| `Ctrl+S` / `Ctrl+Shift+S` | Export this photo / export all keepers |
| `Ctrl+O` | Import photos |
| `Ctrl+0`, `Ctrl+=`, `Ctrl+-` | Fit, zoom in, zoom out |

On a Mac use `⌘` instead of `Ctrl`.

## 9. Troubleshooting

| Problem | What to do |
|---|---|
| A RAW file shows "No readable preview" | Some RAW files store only a small preview. Convert that file to JPEG or TIFF with your camera's software first |
| The export is slow or fails at a huge size | Pick 75% or 50%, or type a smaller long edge |
| My photos were gone after reopening | Photos are kept only while the tab is open. Export your keepers before closing |
| Fonts look plain | You are offline. Everything still works; the custom fonts just need an internet connection |
| "Review with Claude" is missing | That button only works when the page is opened inside Claude; the on-device scores still apply |
| Sky or Subject missed an area | Select that layer, press `K`, and paint the missing part (Alt to erase) |
