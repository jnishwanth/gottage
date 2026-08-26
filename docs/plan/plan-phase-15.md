# Phase 15 — Runtime: Tick Decoupling, Idle Gating, GIF Timing

> **Session size:** Medium · **Depends on:** Phase 13 · **Touches:** `pkg/ui/model.go`, `pkg/art/gif_loader.go`

---

## Goal

Stop burning CPU when nothing is happening, separate simulation rate from render rate, and make GIFs play at their authored speed.

## Why this phase is here

It needs the settled `Model` (12) and command table (13) — tick handling lives in `Update()`, and Phase 13 rewrote its structure. It needs Phase 14's layout so a resize can be treated as a wake event. And `Delays` was deliberately preserved back in Phase 1 for exactly this moment.

---

## The problems

### Three unsynchronized tickers, none gated

`pkg/ui/model.go:200-227`:

```go
func (m Model) Init() tea.Cmd {
	return tea.Batch(m.animateTick(), m.pollCLIAMP(), m.timerTick())
}
func (m Model) animateTick() tea.Cmd {
	return tea.Tick(60*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
}
func (m Model) pollCLIAMP() tea.Cmd { return tea.Tick(1*time.Second, ...) }
func (m Model) timerTick() tea.Cmd  { return tea.Tick(1*time.Second, ...) }
```

- **~16.7 FPS fixed**, self-rescheduling at `:274`, `:277`, `:297`.
- **Zero idle detection.** The tick fires unconditionally forever. Nothing checks whether particles exist, whether audio is playing, whether the active tab is even the art tab, or whether the terminal is focused. On the Timer/Tasks/Ambiance tabs the scene is **still simulated and fully re-rendered 16.7×/second.**
- **Simulation and render are not decoupled.** `tickMsg` advances simulation (`m.Scene.Update()`, `:247`); Bubbletea then calls `View()`, which **rasterizes** (`m.Scene.RenderFrame()` at `view.go:250`, and again at `:695` for Zen view). One rate for both.
- **`AnimSpeed` is ignored by the renderer.** `art/scene.go:615-621` returns early on `SpeedPaused` and halves on `SpeedSlow` — but `RenderFrame()` in `View()` ignores it entirely, so **the full raster cost is paid even when paused.**
- **Rendering is also driven by mouse motion.** `tea.WithMouseCellMotion()` (`main.go:121`) emits a message per cell traversed, each triggering a full re-raster.
- **`pollCLIAMP` forks a subprocess every second, forever** — `exec.Command(cliampPath, "status").Run()` (`audio/cliamp.go:116`) — regardless of whether CLIAMP is installed or playing.

### GIFs ignore their authored timing

`AnimatedGif.Delays` (`pkg/art/gif_loader.go:20`) is populated at `:95` and **never read**. Playback uses `tick % FrameCount` at `:138`, so **every GIF plays at the 60ms UI tick.** A 10fps GIF renders at 16.7fps; a 4fps one runs 4× too fast.

Preserved deliberately in Phase 1 as an unimplemented feature. Given the user's complaint is about the art, this is plausibly the **highest art-quality-per-line change in the whole plan.**

### Double resampling

GIFs are pre-scaled to `(c.Width, c.Height)` at load time. With Phase 14 wiring `RenderTo`, GIF frames now resample **twice** (source → 74×52 → target), which is visibly mushier than resampling once. Thread the *final* target dims to the loader.

This widens the cache key space — which is exactly why Phase 6's LRU bound had to land first.

---

## Workstreams

### 15.1 — Separate simulation from render

- `tickMsg` advances simulation only.
- `View()` renders from the last simulated state and **must not mutate**. Phase 5 already made `RenderFrame` non-mutating; enforce it here by caching the rendered frame string and invalidating on simulation step, resize, or theme change.
- A keypress or mouse move then costs a cache read, not a full 3,848-pixel raster.

### 15.2 — Idle gating

Do not schedule the next animation tick when nothing needs it:

