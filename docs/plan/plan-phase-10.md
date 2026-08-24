# Phase 10 — Scene Migration B + Split `scenes_modular.go`

> **Session size:** Large · **Depends on:** Phase 9 · **Touches:** `pkg/art/scenes_modular.go` → several files

---

## Goal

Finish migrating the remaining handcrafted scenes, then break up the file that has been 45% of this codebase since the beginning.

## Why the split comes last

`scenes_modular.go` was **7,909 lines** — 45% of all Go in the repo, with a **2,117-line `init()`**. Splitting it earlier would have meant splitting it twice: Phase 1 deleted ~half, and Phases 9/10 rewrite the rest. Splitting a file you are about to rewrite is wasted work and a guaranteed merge headache.

By now the file is roughly 3,000 lines and the registration table has fallen from 2,064 lines to ~250. The split is now small and obvious.

---

## Workstream 10.1 — Finish migration

Same per-scene procedure as Phase 9 (see `plan-phase-9.md`). Carry forward:

- The list of remaining scenes from Phase 9's re-validation.
- Any components Phase 9 added.
- The "how to migrate a scene" note.

**Deleting a weak scene is a legitimate outcome.** The confirmed direction is *procedural first, curate* — a scene that is still not worth looking at after migration should go, not be shipped to raise a count. This is the same judgment Phase 1 applied to GIFs.

## Workstream 10.2 — Split the file

Suggested layout:

```
pkg/art/
  scene_registry.go      # RegisterScene, GetScene, SceneOrder (moved from scene_interface.go if cleaner)
  scenes_gif.go          # AnimatedGifScene type + its methods
  scenes_manifest.go     # generated or loaded manifest (see 10.3)
  scenes/                # one file per handcrafted scene, or grouped by category
    interior_window.go
    exterior_landscape.go
    night_artificial.go
  components.go          # Phase 8 (already exists)
```

Grouping by **archetype** (matching Phase 8's reference frames) beats grouping alphabetically — scenes sharing an archetype share components, so related code sits together.

## Workstream 10.3 — Registration table → manifest

The remaining ~40 GIF registrations are still structurally identical literals. Move to a data file (JSON/TOML next to the assets) walked at init, or `embed.FS` + a small manifest. Registration becomes a loop.

Keep `RegisterScene`'s **panic on duplicate** (Phase 1) — it now guards the manifest too, which is exactly where a copy-paste ID mistake would otherwise reappear.

Add `TestManifestMatchesAssets` — every manifest entry resolves through `pkg/assets`, and every embedded GIF appears in the manifest. Closes the orphan-asset loop permanently (19 GIFs were orphaned before Phase 1).

## Workstream 10.4 — Retire the dead flags

By now these should be resolved; if any survive, close them out:

- `SmoothDayNightBlend` (`scene.go:298`) — has a toggle at `:499` that does nothing, and is rendered in the sidebar (`pkg/ui/view.go:1025`). Implement or delete both sides.
- `BeatReactiveParticles` (`scene.go:297`) — same shape, rendered at `view.go:1026`.
- `SpeedFast` (`scene.go:130`) — displayed as "Fast Focus (2.0x)" and has **no branch** in `Update()` (`scene.go:615-621` handles only `SpeedPaused` and `SpeedSlow`), so it behaves identically to Normal. Implement it or remove the option.

Each is a switch the UI advertises and the engine ignores — the same class of lie as the 20 dead scenes and the silent ambient audio.

---

## Do NOT touch in this phase

- The component library, except to add a genuinely missing component.
- `SceneInterface`, tokens, canvas.
- `pkg/ui` — except deleting sidebar rows for flags removed in 10.4.
- **Do not reorder the `SceneType` enum.** Phase 3 made persistence string-keyed, so it is survivable now, but there is no reason to spend the risk.

---

## Verification

```bash
wc -l pkg/art/scenes*.go pkg/art/scenes/*.go     # no file over ~800 lines
go test ./... -count=1
go test ./pkg/art -run TestSceneGoldens
COLUMNS=80 LINES=24 go run .
grep -c 'wx, wy, ww, wh := ' pkg/art/            # ~0  (was 55)
grep -rc 'BlendRect(0, 0, c.Width, c.Height, PicoWhite' pkg/art/   # 1  (was 71)
grep -c 'hex("#' pkg/art/                        # 0
```

New tests: `TestManifestMatchesAssets`, `TestAllScenesAreResolutionIndependent` (extends Phase 9's to every scene), `TestNoDeadSceneFlags` (each `CottageScene` bool flag is read by at least one render path).

**Manual, and it is the real test:** run the app, walk every scene, at 80×24 and at full size. Every scene should be one you would leave on screen.

## Exit criteria

- [ ] All handcrafted scenes migrated (or deliberately deleted)
- [ ] `scenes_modular.go` gone; no file over ~800 lines
- [ ] GIF registrations manifest-driven; manifest ↔ assets tested both ways
- [ ] Duplication counters at ~0
- [ ] No flag advertised in the UI that the engine ignores
- [ ] Every scene resolution-independent and reviewed by eye
- [ ] `make check` green

## Re-validation checklist for Phase 11

1. **Final scene count and list** — Phase 17's README rewrite needs it, and `TestSceneCountMatchesReadme` will assert it.
2. **Is the art actually good now?** Compare against Phase 8's references. If not, that is a signal to iterate here rather than proceeding — the remaining phases are UI and runtime, and none of them improve the picture.
3. Phase 11 is `pkg/audio`-only and independent of everything above. The only cross-check: **Phase 3 persisted volume and active flags** — confirm the schema matches what the new audio engine will expose, or Phase 11 forces a second config change.
