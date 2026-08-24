# Phase 11 — Audio Engine on `oto`

> **Session size:** Medium · **Depends on:** Phase 3 · **Touches:** `pkg/audio/ambient.go`, `cliamp.go`, thin `pkg/ui` seam

---

## Goal

Make the ambient soundscape actually produce sound — in-process, cross-platform, race-free — and stop CLIAMP from freezing the UI.

## Why this phase is here

`pkg/audio` is a leaf package; nothing imports it but `pkg/ui`. It could run almost anywhere in the sequence. It sits here because **Phase 12 restructures `Model`**, and settling the audio API first means the decomposition sees its final shape instead of restructuring around a moving target.

---

## The headline finding

**The entire ambient sound engine is dead code. It has never produced sound.**

```bash
grep -rn "StartAmbientLoop" --include="*.go" .
# pkg/audio/ambient.go:187:  // StartAmbientLoop launches background looped audio playback.
# pkg/audio/ambient.go:188:  func (am *AmbientManager) StartAmbientLoop() {
```

Two hits: the doc comment and the declaration. **Zero call sites.** Not in `main.go`, not in `Model.Init()` (`pkg/ui/model.go:199-207`, which returns `tea.Batch(m.animateTick(), m.pollCLIAMP(), m.timerTick())`), nowhere.

What actually happens: `NewAmbientManager()` (called from `NewModel()`) creates a temp dir and **synthesizes five 4-second 22050Hz 16-bit mono PCM WAVs from scratch** — genuine DSP: low-pass-filtered noise for rain, sine rumble for fire, AM-modulated sine for wind, hum plus 3.2kHz clinks for cafe, 6Hz-gated 4.5kHz chirps for crickets (`ambient.go:102-145`). It writes them to `$TMPDIR/gottage_audio*/`. **And nobody ever plays them.**

Every `[r]`/`[m]`/`[W]`/`[e]` keypress flips a bool that only changes what the sidebar renders.

**Keep the synthesis.** It is real, working signal processing and the most valuable thing in the package. Replace only the playback layer.

### What wiring it up naively would turn on

`ambient.go:200-226`:

