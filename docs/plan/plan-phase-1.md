# Phase 1 — Content Triage: Delete, Curate, Fail Loudly

> **Session size:** Large (mostly deletion) · **Depends on:** Phase 0 · **Touches:** `pkg/art/`, `assets/`

---

## Prerequisites — findings from Phase 0

> Written at the end of Phase 0. Read before starting; items 1 and 2 are blockers.

1. **`git-filter-repo` is NOT installed.** `brew install git-filter-repo` is a hard
   prerequisite for §1.8. Install and confirm it runs *before* the first deletion
   commit, not when you reach the purge — the whole phase's deletions are the
   purge list, and discovering the tool is missing at the end is a bad time.
2. **Tag before deleting anything:** `git tag phase-0-end` (§1.8 step 0). The purge
   list is derived from `phase-0-end~1..HEAD`, so the tag must predate every
   deletion commit in this phase.
3. **Backup verified.** `/Users/nish/proj/gottage-backup.git` is a mirror clone
   holding all 95 pre-Phase-0 commits and all three original refs. After §1.8 it
   is the only copy of the old history; Phase 17 §17.5 decides its fate.
4. **Collision count re-derived: exactly 20.** §1.1's extraction was run verbatim
   and gives 344 GIF ids, 36 handcrafted ids, 20 in the intersection — matching
   the table in §1.2 entry-for-entry. `grep -c 'RegisterScene('
   pkg/art/scenes_modular.go` = **380**. Re-run if scenes were hand-edited since.
   Note the check must be **cross-group** (handcrafted `ID()` returns vs. GIF
   `id:` literals); a within-group uniqueness check finds nothing and is
   misleading.
5. **Anchors held — Phase 0 shifted no `.go` line numbers.** `RegisterScene` is at
   `pkg/art/scene_interface.go:30` (doc cites 29-34, correct). One correction:
   `SceneOrder` is declared at `scene_interface.go:27`, **not** in
   `categories.go`.
6. **`gofmt` is untouched, by design.** Still 16 of 34 files, 1,745 diff lines.
   `make check` fails on `fmt-check` and passes `vet`, `tidy-check` and `test`
   individually. §1.7 (`gofmt -w`, last commit of the phase) is what turns it
   green. The pre-commit hook gates **staged** files only, so it will not block
   your deletion commits.
7. **Lint is not blocking.** `golangci-lint` is not installed and `make lint`
   degrades gracefully; `check` does not depend on it. `.golangci.yml` excludes
   `errcheck` on `pkg/audio/`, `pkg/features/presets.go` and `pkg/(ui|art)/` with
   `TODO(phase-N)` markers, so Phase 1 inherits zero lint noise. The config is
   v2-schema and **unvalidated against the real binary** — whoever installs it
   first should confirm it parses.
8. **`.gitignore` now covers `/gottage`.** §1.8's safety gate (`comm -12` against
   `git ls-files`) is unaffected since the binary is no longer tracked, but the
   binary must still appear in the purge list — keep §1.8 step 1's
   `grep -x gottage /tmp/purge-paths.txt` check. Note `git rm --cached` in Phase 0
   registers as a deletion in `git diff --diff-filter=D`, so it should appear
   automatically given the tag is placed per item 2.
9. **Write the curation aesthetic bar first.** `AGENTS.md` §7 has a placeholder
   for it. §1.3 says it must exist before cutting starts, since it is the
   criterion Phase 8's reference frames get judged against. Fill it in as the
   first act of §1.3, not retroactively.
10. **Asset inventory as of Phase 0 end:** 363 GIFs + 4 JPGs = 367 tracked files,
    31MB. `.git` is 89MB. 430 tracked files total.

---


## Goal

Delete everything that is dead, curate what stays, and make the registry **panic** on the class of mistake that hid 3,177 lines of art. Net effect: ~7,900-line `scenes_modular.go` down to roughly 3,500; assets 31MB → ~6MB; `.git` 88MB → ~8MB; 20 handcrafted scenes that never rendered start rendering.

## Why this phase is second

