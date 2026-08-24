# Phase 6 — `SceneInterface` v2

> **Session size:** Large (37 implementations) · **Depends on:** Phase 5 · **Touches:** `pkg/art/scene_interface.go`, `scenes_modular.go`, `gif_loader.go`

---

## Goal

Replace the scene contract with one that can carry state, has a lifecycle, and expresses resolution independence. Mechanical and wide — **verified by Phase 5's goldens staying byte-identical.**

## Why this phase is here

Every later phase touching scenes — the component library (8), the 36 migrations (9/10) — writes code against this interface. Changing it *after* they land means rewriting them. Changing it now is a mechanical sweep protected by goldens.

It must come after Phase 5 for exactly that reason: without goldens, a 37-implementation sweep is unverifiable.

---

## The current contract

`pkg/art/scene_interface.go:5-22`:

```go
type SceneInterface interface {
	ID() SceneType
	Name() string
	WindowBounds() Rect
	RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64,
	          stars []struct{ X, Y int; Phase float64 })
	RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy,
	                 lampOn, fireOn bool, catBreed CatBreed)
	ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool)
}
```

Problems, each verified:

1. **An anonymous struct type in the signature.** `[]struct{ X, Y int; Phase float64 }` is repeated verbatim across **37 method signatures** — an un-nameable type no caller can declare a variable of. Worse, it is **mostly ignored**: only 7 of 36 handcrafted scenes read `stars`; the rest synthesize their own field from `frameCount` (three different star loops at `:3658-3663`, `:4085-4086`, `:5230-5231`, each with a different magic modulus).
2. **Positional booleans and enums.** `tod`, `lightning`, `frameCount`, `buddy`, `lampOn`, `fireOn`, `catBreed`. Adding one state flag means editing 37 signatures × 3 methods. This is why `FairyLightsOn`, `SmoothDayNightBlend`, and `BeatReactiveParticles` are rendered in the sidebar and **never reach any scene**.
3. **No lifecycle.** There is nowhere to load a resource on scene entry, which is why `AnimatedGifScene.RenderSky` calls `LoadAnimatedGif` **every single frame** — taking a global mutex and building a `fmt.Sprintf` cache key 16.7 times per second.
4. **No resolution contract.** `WindowBounds()` returns absolute pixels against an assumed 74×52.

---

## Workstreams

### 6.1 — A render context struct

```go
type RenderCtx struct {
	Canvas     *Canvas
	Frame      uint64
	TimeOfDay  TimeOfDay
	Lightning  bool
	LampOn     bool
	FireOn     bool
	FairyLights bool          // currently unreachable — wire it here
	CatBreed   CatBreed
	Buddy      *FocusBuddy
	Stars      []Star         // named type, replacing the anonymous struct
	Bounds     Rect           // canvas extent; scenes stop assuming 74x52
}
```

```go
type SceneInterface interface {
	ID() SceneType
	Name() string
	WindowBounds() Rect
	Activate(w, h int) error       // NEW: lifecycle
	RenderSky(ctx *RenderCtx)
	RenderForeground(ctx *RenderCtx)
	ApplyLighting(ctx *RenderCtx)
}
```

Add `type Star struct { X, Y int; Phase float64 }`. Adding future state becomes a struct field, not a 111-signature edit.

Wire `FairyLightsOn` through while you are here — it is a toggle rendered in the sidebar (`pkg/ui/view.go:465`) that no render code has ever read. Either make it work or delete the UI; **Phase 16 must not still be looking at it.**

### 6.2 — `Activate(w, h)` and the GIF cache

This is the highest-value part of the phase.

`CottageScene.SetSceneWithTransition` (and the picker path) calls `Activate` on entry. `AnimatedGifScene.Activate` loads and holds a resolved handle; `RenderSky` uses the handle and does no I/O.

Then bound the cache. `AnimatedGifLoader.cache` (`pkg/art/gif_loader.go:25-33`) is a `map[string]*AnimatedGif` with **no eviction, no size cap, no Close**. Each decoded GIF stores every frame as `[][]colorful.Color` at 24 bytes/pixel. Across the pre-curation asset set that measured **~0.96 GB if a user browsed the full scene list**. Phase 1's curation reduces the ceiling; an LRU with an explicit frame-count or byte cap removes it.

Note for Phase 15: GIFs are pre-scaled at load time, so if `RenderTo` scaling and GIF loading both resample, frames get filtered twice and look visibly mushier. Phase 15 threads final target dims to the loader. That **widens the cache key space**, which is precisely why the LRU bound must exist before then.

