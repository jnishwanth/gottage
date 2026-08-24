# Phase 2b — Simplify: Remove the Scene-Category System

> **Session size:** Small · **Depends on:** Phase 2 · **Touches:** `pkg/art/categories.go`, `pkg/ui/model.go`, `pkg/ui/view.go`, `pkg/ui/command_palette.go`

> **Numbered 2b, not 3.** Inserting a real Phase 3 would renumber ten downstream
> docs and every cross-reference in them, for no benefit. It runs after Phase 2
> and before Phase 3.

---

## Goal

Delete the scene-category system outright: the seven-entry `SceneCategories`
picker taxonomy, the substring matcher that populates it, and the per-category
`SceneMeta` blurbs. One fewer concept, roughly 90 lines gone, and a picker that
lists scenes instead of sorting them into buckets that no longer divide anything.

## Why this phase exists, and why here

Two reasons, one product and one mechanical.

**Product.** The app is meant to be a cozy TUI in the spirit of *Virtual Cottage*:
an animation and some music. Categories are shelving for a catalogue that no
longer exists — Phase 1 took the registry from 360 scenes to 46, and at that size
a flat list is not merely adequate, it is better. This is the first cut of a
larger simplification whose scope is deliberately **deferred to Phase 8**, where
the reference frames force an explicit statement of what the app is for.

**Mechanical.** Phase 1 emptied a category and exposed that the taxonomy fails
silently when it stops dividing the set. `CategorizeScene` sorted 0 of 46 scenes
into "⚙️ Steampunk & Fantasy" — all 57 members were among the 314 GIFs cut — and
`GetScenesByCategoryAndFavorites` answers an empty category by returning
`SceneOrder`, so the picker rendered *"Category: ⚙️ Steampunk & Fantasy
(46 scenes)"* over every scene in the app. No error, no log.

Phase 1 dropped that one category and added
`TestEveryCategoryMatchesARegisteredScene` so the next prune fails loudly. Both
are patches on a structure this phase deletes, which is exactly the situation
`AGENTS.md` §6's ordering rule anticipates: *a phase may replace a structure; it
may not fix bugs inside a structure a later phase replaces.* The patch bought
safety in the interim. The deletion is the fix.

## Why the substring matcher cannot be repaired in place

It is worth recording why "add better keywords" was rejected, so it is not
revisited.

`CategorizeScene` (`pkg/art/categories.go`) routes a scene by lower-casing its
**display name** and testing it against keyword lists in a fixed `switch` order.
Three failure modes, all live at the end of Phase 1:

1. **Order decides membership.** The cyber/neon case precedes the synthwave case,
   so "Kirokaze Retro City 🌃" and "Neon Synthwave Runner 🏎️" — the two genuinely
   retro scenes — land under "⚡ Cyberpunk & Neon" and never reach the synthwave
   branch.
2. **Keywords catch the wrong things.** The only scene that *does* reach the
   synthwave branch is "Coastal Ocean Lighthouse Study 🌊", matched on `coastal`.
   The retro category contains exactly one scene and it is a lighthouse.
3. **The input is not stable.** `SceneType.String()` prefers
   `SceneRegistry[s].Name()`, so a scene's category is a function of its display
   string. Renaming a scene silently recategorises it.

Any keyword fix is a fix to (2) that leaves (1) and (3) standing.

---

## Workstreams

### 2b.1 — Delete `pkg/art/categories.go`

`SceneCategories`, `CategorizeScene`, `GetScenesByCategory`,
`GetScenesByCategoryAndFavorites`, `SceneMeta`, `GetSceneMeta`. The whole file.

**Favorites must survive.** They are a real feature and are entangled with
categories only by accident: category 0 is "⭐ Favorites" and
`GetScenesByCategoryAndFavorites` special-cases `catIdx == 0` to filter
`SceneOrder` by the favorites map. Lift that filter into a standalone
`FavoriteScenes(favorites map[SceneType]bool) []SceneType`, and keep its
fall-back-to-all-scenes-when-empty behaviour, which is honest for favorites in a
way it was not for categories: an empty favorites list has no other sensible
answer, whereas an empty category does.

### 2b.2 — Collapse the picker to a flat list

`pkg/ui/model.go`: remove `PickerCategoryIdx` and the `[` / `]` category-tab
handling. `pkg/ui/view.go`: the picker renders `SceneOrder` (or the favorites
subset when that toggle is on), and the header loses the category name and tab
strip. Keep the `/` search filter — with 46 scenes it is now the primary way to
find one, and it does the job categories were nominally for.

