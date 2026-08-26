# Phase 2b — Simplify: Remove the Scene-Category System

> **Session size:** Small · **Depends on:** Phase 2 · **Touches:** `pkg/art/categories.go`, `pkg/art/categories_test.go`, `pkg/ui/model.go`, `pkg/ui/view.go`, `pkg/ui/command_palette.go`, `pkg/ui/view_test.go`, `pkg/ui/scene_count_test.go`

> **Numbered 2b, not 3.** Inserting a real Phase 3 would rename
> `plan-phase-3.md` through `plan-phase-17.md` — fifteen files — and every
> cross-reference in them, for no benefit. It runs after Phase 2 and before
> Phase 3.

---

## Prerequisites — findings from Phase 2

> Written at the end of Phase 2. Read before starting.
>
> `plan-phase-2.md`'s re-validation checklist forwards its three questions here,
> because 2b now runs between Phase 2 and Phase 3. They are answered in items
> 6-8 below.

1. **`assets/` is a Go package now.** Phase 2 added `assets/assets.go`
   (`package assets`, import `gottage/assets`) with `//go:embed animations/*.gif`,
   exposing `Open`, `ReadFile` and `List`. **This deviates from
   `plan-phase-2.md`'s "a new `pkg/assets/`"**: `//go:embed` cannot traverse
   `../`, so the package has to be an ancestor of the GIFs, and siting it at
   `assets/` needed no `git mv` and kept every existing reference to
   `assets/animations/` true.
2. **There is no CWD fallback, by design.** Resolution is `$GOTTAGE_ASSETS` then
   the embedded FS. A file absent from the override falls through to the embedded
   copy; an override directory that does not exist errors on every `Open`.
3. **`pkg/art` is CWD-independent** — `grep -rn '"\.\."' pkg/` is 0, and
   `TestScenesRenderFromArbitraryCWD` (`pkg/art/asset_resolution_test.go`) holds
   the line. Any test 2b adds must resolve through `assets.Open`, never by
   climbing to the source tree.
4. **The registry counts are asserted**: `wantRegistrations = 46` and
   `wantGifScenes = 10` in `pkg/art/asset_resolution_test.go`. 2b deletes the
   category taxonomy, not scenes, so both must still hold when it lands.
5. **`art.SceneByName` and `art.SceneNames` are new** (`scene_interface.go`),
   added for `--demo --scene`. `SceneByName` resolves against `SceneRegistry`
   directly and errors rather than guessing — 2b's flat picker may want either.
   **`GetScene`'s silent fallback to `SceneRainyCity` is still there**; Phase 2
   routed around it rather than removing it, so it remains a live silent-failure
   vector for any id→scene path.
6. **Test runtime: `pkg/art` went 0.55s → 0.85s**, four runs each side. The cost
   is the CWD test decoding all ten GIFs at 73×51; `-short` brings it to 0.57s.
   Whole suite is ~2.3s. There is headroom for Phase 3 to add tests without
   reaching for `testing.Short()` guards.
7. **Binary size: 7.46 MB → 9.30 MB.** The 1.78 MB of assets, near enough
   exactly. Comfortably inside §2.4's "acceptable for a personal tool", and
   nowhere near the ~10 MB-of-assets line at which that section says to
   reconsider embedding.
8. **Phase 1 renumbered no `SceneType`.** The block in `pkg/art/scenes.go` still
   holds **360** constants, `SceneRainyCity` first through
   `SceneAstralZodiacChronosAetherPortalStellar` last. Phase 3 migrates rather
   than repairs; no existing `user_presets.json` is corrupt.
9. **`art.AssetErrors()` records failed loads** (`pkg/art/asset_errors.go`),
   captured in a wrapper around `AnimatedGifLoader.load` so open and decode
   failures are both caught. **`pkg/art` writes nothing to stderr** — the rule is
   scoped to the TUI, where bubbletea owns the terminal in alt-screen mode and a
   log line at frame rate would corrupt the render it reports on. That path is
   still stderr-free. The one-shot `--demo` path is not the TUI and never enters
   alt-screen, so `main.go` reads `art.AssetErrors()` after rendering and, if it
   is non-empty, prints `gottage: N asset(s) failed to load…` plus one line per
   asset to stderr, suppresses the success trailer, and **exits 1**. Phase 16
   still owns surfacing inside the running app, and the `c.Clear(PicoNavy)`
   fallback is unchanged and still deliberate.
