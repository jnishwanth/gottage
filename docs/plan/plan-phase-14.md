# Phase 14 — Responsive Layout

> **Session size:** Medium · **Depends on:** Phase 13 (and Phase 4's `RenderTo`) · **Touches:** `pkg/ui/view.go`, `styles.go`

---

## Goal

Make the app usable on a standard 80×24 terminal, and replace hand-computed magic numbers with a real layout system.

## Why this phase is here

It needs Phase 4's `RenderTo(cols, rows)` (the scaling primitive) and Phase 13's settled `view.go` (which loses the 63-line help literal). Doing it earlier would mean laying out around code that Phases 12 and 13 then move.

---

## The problems

### The layout floor exceeds a standard terminal

`pkg/ui/view.go:14-34`:

```go
func (m Model) mainWidth() int {
	w := m.Width - 36
	if w < 78  { return 78 }      // hard floor
	if w > 120 { return 120 }
	return w
}
func (m Model) mainHeight() int {
	h := m.Height - 6
	if h < 32 { return 32 }
	if h > 52 { return 52 }
	return h
}
```

Minimum footprint = **78** (main box) + **2** (spacer, `:67`) + **28** (sidebar, hardcoded at `:619`) ≈ **108 columns**, plus two rounded borders ≈ **112 columns**. Vertically ~32 rows plus chrome ≈ **39 rows**.

**On an 80×24 terminal the main box alone overflows.** Lipgloss does not truncate, so every line wraps and the frame becomes scrolled garbage. There is **no breakpoint logic, no minimum-size guard, no "terminal too small" screen, no sidebar-collapse path.**

Reproduce: `COLUMNS=80 LINES=24 go run .`

Note the irony worth recording: a previous iteration explicitly standardized the canvas to 74×52 *"to eliminate line wrapping in standard terminal windows."* The canvas fits; the chrome around it never did.

### `truncate` is byte-sliced

`view.go:970-975`:

```go
func truncate(str string, maxLen int) string {
	if len(str) <= maxLen { return str }
	return str[:maxLen-3] + "..."
}
```

`len()` is bytes, and it is applied to emoji-bearing strings — `truncate(trackName, 22)` (`:536`), `truncate(drink.Name, 22)` (`:584`), `truncate(selScene.String(), 28)` (`:925`), `truncate(meta.Atmosphere, 32)` (`:928`). Any cut landing mid-rune emits a replacement character. Must use `ansi.StringWidth` / rune-aware truncation — display width, not byte count.

### `bg` is dead in all eight modal renderers

`renderHelpModal(bg string)` (`:622`), `renderPresetModal` (`:748`), `renderThemePickerModal` (`:810`), `renderScenePickerModal` (`:848`), `renderSceneInfoModal` (`:977`), `renderCommandPaletteModal` (`:1046`), `renderSoundMixerModal` (`:1102`), `renderSpectrumModal` (`:1171`). `grep -c '\bbg\b' view.go` = 8 — the signatures only. Each ends with `lipgloss.Place(m.Width, m.Height, ..., modal)`, discarding the composed background. **Modals replace the UI rather than overlaying it.** Either implement overlay or drop the parameter.

### Magic numbers everywhere

Sidebar `Width(28)` (`:619`); `strings.Repeat("─", 24)` ×4 (`:562`, `:578`, `:590`, `:599`); `barWidth := 18` (`:569`); `w - 24` (`:308`); `w - 42` (`:371`); modal widths `76`, `65`, `34`, `96`, `86`, `88`, `74`, `76`; mixer slider `20` (`:1129`); scroll windows `PickerSceneIdx > 8 … startIdx + 11` (`:875-878`) and `CommandIdx > 5 … startIdx + 7` (`:1054-1057`).

Worst: **mouse hit-boxes are hand-computed against the sidebar layout** — `msg.Y <= 2`, `tabSpan := 12`, `Y >= 3 && Y <= 10`, `12..17`, `18..22`, `>= 23` (`model.go:301-338`). Any sidebar edit silently breaks click targets, with no test to catch it.

### `styles.go` is a 6-slot palette, not a token system

Better than average — there **is** a `Theme` struct (`:8-16`), ten themes, `AllThemes` (`:119`), and 18 pre-built lipgloss styles built once by `NewStyles(t)` (`:167-237`). But:

- Slots are presentational (`Border/Title/Accent/Highlight/Subdued/Background`), not roles (`surface`, `danger`, `success`, `focus-ring`). No spacing scale, no radius tokens.
- **`Background` is declared on every theme and never used** — no style calls `.Background(t.Background)`.
- **Hardcoded off-theme hex inside the theme system**: `#ffffff` (`:188`), `#282c34` (`:197`), `#101424` (`:208`, `:223`), `#2d3139` (`:229`). A light theme would be unreadable.
- **Style caching is partly undone.** `m.Styles` is built once per theme change (good), but nine `lipgloss.NewStyle()` calls sit in render paths: `view.go:342` (big clock, every timer frame), `:409` (per task row), `:768` (per preset row), `:826-829` (4 per theme row × 10 themes = 40 per frame), `:932`, `:990`, and `visualizers.go:129-132` (4 per 2D-spectrum frame).

Also `daemonStr` (`view.go:963-968`) is dead code.

---

## Workstreams

### 14.1 — Breakpoints

Define named breakpoints and a `Layout` struct computed once per `WindowSizeMsg`:

- **Compact** (< ~90 cols): sidebar hidden, art scaled via `RenderTo`, chrome minimal. **80×24 must work here.**
- **Standard** (~90-130): sidebar collapsed to icons or narrowed.
- **Wide** (> ~130): current layout; consider 2× integer art scaling — Phase 4 noted `mainWidth()` already allows 120 columns against a 74-wide canvas, so 46 columns of slack sit unused today.
- **Below the floor** (< ~60 cols or < ~16 rows): the "terminal too small" card.

Every width comes from `Layout`. No literal in a render function.

### 14.2 — Wire `RenderTo`

The canvas is created once at `art.NewCottageScene(74, 52)` (`model.go:117`) and `WindowSizeMsg` (`:241-243`) only sets `m.Width`/`m.Height`. Pass the computed art viewport into `RenderTo(cols, rows)`.

Keep the logical canvas at 74×52 — Phase 4 settled that scaling, not reflow, is the mechanism. Scenes migrated in Phases 9/10 are resolution-independent and will look better at scale; un-migrated ones still work.

### 14.3 — Rune-safe truncation

Replace `truncate` with an `ansi.StringWidth`-based implementation. Add `TestTruncateIsRuneSafe` with emoji input asserting `utf8.ValidString`.

### 14.4 — Resolve `bg`

Implement overlay via `lipgloss.Place` over the composed background, or delete the parameter from all eight. Overlay is the better UX (modals stop feeling like page navigations); deleting is honest. Pick one.

### 14.5 — Derive mouse hit-boxes from layout

Hit regions computed from the same `Layout` that renders them. Add `TestMouseHitBoxesMatchLayout` so a sidebar change can never again silently break clicking.

### 14.6 — Finish the style system

- Add semantic roles; use `Background` or remove it.
- Move the five hardcoded hex values into theme slots.
- Hoist the nine in-render `lipgloss.NewStyle()` calls into `Styles`.
- Delete `daemonStr`.

---

## Do NOT touch in this phase

- Keybindings, help content, the command table — Phase 13, settled.
- `Model`'s shape — Phase 12, settled.
- Tick cadence, idle gating — Phase 15.
- `pkg/art` internals. Calling `RenderTo` is wiring, not modification.

---

## Verification

```bash
COLUMNS=80  LINES=24 go run .        # fully usable
COLUMNS=60  LINES=20 go run .        # degrades gracefully
COLUMNS=40  LINES=12 go run .        # "terminal too small" card
COLUMNS=200 LINES=60 go run .        # no stretched chrome
go test ./pkg/ui -count=1
grep -c 'lipgloss.NewStyle()' pkg/ui/view.go pkg/ui/visualizers.go   # 0 in render paths
```

New tests:

- **`TestViewFitsTerminal`** — table over {40×12, 60×20, 80×24, 100×30, 120×40, 200×60} × all tabs × all modals, asserting **every** line satisfies `ansi.StringWidth(line) <= cols` and the line count `<= rows`. **The regression test for the phase**, and the one that would have caught this from day one.
- `TestTruncateIsRuneSafe`
- `TestMouseHitBoxesMatchLayout`
- `TestNoHardcodedHexInStyles`

## Exit criteria

- [ ] 80×24 fully usable; below-floor shows a card, never wrapped garbage
- [ ] `TestViewFitsTerminal` green across all sizes, tabs, and modals
- [ ] Art viewport driven by `RenderTo`
- [ ] Rune-safe truncation
- [ ] `bg` implemented or removed
- [ ] Hit-boxes derived from layout
- [ ] Zero `NewStyle()` in render paths; zero hardcoded hex in `styles.go`
- [ ] `make check` green

## Re-validation checklist for Phase 15

1. **Did wiring `RenderTo` change per-frame cost?** Downsampling adds a pass. Re-run Phase 4's benchmarks against the p99 < 8ms budget before Phase 15 tunes cadence.
2. **Does the art viewport change on resize?** Phase 15's idle gating must treat a resize as a wake event.
3. **Is `View()` still calling `RenderFrame()`?** (`view.go:250`, `:695`) Phase 15 decouples simulation from render and needs to know both call sites still exist.
4. Re-derive `view.go` anchors — this phase moves them again.
