# Phase 8 — Art Direction Reference + Component Library

> **Session size:** Medium · **Depends on:** Phase 7 · **Touches:** new `pkg/art/components.go`, reference goldens

---

## Goal

Two deliverables, in order: **an agreed picture of what good looks like**, then the shared component vocabulary that Phases 9/10 rebuild all 36 scenes with.

## Why this phase is gated on reference frames

The prompt that drove the last 88 commits said:

> *"make the tui design better... do some research on what would be actually good art/animation to display here. **what we have is not quite what i want.** going for something that would be actually what humans would want... it needs to be cozy and **something people will actually look at**"*

There is still no artifact anywhere in this repo that says what "what i want" *is*. Phases 9/10 rebuild 36 scenes — doing that against an unstated target repeats the exact mistake that produced 88 commits of unsatisfying art.

**Do not start 8.2 until 8.1 is signed off.**

---

## Workstream 8.1 — Reference frames (gate)

Produce **2-3 fully-realized reference frames** at 74×52, rendered and viewed in a real terminal, and get explicit sign-off.

Suggested spread: one interior with a window (the dominant archetype), one exterior/landscape, one night scene with strong artificial light. These three exercise sky ramps, window framing, warm-light falloff, and silhouette work — the shared vocabulary's whole surface.

For each, write down **why it works**: light direction and falloff, contrast ratio between the window and the interior, how many distinct tones carry a surface, how animation reads at 16.7fps, how much detail survives half-block resolution.

Two known-real quality inputs to fold in:

- **GIF frame timing.** `AnimatedGif.Delays` is populated and never read, so every GIF plays at the 60ms UI tick, ignoring authored timing (a 4fps GIF runs 4× too fast). Phase 15 implements it — but judge reference GIF scenes at *correct* timing, or the reference is set against a bug.
- **The companion.** `showBuddy`/`buddyX`/`buddyY` are never set true on any GIF scene, so those scenes contain no character at all. A cozy scene with nobody in it is a large part of "not quite what i want." Decide here whether the buddy belongs in GIF scenes and where.

Record the sign-off in the phase's commit message and check the reference goldens in. Phases 9/10 are judged against them.

## Workstream 8.2 — The component library

The duplication these components replace (**pre-Phase-1 counts — re-derive**):

| Pattern | Occurrences | Component |
|---|---|---|
| `wx, wy, ww, wh := r.X, r.Y, r.W, r.H` | 55 | `Window(ctx, rect, style)` |
| `if lightning { c.BlendRect(0,0,c.Width,c.Height, PicoWhite, α) }` | **71** | `LightningFlash(ctx)` |
| `switch tod { case TimeDawn: skyTop, skyBot = ... }` | 34 | `Sky(ctx, rect)` — 3-stop ramp from Phase 7 |
| `c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ...)` | 17 | folded into `Window` |
| `c.FillRect(wx+ww/2-1, wy, 2, wh, ...)` (mullion) | 16 | `Window` option |
| `c.FillRect(0, 3N, c.Width, c.Height-3N, ...)` (floor) | 25 | `Floor(ctx, rect, material)` |
| `c.ApplyRadialLight(...)` in `ApplyLighting` | 79 | `WarmLight(ctx, pos, radius, token)` |
| inline steam animation | ~25 distinct | `Steam(ctx, emitter)` |
| twinkling star loops | ~13, 3 different moduli | `StarField(ctx, rect)` |

Only **6 shared helpers exist today** (`renderBookshelfHelper`, `renderPottedSucculent`, `renderPottedMonstera`, `renderDeskHelper`, `renderFireplaceHelper`, `renderWoodStoveHelper`, `renderCatHelper`), with thin adoption — `renderCatHelper` 38 uses, `renderFireplaceHelper` 6, several at 3. They are the right idea; extend the pattern rather than replacing it.

### Design rules

**1. Every component takes a `Rect`, never absolute coordinates.** This is what makes migrated scenes resolution-independent *for free* — and it is the entire reason genuine reflow does not need its own phase. Un-migrated scenes keep rendering at 74×52 and get scaled by Phase 4's `RenderTo`, so migration stays incremental and safe.

**2. Components take `*RenderCtx`** (Phase 6), so they see toggles and frame count without new parameters.

**3. Components are built on Phase 7 tokens.** If a component needs a raw hex, the token vocabulary is incomplete — go back and extend it.

**4. Options over variants.** `Window(ctx, rect, WithMullion(2), WithArch(), WithRainStreaks())` beats seventeen near-identical window functions.

**5. Animation is a function of `ctx.Frame`**, never of package state. Phase 5's determinism guarantee has to survive.

### Verify components against the references

Rebuild the 2-3 reference frames **using only components**. If they cannot reproduce the signed-off look, the library is wrong — fix it here, before 36 scenes depend on it.

---

## Do NOT touch in this phase

- **The 36 scenes.** Phases 9/10 own migration. Migrating "just a couple while I'm here" makes the library's design feedback loop invisible.
- `SceneInterface` (Phase 6, settled). If `RenderCtx` turns out insufficient, that is a red flag — extending it now means re-touching 37 implementations, so weigh it deliberately rather than drifting into it.
- Palette tokens (Phase 7, settled) — except to *extend* where a component proves a gap.
- `pkg/ui`.

---

## Verification

```bash
go test ./pkg/art -run 'TestComponent|TestReference' -count=1
go test ./pkg/art -run TestSceneGoldens -count=1     # UNCHANGED — no scene was touched
go test ./... -count=1
```

New tests:

- `TestReferenceFrameGoldens` — the signed-off frames, locked.
- `TestComponentsAreResolutionIndependent` — render each component into 74×52, 120×80, and 40×28 canvases; assert it stays inside its `Rect` and produces sane output at every size. **The regression test for the phase.**
- `TestComponentsAreDeterministic` — same `ctx`, same bytes.
- `TestComponentsUseOnlyTokens` — source-level guard, no `hex(` in `components.go`.

**Existing scene goldens must be untouched.** A diff means the phase leaked into Phase 9/10.

**Manual:** render the reference frames in a real terminal at several sizes and look at them. This phase's deliverable is aesthetic; a passing test proves nothing about whether it is cozy.

## Exit criteria

- [ ] 2-3 reference frames rendered, reviewed, **signed off**, and locked as goldens
- [ ] Written rationale for what makes them work
- [ ] Component library covering window / sky / floor / warm-light / steam / stars / lightning
- [ ] Every component `Rect`-parameterized and resolution-independent
- [ ] Reference frames reproducible **using only components**
- [ ] Scene goldens unchanged
- [ ] `make check` green

## Re-validation checklist for Phase 9

1. **Was sign-off actually obtained?** If 8.1 was skipped or self-approved, stop. Phase 9 rebuilds 18 scenes against it, and a wrong reference costs two phases.
2. **Did building the references reveal missing components?** Add them now — discovering a gap at scene 14 of 18 means going back and re-touching thirteen.
3. **Is `renderCatHelper` (38 call sites) folded into the library or left alone?** Decide before migration; the answer changes every scene's diff.
4. **Pick the Phase 9 scene batch** — start with 2-3 scenes closest to the reference archetypes so the migration pattern is established on easy cases.
