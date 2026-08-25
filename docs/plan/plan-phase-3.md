# Phase 3 — Config & Persistence Format

> **Session size:** Medium · **Depends on:** Phase 2b · **Touches:** `pkg/features/presets.go`, a new `pkg/config/`, `pkg/ui/model.go` (call sites only)

---

## Prerequisites

Phase 2b sits between Phase 2 and this phase, so **`plan-phase-2b.md`'s
"Re-validation checklist for Phase 3" is the one to run** — it carries Phase 2's
three items forward alongside its own. `plan-phase-2.md`'s checklist is the
earlier link in that chain and is preserved there; 2b's supersedes it as the
entry point.

**Persist conservatively — this is a constraint on §3.2's schema.** The
*Virtual Cottage* simplification (`plan-phases.md`, Confirmed decisions) defers
**how far to cut** to Phase 8, which is gated on signed-off reference frames.
Phase 8 runs after this phase and cannot be waited for, so this phase does not
guess. **These two lists are exhaustive over §3.2's schema — every one of its 17
entries appears in exactly one of them.** Anything added to that schema later
must be classified into one of them at the time it is added; an unlisted entry is
an unanswered question, not a default.

**Persist** — state whose survival is not in question:

> active scene, theme, favorites, weather / time-of-day and lamp/fire toggles,
> **anim speed**, per-channel audio volumes and active flags, **CLIAMP preset
> index**, timer session count and phase, **zen mode**.

The three in bold are the ones the first draft of this list left unstated:

- **anim speed** — a property of how the scene animates. Scenes survive any
  Phase 8 outcome, and §3.1 already gives `AnimSpeed` a stable string identity.
- **CLIAMP preset index** — which curated lofi track is selected
  (`pkg/audio/cliamp.go:49`, indexing `CuratedLofiPresets`). Music is half the
  stated direction, and Phase 11 rewrites the audio engine without removing the
  feature.
- **zen mode** — a view-level toggle over the scene, not a separate feature
  whose survival is in question.

**Defer** — state whose survival Phase 8 decides:

> tasks, beverage state, color filter, particle overlay, shader effect,
> vignette/CRT, render mode, GPU render flag.

Also deferred, though §3.2 does not list them: **inspiration quotes** and the
**34 built-in atmosphere presets** (`loadBuiltinPresets`, `pkg/features/presets.go`).

**Presets are deferred, but not because built-ins and user data are separable
today — they are not.** `loadBuiltinPresets` (`presets.go:41`) fills `pm.Presets`
with the 34 built-ins; `SaveToFile` (`presets.go:451`) marshals **all** of
`pm.Presets`; and both `SavePreset` (`presets.go:438`) and `DeletePreset`
(`presets.go:447`) call it, reachable from the UI at `model.go:513` and
`model.go:551`. So saving a single custom preset writes all 34 built-ins into
`user_presets.json`, and `LoadFromFile` (`presets.go:475`) then does
`pm.Presets = loaded` wholesale, discarding the in-code copies.

That is the same defect as **bug 4 below** ("Load clobbers built-ins",
`presets.go:474-476`) seen from the other side, and Phase 3 owns both. It also
means §3.0's framing understates the exposure: a `user_presets.json` written
before Phase 1 pins the *old* built-ins too, so "⚡ Cyberpunk Hackathon Deck" and
"🌌 Stargazing Observatory" come back at their pre-curation `SceneType`s and hit
the same silent fallback — and `TestBuiltinPresetScenesAreRegistered` cannot
catch it, because it builds `NewPresetManager("")` and `LoadFromFile` returns
early on an empty path, so the test only ever sees the in-code copies.

**Consequence for this phase:** presets cannot be deferred by simply leaving them
out of the schema, because the file already contains state on both sides of the
line. Phase 3 has to decide how built-in and user presets are separated — on
migration, are persisted built-ins re-seeded from code by name, or carried
forward as opaque user data? Answer that here; the *contents* of the built-in
library remain Phase 8's call. `user_presets.json` (`pkg/ui/model.go:147`) is
user data and **is in scope**: §3.2's "keep user presets here or in a sibling
`presets.json`" decision still stands, and the int→string migration applies to it.

Deferring them costs nothing, and the asymmetry is the whole reason: *adding* a
field to a versioned config later is backward-compatible and **is not a format
break**, whereas *removing* one is. Persisting an undecided feature now is the
only option that buys a second break. This does not weaken "one on-disk format
break" — the break is §3.1's enum→string identity change, not the field list.

§3.2's schema list below predates this constraint and names several deferred
entries; the constraint wins.

---

## Goal

One on-disk format break, done once, covering everything the app should remember. After this phase, no later phase changes the persistence format.

## Why this phase is here — and why it must not be later

This is the **last free format change**. Right now there are effectively no real user files, so the enum→string migration is a rewrite, not a migration. Every phase after this that adds persisted state (Phase 12 persists UI state, Phase 11 persists volumes) writes into a schema that is already correct.

Deferring it means either shipping a migration path forever, or breaking user data twice.

## The bugs

