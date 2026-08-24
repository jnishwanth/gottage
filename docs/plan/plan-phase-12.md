# Phase 12 — `Model` Decomposition + Modal Focus Stack

> **Session size:** Large · **Depends on:** Phase 11 · **Touches:** `pkg/ui/model.go`, `view.go`

---

## Goal

Break the 40-field god object into composed sub-models, and replace eight parallel modal booleans with one focus stack — which makes a live user-visible bug impossible by construction.

## Why this comes before the command table (Phase 13)

Phase 13 replaces `Update()`'s dispatch with a command table whose handlers bind to `Model`. If `Model` decomposes *after* that, every handler is rewritten.

The order is: decompose first (mechanical field-path edits — `m.ShowHelp` → `m.modals.Top() == ModalHelp` — while the switch itself survives), then Phase 13 replaces the switch. The reverse order does the same work twice.

---

## The problems

### A 40-field flat struct

`pkg/ui/model.go:47-92`. **~24 of the 40 fields are modal bookkeeping** — eight `Show*Modal bool` plus their index/input companions:

```go
ShowScenePicker    bool
PickerCategoryIdx, PickerSceneIdx int
PickerSearchMode   bool
PickerSearchInput  textinput.Model
ShowThemePicker    bool; PickerThemeIdx int
ShowSceneInfoModal bool
ShowCommandPalette bool
CommandInput       textinput.Model
CommandIdx         int
ShowPresetModal    bool
PresetIdx          int
PresetSaving       bool
PresetInputText    textinput.Model
ShowSoundMixerModal bool; MixerChannelIdx int
ShowSpectrumModal  bool; SpectrumMode SpectrumStyle
ShowHelp           bool
```

There is **no `Modal` interface, no stack, no focus enum**. Any two modals can be true simultaneously.

### The live bug: Update and View disagree on precedence

**Update's cascade** (`model.go:351-682`): Spectrum → SceneInfo → CommandPalette → SoundMixer → Theme → Preset → ScenePicker → ZenMode → AddingTask → main.

**View's chain** (`view.go:139-176`): CommandPalette → Spectrum → SoundMixer → Preset → SceneInfo → Theme → ScenePicker → Help.

Different orders. **With two modals open you type into one and see another.**

And **`ShowHelp` has no guard in Update at all** — it renders at `view.go:174`, but no branch in the key cascade checks it, so every keystroke while help is open fires the real action behind it.

A focus stack makes both impossible: one source of truth for who has focus, read by Update and View alike.

### Untestable construction

`NewModel()` (`model.go:95`) unconditionally constructs `audio.NewCLIAMPController()` (which shells out to `--version`, `cliamp.go:66`) and `audio.NewAmbientManager()` (which `os.MkdirTemp`s and writes WAVs). **Every test in `view_test.go` has real filesystem and subprocess side effects.** Phase 11 introduced an injectable `Player`; this phase uses it.

### Value receiver

`func (m Model) Update(...)` copies all 40 fields including four `textinput.Model` values, on every message — and `tea.WithMouseCellMotion()` (`main.go:78`) emits a message per mouse cell traversed.

**Do not frame this as performance.** It is ~500 bytes a few hundred times a second — noise. Decompose for comprehensibility and testability; those are the real wins, and letting a perf framing drive it would rank it above the focus stack, which fixes an actual user-visible bug.

---

## Workstreams

### 12.1 — Compose sub-models

Sketch:

```go
type Model struct {
	layout  LayoutState     // Width, Height, breakpoint
	nav     NavState        // ActiveTab, ZenMode
	modals  ModalStack      // replaces 24 fields
	toast   ToastState      // StatusMsg, StatusMsgTime
	tasks   TasksState      // AddingTask, TextInput, cursor
	deps    Deps            // Scene, CLIAMP, Ambient, Timer, Tasks, Presets, Inspiration, Beverage
	cfg     *config.Config  // Phase 3
}
```

`FavoriteScenes map[art.SceneType]bool` is user data, not render state — it belongs in `cfg` (Phase 3 already persists it).

### 12.2 — The modal stack

