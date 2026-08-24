# Phase 13 — Command Table

> **Session size:** Large · **Depends on:** Phase 12 · **Touches:** `pkg/ui/model.go`, `view.go`, `command_palette.go`

---

## Goal

Collapse three drifting sources of keybinding truth into one table that dispatch, help, and the command palette all read. Every documented-but-broken binding is fixed **by construction**, and the drift class is closed permanently.

## Why every keybinding fix waited for this phase

The `[t]`/`[s]`/`[K]` bugs have been known since the first audit. They were **deliberately not fixed earlier**: this phase regenerates all of it from one source, so any earlier point fix would be thrown away. That is the plan's core ordering rule, and this is its clearest instance.

---

## The problems

### `Update()` is 794 lines

`pkg/ui/model.go:236-1029`. The `tea.KeyMsg` arm alone is `:349-1026` — **678 lines**. Inside: **14 separate switch statements** and **102 `case "..."` clauses**, with the main keybinding switch at `:685-1025` (341 lines, 56 cases).

**Zero use of `bubbles/key`** — verified: `grep -rn 'bubbles/key\|key.Binding\|key.Matches\|help.Model' .` returns nothing. Only `bubbles/textinput` is used. Every binding is a raw string comparison.

### Truth lives in three places, and has already drifted

1. The `Update()` switch (`model.go:685-1025`)
2. A **hand-maintained 63-line raw string literal** (`view.go:623-686`) rendering the help modal
3. `CommandItem.Shortcut` in `command_palette.go` — **also hand-written display strings**

Confirmed drift:

| Documented | Reality |
|---|---|
| `[t]` Time of Day (`view.go:661`) | **Unbound.** `grep 'case "t"' pkg/ui/` → empty |
| `[s]` Fairy Lights (`view.go:659`) | **Unbound** |
| `[K]` Toggle Cafe Chatter (`view.go:672`, `:474`) | Pets the cat (`model.go:966`) |
| `[S]` Animation Speed (`view.go:453`) | Opens the Preset modal (`model.go:806`) |
| — | `CycleAnimSpeed()` (`art/scene.go:474`) reachable only from tests |
| — | `ToggleCafe()` has no main binding at all |

README omits ~25 bindings that exist. Three sources, no reconciliation, no test.

### Context-dependent overloads hidden in `if` bodies

Not expressed as modes — buried in conditionals:

- `f` = task filter **or** fireplace (`model.go:882-896`)
- `p` = task priority **or** previous track (`:932`)
- `d` = delete task **or** sip beverage (`:1017`)
- `k`/`up` = cursor up **or** pet the cat (`:1000-1006`)

### Magic moduli inline

`(m.ActiveTab + 1) % 4` (`:819`), `(m.SpectrumMode + 1) % 5` (`:356`), `(m.Scene.Weather + 1) % 5` (`:899`), `(m.Scene.CatBreed + 1) % 4` (`:910`), `(m.MixerChannelIdx + 1) % 5` (`:430`). Raw enum arithmetic in the UI, with hand-maintained counts that silently break when an enum grows.

---

## Workstreams

### 13.1 — Define the table

```go
type Command struct {
	ID       CommandID
	Keys     []string        // "t", "shift+m", "ctrl+p"
	Name     string          // palette + help label
	Desc     string          // help text
	Category string          // help + palette grouping
	Scope    Scope           // Global | TabArt | TabTasks | ModalScenePicker | ...
	Enabled  func(*Model) bool
	Handler  func(*Model) tea.Cmd
}
```

`Scope` is what models the overloads honestly: `f` becomes two commands with different scopes instead of one `if` inside a case.

**Start from the existing `CommandItem`** (`command_palette.go`) rather than inventing a parallel type — it already has `Name`/`Desc`/`Shortcut`/handler-ish shape and is already wired into the palette. Promoting it is less work and leaves no second registry behind.

### 13.2 — Prefer this over a full `bubbles/key` migration

`bubbles/key` gives `Binding`/`Matches`/`help.Model`. Useful, but it does **not** solve the palette's third copy of the truth. One table that all three consumers read does. Use `key.Binding` inside `Command` if convenient; do not let the migration become the goal.

### 13.3 — Dispatch from the table