**1. Raw `iota` serialization.** `AtmospherePreset` (`pkg/features/presets.go:10-23`) has JSON tags and serializes six `art` enums as bare ints:

```go
type AtmospherePreset struct {
	Scene          art.SceneType       `json:"scene"`           // -> 0, 1, 2...
	ColorFilter    art.ColorFilter     `json:"color_filter"`
	ParticleOverlay art.ParticleOverlay `json:"particle_overlay"`
	ShaderEffect   art.ShaderEffect    `json:"shader_effect"`
	AnimSpeed      art.AnimSpeed       `json:"anim_speed"`
	// ...
}
```

**Reordering any enum in `pkg/art` silently corrupts every saved preset file.** Phase 1 just pruned the `SceneType` enum; Phases 6-10 will touch these types further. This must be closed now.

**2. Writes to the current working directory.** `features.NewPresetManager("user_presets.json")` (`pkg/ui/model.go:145`) — presets appear and disappear depending on where you launch, and the app litters the user's cwd.

**3. Errors swallowed.** `_ = pm.SaveToFile()` at `presets.go:438` and `:447`; `LoadFromFile`'s error discarded at `:37`. Preset persistence fails silently.

**4. Load clobbers built-ins.** `presets.go:474-476` does `if len(loaded) > 0 { pm.Presets = loaded }` — *replaces* rather than merges, so any built-in shipped in a future version is invisible to existing users.

**5. All 34 built-in presets never apply their theme.** `pkg/ui/model.go:172`:

```go
for _, theme := range AllThemes {
    if strings.EqualFold(theme.Name, p.Theme) { ... }
}
```

Theme names carry emoji (`pkg/ui/styles.go:20` → `"Cozy Autumn 🍁"`); built-in presets store bare names (`presets.go:47` → `"Cozy Autumn"`). `EqualFold("Cozy Autumn 🍁", "Cozy Autumn")` is false, so **no built-in preset has ever changed the theme.** User-saved presets *do* work (`model.go:497` stores the emoji-bearing name), so the two behave differently. `presets.go:393` is wrong even ignoring emoji: `"Cyberpunk Neon"` vs the theme's `"Cyberpunk 8-Bit ⚡"`.

**6. Almost nothing persists.** Repo-wide, the only persistence is presets (plus temp WAVs). **Tasks, timer sessions, favorites, active theme, active scene, beverage state, volumes, and every toggle are lost on restart.** No config file of any kind, no `UserConfigDir`.

---

## Workstreams

### 3.0 — An unguarded load path Phase 1 left behind

> Recorded at the end of Phase 1 so it is not rediscovered as a bug report.

Phase 1 cut 314 GIF registrations without pruning the `SceneType` enum, so **314
constants are now valid enum values that nothing registers.**

> **Three GIF counts circulate; they measure different moments and all three are
> right.** **344** GIF registration calls existed when Phase 1 began
> (`plan-phase-1.md:22,84`) — with the 36 handcrafted calls that is the **380
> registrations** of `AGENTS.md` §4, of which 20 IDs were claimed twice, so 360
> distinct entries reached the registry. **324** is the GIF count once those 20
> shadowing registrations were deleted (360 total). **314** is what the curation
> pass then cut, leaving 10. So `presets_test.go:52`'s "from 344 to 10" is
> correct as written — Phase 1 removed both the 20 and the 314. Check which
> moment a count refers to before treating it as a discrepancy.

A saved
`user_presets.json` written before that curation can name one, and nothing on the
load path notices:

- `LoadFromFile` unmarshals the raw int with no range or liveness check.
- `ApplyPreset` (`pkg/ui/model.go`) assigns it to `Scene.CurrentScene` unvalidated.
- `GetScene` (`pkg/art/scene_interface.go`) falls back to `SceneRainyCity` for an
  unregistered type — no error, no log.
- `SceneType.String()` misses the registry and returns the 4-case switch default,
  `"8-Bit Pixel Animation 🎨"`.

Net effect: the app silently shows the rainy-city scene under a header naming
neither that scene nor the one the preset asked for. Phase 1 added
`TestBuiltinPresetScenesAreRegistered` covering the **built-in** presets, which is
what caught two of them at the time; **user-saved presets have no equivalent
guard**, and a test cannot supply one because the input is a user's file.

That guard is also narrower than it looks. It builds `NewPresetManager("")`, and
`LoadFromFile` returns early on an empty path, so it only ever checks the in-code
built-ins. Once a user has saved any preset, `user_presets.json` holds a *copy* of
all 34 (see the Prerequisites section), and `LoadFromFile` prefers that copy — so
built-ins pinned at pre-curation `SceneType`s reach the same fallback while the
test stays green.

This is not a bug to patch here — §3.1's `UnmarshalJSON` over stable string keys
is the fix, provided it treats an unknown key as the documented "fall back to a
default **and record a warning**" case rather than silently zeroing it. The int
form must get the same treatment during migration: an int with no live
registration is exactly as unknown as an unrecognised string.

### 3.1 — Stable string identities for `art` enums