```go
func (am *AmbientManager) runLoop(filename string, activeFlag *bool, volVal *float64, cmdVar **exec.Cmd) {
	for {                                    // no exit condition, ever
		if *activeFlag && am.playerBinary != "" {
			// ... build afplay/mpv/ffplay command ...
			if cmd != nil { *cmdVar = cmd; _ = cmd.Run() }   // blocking, no ctx, no timeout
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

- **5 goroutines that never stop.** `Cleanup()` (`:347`) kills *processes* and leaves every goroutine running until process exit.
- **Data races.** `runLoop` reads `*activeFlag`/`*volVal` **unlocked**; `ToggleRain/Fire/Wind/Cafe/Crickets` (`:292-344`) write them **under `am.mu`**; and `pkg/ui` writes them **directly, bypassing the mutex entirely** — `model.go:285-286` (`m.Ambient.RainActive = true; m.Ambient.FireActive = true`), plus `:889`, `:902`, `:904`. `View()` reads them unlocked at `view.go:79-91` and elsewhere. `Cleanup()` itself takes **no lock at all** while every `Toggle*` does.
- **Linux busy-spin.** `LookPath("aplay")` sets `playerBinary` (`:74-82`) but `runLoop` has **no `aplay` case**, so `cmd` stays nil and the `time.Sleep` is in the *other* branch. On an ALSA-only box, an active channel is a **100% CPU goroutine playing nothing**. Five of them.
- **Volume by process-kill.** `AdjustChannelVolume` (`:263-265`) kills and respawns, so every `+`/`-` restarts the 4-second clip from zero.
- **afplay churn.** afplay exits after 4s and the loop respawns immediately — ~15 process spawns/minute/channel with an audible gap each cycle.
- **19 ignored errors in this one file**, including 10 unchecked `binary.Write` calls in the WAV header writer (`:158-181`).

---

## Workstreams

### 11.1 — Replace playback with `oto`

Confirmed decision: rewrite on a real Go audio library (`github.com/ebitengine/oto/v3`, or `gopxl/beep` if a higher-level mixer is wanted).

- Keep `createProceduralWAV`'s DSP; retarget it to produce **in-memory PCM**, not temp files. The temp dir disappears entirely.
- One `oto.Context`; one player per channel; a mixer applying per-channel gain **and** the master `am.Volume` (which today is set to 0.5 and applied to nothing).
- **Real gapless looping** — loop the PCM buffer, not the process.
- **Real volume** — scale samples; never restart playback.
- `context.Context` cancellation. Every goroutine exits on `Cleanup()`.
- Guard construction: if no audio device is available (headless, CI, no ALSA), degrade to silent-but-honest. **Never claim to be playing when it is not.**

### 11.2 — Make the state race-free by construction

Do not just add locks — **remove the shared mutable fields**. Channel state moves behind methods (`SetRain(bool)`, `SetVolume(SoundType, float64)`); the exported bools go away.

This also kills the awkward workaround in `ApplyPreset` (`model.go:178-194`), which does *"if desired && !current { Toggle() } else if !desired && current { Toggle() }"* three times over because there is no setter.

Note `SoundClock`/`ClockActive` are declared, initialized, and never read as state (`SoundClock` appears only as a defensive guard at `model.go:435`). Remove or implement.

### 11.3 — Unblock the UI from CLIAMP

`pkg/audio/cliamp.go` is subprocess-only — every operation is `exec.Command(cliampPath, verb)`. Problems:

- **Blocking calls run inline in `Update()`.** `model.go:915-916` (space), `:927` (n), `:937` (p), `:872/:878` (+/-) call `TogglePlayPause`/`NextTrack`/`PrevTrack` synchronously. With cliamp installed, one spacebar = `EnsureDaemon()` + **`time.Sleep(300 * time.Millisecond)` while holding `c.mu`** (`cliamp.go:183`) + a blocking `cliamp toggle` + a blocking `cliamp status`. **A guaranteed ≥300ms UI freeze per keypress**, unbounded if cliamp hangs.
- **Zero `context` usage repo-wide.** No `exec.CommandContext`, no timeouts. A wedged `cliamp status` blocks forever — and `pollCLIAMP` re-arms itself from its own result, so if it ever blocks, status polling dies silently with no recovery.
- **Zombie processes.** `EnsureDaemon` (`:179`) and `startNativeStream` (`:268`) call `Start()` and never `Wait()`. `Process.Kill()` without `Wait()` (`:205`, `:252`) leaves them unreaped.
- **Unlocked mutation.** `c.ActivePreset` is written at `:217`, `:223`, `:232` without the lock, while `PollStatus` reads it **under** the lock from a different goroutine. A genuine cross-goroutine race.
- **Hardcoded Apple-Silicon path** at `:65`: `/opt/homebrew/bin/cliamp`.
- **Fragile parsing.** `parseStatusOutput` (`:137`) scrapes stdout with `strings.Contains(raw, "playing")` and `strings.Contains(raw, "▶")`.
- **Silent false-positive.** With cliamp absent, `startNativeStream` falls back to `exec.Command("sh", "-c", fmt.Sprintf("curl -sL %s | afplay -", preset.URL))` (`:265`) — which `Start()`s successfully even if neither `curl` nor `afplay` exists, so the UI reports "Playing Lofi: …" while nothing plays. Also a command-injection shape that is safe only because URLs are currently compile-time constants: one "custom stream URL" feature from being RCE. **Build the URL as argv, not a shell string.**

Fixes: `exec.CommandContext` with timeouts everywhere; every call wrapped in a `tea.Cmd` so `Update` never blocks; `Wait()` or `Release()` every child; all state under the lock or behind methods; probe binaries before claiming success; drop the `sh -c` pipeline.

> **Boundary with Phase 15:** this phase makes the calls *non-blocking and cancellable*. **Phase 15** owns tick cadence and idle gating — including the fact that `pollCLIAMP` forks a subprocess every second forever regardless of whether CLIAMP is even installed.

### 11.4 — Test the package that has 0% coverage

`pkg/audio` has **zero test files**. Introduce a `Player` interface with a fake so behavior is testable without a sound device.

---

## Do NOT touch in this phase

- `Model`'s shape — **Phase 12**. Keep the UI seam to swapping direct field writes for method calls.
- Tick cadence, idle gating, the 1s poll interval — **Phase 15**.
- The sound mixer modal's layout — **Phase 14**. Make its controls *work*; do not redesign it.
- Rendering.

---

## Verification

```bash
go test -race ./pkg/audio -count=1
go test ./pkg/audio -cover                      # > 0%, target > 60%
grep -rn 'for {' pkg/audio/                     # no unbounded loops
grep -rn 'exec.Command(' pkg/audio/ | grep -v CommandContext   # 0
grep -rn 'time.Sleep' pkg/audio/                # 0
grep -c 'RainActive\|FireActive' pkg/ui/        # 0 direct field access
```

New tests: `TestAmbientProducesPCM` (rain PCM is non-silent and correctly shaped), `TestChannelsStopOnCleanup` (goroutine count returns to baseline), `TestVolumeDoesNotRestartPlayback`, `TestCLIAMPAbsentDegradesHonestly` (`Playing` stays false when no player exists), `TestUpdateNeverBlocks` (a CLIAMP-triggering `KeyMsg` returns in < 5ms).

**Manual:** press `[r]`. **You should hear rain.** That has never once been true in this repo.

## Exit criteria

- [ ] Ambient audio audibly plays; gapless; volume changes without restarting
- [ ] Zero leaked goroutines; `-race` clean; no exported mutable channel state
- [ ] All subprocess calls context-bounded and off the `Update` path
- [ ] No zombies; no `sh -c` string interpolation
- [ ] Honest degradation with no device and no cliamp
- [ ] `pkg/audio` coverage > 60%
- [ ] `make check` green

## Re-validation checklist for Phase 12

1. **Does the audio API shape match what Phase 3 persisted?** If channel identity or volume representation changed, fix the config schema now — Phase 3 exists so this is the *only* format change.
2. **How large is the remaining `pkg/ui` audio seam?** Phase 12 decomposes `Model`; a small, method-based audio surface makes that clean.
3. **Is `Player` injectable?** Phase 12's `TestNewModelIsHermetic` requires `NewModel()` to construct **zero** real audio devices and write **zero** files. If the fake cannot be injected, add that seam now.
4. Confirm `SoundClock`/`ClockActive` were resolved, so Phase 12 does not carry a dead enum into the new model.