`Update`'s key arm becomes: resolve scope from `modals.Top()` + active tab (Phase 12's stack), look up the key, run the handler. The 341-line switch and its 14 nested siblings collapse into a lookup.

Handlers move to named functions — testable in isolation, unlike a `case` body.

### 13.4 — Generate help and the palette from the table

Delete the 63-line literal at `view.go:623-686` (re-derive the anchor; Phase 1 moved it). Render help by grouping the table by `Category`, filtered by `Scope`.

One thing the literal now carries that generated help must keep: its two scene-count lines interpolate `len(art.SceneOrder)`, and `TestSceneCountStringsTrackRegistry` (`pkg/ui/scene_count_test.go`) asserts the rendered help shows the live registry count. Regenerate the count, do not drop it.

`GetCommandList()` returns the table filtered by scope. This also fixes a real inefficiency: today `GetCommandList(&m)` is called **on every keystroke** (`model.go:393`) *and* **on every render** (`view.go:1047`), each rebuilding ~46 items with fresh closures — including a closure per preset across all 34 presets (`command_palette.go:26-37`). Nothing is memoized. A static table with a scope filter removes the churn.

### 13.5 — Fix the drift by construction, then assert it

With one source, these resolve themselves:

- `[t]` and `[s]` get real handlers (or leave the help).
- `[K]` gets one meaning; cafe chatter gets a binding or stops being advertised.
- `[S]` label matches behavior.
- `CycleAnimSpeed` becomes reachable. *(Note `SpeedFast` still has no branch in `art/scene.go:615-621` — Phase 10 §10.4 should have resolved it; if not, do it here rather than exposing a "2.0x" that does nothing.)*

Then lock it: **`TestEveryDocumentedKeyIsBound`** and **`TestNoDuplicateKeyInScope`**. Drift becomes a test failure.

### 13.6 — Named cycle helpers

Replace `(x + 1) % N` with `NextWeather()`, `NextCatBreed()` etc., defined next to their enums with the count derived from the enum. Add `TestEnumCyclesCoverAllValues`.

---

## Do NOT touch in this phase

- `Model`'s shape — Phase 12, settled.
- **Layout, widths, modal visuals, `truncate`** — Phase 14. This phase changes what keys *do* and how help is *sourced*, not how anything looks.
- Tick cadence, idle gating — Phase 15.
- Art, audio internals, config.

---

## Verification

```bash
go test ./pkg/ui -count=1
grep -c 'case "' pkg/ui/model.go       # near 0 for key handling
wc -l pkg/ui/model.go                  # Update() well under 200 lines
grep -n 'Keyboard Shortcuts' pkg/ui/view.go   # help renders from the table, no literal
go run .    # press ? and try every documented key
```

New tests:

- `TestEveryDocumentedKeyIsBound` — every table entry has a handler and appears in help. **The regression test for the phase.**
- `TestNoDuplicateKeyInScope` — no key bound twice within one scope.
- `TestHelpMatchesCommandTable` — help text is generated, not authored.
- `TestPaletteMatchesCommandTable` — the third source is gone.
- `TestEnumCyclesCoverAllValues` — cycling any enum visits every value and returns to start.
- `TestScopedOverloads` — `f` in the tasks tab filters; `f` in the art tab toggles the fireplace.

## Exit criteria

- [ ] One command table; dispatch, help, and palette all read it
- [ ] `Update()`'s key arm is a lookup, not a switch cascade
- [ ] Zero hand-written help strings; zero hand-written palette shortcuts
- [ ] `[t]`, `[s]`, `[K]`, `[S]`, cafe chatter, anim speed all correct
- [ ] Overloads expressed as scopes, not `if` bodies
- [ ] No inline `% N` enum arithmetic in `pkg/ui`
- [ ] `make check` green

## Re-validation checklist for Phase 14

1. **Is `view.go` materially smaller** after the help literal is gone? Phase 14 works in that file; note the new size and anchors.
2. **Did the palette rebuild churn actually go away?** Phase 14 touches render paths — confirm `GetCommandList` is no longer called per render.
3. **Any binding still needing a UI affordance?** If `[t]`/`[s]` are newly live, Phase 14 should surface them in the sidebar/quick deck.
4. Re-derive `pkg/ui/view.go` line anchors before trusting Phase 14's citations — this phase moves them substantially.
