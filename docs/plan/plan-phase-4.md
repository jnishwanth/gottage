# Phase 4 — Canvas Core: Golden, Encoder, `RenderTo`

> **Session size:** Large · **Depends on:** Phase 0 (independent of 1-3) · **Touches:** `pkg/art/canvas.go`, `pkg/art/terminal_adapter.go`

---

## Goal

Rewrite the bottom rendering layer: one parameterized serializer instead of four, an allocation-free hot loop, and `RenderTo(cols, rows)` so the app works on an 80×24 terminal. Gate the whole thing on a byte-level golden written **first**.

## Why this phase is here

`Canvas` is the layer everything else renders through. Every scene, every component, every UI surface. Settling it before Phases 6-10 touch scene code means no scene is ever written against the old API and then rewritten.

It depends only on Phase 0 — it is deliberately independent of the content and config work, so it could run in parallel if you ever want two sessions going.

---

## Workstream 4.1 — The byte golden, FIRST

**Nothing else in this phase starts until this is committed and green.**

Build a `Canvas`, fill it with a fixed synthetic pattern (gradients, glyph overlays, sharp color edges, a full-black region, a full-white region), serialize, compare bytes against a checked-in golden. Add `-update` to regenerate deliberately.

This depends on **none** of the determinism blockers (no RNG, no clock, no scene), which is exactly why it can lead. Phase 5's *scene*-level golden needs the determinism work; this one does not.

**The specific trap it exists to catch.** In `renderHalfBlock` (`pkg/art/canvas.go:386-440`), the run-length color trackers reset to `-1` at every row end (`:435-436`) *and* the row terminates with `\x1b[0m`. Those two facts are **coupled**: the reset is correct precisely because `\x1b[0m` clears SGR state. Remove the reset without removing the `\x1b[0m` and you emit stale-state garbage that still looks plausible in a terminal. A byte golden catches it instantly; a visual check will not.

## Workstream 4.2 — Profile before optimizing

The obvious target is `fmt.Sprintf`. **Measure before assuming.**

`renderHalfBlock` emits up to 2 `fmt.Sprintf` per cell for truecolor SGR:

```go
sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", fgR, fgG, fgB))   // canvas.go:410
sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", botR, botG, botB)) // canvas.go:414
```

At 74×26 cells that is up to **3,848 allocations per frame ≈ 64k/second** at the 60ms tick. And `strings.Builder` is never `Grow`n (`canvas.go:387,443,492,565` are all bare `var sb strings.Builder`) against 20-80KB of output, so it doubles ~15 times per frame, copying the whole buffer each time.

But `RenderFrame` also runs up to **four additional full-canvas sweeps** in `colorful.Color` space — `applyColorFilter` (`scene.go:1074`), `ApplyVignette` (`canvas.go:670`), `ApplyCRTScanlines` (`canvas.go:731`), plus `BlendCanvas` during transitions. `colorful` operations can involve Lab conversions that may dwarf `Sprintf`.

Write `BenchmarkRenderFrame` and `BenchmarkRenderHalfBlock` and take a CPU profile **before** touching the encoder. Optimizing the visible problem first is how a phase spends itself for a 15% win.

## Workstream 4.3 — One serializer instead of four

`renderHalfBlock` (`:386`), `renderASCII` (`:442`), `renderBraille` (`:491`), `renderPhosphor` (`:564`) are four near-identical ~55-line loops differing only in glyph selection. Collapse to one emitter parameterized by a glyph-selection strategy.

Rewrite the hot loop:

- Reused `[]byte` scratch buffer on the `Canvas`, not a fresh `strings.Builder` per frame.
- `strconv.AppendInt` instead of `fmt.Sprintf`.
- `Grow`/pre-size to the measured typical output.
- Carry run-length suppression **across** rows — but only together with removing the per-row `\x1b[0m` (see 4.1).

**Budget: p99 frame < 8ms at 74×52, and < 10 allocations per full frame** (from ~3,800). Record before/after benchmark output in the commit message.

## Workstream 4.4 — `RenderTo(cols, rows)`

Today the canvas is created once at `art.NewCottageScene(74, 52)` (`pkg/ui/model.go:117`) and **never resized** — there is no `Canvas.Resize` method at all. `WindowSizeMsg` (`model.go:241-243`) only sets `m.Width`/`m.Height`. The layout floor is ~112 columns × 39 rows, so **a standard 80×24 terminal renders wrapped garbage.**

**Keep the logical canvas fixed at 74×52** and scale at serialize time:

```go
func (c *Canvas) RenderTo(cols, rows int) string
```