10. **`main.go` gained `--scene`**, which `plan-phase-2.md`'s Touches line does
    not list. It exists because exit criterion 1 was otherwise unverifiable:
    `--demo` renders `SceneRainyCity`, a handcrafted scene that reads no asset.
11. **Unparseable flag values now exit 2 on both paths.** `--scene`, `--weather`
    and `--time` reject an unrecognised *or empty* value with a diagnostic
    instead of falling back to a default — `--weather hail` and `--weather ""`
    exit 2 rather than launching the TUI with the default weather, where before
    they fell off the end of a `switch` with no default arm. `--time` also
    reaches the `--demo` frame now; it used to be read only after the demo block
    had already exited. **This was reviewed and accepted deliberately, not
    inherited** — a typo'd flag value silently doing the wrong thing is the
    failure class this whole rehabilitation removes, and it matches the `--scene`
    decision. Do not re-litigate it or restore a silent fallback. Detecting
    "flag passed empty" rather than "flag absent" needs `flag.FlagSet.Visit`;
    `main.go`'s `flagPassed` is the helper.

---

## Goal

Delete the scene-category system outright: the six-entry `SceneCategories`
picker taxonomy (`pkg/art/categories.go:8-15`), the substring matcher that
populates it, and the per-category `SceneMeta` blurbs. One fewer concept, roughly
90 lines gone, and a picker that lists scenes instead of sorting them into
buckets that no longer divide anything.

Because entry 0 of that taxonomy is the app's only route to favorites, the phase
also **adds a picker-level favorites filter**. That is not scope creep; see 2b.1.

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
into "⚙️ Steampunk & Fantasy" — all **48** of its members were among the 314 GIFs
cut, and not one was handcrafted — and `GetScenesByCategoryAndFavorites` answers
an empty category by returning `SceneOrder`, so the picker rendered *"Category:
⚙️ Steampunk & Fantasy (46 scenes)"* over every scene in the app. No error, no log.

**48, not 57.** 57 is the raw count of pre-curation display names containing a
steampunk-case keyword; nine of them never reached that case, because the cosmic
case sits above it and claimed them on `observatory`, `astral` or `abyssal` —
"Steampunk Sky Port Citadel Observatory ⚙️" and "Astral Clockwork Universe 🌌"
among them. Keep the two numbers distinct: the 9-name gap between them is the
order-decides-membership failure of the section below, measured rather than
asserted.

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

**Favorites must survive, and today the category system is their only route.**
Category index 0 is "⭐ Favorites", and `GetScenesByCategoryAndFavorites`
special-cases `catIdx == 0` to filter `SceneOrder` by the favorites map
(`pkg/art/categories.go:42-52`). Exactly one caller passes the favorites map into
it — `getFilteredScenes` (`pkg/ui/model.go:1035`), keyed on `m.PickerCategoryIdx`
— so deleting `PickerCategoryIdx` deletes browsing by favorite outright unless
this phase replaces it. `m.Scene.AutoCycleFavsOnly` is **not** a replacement: it
is read only by the auto-cycle tick (`pkg/ui/model.go:256`) and the ambiance
view's scope label (`pkg/ui/view.go:442`), never by the picker.

**A second reason to delete rather than repair.** The picker's *render* path does
not pass the map at all: `renderScenePickerModal` calls
`GetScenesByCategory(m.PickerCategoryIdx)` (`pkg/ui/view.go:853`), which forwards
`nil` favorites, so on the "⭐ Favorites" tab it renders the Cozy fallback while
the key handler (`model.go:599`) selects out of `getFilteredScenes`, the real
favorites list. Measured on a model with 2 favorites marked: 35 scenes rendered,
2 scenes selectable — `enter` applies a scene the user is not looking at. Two
call sites, two different answers to the same question. The replacement must
read its list once and render and select from that same slice.