- Active tab is not the art tab **and** no visible animated element.
- `AnimSpeed == SpeedPaused`.
- No active particles, no transition in flight, no active shader.

Wake on: keypress, mouse, resize, tab change, scene change, timer transition.

Keep a slow heartbeat (1-2s) for the clock badge — `View()` reads `time.Now()` at `view.go:204`, `:708`, so output legitimately changes at least once a second. Phase 5 made the clock injectable; use it rather than re-adding a raw `time.Now()`.

**Respect `AnimSpeed` in the render path too**, closing the "paused still rasters" bug.

### 15.3 — Gate `pollCLIAMP`

Do not poll when CLIAMP is not installed. Back off when idle or not playing. Phase 11 made the calls context-bounded and non-blocking; this phase makes them *infrequent*.

Note the existing failure mode: `pollCLIAMP` re-arms from its own result, so if the call ever blocks, **polling dies silently with no recovery.** Re-arm unconditionally.

### 15.4 — Honor GIF `Delays`

Accumulate elapsed simulated time and advance the GIF frame index by its authored delay, not by `tick % FrameCount`. Handle zero/missing delays with the conventional 100ms fallback.

Add `TestGifHonorsDelays` — a synthetic 2-frame GIF at 100cs advances exactly once per simulated second.

### 15.5 — Thread target size to the loader

Pass final render dims through `Activate(w, h)` (Phase 6) so each GIF is decoded once at the size it will actually be shown. Re-`Activate` on resize. Confirm the LRU cap still holds with the wider key space.

### 15.6 — Mouse motion

`tea.WithMouseCellMotion()` emits per-cell messages. With 15.1's frame cache this is cheap, but if cell-level motion is not actually used, `tea.WithMouseAllMotion()` or plain click tracking is cheaper still. Check what the hit-box code (Phase 14) needs.

---

## Do NOT touch in this phase

- The command table, `Model` shape, layout — Phases 12-14, settled.
- Scene art, components, tokens.
- Audio internals — Phase 11, settled. This phase changes *when* audio calls happen, not what they do.
- The canvas serializer — Phase 4, settled.

---

## Verification

```bash
go test ./pkg/ui ./pkg/art -count=1
go test ./pkg/art -bench BenchmarkRenderFrame -benchmem     # still within budget
```

Measure idle CPU directly — the headline number for this phase:

```bash
go run . &            # sit on the Tasks tab, do nothing
top -pid $!           # expect near-zero, not a steady 16.7fps burn
```

New tests:

- `TestNoTickWhenIdle` — paused, non-art tab, no particles ⇒ `Update` returns no animation `Cmd`. **The regression test for the phase.**
- `TestTickResumesOnWake` — keypress/resize/tab-change each re-arm the tick.
- `TestViewDoesNotMutateScene` — `FrameCount` and `TransitionStep` unchanged across two `View()` calls.
- `TestGifHonorsDelays`
- `TestPollCLIAMPSkippedWhenAbsent`
- `TestFrameCacheInvalidatedOnResize`

**Manual, and it matters:** watch a GIF scene that was obviously too fast. It should now play at a natural pace.

## Exit criteria

- [ ] Idle CPU near zero on a non-animating tab
- [ ] `View()` never mutates; frame cache with correct invalidation
- [ ] Paused means paused — no raster cost
- [ ] `pollCLIAMP` gated and unconditionally re-armed
- [ ] GIFs play at authored speed; decoded once at final size
- [ ] LRU cap still holds with the wider key space
- [ ] `make check` green

## Re-validation checklist for Phase 16

1. **What bugs from the original audit are still open?** Most should be fixed by construction. Phase 16's scope is exactly the remainder — enumerate it explicitly rather than assuming.
2. **Did the frame cache introduce staleness?** Any "screen didn't update" report is an invalidation gap; hunt it before Phase 16 adds lint pressure.
3. **Is `-race` still clean** with the new wake/sleep paths?
4. **Re-measure against Phase 4's budget** (p99 < 8ms, < 10 allocs/frame) — Phase 14's scaling and this phase's caching both moved it.
