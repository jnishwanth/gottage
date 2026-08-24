# Phase 9 — Scene Migration A (first ~18 scenes)

> **Session size:** Large · **Depends on:** Phase 8 · **Touches:** `pkg/art/scenes_modular.go`

---

## Goal

Rebuild the first half of the handcrafted scenes onto the Phase 8 component library, judged against the Phase 8 reference frames. Establish the migration pattern so Phase 10 is mechanical.

## Why this is split from Phase 10

36 scenes, each needing individual aesthetic judgment, does not fit one session honestly. The work is uniform per scene, so it splits cleanly at any boundary — **and if this session runs long, split again.** "Scenes 1-9" then "10-18" is a legitimate cut for the executing agent to make on its own.

---

## Prerequisites — verify before starting

- Phase 8 reference frames exist and were **signed off** (not self-approved).
- Component library covers window / sky / floor / warm-light / steam / stars / lightning.
- Phase 5 goldens exist for every scene and are stable under `-count=5`.

If any is missing, stop and go back. This phase is where the payoff lands and where unverified foundations cost the most.

---

## Ordering within the phase

**Start with 2-3 scenes closest to a reference archetype.** Get the pattern right on easy cases, then work outward. Prioritize:

1. Scenes matching a reference archetype (interior-with-window first — it is the dominant form).
2. **The 20 scenes that never rendered before Phase 1.** They have never been seen by anyone, so they are the most likely to contain genuinely broken art. Their line ranges are listed in `plan-phase-1.md`. Expect surprises.
3. Everything else.

## Per-scene procedure

For each scene:

1. **Look at it first.** `go run .`, navigate to it, at more than one terminal size. Note what works and what does not, against the reference rationale.
2. Replace hand-rolled patterns with components: window destructure → `Window`; ToD sky switch → `Sky`; lightning blend → `LightningFlash`; radial light → `WarmLight`; inline steam → `Steam`; star loop → `StarField`; floor fill → `Floor`.
3. Replace literal coordinates with `Rect`-derived ones. **This is the resolution-independence payoff** — 463 of 1,159 draw calls pass a literal first coordinate today (`ZenTeahouseScene` hardcodes room height `34`; `CityCafeScene` lays out 8 buildings at `bx := wx + b*7 + 1` against `ww = 58`, which walk off the edge at `ww = 40`).
4. Improve the art deliberately where it falls short of the reference. **This phase is allowed and expected to change how scenes look** — unlike Phases 6 and 7.
5. Regenerate that scene's golden, **look at the before/after**, commit per scene or per small batch.

**Commit granularity matters.** One commit per scene (or per 2-3) keeps the golden diff reviewable. A single 18-scene commit is unreviewable, and unreviewed golden regeneration is how bad art ships — which is the failure mode this entire plan exists to correct.

## What "done" means for a scene

- Uses components rather than hand-rolled equivalents.
- No literal absolute coordinates; everything derives from `Rect`.
- Renders correctly at 74×52 **and** at a downscaled size via `RenderTo`.
- Has a regenerated, visually-reviewed golden.
- Meets the reference bar — the honest test is *would you leave this on screen while working?*

---

## Do NOT touch in this phase

- The component library, **except** to add a genuinely missing component. If you add one, note it — Phase 10 gets it for free, and it is a signal Phase 8 was under-scoped.
- `SceneInterface`, palette tokens, the canvas.
- `pkg/ui`.
- The GIF scenes. They are data, not code; Phase 15 fixes their timing.

---

## Verification

```bash
go test ./pkg/art -count=1
go test ./pkg/art -run TestComponentsAreResolutionIndependent
COLUMNS=80 LINES=24 go run .        # migrated scenes readable at small size
go run .                             # walk every migrated scene by eye
```

Per-scene: golden regenerated and reviewed. Repo-wide, track the duplication counters falling:

```bash
grep -c 'wx, wy, ww, wh := ' pkg/art/scenes_modular.go
grep -c 'BlendRect(0, 0, c.Width, c.Height, PicoWhite' pkg/art/scenes_modular.go
```

New test: `TestMigratedScenesAreResolutionIndependent` — for each migrated scene, render at 74×52 and at 60×40 natively (not via `RenderTo`); assert nothing draws outside bounds. This is the assertion that proves migration actually bought resolution independence rather than just moving code around.

## Exit criteria

- [ ] ~18 scenes migrated, each with a reviewed golden
- [ ] Zero literal absolute coordinates in migrated scenes
- [ ] Migrated scenes render correctly at 74×52 and downscaled
- [ ] Duplication counters materially down
- [ ] Every migrated scene passes the *would you leave it on screen* test
- [ ] `make check` green

## Re-validation checklist for Phase 10

1. **Which scenes remain?** List them explicitly — Phase 10's scope is exactly this remainder.
2. **Did migration reveal missing components?** If Phase 9 added several, Phase 10 should re-check the library before starting rather than adding piecemeal.
3. **Is the migration pattern written down?** A short "how to migrate a scene" note in `AGENTS.md` makes Phase 10 (and any future scene) mechanical.
4. **Are any remaining scenes not worth migrating?** If a scene is weak even after Phase 1 rescued it, **deleting it is a legitimate outcome.** The confirmed direction is a curated set people actually want to look at, not maximum scene count.
5. **Duplication counters** — record current values so Phase 10 can show the delta.