```go
type ModalID int
const (ModalNone ModalID = iota; ModalHelp; ModalScenePicker; ModalThemePicker;
       ModalSceneInfo; ModalCommandPalette; ModalPreset; ModalSoundMixer; ModalSpectrum)

type ModalStack struct { stack []ModalID; state map[ModalID]any }
func (s *ModalStack) Top() ModalID
func (s *ModalStack) Push(ModalID); func (s *ModalStack) Pop()
```

**Update and View both dispatch on `Top()`.** Divergence stops being expressible.

Each modal's index/input state moves into its own struct keyed in the stack. `esc` pops. `ShowHelp` gets a real guard for free.

### 12.3 — Injectable dependencies

`NewModel(deps Deps)` with a default constructor for production. Tests inject Phase 11's fake `Player`, a fake CLIAMP, a temp config dir. **`NewModel()` must construct zero real audio devices, spawn zero subprocesses, and write zero files.**

### 12.4 — Pointer receiver

Switch `Update` to a pointer receiver, or keep the value receiver with a compact struct. Either is fine; do not spend the phase on it.

### 12.5 — Sweep, don't redesign

The 794-line `Update()` **survives this phase**, with mechanical field-path edits. That is intentional: it keeps the diff reviewable and hands Phase 13 a clean target. Resist dismantling the switch here — that is Phase 13's whole job, and doing half of it in each phase makes both unreviewable.

---

## Do NOT touch in this phase

- **Keybindings.** Not one `case` string, not the help text, not `[t]`/`[s]`/`[K]`. **Phase 13** regenerates all of it from one table; fixing bindings here is work Phase 13 throws away. *(This is the single most important no-rework boundary in the plan.)*
- Layout, widths, `truncate`, modal *visual* design — **Phase 14**.
- Tick cadence, idle gating — **Phase 15**.
- Audio internals (Phase 11), art, config format (Phase 3).

---

## Verification

```bash
go test ./pkg/ui -count=1
go test -race ./pkg/ui -count=1
grep -c 'Show[A-Za-z]*Modal\|ShowHelp\|ShowScenePicker' pkg/ui/*.go   # 0
go run .    # open modals, esc out, confirm keys go where you see focus
```

New tests:

- `TestModalPrecedenceConsistent` — for every pair of modals, push both and assert the one receiving keys **is** the one rendering. **The regression test for the phase.**
- `TestHelpModalSwallowsKeys` — with help focused, a key that would otherwise cycle scenes does nothing.
- `TestNewModelIsHermetic` — construct under a temp `HOME`; assert zero subprocesses spawned and zero files written outside it.
- `TestModalStackPushPop` — `esc` unwinds correctly, including nested.

Also worth doing here: `pkg/ui/view_test.go`'s `TestModelViewModals` is a **285-line single function with ~45 unrelated assertions** and mutating shared state, so a failure midway skips everything after it. Two assertions are inert (`view_test.go:158-161` has an empty `if` body; `:62` is `_ = origScene`). Split it as you touch the code it covers.

## Exit criteria

- [ ] `Model` composed; no flat `Show*` booleans
- [ ] One modal stack driving both Update and View
- [ ] Help modal swallows keys
- [ ] `NewModel` hermetic under injected deps
- [ ] `TestModalPrecedenceConsistent` green over all pairs
- [ ] `Update()` still functionally identical (bindings untouched)
- [ ] `make check` green

## Re-validation checklist for Phase 13

1. **Is `Update()` still one big switch?** It should be — Phase 13 needs an intact target. If it was partly dismantled, re-scope Phase 13 to the remainder.
2. **What is the handler signature?** Phase 13's command table binds to it. Nail it down now: `func(*Model) tea.Cmd` or similar.
3. **Do modal-scoped keys need scoped commands?** The stack means a key can mean different things per focus. Phase 13's table must express that — decide the shape before building it.
4. **List the context-dependent overloads that survived:** `f` = task-filter *or* fireplace (`model.go:882-896`), `p` = priority *or* prev-track (`:932`), `d` = delete task *or* sip beverage (`:1017`), `k`/`up` = cursor-up *or* pet the cat (`:1000-1006`). These are hidden inside `if` bodies rather than expressed as modes — Phase 13 has to model them explicitly or reassign them.