Give `SceneType`, `ColorFilter`, `ParticleOverlay`, `ShaderEffect`, `AnimSpeed`, `TimeOfDay`, `Weather`, `CatBreed` `MarshalJSON`/`UnmarshalJSON` over stable string keys (`"SceneRainyCity"`, not `0`).

Unknown key on load → fall back to a default **and record a warning**; do not fail the load and do not silently zero it.

Add `TestEnumStringIdentitiesAreStable` with a golden fixture. This is the test that makes future enum edits safe.

### 3.2 — `pkg/config` on `os.UserConfigDir()`

One file, e.g. `~/.config/gottage/config.json`, with a `"version"` field from day one. Atomic write (temp + rename). Errors returned, never `_ =`.

Schema covers: active scene, theme, favorites, weather/ToD/lamp/fire toggles, anim speed, render mode, color filter, overlay, shader, vignette/CRT, per-channel audio volumes and active flags, CLIAMP preset index, tasks, timer session count and phase, beverage state, zen mode, GPU render flag.

> **Narrow this list against the Prerequisites section above.** It was written before the *Virtual Cottage* simplification was recorded. The Prerequisites persist/defer lists are exhaustive over these 17 entries and are authoritative where the two disagree: eight of them — tasks, beverage state, color filter, overlay, shader, vignette/CRT, render mode, GPU render flag — are state whose survival Phase 8 decides and are deferred; the other nine persist. Adding a deferred field in a later phase is backward-compatible; removing one is not. Persist only what is settled.

Keep user presets either in the same file or a sibling `presets.json` — one decision, made here.

### 3.3 — Move `AtmospherePreset` out of `pkg/features`

`pkg/features/presets.go:7` imports `gottage/pkg/art`, and 6 of 11 preset fields are `art` types. This is an **inverted dependency**: `presets.go` is serialized renderer configuration misfiled into the domain package, and it drags the 7,900-line art package into anything wanting the task list or timer.

Move it to `pkg/config` (or `pkg/art`). `pkg/features` should end this phase importing nothing from `pkg/art`.

### 3.4 — Fix the theme match and merge semantics

- Match presets to themes by a **stable ID**, not display name. Add `Theme.ID` (`"cozy-autumn"`) and key on it. Emoji live only in `Theme.Name`, which is display-only. This kills the bug class, not just the instance.
- Merge built-ins with user presets on load instead of replacing (`presets.go:474-476`); dedupe by ID, user wins.
- Return and surface save errors.

### 3.5 — Wire persistence into the app

Load in `NewModel()`; save on quit and on meaningful state change (debounced). Keep the seam small — **Phase 12 restructures `Model`**, so do not restructure it here. A single `m.Config` field plus load/save calls is the right footprint.

---

## Do NOT touch in this phase

- **`Model`'s shape.** Phase 12 decomposes it. Add one field; change nothing else.
- Keybindings, help text, modal handling — Phases 12/13.
- Audio *behavior*. Persist volume/active fields; **Phase 11** rewrites the engine. Persisting a flag that does nothing yet is correct — Phase 11 makes it real.
- Rendering.

---

## Verification

```bash
go build ./... && go test ./... -count=1
grep -rn 'gottage/pkg/art' pkg/features/          # 0 hits
grep -rn '_ = pm.Save\|_ = .*SaveToFile' pkg/     # 0 hits
grep -rn '"user_presets.json"' pkg/               # 0 hits

# round-trip by hand
go run .            # change scene + theme, add a task, quit
cat ~/.config/gottage/config.json | head -30      # string enum values, not ints
go run .            # state restored
```

New tests:

- `TestPresetJSONUsesStrings` — marshal a preset, assert `"scene":"SceneRainyCity"` and **not** `"scene":0`.
- `TestPresetSurvivesEnumReorder` — unmarshal a checked-in golden JSON fixture; assert correct resolution. **The regression test for the phase.**
- `TestBuiltinPresetsResolveThemes` — every built-in's theme ID matches exactly one entry in `AllThemes`. Catches bug 5 permanently.
- `TestConfigRoundTrip`, `TestConfigPathIsXDG` (under a temp `HOME`), `TestConfigUnknownEnumFallsBack`, `TestSaveErrorsAreReturned`.

## Exit criteria

- [ ] No `art` import in `pkg/features`
- [ ] All enums serialize as stable strings; golden fixture committed
- [ ] Config lives under `os.UserConfigDir()`; atomic writes; errors returned
- [ ] Applying a built-in preset visibly changes the theme
- [ ] Tasks, timer, favorites, theme, scene, volumes survive restart
- [ ] `make check` green

## Re-validation checklist for Phase 4

1. Phase 4 is rendering-only and independent of this phase. The main cross-check: **did adding `pkg/config` change `NewModel()`'s side effects?** Phase 12's `TestNewModelIsHermetic` will require config writes to be redirectable via `HOME` — confirm that holds now rather than discovering it in Phase 12.
2. **Is the config schema genuinely complete?** Walk `Model`'s 40 fields and confirm every one that should survive restart is covered. Anything missed becomes a second format change later, which this phase exists to prevent.