So the phase must, in the same change:

- Lift the filter into a standalone
  `FavoriteScenes(favorites map[SceneType]bool) []SceneType`.
- Add a **dedicated picker-level favorites toggle** (2b.2), separate from
  `AutoCycleFavsOnly`. Do not couple them — one scopes what the picker lists, the
  other scopes what auto-cycling rotates through, and they are independently
  useful.

Shipping the deletion without the toggle removes a working feature silently,
which is the failure this phase exists to stop repeating.

**An empty favorites filter renders an explicit empty state**, rendered where the
picker already handles an empty list (`pkg/ui/view.go:883-884`).

**The copy must name an action that works from that state.** *"No favorites yet —
press `f` to add one"* does not. `f` marks the *highlighted* scene
(`model.go:607`) and its body is guarded by
`len(catScenes) > 0 && m.PickerSceneIdx < len(catScenes)`, so with the filter on
and no favorites there is nothing highlighted, and the key does nothing at all —
not even a status message. The only route out is to turn the filter off, then
mark a scene. So the copy has to point at the filter key first: *"No favorites
yet — press `F` to show all scenes, then `f` to add one"*, substituting whichever
key 2b.2 binds the filter to. An empty state that tells the user to press an
inert key is the same defect as a header that names the wrong list.

Do **not** carry the current behaviour forward, and do not replace it with
fall-back-to-all-scenes. What the code does today is fall back to the **Cozy
bucket**: `if len(list) == 0 { return GetScenesByCategoryAndFavorites(1, nil) }`
(`pkg/art/categories.go:48-50`), so a user with no favorites sees the header
*"⭐ Favorites (35 scenes)"* over 35 cozy scenes. That is the same silent
substitution that put 46 scenes under a Steampunk header — one instance of it is
being deleted here, and the phase must not introduce another in its place.
An empty result is information; answering it with a different list destroys that
information.

### 2b.2 — Collapse the picker to a flat list

**`pkg/ui/model.go`** — remove the `PickerCategoryIdx` field (`model.go:54`) and
the category-tab handling. The tabs are bound to **`h`/`left` and `l`/`right`**
(`model.go:617-622`), *not* to `[` and `]`; `[`/`]` have no key case anywhere in
the repo, so "unbind `[`/`]`" would be vacuously true and is not the work. Those
four bindings are the ones that must go, along with the `len(art.SceneCategories)`
modular arithmetic behind them. `getFilteredScenes` (`model.go:1032-1045`) drops
its `GetScenesByCategoryAndFavorites` call and returns `SceneOrder`, the search
matches, or `FavoriteScenes(m.FavoriteScenes)` depending on picker state.

**The favorites toggle.** Bind it inside the picker's non-search key switch
(`model.go:600-637`). Recommended: **`F`** — `f` marks the highlighted scene a
favorite (`model.go:607`), `F` filters the list to them. Record the one caveat:
`F` is globally Zen Focus Mode (`model.go:709`). It is not a functional conflict,
because the picker branch returns at `model.go:638` before the global switch is
reached, and Phase 13's command table is what reconciles per-modal bindings into
one source. If the phase would rather take zero collision, `o`, `y`, `N`, `L` and
`Z` are unbound everywhere in `model.go`; `s` and `t` are **not** available —
they are advertised-but-unbound and reserved for Phase 13's fix.

**Leave `f`'s empty-list guard in place.** `f`'s body is guarded by
`len(catScenes) > 0 && m.PickerSceneIdx < len(catScenes)` (`model.go:607-616`),
which stays correct once the filter exists: with the filter on and no favorites
there is genuinely nothing highlighted to toggle. That guard is what makes `f`
inert in the empty state, and it is why 2b.1's empty-state copy sends the user to
the filter key rather than to `f`. The two are one decision — if this phase
relaxes the guard so `f` does something from an empty list, the empty-state copy
has to change with it.

**`pkg/ui/view.go`** — four sites:

- `renderScenePickerModal` (`view.go:848-855`) drops the
  `GetScenesByCategory(m.PickerCategoryIdx)` branch and the tab strip
  (`view.go:856-864`); the header at `view.go:870` loses `"Category: %s"` and
  states the active scope instead (all scenes, or favorites).