Deletion is the only operation no later phase can obsolete. Doing it now means every subsequent phase — the canvas rewrite, the scene API change, the component library, the migration — operates on ~55% less material. Curating before Phase 2's `//go:embed` is what makes embedding a 6MB proposition instead of a 31MB one.

---

## The headline bug

`pkg/art/scene_interface.go:29-34`:

```go
func RegisterScene(s SceneInterface) {
	if _, ok := SceneRegistry[s.ID()]; !ok {
		SceneOrder = append(SceneOrder, s.ID())
	}
	SceneRegistry[s.ID()] = s     // <-- unconditional overwrite
}
```

`init()` in `pkg/art/scenes_modular.go` registers 36 handcrafted scenes (lines 10-45), then 344 `AnimatedGifScene` literals (lines 47-2110). **20 SceneType IDs are claimed by both.** The GIF registers later and silently wins.

Reproduce:

```bash
grep -oE '^\s*id:\s*(Scene[A-Za-z0-9]+)' pkg/art/scenes_modular.go \
  | grep -oE 'Scene[A-Za-z0-9]+' | sort -u > /tmp/gifids.txt
awk '/^func \([a-z]+ \*[A-Za-z]+\) ID\(\) SceneType/{
  if (match($0, /return Scene[A-Za-z0-9_]+/)) { print substr($0, RSTART+7, RLENGTH-7); next }
  getline; if (match($0, /return Scene[A-Za-z0-9_]+/)) print substr($0, RSTART+7, RLENGTH-7)
}' pkg/art/scenes_modular.go | sort -u > /tmp/handids.txt
comm -12 /tmp/handids.txt /tmp/gifids.txt | wc -l    # -> 20
```

The last 8 commits were all *"Implemented 4 new handcrafted cozy procedural scenes."* Most shipped dead on arrival. `TestHandcraftedProceduralScenes` (`pkg/art/scene_test.go:112`) names all 36 and passes anyway, because it asserts `len(rendered) != 0` — and when the GIF path fails to load from the test's CWD it falls back to `c.Clear(PicoNavy)`, so **the test asserts that a solid navy rectangle is non-empty.**

### The 20 shadowed scenes

| SceneType ID | Shadowed Go type | Lines |
|---|---|---|
| `SceneZenTeahouse` | `ZenTeahouseScene` | 2849-3042 |
| `SceneCozyCampfire` | `CozyCampfireScene` | 3043-3218 |
| `SceneCozyAttic` | `CozyAtticScene` | 3219-3339 |
| `ScenePixelWinterCabin` | `WinterCabinScene` | 3340-3448 |
| `SceneRetroLofiTrain` | `RetroLofiTrainScene` | 4056-4232 |
| `SceneZenWaterfallShrine` | `ZenOnsenScene` | 4233-4385 |
| `SceneNeonBistro` | `VinylLoungeScene` | 4386-4537 |
| `SceneStarlightCamp` | `SunkenHearthScene` | 4538-4687 |
| `SceneCoffeeShop` | `AntiqueBookshopScene` | 5359-5516 |
| `SceneEnchantedForest` | `GlasshouseBotanicalScene` | 5517-5663 |
| `SceneRiverBridge` | `RiversideCampScene` | 5664-5827 |
| `SceneOceanWaves` | `CoastalBalconyScene` | 5828-5980 |
| `SceneSunsetStation` | `LofiStudyDeskScene` | 5981-6157 |
| `SceneBambooGrove` | `ZenRainGardenScene` | 6158-6326 |
| `SceneWinterCafe` | `CozyCatCafeScene` | 6327-6460 |
| `SceneAutumnPark` | `AutumnLakeDockScene` | 6461-6599 |
| `SceneZenBonsai` | `MachiyaGardenScene` | 7207-7380 |
| `SceneCyberDiner` | `CyberLofiDinerScene` | 7381-7572 |
| `SceneCozyIgloo` | `NordicBayWindowScene` | 7573-7734 |
| `SceneSkyIsland` | `ForestTreehouseScene` | 7735-7910 |