### 6.3 — Sweep all 37 implementations

Mechanical. After each batch, run the golden suite — **output must stay byte-identical.** Any diff is a mistake in the sweep, not an improvement.

Do **not** improve art, share code, or replace magic numbers while sweeping. That is Phases 8-10, and mixing it in destroys the byte-identical gate that makes this phase safe.

### 6.4 — Window bounds sanity

All GIF registrations shared `winBounds: Rect{X:10, Y:10, W:140, H:60}` — describing a window **larger than the 74×52 canvas**. `Update()` respawns rain at `X = 10 + rng.Intn(140)` and only resets at `X ≥ 150` or `Y ≥ 70` (`scene.go:647-650`), so roughly **half the rain particles sit off-canvas at any instant**, and each spends about a third of its fall below the visible area.

Fix the bounds, and derive particle spawn ranges from the canvas and the active scene's `WindowBounds` rather than the hardcoded `10 + rng.Intn(50)` (`scene.go:574,585`). Add `TestWindowBoundsWithinCanvas` — the existing `TestWindowBounds` (`scene_test.go:25`) only asserts `W > 0 && H > 0`, which is why this was never caught.

Note: this **changes goldens deliberately** (more visible rain). Regenerate with `-update` in a separate, clearly-labelled commit from the mechanical sweep.

### 6.5 — Retire the residual scene switch

`updateSteamParticles` (`scene.go:666-681`) still hardcodes per-scene positions:

```go
mugX := 22.0
if cs.CurrentScene == SceneRainyCity      { mugX = 30.0 }
else if cs.CurrentScene == SceneForestCabin { mugX = 28.0 }
```

A switch over 2 of ~76 scenes; every other scene draws mug steam inline in `RenderForeground` (~25 distinct implementations). Move emitter positions into scene data. Phase 8 turns steam into a shared component.

---

## Do NOT touch in this phase

- **Scene art.** No new visuals, no shared helpers, no magic-number cleanup. Byte-identical goldens are the safety net; art changes remove it.
- Palette (Phase 7) — do not replace `hex()` literals during the sweep.
- `Delays` (Phase 15), `showBuddy` (Phase 8/9).
- `pkg/ui`, except where `FairyLightsOn` wiring requires a one-line call-site change.

---

## Verification

```bash
go test ./pkg/art -run TestSceneGoldens -count=1   # byte-identical after 6.1/6.3
go test ./... -count=1
grep -c 'struct{ X, Y int; Phase float64 }' pkg/art/*.go   # 0
grep -n 'LoadAnimatedGif' pkg/art/                          # only inside Activate
```

New tests:

- `TestSceneActivateLoadsOnce` — count loader invocations across 100 frames; must be 1.
- `TestGifCacheBounded` — activate more distinct scenes than the cap; assert `len(cache) <= cap`. **The regression test for the ~1GB problem.**
- `TestWindowBoundsWithinCanvas` — every scene's bounds fit the canvas.
- `TestRenderCtxCarriesAllToggles` — `FairyLights` reaches a scene.

Memory check:

```bash
# activate every scene in sequence, then assert heap stays bounded
go test ./pkg/art -run TestGifCacheBounded -memprofile /tmp/m.out && go tool pprof -top /tmp/m.out | head
```

## Exit criteria

- [ ] `RenderCtx` + named `Star`; zero anonymous structs in signatures
- [ ] `Activate(w, h)` on all 37; zero `LoadAnimatedGif` calls in any render path
- [ ] GIF cache bounded and proven by test
- [ ] Goldens byte-identical after 6.1/6.3; the 6.4 change regenerated in its own labelled commit
- [ ] Every scene's `WindowBounds` fits the canvas; particles derive spawn from bounds
- [ ] `FairyLightsOn` either functional or removed from the UI
- [ ] `make check` green

## Re-validation checklist for Phase 7

1. **Did the sweep stay byte-identical?** If goldens changed outside 6.4, find out why before Phase 7 layers palette changes on top of an unexplained diff.
2. **Is `RenderCtx` genuinely sufficient?** Phase 8's components will want scene-local state (emitter positions, animation phase). If it is not there, adding it in Phase 8 means touching all 37 again — extend `RenderCtx` **now**.
3. **Cache cap value** — confirm the chosen bound against the curated asset set, and that Phase 15's target-size threading will not blow the key space past it.
4. Phase 7 replaces 376 inline `hex()` calls in `scenes_modular.go`. Confirm the sweep did not add more.
