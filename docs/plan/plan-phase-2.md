# Phase 2 — Asset Pipeline & Embed

> **Session size:** Small · **Depends on:** Phase 1 · **Touches:** `pkg/art/gif_loader.go`, `pkg/art/scenes_modular.go`, `pkg/art/scene_interface.go`, `assets/`, `main.go`

---

## Prerequisites — findings from Phase 1

> Written at the end of Phase 1. Read before starting.

1. **Asset footprint is 1.78MB across 10 GIFs** — well inside §2.4's comfortable
   range, not the ~6MB the goal line predicts. Embed is unambiguously the right
   call: measured, the binary is 7.5MB and lands at **~9MB embedded**. No Git LFS
   was introduced; `.gitattributes` still marks binaries as binary.
2. **Registration count is 46** — 36 handcrafted + **10** `AnimatedGifScene`. That
   is the number `TestEveryRegisteredAssetResolves` should expect, and the type
   switch to reach `gifPath` is `sc.(*AnimatedGifScene)` from inside `package art`.
3. **`SceneType` was NOT pruned.** All 360 constants remain, no renumbering, so
   Phase 3 migrates rather than repairs — no existing `user_presets.json` is
   corrupt. The 314 cut ids are simply unregistered.
4. **`GetScene` silently falls back to `SceneRainyCity`** for an unregistered type
   (`scene_interface.go`). That masked two built-in presets pointing at cut scenes;
   `TestBuiltinPresetScenesAreRegistered` (`pkg/features/presets_test.go`) now
   guards it. Phase 2's §2.3 error-surfacing should treat this fallback as another
   instance of the same silent-failure class.
5. **Every asset path still resolves** and there are zero orphans, asserted both
   directions by `TestEverySceneHasAsset` (`pkg/art/scene_test.go`). Paths are
   still CWD-relative and the test joins `"..", ".."` — Phase 2 is what fixes that,
   and should update the test to resolve through `assets.Open` instead.
6. **Test runtime dropped from ~21s to ~2s.** `TestVerifyDownloadedGifs` was 15.75s
   of it; it now decodes 10 GIFs rather than 363 and is `testing.Short()`-guarded.
   There is headroom for Phase 2 to add tests.
7. **Line anchors in `scenes_modular.go` are all wrong now.** The file went
   7909 → 6035 lines via a repo-wide `gofmt` plus 334 deleted registration blocks.
   Re-derive with `grep -n`.
8. **History was purged** — `.git` 92MB → **2.2MB**, zero `gottage` blobs. Two
   deviations from `plan-phase-1.md` §1.8 worth knowing:
   - Its purge base (`phase-0-end~1..HEAD`) yields an **empty list** and would have
     left the binary in history. The correct base was `65ba4e3`.
   - Path-only purging leaves the 10 keepers' pre-re-encode blobs (8.40MB) alive,
     since those paths still exist. They were removed with
     `--strip-blobs-with-ids` after verifying none was reachable from HEAD.
   All SHAs are rewritten. One assets-only commit was pruned as empty. Backups:
   `gottage-backup.git` (95 pre-Phase-0 commits) and `gottage-prepurge.git`
   (111 commits, immediately pre-purge).

---

## Goal

Make the binary run correctly from **any** working directory by embedding the curated assets, and route every asset read through one resolver.

## Why this phase is here

It depends on Phase 1 (embed only survivors — 6MB, not 31MB) and blocks nothing structural, so it slots in early and cheaply. Every later phase's tests benefit: today, running `go test ./pkg/art` from the package directory silently exercises the GIF *failure* path.

## The bug

> **Fixed by this phase.** The section below describes the tree as it stood
> *before* Phase 2, including a `grep` that now returns hits. Kept in that voice
> because it is the argument for why the phase exists.

There is **no `//go:embed` anywhere** in the repo (verified: `grep -rn 'embed' --include='*.go' .` → zero hits). All GIF paths are relative string literals like `"assets/animations/kirokaze_room.gif"`, opened with plain `os.Open` at `pkg/art/gif_loader.go:50`.

So **every GIF scene silently renders as flat navy unless the process CWD is the repo root**:

```go
// pkg/art/scenes_modular.go, AnimatedGifScene.RenderSky
anim, err := LoadAnimatedGif(s.gifPath, c.Width, c.Height)
if err == nil && anim != nil {
    anim.DrawFrame(c, frameCount)
} else {
    c.Clear(PicoNavy)          // <-- silent failure, no error surfaced
}
```

Reproduce:

```bash
go build -o /tmp/gottage . && cd /tmp && ./gottage    # GIF scenes are solid navy
```

This is also why `TestHandcraftedProceduralScenes` was able to assert that a navy rectangle is non-empty — the test runs from `pkg/art/`, so the relative path never resolved.

---

## Workstreams

### 2.1 — `assets` with an embedded FS