`pkg/ui/command_palette.go`: the catalog command's description still says
*"categorized scene catalog"*.

### 2b.3 — Re-home the `SceneMeta` blurbs, or drop them

`GetSceneMeta` returns a `MoodTag` / `RecommendedWeather` / `Atmosphere` triple
keyed on category index, rendered at `view.go` in the picker detail pane and the
scene header. With categories gone there is no key.

Two options, and this is a **product call for whoever runs the phase**:

- **Drop them.** Simplest, consistent with the phase's direction, and honest —
  six generic blurbs stretched across 46 scenes were never scene-specific.
- **Move them onto the scene.** Add the fields to `SceneInterface` so each scene
  states its own mood and recommended weather. Better information, but it is
  46 × 3 strings to write and it widens an interface that **Phase 6 rewrites**.

Recommendation: **drop them**, and let Phase 6 add per-scene metadata to
`SceneInterface` v2 if it is still wanted. Deleting is reversible from git;
writing 138 strings against an interface about to change is not free.

### 2b.4 — Tests

- Delete `pkg/art/categories_test.go`, including
  `TestEveryCategoryMatchesARegisteredScene` — it guards a structure that no
  longer exists.
- Keep `pkg/ui/scene_count_test.go`. It asserts rendered output carries
  `len(art.SceneOrder)` rather than a hardcoded number, which is orthogonal to
  categories and stays valuable.
- Add a picker test asserting the flat list renders every registered scene, so
  a future filter bug cannot silently hide scenes the way the category fallback
  silently showed all of them.

---

## Known live issue this phase removes

`GetScenesByCategoryAndFavorites` still contains the fallback that caused the
original failure:

```go
if len(list) == 0 {
    return SceneOrder      // an empty category answers with every scene
}
```

Phase 1 deliberately left it. Dropping the empty category removed the *instance*;
the *mechanism* is still there, and after Phase 1 two categories hold exactly one
scene each ("🌌 Sci-Fi & Cosmic" and "🌅 Synthwave & Retro"), so either could
re-trigger it. Until this phase runs,
`TestEveryCategoryMatchesARegisteredScene` is what stands between a curation pass
and a silently mislabelled picker. Deleting the function deletes the mechanism —
do not "fix" the fallback first and call the phase done.

## Do NOT touch in this phase

- **Themes, weather, time-of-day, the focus buddy, presets, the timer.** The
  wider question of which features survive is **Phase 8's**, gated on reference
  frames. This phase removes one structure that is already provably broken; it is
  not the vehicle for the larger cut.
- `SceneOrder`, `SceneRegistry`, `RegisterScene` — Phase 6 owns the scene API.
- The `SceneType` enum. Still append-only until Phase 3.

---

## Verification

```bash
make check
test ! -f pkg/art/categories.go
grep -rn 'CategorizeScene\|SceneCategories\|GetSceneMeta\|PickerCategoryIdx' --include='*.go' .   # 0 hits
grep -rn 'categoriz' --include='*.go' .                                                           # 0 hits
go run .    # picker lists all 46 scenes; [ and ] no longer switch tabs; / still filters
```

- Favorites still work: mark one, toggle the favorites scope, confirm the picker
  narrows to it and that an empty favorites list still shows every scene.
- `pkg/ui/scene_count_test.go` still passes — the rendered count tracks the
  registry.

## Exit criteria

- [ ] `pkg/art/categories.go` deleted; zero references remain repo-wide
- [ ] Picker renders a flat scene list; `/` search intact; `[`/`]` unbound
- [ ] Favorites preserved as a standalone filter
- [ ] `SceneMeta` dropped (or explicitly re-homed with a recorded decision)
- [ ] `make check` green

## Re-validation checklist for Phase 3

1. **Did `PickerCategoryIdx` reach the persistence schema?** Phase 3 §3.2 lists
   UI state to persist. If categories are gone before Phase 3 runs, that field
   must not appear in the config schema — persisting a field for a deleted
   feature is the format break paying for nothing.
2. **`SceneMeta` decision?** If it moved onto `SceneInterface`, Phase 6's v2
   interface has three more methods to carry.
3. Confirm the scene count Phase 3's tests should expect: **46** at the end of
   Phase 1, unchanged by this phase.