- **Larger terminal:** center at 1×, or **integer** scale 2×. Never non-integer upscale — it destroys pixel art. Note `mainWidth()` (`pkg/ui/view.go:14-21`) already permits up to 120 columns while the canvas uses 74, so 46 columns of slack already exist unused.
- **Smaller terminal:** box-downsample. 74×52 → 60×40 is clean averaging and degrades gracefully.
- **Below ~40 cols × 12 canvas rows:** render a "terminal too small" card.

**Why scaling and not reflow.** 463 of 1,159 draw calls in `scenes_modular.go` pass a literal first coordinate — `ZenTeahouseScene` hardcodes a room height of `34` with no relation to `c.Height`; `CityCafeScene` lays out exactly 8 buildings at `bx := wx + b*7 + 1` against `ww = 58`, so at `ww = 40` buildings 6-8 walk off the edge. Genuine reflow means 36 individual **art** redraws, each needing aesthetic judgment. Box-filtering a `[][]colorful.Color` grid is ~30 lines written once and is correct for every scene simultaneously, forever. The asymmetry is not close.

Phases 8-10 make migrated scenes resolution-independent as a *by-product* of the component library. `RenderTo` is what lets that happen incrementally instead of as a prerequisite.

## Workstream 4.5 — 256-color fallback

`pkg/art/terminal_adapter.go` detects Kitty and iTerm2 graphics protocols but **never checks `COLORTERM` for truecolor**. On 256-color terminals — tmux's default, older Windows consoles — the 24-bit SGR output is garbage. Add detection and a 256-color quantized path through the same parameterized emitter.

Also fix two bugs here while the file is open:
- `EncodeKittyGPU(cs.Canvas, 74, 34)` (`pkg/art/scene.go:1158`) hardcodes `rows=34` while the canvas has 26 char rows, and `terminal_adapter.go:110-114` emits 33 newlines to match the wrong count.
- `TestEncodeKittyGPU` (`terminal_adapter_test.go:40`) asserts `s=320,v=180,c=74,r=34` — a header shape the app **never emits** (it passes a 74×52 canvas, producing `s=148,v=104,...`), which is why it never caught the row mismatch.
- The GPU path allocates ~250KB/frame (~4MB/s): an RGBA buffer, a base64 string, then a re-chunked builder with a `Sprintf` per 4KB chunk (`:104,106`).

---

## Do NOT touch in this phase

- **Scene code.** `RenderFrame`'s step order, any `Render*` method, any scene art. Phase 6 owns the interface, 9/10 own the art.
- **`pkg/ui` layout.** Adding `RenderTo` does not mean wiring it up — **Phase 14** owns responsive layout and will call it. Leave the current 74×52 call site alone.
- Determinism (Phase 5), palette (Phase 7), GIF caching (Phase 6).

---

## Verification

```bash
go test ./pkg/art -run TestCanvasGolden -count=1      # byte-identical
go test ./pkg/art -bench 'BenchmarkRender' -benchmem  # record before/after
go test ./... -count=1
```

New tests:

- `TestCanvasGolden` — the encoder contract. **Regenerating it must be a deliberate `-update` commit, never incidental.**
- `TestRenderToScaling` — table over target sizes; assert output line count and `ansi.StringWidth` per line match the request exactly.
- `TestRenderToTooSmall` — below the floor, returns the card, never a wrapped frame.
- `TestRenderModesAgree` — all four glyph modes produce correctly-sized output through the unified emitter.
- `Test256ColorFallback` — with `COLORTERM` unset, output contains no `38;2;` sequences.

Target: `BenchmarkRenderHalfBlock` allocs/op from ~3,800 to **< 10**.

## Exit criteria

- [ ] `TestCanvasGolden` committed and green; `-update` documented
- [ ] Four serializers collapsed to one; allocs/op < 10; p99 frame < 8ms
- [ ] Profile captured and the actual bottleneck named in the commit message
- [ ] `RenderTo(cols, rows)` handles upscale, downscale, and too-small
- [ ] 256-color fallback; Kitty row-count bug fixed; its test asserts a real shape
- [ ] `make check` green

## Re-validation checklist for Phase 5

1. **Is `TestCanvasGolden` genuinely byte-stable across runs?** If it flakes, something non-deterministic leaked into the canvas layer and Phase 5's scope grows.
2. **What did the profile actually say?** If post-processing sweeps dominated rather than `Sprintf`, Phase 5 and 15 should re-weight accordingly — and `applyColorFilter`/`ApplyVignette`/`ApplyCRTScanlines` may deserve their own optimization pass.
3. **Did the encoder rewrite change output bytes?** If yes, that was deliberate — confirm the golden was regenerated with `-update` in its own commit and visually reviewed, not silently.
4. Phase 5 injects a seed and clock into `CottageScene`. Confirm this phase left `scene.go` untouched apart from the `EncodeKittyGPU` call.
