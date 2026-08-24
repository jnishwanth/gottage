# Phase 2 — Asset Pipeline & Embed

> **Session size:** Small · **Depends on:** Phase 1 · **Touches:** `pkg/art/gif_loader.go`, a new `pkg/assets/`, `assets/`

---

## Goal

Make the binary run correctly from **any** working directory by embedding the curated assets, and route every asset read through one resolver.

## Why this phase is here

It depends on Phase 1 (embed only survivors — 6MB, not 31MB) and blocks nothing structural, so it slots in early and cheaply. Every later phase's tests benefit: today, running `go test ./pkg/art` from the package directory silently exercises the GIF *failure* path.

## The bug

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

### 2.1 — `pkg/assets` with an embedded FS

```go
package assets

//go:embed animations/*.gif
var embedded embed.FS

// Open resolves an asset by its registration path, e.g. "assets/animations/x.gif".
func Open(name string) (fs.File, error)
```

Resolution order:

1. `$GOTTAGE_ASSETS` if set — an override directory. Keeps iteration fast during Phases 8-10 (drop in a GIF without rebuilding) and gives a supported escape hatch.
2. The embedded FS.

Do **not** fall back to CWD. Silent CWD resolution is the bug being fixed; an explicit env override replaces it.

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

If it landed materially above ~10MB of assets, stop and reconsider a runtime data directory before committing to embed — but note that going the other way (runtime dir) reopens the CWD-dependency bug this phase exists to close, so prefer tightening curation over abandoning embed.

---

## Do NOT touch in this phase

- The per-frame `LoadAnimatedGif` call or the unbounded cache — **Phase 6** owns both, via the `Activate(w, h)` lifecycle hook. Changing the call site here would be redone there.
- `Delays` (Phase 15). Scaling/target dimensions (Phase 4/15).
- Asset selection. Phase 1 owned curation.

---

## Verification

```bash
go build -o /tmp/gottage . && cd /tmp && ./gottage --demo | head -5   # real art, not solid navy
ls -lh /tmp/gottage                                                   # ~13MB
go test ./... -count=1
grep -rn 'os.Open' pkg/art/ | grep -v _test    # 0 direct asset opens
```

New tests:

- `TestEveryRegisteredAssetResolves` — iterate `SceneRegistry`, resolve every `gifPath` through `assets.Open`.
- `TestScenesRenderFromArbitraryCWD` — `t.Chdir(t.TempDir())`, render a GIF scene, assert distinct-color-count > 1. **This is the regression test for the phase**, and it is what makes every later phase's tests trustworthy.
- `TestAssetOverrideEnv` — `GOTTAGE_ASSETS` pointing at a temp dir takes precedence.
- `TestEmbeddedAssetsAreRealGifs` — every embedded asset has GIF magic bytes and a plausible size. Guards against LFS pointers, truncation, and `.gitignore` accidents.

## Exit criteria

- [ ] Binary renders GIF scenes correctly from `/tmp`
- [ ] Zero direct `os.Open` of assets outside `pkg/assets`
- [ ] `GOTTAGE_ASSETS` override works
- [ ] Asset load failures recorded rather than silently swallowed
- [ ] `TestEmbeddedAssetsAreRealGifs` green — no pointer files, no truncated assets
- [ ] `make check` green

## Re-validation checklist for Phase 3

1. **Did embed change test runtime?** `pkg/art` was 17s (222s under `-race`). If embedding pushed it materially higher, Phase 3 should add the `testing.Short()` guards before adding more tests.
2. **Binary size acceptable?** If not, revisit before Phase 9/10 add more assets.
3. Phase 3 is independent of this phase's code — it only needs Phase 1's enum decision. Confirm whether Phase 1 renumbered any `SceneType`.