Line ranges will drift as you delete — re-derive rather than trusting them.

---

## Workstreams

### 1.1 — Make `RegisterScene` panic on duplicate

```go
func RegisterScene(s SceneInterface) {
	if existing, ok := SceneRegistry[s.ID()]; ok {
		panic(fmt.Sprintf("duplicate scene ID %v: %T would shadow %T", s.ID(), s, existing))
	}
	SceneOrder = append(SceneOrder, s.ID())
	SceneRegistry[s.ID()] = s
}
```

**Panic, not log-and-skip.** Registration happens in `init()`, so a panic fails the process *and* every test binary immediately and unmissably. Log-and-skip merely converts a silent-overwrite bug into a silent-skip bug — the identical failure mode that produced this state.

Do this **first**. The panic then drives the rest of the phase: run the tests, read the panic, delete the offending registration, repeat.

### 1.2 — Delete the 20 colliding `AnimatedGifScene` registrations

Per the confirmed decision (**procedural first, curate GIFs**), the handcrafted scene wins every collision. When a hand-drawn scene and a bulk GIF collide, the hand-drawn one is what 8 commits of intent were aimed at.

Delete the 20 `RegisterScene(&AnimatedGifScene{...})` blocks whose `id:` appears in the collision list, and their GIF files.

**`SceneOrder` is unaffected.** Verified: `RegisterScene` appends to `SceneOrder` only on *first* registration, and the 36 handcrafted scenes register first — so the ordering is already the handcrafted ordering. Only the map *value* was wrong. Picker order, `CategorizeScene` (which buckets by substring-matching `st.String()`, `pkg/art/categories.go:19`), and saved preset ints are all untouched.

### 1.3 — Curate the remaining GIFs

Target ~40 keepers from the 324 remaining. **Write the aesthetic bar down in the commit message or `AGENTS.md`** — it is the criterion Phase 8's reference frames will be judged against, and it needs to exist before you start cutting.

Suggested criteria: reads clearly at 74×52 half-blocks; a coherent cozy interior or view (not an abstract loop); visible micro-animation; a palette compatible with the app's warm tokens; **under ~300KB**.

**The size cap is doing more work than it looks.** Measured distribution across the current 363 GIFs: 27.0MB total, mean 76KB — but **12 files exceed 500KB, and the top 5 alone account for 9.0MB** (`rain_setup.gif` 2.8MB, `ocean_waves.gif` 1.9MB, `snow_cabin.gif` 1.8MB, `stargazing.gif` 1.5MB, `magic_sanctuary.gif` 1.0MB). So the keeper set's footprint depends almost entirely on whether the whales survive:

| 40 keepers chosen as… | Footprint |
|---|---|
| the 40 largest | 18.3 MB |
| a median band | ~1.0 MB |
| the 40 smallest | 0.5 MB |

Enforcing the 300KB cap lands the whole asset set at **1-3MB** rather than 6MB. That matters downstream: it halves Phase 2's embedded binary, and it is the reason Git LFS is unnecessary (see `plan-phase-0.md` §0.2b).

A big GIF is usually big because it has many frames or a wide palette — neither of which survives downsampling to 74×52 half-blocks. **If a scene is genuinely worth keeping but oversized, re-encode it** (fewer frames, reduced palette, cropped) rather than either dropping it or shipping the whale.

```bash
# review candidates by size before deciding
stat -f '%z %N' assets/animations/*.gif | sort -rn | awk '{printf "%6.0f KB  %s\n", $1/1024, $2}' | head -30
```

Delete non-keepers **and their registration blocks together** — an orphan registration is now a panic.

Also delete outright:

- **19 GIFs already orphaned** (on disk, referenced by zero registrations): `cloud_temple`, `cozy_greenhouse`, `cozy_library`, `cyber_bar`, `fairytale_windmill`, `lighthouse_cliff`, `magic_sanctuary` (1.07MB), `mountain_aurora_shack`, `pixel_bakery`, `ramen_shop`, `starlight_treehouse`, `wizard_tower`, `astral_zodiac_chronos_aether_portal_orbit`, `biolum_crystal_abyssal_tide_oasis`, `cyber_neon_matrix_stargazer_deck`, `steampunk_sky_clockwork_citadel_lounge`, `steampunk_sky_clockwork_citadel_terrace`, `synthwave_grid_neon_coastal_dash`, `zen_bamboo_moon_shrine_valley_spring`.
- **All 4 `assets/scenes/*.jpg` (3.5MB)** — their only possible consumer, `art.LoadImageToCanvas`, has zero callers.
- The biggest offenders should be re-encoded or dropped regardless: `rain_setup.gif` (2.9MB), `ocean_waves.gif` (1.9MB), `snow_cabin.gif` (1.8MB), `stargazing.gif` (1.5MB).

The registration table shrinks from **2,064 lines to roughly 250**.

Prune the `SceneType` enum (`pkg/art/scenes.go:12-373`) to match. **Delete constants from the end, or renumber nothing** — reordering the const block silently corrupts every saved preset, which serializes raw iota ints. (Phase 3 fixes that permanently; until then, treat the enum as append-only.)

### 1.4 — Delete unambiguously dead code

Each verified as zero-caller:

| Target | Location | Size |
|---|---|---|
| `LoadImageToCanvas` + whole file | `pkg/art/image_loader.go` | 83 lines |
| `RenderRainyCity`, `RenderForestCabin`, `RenderCottageClassic`, `RenderObservatory`, `render8BitDesk`, `render8BitFireplace`, `render8BitCat` | `pkg/art/scenes.go:403-700` | ~300 lines |
| `Canvas.SetPixelHex` | `pkg/art/canvas.go:103` | |
| `AnimatedGif.DrawFrameAt` | `pkg/art/gif_loader.go:156` | only its own test calls it |
| `Particle.Char` | `pkg/art/scene.go:261` | written at `:440`, `:580`; never read |
| `CLIAMPController.SetVolume` | `pkg/audio/cliamp.go:280` | |
| `AmbientManager.StatusSummary` | `pkg/audio/ambient.go:369` | |

### 1.5 — Do NOT delete these two (they are features, not dead code)

- **`AnimatedGif.Delays`** (`pkg/art/gif_loader.go:20`) is populated at `:95` and never read; playback uses `tick % FrameCount` at `:138`. So **every GIF plays at the 60ms UI tick, ignoring its authored timing** — a 10fps GIF renders at 16.7fps, a 4fps one runs 4× too fast. Given the user's literal complaint is *"what we have is not quite what i want"* about the art, honoring frame delays is plausibly the highest art-quality-per-line change available. **Phase 15 implements it.** Leave a `TODO(phase-15)`.
- **`AnimatedGifScene.showBuddy` / `buddyX` / `buddyY`** (`:2121-2123`) are never set true, so no GIF scene contains the companion character. That is a product gap for **Phase 8/9**, not dead code. Leave a `TODO(phase-8)`.

### 1.6 — Replace the vacuous test

`TestHandcraftedProceduralScenes` (`pkg/art/scene_test.go:112`) starts exercising real handcrafted code the moment 1.2 lands. If you fix the bug and keep `len(rendered) != 0`, you have kept the assertion that could never have caught it.

Replace with a **distinct-RGB-count** assertion: render, count unique pixel colors, require ≥ N (say 12) and reject monochrome. Cheap, needs no golden files, and catches the navy-rectangle case directly.

Add alongside:

- `TestNoDuplicateSceneRegistrations` — `len(SceneRegistry) == len(SceneOrder)`.
- `TestHandcraftedScenesResolveToHandcrafted` — table over all 36, type-asserting e.g. `GetScene(SceneZenTeahouse).(*ZenTeahouseScene)`. **This is the regression test for this entire phase.**
- `TestEverySceneHasAsset` — every registered `gifPath` exists on disk (catches curation mistakes).
- Fix `TestVerifyDownloadedGifs` (`pkg/art/gif_loader_test.go:41`): it decodes **all** GIFs at 160×90 into the unbounded global cache (~3.6GB resident) and asserts `gifCount < 356`, a hardcoded count that breaks the instant anyone prunes an asset — which this phase does. Assert against the registration count instead, and add a `testing.Short()` guard.