- The footer at `view.go:945` renders
  `"Nav: [h/l] Category • [/] Search • [f/*] Favorite • …"`. The brackets there
  are notation *around* the key, not the key — that string must lose its
  `[h/l] Category` segment and gain the new favorites-filter key.
- The help modal at `view.go:635` reads
  `"[V] Open Categorized Scene Catalog Modal (%d Scenes)"`. It is the most
  prominent "Categorized" in the app and must change.
- `view.go:1001` builds `catName := art.SceneCategories[art.CategorizeScene(st)]`
  for the scene-overview `"Category:"` row. That row goes with the taxonomy.

Keep the `/` search filter — with 46 scenes it is now the primary way to find
one, and it does the job categories were nominally for.

**`pkg/ui/command_palette.go`** — the catalog command's description at
`command_palette.go:292` still says *"Open categorized scene catalog with %d
pixel-art presets"*.

### 2b.3 — Re-home the `SceneMeta` blurbs, or drop them

`GetSceneMeta` returns a `MoodTag` / `RecommendedWeather` / `Atmosphere` triple
keyed on category index, rendered in the picker preview card (`view.go:916`) and
the scene-overview pane (`view.go:979`). With categories gone there is no key.

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
- `pkg/ui/view_test.go:111` sets `m.PickerCategoryIdx = 0` to reach the favorites
  list. It **will not compile** once the field is removed; repoint it at the new
  favorites toggle, which keeps the assertion below it meaningful.
- Keep `pkg/ui/scene_count_test.go` — it asserts rendered output carries
  `len(art.SceneOrder)` rather than a hardcoded number, which is orthogonal to
  categories and stays valuable — but **one assertion in it must change**:
  `scene_count_test.go:43` pins the exact help-modal string
  `"Open Categorized Scene Catalog Modal (" + count + " Scenes)"`, which
  `view.go:635` is being reworded. Update the expectation, do not delete the case.
- Add a picker test asserting the unfiltered list **is** every registered scene —
  `getFilteredScenes()` returning `len(art.SceneOrder)` entries, the list the
  picker reads. Do **not** assert this against `Model.View()`: the picker windows
  its output to at most 11 rows (`endIdx := startIdx + 11`, `view.go:874-881`),
  so no rendered surface ever carries all 46 and such a test cannot pass. The
  goal is unchanged — a future filter bug must not silently hide scenes — but the
  list is where that is observable. The `SCENE CATALOG BROWSER (%d Total Scenes)`
  count (`view.go:942-943`) is the rendered half, and `scene_count_test.go`
  already pins it.
- Add a picker test for the empty favorites case: with no favorites marked, the
  filter renders the empty-state message and **not** a substituted list. Assert
  on rendered output through `Model.View()` here — an empty state *is* a
  rendering concern, the 11-row window does not affect a zero-row list, and the
  rendered surface is what the category bug got wrong.

---

## Known live issue this phase removes

`GetScenesByCategoryAndFavorites` still contains **two** empty-result fallbacks,
one of which caused the original failure:

```go
if len(list) == 0 {                                  // categories.go:48-50
    return GetScenesByCategoryAndFavorites(1, nil)   // no favorites → the Cozy bucket
}
...
if len(list) == 0 {                                  // categories.go:58-60
    return SceneOrder                                // an empty category → every scene
}
```

The second is the one that shipped the Steampunk header over 46 scenes. The first
is the same bug with a smaller blast radius, live right now: with no favorites
marked the picker reads *"⭐ Favorites (35 scenes)"* over the Cozy bucket.
2b.1's explicit empty state is what replaces it — both go, neither is re-created.