**Landed at `assets/assets.go`, not `pkg/assets/`.** `//go:embed` cannot traverse
`../`, so the package has to be an ancestor of the files it embeds. Making the
existing `assets/` directory a Go package needed no `git mv` and kept every
existing reference to `assets/animations/` true — `AGENTS.md` §5 and §7, five
test path joins, `README.md`, and the `assets/` naming in the Phase 8-10 docs.

```go
package assets   // import "gottage/assets"

//go:embed animations/*.gif
var embedded embed.FS

// Open resolves an asset by its registration path, e.g. "assets/animations/x.gif".
func Open(name string) (fs.File, error)
func ReadFile(name string) ([]byte, error)
func List() ([]string, error)
```

Resolution order:

1. `$GOTTAGE_ASSETS` if set — an override directory. Keeps iteration fast during Phases 8-10 (drop in a GIF without rebuilding) and gives a supported escape hatch.
2. The embedded FS.

Do **not** fall back to CWD. Silent CWD resolution is the bug being fixed; an explicit env override replaces it.

A file *absent* from the override falls through to the embedded copy, so
overriding one asset does not mean copying all ten. The silent-failure risk that
opens — a typo'd directory quietly reading the embedded set — is closed by making
a `$GOTTAGE_ASSETS` that is not an existing directory an error on every `Open`.
The env var is read per call rather than cached, so `t.Setenv` can reach it.

### 2.2 — Route `gif_loader.go` through the resolver

Replace `os.Open(filePath)` at `pkg/art/gif_loader.go:50`. Keep the registration paths (`"assets/animations/x.gif"`) as the canonical key so nothing else changes.

### 2.3 — Surface the failure instead of swallowing it

`c.Clear(PicoNavy)` on load failure is how this bug survived. At minimum, log once per path (not per frame — this runs 16.7×/sec) behind a package-level "asset error" record that Phase 16's status-line error channel can read. A `TODO(phase-16)` is fine; a silent fallback is not.

### 2.3b — Why the assets must NOT be Git LFS objects

**Recording this here because Phase 2 is where it would break, and the breakage is silent.**

A Git LFS pointer file is about 130 bytes of plain text:

```
version https://git-lfs.github.com/spec/v1
oid sha256:4d7a2f...
size 28419
```

`//go:embed` reads whatever is **on disk at build time**. If the LFS smudge filter has not run — a clone without `git lfs pull`, LFS not installed, a CI or container image lacking it, a `GIT_LFS_SKIP_SMUDGE` env var — then the file on disk *is* that pointer text.

So `go build` would embed **130 bytes of text where a GIF should be**. It would not error. `assets.Open` would succeed. `LoadAnimatedGif` would fail to decode, and the scene would fall back to `c.Clear(PicoNavy)`.

That is precisely the failure this phase exists to eliminate, reintroduced through a different door — and it is exactly the class of silent failure (`len(rendered) != 0` passing on a navy rectangle) that let 20 broken scenes ship across 8 commits.

If LFS is ever adopted, embedding must be abandoned first, not simultaneously. See `plan-phase-0.md` §0.2b for the full decision and its revisit conditions.

**Guard it:** `TestEmbeddedAssetsAreRealGifs` — assert every embedded asset starts with the GIF magic bytes (`GIF87a`/`GIF89a`) and exceeds a plausible minimum size. Ten lines, and it makes this failure mode impossible to ship.

### 2.4 — Binary size sanity

If Phase 1's size cap held, assets should be **1-3MB**, putting the binary at ~8-10MB. Acceptable for a personal tool.

**Measured: 7.46 MB → 9.30 MB**, the 1.78 MB of assets near enough exactly.

If it landed materially above ~10MB of assets, stop and reconsider a runtime data directory before committing to embed — but note that going the other way (runtime dir) reopens the CWD-dependency bug this phase exists to close, so prefer tightening curation over abandoning embed.

### 2.5 — `--scene`, because exit criterion 1 was unverifiable

`--demo` renders `SceneRainyCity`, which is the handcrafted `CityCafeScene` and
reads no asset at all, so `cd /tmp && ./gottage --demo` could never observe this
phase's bug. `main.go` gains `--scene <name>`, resolved by a new
`art.SceneByName`.

That lookup goes to `SceneRegistry` directly and **not** through `GetScene`,
which answers an unregistered type with `SceneRainyCity` — the fallback that let
two built-in presets point at cut scenes. An unknown name exits 2 and lists what
is available; an ambiguous substring names its candidates rather than picking
one. A unique substring is accepted because every display name carries an emoji.

**Exit codes.** An unrecognised *or empty* `--scene`, `--weather` or `--time`
value exits **2** with a diagnostic, on the interactive path as well as under
`--demo`; the two `switch` statements that used to drop such a value silently are
gone. `--scene` without `--demo` is also a 2. Separately, `--demo` exits **1**
and reports to stderr when `art.AssetErrors()` is non-empty after the frame is
drawn, because a frame made of the navy fallback is not evidence that the binary
resolved its assets — see the Verification block below.