### 1.7 — `gofmt -w` as the final commit

16 of 34 files are unformatted (1,745 diff lines). Run it **last**, as its own commit, after all deletions — reformatting lines you then delete is wasted diff and guarantees conflicts with every other edit in this phase.

### 1.8 — One history purge, covering the binary *and* the dropped assets

**The last act of the phase. Irreversible — it rewrites all 95 commits.**

Confirmed with the user: rather than purging the binary in Phase 0 and the assets later, both go in **one** `git-filter-repo` pass here. This is the correct order — you cannot purge assets from history until curation has decided which survive — and it means all SHAs are rewritten exactly once.

**Preconditions:** Phase 0's backup mirror exists; working tree clean; every deletion above committed; the `gofmt` commit landed.

#### Step 0 — Tag the starting point (do this at the *start* of Phase 1)

```bash
git tag phase-0-end        # before any Phase 1 deletion commits
```

Everything deleted from here to the end of the phase is exactly the purge list.

#### Steps 1-4 — Build the list, check it, purge, verify

```bash
# 0. Verify the backup before doing anything irreversible
git -C /Users/nish/proj/gottage-backup.git log --oneline | wc -l   # -> 95
git status --porcelain                                             # must be empty

# 1. The purge list = every path deleted across Phase 0 + Phase 1, net.
#    --name-only lists files only, never directories. Verified against this repo.
git diff --diff-filter=D --name-only phase-0-end~1..HEAD > /tmp/purge-paths.txt

wc -l /tmp/purge-paths.txt              # ~330 (≈304 curated GIFs + 19 orphans + 4 JPGs + binary)
grep -x gottage /tmp/purge-paths.txt    # the binary must be in there

# 2. SAFETY GATE — a surviving asset in this list would be destroyed.
git ls-files | sort -u > /tmp/keep.txt
comm -12 <(sort -u /tmp/purge-paths.txt) /tmp/keep.txt     # MUST be empty
grep -vE '^(assets/|gottage$)' /tmp/purge-paths.txt        # review anything unexpected

# 3. Purge  (brew install git-filter-repo)
git filter-repo --paths-from-file /tmp/purge-paths.txt --invert-paths --force

# 4. Verify
du -sh .git                                        # ~88M -> ~8M
git log --all --oneline -- gottage | wc -l         # 0
git ls-files assets/ | wc -l                       # unchanged from before the purge
go build ./... && go test ./... -count=1           # working tree intact
```

#### Gotchas that matter here

- **Step 2 is not optional.** `--invert-paths` removes every listed path from all history. One surviving asset in that list and it is gone from the working tree too — recoverable only from the backup mirror.
- **Do not build the list from `git rev-list --objects --all`.** That prints *tree* objects as well as blobs, so bare directory paths (`assets`, `assets/animations`, `assets/scenes`) land in the list — and since `git ls-files` never lists directories, a naive `comm` diff will not filter them out. Purging `assets/animations` deletes **every** GIF including the curated keepers. (Verified against this repo: those three tree paths do appear.) `git rev-list` also dedups by blob SHA, so byte-identical files report only one of their paths — 367 tracked assets currently resolve to just 349 reported paths. The `git diff --diff-filter=D --name-only` form has neither problem.
- `git-filter-repo` refuses to run with uncommitted changes, and needs `--force` because this is not a fresh clone.
- It rewrites **every** SHA, including Phase 0's. Harmless — no phase doc cites a commit SHA, only `file:line` anchors, which are untouched. The `phase-0-end` tag is consumed by the rewrite; delete it afterwards.
- It drops configured remotes and runs its own `reflog expire` + `gc --prune=now`. No remote here, so nothing to restore.
- **Afterwards the backup mirror is the only copy of the old history.** Keep it until confident; Phase 17 §17.5 decides its fate.

Expected result: `.git` from **88MB to roughly 8MB** — the in-pack total was 83.4 MiB, dominated by 39 copies of a 7MB binary plus ~25MB of GIFs that do not deltify.