Phase 1 deliberately left them. Dropping the empty category removed the *instance*;
the *mechanism* is still there, and after Phase 1 two categories hold exactly one
scene each ("🌌 Sci-Fi & Cosmic" and "🌅 Synthwave & Retro"), so either could
re-trigger it. Until this phase runs,
`TestEveryCategoryMatchesARegisteredScene` is what stands between a curation pass
and a silently mislabelled picker. Deleting the function deletes the mechanism —
do not "fix" the fallbacks first and call the phase done.

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
grep -rni 'categoriz' --include='*.go' .                                                          # 0 hits
go run .    # picker lists all 46 scenes; h/l and left/right no longer switch tabs; / still filters
```

The `categoriz` grep is **case-insensitive on purpose**. Case-sensitive it finds
`command_palette.go:292` but misses `view.go:635`
(`"Open Categorized Scene Catalog Modal"`) and the assertion pinning that string
at `scene_count_test.go:43` — it would report 0 hits while the help modal still
called a flat list "Categorized". A check that passes with the most visible
wrong string still on screen is worse than no check.

Manual, in `go run .`:

- Press `V`, then `h`, `l`, `left`, `right`: none of them change the list. This
  is the real tab check — `[` and `]` were never bound, so testing those proves
  nothing.
- Mark a scene with `f`, toggle the favorites filter, confirm the picker narrows
  to exactly the marked scenes and the header says so.
- With the filter on, press `enter` on the highlighted row and confirm the scene
  that activates is the one on screen. This is the render/select divergence from
  2b.1; it must be gone, not merely relocated.
- Unmark everything, toggle the filter again: the empty state renders, and **no**
  substituted list appears. Cross-check the count against the header.
- `U` still toggles auto-cycle scope independently of the picker filter — the two
  are deliberately not wired together.

`pkg/ui/scene_count_test.go` still passes, with its `view.go:635` expectation
updated — the rendered count tracks the registry.

## Exit criteria

- [ ] `pkg/art/categories.go` deleted; zero references remain repo-wide
- [ ] `grep -rni 'categoriz' --include='*.go' .` returns 0 hits — including
      `view.go:635` and its assertion at `scene_count_test.go:43`
- [ ] Picker renders a flat scene list; `/` search intact; the category tabs'
      `h`/`left` and `l`/`right` bindings removed, and `Nav: [h/l] Category`
      gone from the footer at `view.go:945`
- [ ] Favorites reachable from the picker via a dedicated filter key, decoupled
      from `AutoCycleFavsOnly`
- [ ] Empty favorites renders an explicit empty state, not a substituted list,
      and a test asserts it
- [ ] `SceneMeta` dropped (or explicitly re-homed with a recorded decision)
- [ ] `make check` green

## Re-validation checklist for Phase 3

2b sits between Phase 2 and Phase 3, so **this is the checklist Phase 3 runs** —
`plan-phase-3.md`'s Prerequisites section points here. Phase 2's own checklist
(`plan-phase-2.md`, "Re-validation checklist for Phase 3") is not skipped: its
three items are carried forward as 1-3 below, because 2b does not touch the
embed, the binary, or the enum and cannot answer them itself.

**Carried forward from Phase 2:**

1. **Did embed change test runtime?** `pkg/art` was 17s (222s under `-race`). If
   embedding pushed it materially higher, Phase 3 should add the
   `testing.Short()` guards before adding more tests.
2. **Binary size acceptable?** If not, revisit before Phase 9/10 add more assets.
3. **Did Phase 1 renumber any `SceneType`?** Phase 3 needs Phase 1's enum
   decision, not Phase 2's or 2b's code.

**From 2b:**

4. **Did `PickerCategoryIdx` reach the persistence schema?** Phase 3 §3.2 lists
   UI state to persist. If categories are gone before Phase 3 runs, that field
   must not appear in the config schema — persisting a field for a deleted
   feature is the format break paying for nothing. The favorites filter that
   replaces it is transient picker state and should not be persisted either; the
   favorites *map* already is.
5. **`SceneMeta` decision?** If it moved onto `SceneInterface`, Phase 6's v2
   interface has three more methods to carry.
6. Confirm the scene count Phase 3's tests should expect: **46** at the end of
   Phase 1, unchanged by this phase.
7. **Schema scope.** Phase 3 persists conservatively — see its Prerequisites
   section. The *Virtual Cottage* cut is deferred to Phase 8, which runs later,
   so §3.2 must not persist state whose survival Phase 8 decides.
