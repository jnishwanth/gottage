# Phase 16 — Bug Sweep & Error Surfacing

> **Session size:** Medium · **Depends on:** Phase 15 · **Touches:** repo-wide

---

## Goal

Close out whatever survived, install the error-surfacing channel that would have made this entire plan unnecessary, and widen the lint gate now that the code can pass it.

## Why this phase is last-but-one

By now most original bugs are fixed **by construction** rather than individually: bindings by Phase 13's table, modal precedence by Phase 12's focus stack, preset themes by Phase 3's ID matching, dead scenes by Phase 1's panic. What remains is genuinely point work that nothing later will obsolete — which is exactly why it waited.

**Start by re-deriving the open list.** Do not work from this document's assumptions; verify each item still reproduces.

---

## Workstream 16.1 — Error surfacing

**This is the most important thing in the phase.** The recurring theme of this entire audit is that the app lies quietly:

- `c.Clear(PicoNavy)` when a GIF fails to load — how 20 dead scenes survived 8 commits
- `_ = pm.SaveToFile()` (`features/presets.go:438`, `:447`) — preset saves fail silently
- `hex()` returning black on a bad literal (`palette.go:352`)
- `sh -c "curl ... | afplay -"` reporting "Playing" when neither binary exists (`cliamp.go:265`)
- Ambient toggles flipping bools with no engine behind them
- **42 ignored errors** repo-wide (19 in `ambient.go` alone, including 10 unchecked `binary.Write` calls at `:158-181`)

Build a real channel: an error/warning record the status line surfaces, plus an optional debug log behind a flag or env var. Rate-limit — a per-frame failure must not spam.

Then audit every remaining `_ =` and decide, per site: handle, surface, or document why it is genuinely ignorable. `errcheck` (enabled narrowly in Phase 0) should end this phase with **zero suppressions that lack a reason**.

## Workstream 16.2 — Re-derive and close the open list

Verify each; expect most to be already fixed:

| Original finding | Should have been fixed by |
|---|---|
| `[t]`/`[s]` unbound, `[K]` mislabeled, `[S]` label wrong | Phase 13 |
| `ShowHelp` no key guard | Phase 12 |
| Modal precedence divergence | Phase 12 |
| Built-in presets never apply theme | Phase 3 |
| `truncate` byte-slices emoji | Phase 14 |
| Dead `bg` param ×8 | Phase 14 |
| `FairyLightsOn` unreachable | Phase 6 |
| `SmoothDayNightBlend`, `BeatReactiveParticles`, `SpeedFast` | Phase 10 §10.4 |
| `Particle.Vx/Vy` discarded for overlay/shader particles | Phase 5 §5.3 |
| Preset load clobbers built-ins | Phase 3 |
| Scene picker bypasses `SetSceneWithTransition` (`model.go:582`, `:631`) | Phase 12/13 |
| Unbounded GIF cache | Phase 6 |
| Kitty `r=34` vs 26 char rows | Phase 4 |
| Wide-glyph column shift | Phase 5 §5.6 |

Anything still open is this phase's work.

Then hunt for what the audit never covered — `pkg/features` got the least attention. Known: `NewTaskList` seeds four hardcoded sample tasks (`tasks.go:65-68`), which is odd once tasks persist (Phase 3). `BeverageManager.TakeSip` floors warmth at 25% (`beverage.go:168-177`) with no documented reason. `FocusTimer.Tick` uses wall-clock delta (`timer.go:82-109`) and needs Phase 5's injectable clock to be testable.

## Workstream 16.3 — `main.go`

- `os.Exit(1)` at `main.go:90` **skips the deferred `model.Ambient.Cleanup()`** at `:82-86`, leaking the temp dir on every error exit. (Phase 11 removed temp WAVs, but the pattern is still wrong.)
- `time.Sleep(100 * time.Millisecond)` at `:94`, commented *"Sleep briefly for cleanup"* — a race-y hack. Phase 11 gave `Cleanup()` real cancellation; wait on it properly.
- The `--weather`/`--time` flag handling duplicates a nested string switch **twice** (`:24-35` and `:49-72`). With Phase 3's enum string parsing, both collapse to one call.

## Workstream 16.4 — Widen the lint gate

Phase 0 deliberately enabled only `gofmt`, `govet`, `errcheck`, `ineffassign`, because `funlen`/`gocyclo` could not pass against a 794-line `Update()` and a 2,117-line `init()`. Both are gone.

Enable `funlen`, `gocyclo`, `gocritic`, `revive`, `unparam`, `unused`. Fix what they find, or record an explicit exclusion with a reason. Update `AGENTS.md`.

## Workstream 16.5 — Test suite health

- **Split `TestModelViewModals`** (`pkg/ui/view_test.go:37-321`) if Phase 12 did not — a 285-line function with ~45 unrelated assertions and mutating shared state, where a mid-function failure skips everything after. Two assertions are inert: `:158-161` has an empty `if` body, `:62` is `_ = origScene`.
- Fix env mutation without `t.Setenv` in `terminal_adapter_test.go:9`.
- Confirm no test still asserts `len(output) != 0` as a rendering check:
  ```bash
  grep -rn 'len(.*) == 0\|len(.*) != 0' --include='*_test.go' . | grep -i 'render\|view\|draw'
  ```
- **Runtime budget.** `pkg/art` was 17s and **222s under `-race`**. After Phase 1's curation and Phase 6's caching it should be far better — measure, and add `testing.Short()` guards so `make check` stays fast enough that people actually run it.
- Coverage: `pkg/audio` was 0% (Phase 11), `main` 0%. Set a floor and enforce it in `make check`.

---

## Do NOT touch in this phase

- Architecture. If something needs restructuring, it belongs in a follow-up phase, not smuggled into a sweep.
- Scene art (Phases 9/10), the component library, tokens.
- **Do not add features.** This phase subtracts lies and adds signal.

---

## Verification

```bash
golangci-lint run                                  # clean with the wide set
go test -race ./... -count=1
go test ./... -cover                               # every package above the floor
grep -rn '_ = ' --include='*.go' . | grep -v _test # each remaining site justified
make check                                         # green and fast
```

New tests: `TestErrorsSurfaceToStatusLine` (force a GIF load failure, assert a visible message), `TestNoSilentFallbacks` (asset failure records an error), `TestMainCleanupOnErrorExit`, `TestTimerUsesInjectedClock`.

## Exit criteria

- [ ] Error/warning channel exists, surfaces to the status line, rate-limited
- [ ] Every `_ =` handled, surfaced, or justified in a comment
- [ ] Full original findings list re-derived; every item closed or explicitly deferred with a reason
- [ ] `main.go` cleanup correct; flag parsing deduplicated
- [ ] Wide lint set green
- [ ] No rendering test asserting `len != 0`
- [ ] `-race` clean; coverage floor enforced; `make check` fast

## Re-validation checklist for Phase 17

1. **Final feature inventory** — Phase 17's README must describe what exists *now*. List: scene count, tabs, working audio channels, every binding (from Phase 13's table — the README can be generated from it).
2. **Anything deferred?** It goes in the README's known-limitations section, not silently omitted. Silent omission is the habit this whole plan exists to break.
3. **Is the app genuinely good to use?** Sit with it for an hour. Phase 17 is documentation; if something is still wrong, fix it here.