**This is a scope deviation**: `main.go` was not in the Touches line above.

---

## Do NOT touch in this phase

- The per-frame `LoadAnimatedGif` call or the unbounded cache — **Phase 6** owns both, via the `Activate(w, h)` lifecycle hook. Changing the call site here would be redone there.
- `Delays` (Phase 15). Scaling/target dimensions (Phase 4/15).
- Asset selection. Phase 1 owned curation.

---

## Verification

```bash
go build -o /tmp/gottage . && cd /tmp
env -u TERM_PROGRAM -u TERM -u KITTY_WINDOW_ID ./gottage --demo --scene "Kirokaze Cozy Room" | head -4
env -u TERM_PROGRAM -u TERM -u KITTY_WINDOW_ID ./gottage --demo --scene "Kirokaze Cozy Room" \
  | grep -o '38;2;[0-9]*;[0-9]*;[0-9]*' | sort -u | wc -l   # 1367; the fallback scores 1
./gottage --demo --scene "Kirokaze Cozy Room"; echo $?      # 0 - a non-zero exit means an asset failed
GOTTAGE_ASSETS=/tmp/nope ./gottage --demo --scene "Kirokaze Cozy Room"; echo $?   # 1, and it names the asset
./gottage --demo --scene nonsense; echo $?                  # 2, and it lists valid names
./gottage --weather hail; echo $?                           # 2 - no silent fallback to the default
ls -lh /tmp/gottage                                         # 9.3MB
go test ./... -count=1
grep -rn 'os.Open' pkg/art/ | grep -v _test                 # 0 direct asset opens
grep -rn '"\.\."' pkg/                                      # 0 - pkg/art is CWD-independent
```

**The `env -u` on the colour count is load-bearing, not tidiness.**
`CottageScene.Render` consults `art.DetectTerminalProtocol`, and from Ghostty,
Kitty or WezTerm it answers with `EncodeKittyGPU`'s base64 payload, which carries
no `38;2;R;G;B` sequences at all. Run the pipeline from one of those terminals
without clearing the variables and it prints **0**, on a binary that is working
perfectly. `art.TerminalProtocolEnv()` is the list to clear;
`main_test.go`'s `runGottageWithAssets` takes it from there rather than restating
it, so the same trap cannot reappear in the suite.

The `--demo` line without `--scene` is deliberately gone: it renders
`SceneRainyCity`, a handcrafted scene that reads no asset, and proves nothing
about asset resolution. See §2.5.

New tests:

- `TestEveryRegisteredAssetResolves` — iterate `SceneRegistry`, resolve every `gifPath` through `assets.Open`.
- `TestScenesRenderFromArbitraryCWD` — `t.Chdir(t.TempDir())`, render a GIF scene, assert distinct-color-count > 1. **This is the regression test for the phase**, and it is what makes every later phase's tests trustworthy.
- `TestAssetOverrideEnv` — `GOTTAGE_ASSETS` pointing at a temp dir takes precedence.
- `TestEmbeddedAssetsAreRealGifs` — every embedded asset has GIF magic bytes and a plausible size. Guards against LFS pointers, truncation, and `.gitignore` accidents.

## Exit criteria

- [x] Binary renders GIF scenes correctly from `/tmp` — 1367 distinct colors via `--demo --scene`, exit 0; a failed asset is exit 1 rather than a silent navy frame
- [x] Zero direct `os.Open` of assets outside `assets/`
- [x] `GOTTAGE_ASSETS` override works, and a bad override directory errors loudly
- [x] Asset load failures recorded rather than silently swallowed — `art.AssetErrors()`
- [x] `TestEmbeddedAssetsAreRealGifs` green — no pointer files, no truncated assets
- [x] `grep -rn '"\.\."' pkg/` → 0
- [x] `make check` green

## Re-validation checklist for Phase 3

> **Phase 2b now runs between this phase and Phase 3.** All three are **answered**
> below and carried into `plan-phase-2b.md`'s prerequisites (items 6-8), which is
> the checklist Phase 3 runs.

1. **Did embed change test runtime?** `pkg/art` went **0.55s → 0.85s**, four runs
   each side; `-short` is 0.57s and the whole suite ~2.3s. The cost is the CWD
   test decoding all ten GIFs at 73×51. No `testing.Short()` guards needed yet.
2. **Binary size acceptable?** **9.30 MB**, up from 7.46 MB. Yes.
3. **Did Phase 1 renumber any `SceneType`?** **No** — 360 constants,
   `SceneRainyCity` first through `SceneAstralZodiacChronosAetherPortalStellar`
   last. Phase 3 migrates rather than repairs.