---

## Do NOT touch in this phase

- **`AnimatedGif.Delays` and `showBuddy`/`buddyX`/`buddyY`** — see 1.5.
- Rendering logic, the canvas serializer, the palette, `SceneInterface`'s shape. This phase deletes and gates; Phases 4/6/7 redesign.
- The 344→~40 GIF *loading* path (per-frame `LoadAnimatedGif`, unbounded cache). Phase 6 owns it via the `Activate` hook.
- Scene art content. Do not "improve" a handcrafted scene while you are here — Phases 9/10 own that, with goldens in place.

---

## Verification

```bash
go build ./... && go vet ./...
go test ./... -count=1                          # panic-free => zero duplicate registrations
gofmt -l . | wc -l                              # 0
make check                                      # now fully green
du -sh assets/                                  # ~1-3M with the size cap held (was 31M)
stat -f '%z' assets/animations/*.gif | sort -rn | head -1   # largest keeper under ~300KB
du -sh .git                                     # ~8M  (was 88M)  -- after 1.8
git log --all --oneline -- gottage | wc -l      # 0
grep -c 'RegisterScene(&' pkg/art/scenes_modular.go   # ~76  (was 380)
wc -l pkg/art/scenes_modular.go                 # ~3500 (was 7909)
test ! -f pkg/art/image_loader.go
test ! -d assets/scenes

# every registered asset exists
grep -oE 'assets/animations/[a-z0-9_]+\.gif' pkg/art/scenes_modular.go | sort -u | while read -r f; do
  test -f "$f" || echo "MISSING: $f"
done
```

New tests must pass: `TestNoDuplicateSceneRegistrations`, `TestHandcraftedScenesResolveToHandcrafted`, `TestEverySceneHasAsset`, and the rewritten `TestHandcraftedProceduralScenes`.

**Manual check that matters:** `go run . ` then press `v` through the 20 formerly-shadowed scenes. They should show hand-drawn art, not a looping GIF. This is the payoff for 8 commits of work.

## Exit criteria

- [ ] `RegisterScene` panics on duplicate; test suite runs panic-free
- [ ] All 36 handcrafted scenes type-assert to their handcrafted type
- [ ] `assets/` ≈ 1-3MB and no keeper over ~300KB; every registered `gifPath` resolves; zero orphan GIFs
- [ ] All 1.4 dead code deleted; `Delays` and `showBuddy` preserved with TODO markers
- [ ] `TestHandcraftedProceduralScenes` asserts distinct-color-count, not `len != 0`
- [ ] `gofmt -l .` empty; `make check` green
- [ ] Aesthetic curation bar written down
- [ ] **History purged in one pass**: `.git` ~8MB, zero `gottage` blobs, zero dropped-asset blobs, every surviving asset still present and tests green

## Re-validation checklist for Phase 2

1. **Final asset footprint and largest keeper?** Phase 2 embeds them. If curation landed well above ~10MB, reconsider embed-vs-runtime-dir — but prefer re-encoding the offenders, since a runtime dir reopens the CWD bug Phase 2 exists to close. Also confirm **no one added Git LFS** (see `plan-phase-0.md` §0.2b) — it would silently break the embed.
2. **Enum pruned or left append-only?** Phase 3's enum→string migration needs to know whether any `SceneType` constant was renumbered — if so, existing `user_presets.json` files are already corrupt and Phase 3 must handle that, not just migrate.
3. **Registration count** — Phase 2's embed test asserts every registered path resolves from an embedded FS. Confirm the count Phase 2 should expect.
4. **Purge went cleanly?** Confirm `git ls-files assets/ | wc -l` matches the pre-purge kept count and `go test ./...` passes. If a kept asset was destroyed, restore from the backup mirror and redo — do not patch around it.
5. **Line anchors shifted heavily.** Every `pkg/art/scenes_modular.go` line number in Phase 4+ docs is now wrong. Re-derive with `grep -n`; do not trust cited line numbers in later docs for this file.
