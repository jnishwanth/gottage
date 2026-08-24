# Phase 7 — Palette → Design Tokens

> **Session size:** Medium · **Depends on:** Phase 6 · **Touches:** `pkg/art/palette.go`, `scenes_modular.go`

---

## Goal

Replace 267 flat object-named color constants and 376 inline hex literals with a semantic token system that has time-of-day ramps and resolves at compile time.

## Why this phase is here — before components, before migration

Phase 8 builds shared scene components. Phase 9/10 rewrite 36 scenes onto them. If components are built on ad-hoc colors, both get re-touched when tokens arrive.

Colors are the vocabulary components are written in. The vocabulary comes first.

---

## The problems

> All counts below are **pre-Phase-1** baselines. Phase 1 deleted roughly half of `scenes_modular.go` — **re-derive before starting.**

**1. It is a name dump, not a system.** `pkg/art/palette.go` is one 342-line `var (...)` block: 16 PICO-8 colors, then ~36 comment-delimited blocks each named after a single scene (`// Shire Hill Hobbit Burrow & Round Hearth`, `// Rainy Midnight Cyber Lofi Diner & Neon Window Booth`). It grows by one block per scene shipped. Names are object names — `ColorHobbitBrassKnob`, `ColorCherryPie`, `ColorSeedCake` — so **nothing is reusable across scenes**, which is a direct cause of the duplication Phase 8 exists to remove.

There are no semantic roles (surface, accent, warm-light, shadow), no light/dark or time-of-day ramps.

**2. Dead and duplicated entries.** ~267 named colors: ~70 used exactly once, **9 never referenced at all** (`ColorHobbitBrassKnob`, `ColorHobbitDoorGreen`, `ColorLanternGlow`, `ColorMachiyaLattice`, `ColorSkyDawnMid`, `ColorSkyDayMid`, `ColorSkySunsetMid`, `ColorSnowFlake`, `ColorWallPaperDay`). Duplicate hex values appear under different names — `#d4a373` is defined 5 times, `#f59e0b` 3 times.

The three unused `…Mid` sky colors are the tell: the palette was *designed* for 3-stop sky ramps, and every scene uses the 2-stop `DrawVerticalGradient` instead. **The ramp system was designed and never built.** This phase builds it.

**3. 376 inline `hex("#...")` calls in scene code**, across ~170 distinct values, **54 of which exactly duplicate an already-named constant**. Top offenders: `#1e293b` ×29, `#0f172a` ×14, `#334155` ×10, `#10141f` ×10.

**4. They re-parse every frame.** `palette.go:352-355`:

```go
func hex(s string) colorful.Color { c, _ := colorful.Hex(s); return c }
```

`colorful.Hex` runs `fmt.Sscanf`. Called from scene render paths, that is ~376 `Sscanf` invocations **per frame** for a scene using them all — and the error is silently discarded, so a typo yields black.

---

## Workstreams

### 7.1 — Define the token layer

Semantic roles, not object names. Sketch:

- **Surfaces:** `SurfaceWall`, `SurfaceFloor`, `SurfaceCeiling`, `SurfaceWood`, `SurfaceStone`, `SurfaceFabric`
- **Light:** `LightWarm`, `LightCool`, `LightHearth`, `LightLamp`, `LightMoon`, `LightNeon`
- **Sky ramps:** `SkyRamp(tod)` returning a **3-stop** ramp (top/mid/bottom) — building what the dead `…Mid` constants were for
- **Foliage / water / accent / shadow** families
- **Materials:** small palettes for glass, brass, ceramic, paper

Ramps are the point. `Warm(base, t)` / `Shade(base, n)` derived programmatically beats 267 hand-picked constants and makes lighting coherent across scenes.

Keep the 16 PICO-8 constants — they are used broadly (`PicoWhite` ~145 references) and are a legitimate base layer.

### 7.2 — Resolve colors at compile time

Every color must be a package-level `var` initialized once, never parsed in a render path.

- `hex()` becomes **init-time only**, and must **panic** on a bad literal instead of returning black. A typo should fail the build-and-test, not silently paint a black pixel. (Same principle as Phase 1's `RegisterScene`.)
- Replace all 376 inline calls with tokens.
- Add `TestNoInlineHexInSceneCode` as a source-level guard: `grep -c 'hex("#' pkg/art/scenes_modular.go` must be 0. Cheap, and it holds the line permanently.

### 7.3 — Map, don't just delete

For each of the ~170 distinct inline values and ~267 constants, map to a token. Expect real consolidation — `#1e293b` appearing 29 times across unrelated scenes is one token used 29 times.

**Where a color is genuinely scene-specific, keep it** — as a scene-local constant near the scene, not in the global palette. The goal is a shared vocabulary plus small local dialects, not one flat namespace with a different shape.

### 7.4 — Guard against regrowth

- `TestNoUnusedPaletteColors` — every exported token referenced at least once.
- `TestNoDuplicatePaletteHex` — no two tokens with the same value (they should alias).
- Note both in `AGENTS.md`: new colors go through tokens.

---

## Do NOT touch in this phase

- **Scene composition and layout.** Swap colors only. Goldens will change (colors consolidate), but *structure* must not — a diff showing moved geometry means the phase leaked into Phase 9/10.
- Shared components (Phase 8). It is tempting to extract the sky gradient while touching all 34 ToD switches — **don't.** That is Phase 8, and mixing them makes both unreviewable.
- `pkg/ui/styles.go`. Distinct system (6-slot lipgloss `Theme`). It has its own problems — `Background` declared on every theme and never used, hardcoded off-theme hex at `styles.go:188,197,208,223,229` — but it is **Phase 14's** territory.

---

## Verification

```bash
grep -c 'hex("#' pkg/art/scenes_modular.go     # 0
grep -rn 'colorful.Hex' pkg/art/               # only inside init-time hex()
go test ./pkg/art -run TestSceneGoldens        # regenerate deliberately with -update
go test ./pkg/art -bench BenchmarkRenderFrame -benchmem   # no per-frame Sscanf
go test ./... -count=1
```

New tests: `TestNoInlineHexInSceneCode`, `TestNoUnusedPaletteColors`, `TestNoDuplicatePaletteHex`, `TestHexPanicsOnBadLiteral`, `TestSkyRampHasThreeStops`.

**Goldens change in this phase.** Regenerate with `-update` in a dedicated commit and **review the diff visually** — consolidating 170 values into a token set is exactly where an unintended color shift hides.

## Exit criteria

- [ ] Token layer with semantic roles and 3-stop ToD sky ramps
- [ ] Zero inline `hex()` in scene code; zero color parsing in render paths
- [ ] `hex()` panics on bad input, init-time only
- [ ] No unused, no duplicate-valued tokens
- [ ] Goldens regenerated in a labelled commit and visually reviewed
- [ ] `make check` green

## Re-validation checklist for Phase 8

1. **Is the token set actually expressive enough to build components from?** Sketch one component (a window frame) against it before declaring the phase done. If you reach for a raw hex, the vocabulary is incomplete — fix it here, not in Phase 8.
2. **Did the golden diff contain only color changes?** Any structural movement means scene composition was touched, and Phase 9/10's baseline is compromised.
3. **Do 3-stop sky ramps look right at all four times of day?** Phase 8's sky component depends on it directly.
