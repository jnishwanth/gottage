# Phase 5 — Determinism & Scene Golden Harness

> **Session size:** Medium · **Depends on:** Phase 4 · **Touches:** `pkg/art/scene.go`, new golden test harness

---

## Goal

Make rendering a pure function of state, then build the scene-level golden harness. This is the **verification substrate** for Phases 6 through 10 — without it, the scene API change and 36 scene migrations cannot be proven safe.

## Why this phase is here

Phases 6-10 rewrite scene code wholesale. There is currently **no way to tell whether a scene refactor changed the picture.** The only art assertions in the repo are `len(rendered) != 0`.

It comes after Phase 4 because the canvas serializer must be settled first — otherwise every scene golden regenerates when the encoder changes.

---

## The three blockers

Rendering cannot be golden-tested today. All three must go:

**1. Wall-clock RNG seed** — `pkg/art/scene.go:312,318`:

```go
rng: rand.New(rand.NewSource(time.Now().UnixNano()))
```

No injection point, no `SetSeed`.

**2. `time.Now()` inside render** — `EffectiveTimeOfDay()` (`scene.go:602`) reads `time.Now().Hour()` when `TimeAuto`, so identical state renders differently depending on the hour. `View()` also reads `time.Now()` for the clock badge (`pkg/ui/view.go:204,708`).

**3. RNG consumed *during* render** — `renderShaderEffect` calls `cs.rng.Intn(10)` at `scene.go:1032` and `:1062`. Two renders of the same state differ.

Plus a fourth, subtler one: **`RenderFrame` mutates state.** `cs.TransitionStep++` at `scene.go:845`. Calling `View()` twice for one tick produces different output *and* burns a transition step — and `renderZenView` (`view.go:695`) and `renderArtView` (`view.go:250`) each call `RenderFrame()` independently.

---

## Workstreams

### 5.1 — Injectable seed

Add a seed to the `CottageScene` constructor (an options struct or `NewCottageSceneWithSeed`). Default stays wall-clock for the real app; tests pass a fixed seed.

### 5.2 — Injectable clock

Introduce a `Clock` interface (`Now() time.Time`) on `CottageScene`, defaulting to wall-clock. `EffectiveTimeOfDay()` reads it. Tests inject a fixed instant.

Do the same for `pkg/ui/view.go`'s clock badge — otherwise Phase 12+ UI goldens are impossible. Keep the UI change minimal: one field, one call site pair.

### 5.3 — RNG out of the render path

Move randomness into `Update()` by precomputing values onto the particle. `renderShaderEffect`'s `cs.rng.Intn(10)` becomes a value stamped on the particle when it is spawned or stepped.

This is also a latent correctness fix: `Particle` already carries fields the render path ignores. `Particle.Vx/Vy` are **discarded** for overlay and shader particles (`scene.go:698-712`, `:739-749` use hardcoded per-mode deltas), so `PetCat`'s carefully computed `Vx: math.Cos(angle)*speed` (`scene.go:436`) does nothing. `Particle.Char` was deleted in Phase 1 for the same reason. Wire the fields the struct already has rather than adding new ones.

### 5.4 — `RenderFrame` becomes non-mutating

Move `TransitionStep++` (`scene.go:845`) into `Update()`. After this, `RenderFrame()` twice with no `Update()` between must produce identical bytes.

### 5.5 — The scene golden harness

Per-scene golden files with `-update`. For each registered scene: fixed seed, fixed clock, N `Update()` steps, render, compare.

Design decisions that matter:

- **Store what a human can inspect.** Raw ANSI is unreadable in a diff. Store a compact indexed-color text form, or a PNG plus a hash. Phases 9/10 involve deliberately changing goldens 36 times; a reviewable format is the difference between real review and rubber-stamping.
- **Cover the axes that matter:** at minimum `{dawn, day, sunset, night} × {lamp on/off} × {fire on/off}` for a representative subset, and one frame per scene for all of them. Do not cross the full product — that is thousands of files.
- **Guard the runtime.** `pkg/art` is already 17s (**222s under `-race`**). Add `testing.Short()` guards and keep the default set small.

### 5.6 — Wide-glyph hazard

`DrawGlyph` assumes one column, but emoji are double-width. `scene.go:903` draws `'💧'` and `scene.go:1053` draws `'🌸'` — each **shifts every subsequent cell on that row by one column**. Goldens will lock this in, so fix it now: either reject wide runes in `DrawGlyph` or account for the advance. Add `TestDrawGlyphRejectsWideRunes`.

---

## Do NOT touch in this phase

- **Scene art content.** Goldens must first capture the *current* picture so Phase 6's mechanical interface change can be proven byte-identical. Changing art here destroys that baseline.
- `SceneInterface`'s shape — Phase 6.
- Palette (Phase 7), canvas serializer (Phase 4, done), tick rate / idle gating (Phase 15).
- `Model`, keybindings, layout.

---

## Verification

```bash
go test ./pkg/art -run 'TestDeterministic|TestGolden' -count=5   # stable across repeats
go test ./pkg/art -count=1
grep -n 'time.Now()' pkg/art/*.go        # only inside the default Clock impl
grep -n 'rng\.' pkg/art/scene.go         # no hits inside any render* function
```

New tests:

- `TestRenderIsDeterministic` — two scenes, same seed and fixed clock, N `Update`+`Render` cycles, byte-identical. **The regression test for the phase.**
- `TestRenderFrameIsIdempotent` — two `RenderFrame()` with no `Update()` produce identical bytes. Proves 5.4.
- `TestSceneGoldens` — the harness, with `-update`.
- `TestEffectiveTimeOfDayUsesInjectedClock`.
- `TestDrawGlyphRejectsWideRunes`.

## Exit criteria

- [ ] Seed and clock injectable; zero `time.Now()` / `rng.` calls inside render paths
- [ ] `RenderFrame` mutates nothing
- [ ] Golden harness covering every registered scene, in a human-reviewable format
- [ ] `-count=5` stable
- [ ] Default `go test ./pkg/art` not materially slower than before
- [ ] `make check` green

## Re-validation checklist for Phase 6

1. **Are goldens actually stable on a second machine / after `-count=10`?** Any residual flake means a nondeterminism source was missed, and Phase 6's "goldens stay byte-identical" gate becomes meaningless.
2. **How long does the golden suite take?** Phase 6 runs it after every one of 37 interface edits. If it is slow, add a fast subset target to the `Makefile` now.
3. **Is the golden format reviewable?** Phases 9/10 will regenerate all of them deliberately. Confirm a human can read the diff before relying on it.
4. Phase 6 changes `SceneInterface` across 37 implementations. Confirm the golden set covers **every** registered scene, not a sample — a scene without a golden is a scene whose refactor is unverified.
