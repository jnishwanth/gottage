# AGENTS.md — working conventions for `gottage`

Read this before writing code. Then read `docs/plan/plan-phases.md`.

**Why this file exists.** 93% of this repo's 95 commits were written by AI agents
over three days with no written conventions. The result passed a green test suite
the entire time while: 20 of 36 handcrafted scenes never rendered, the ambient
audio engine never played a sound, all 34 presets silently failed to apply, and a
6.6MB binary accumulated across 39 commits until `.git` reached 88MB.

None of that was caused by a hard problem. It was caused by the absence of the
rules below. Every rule here names the specific failure it prevents — if a rule
looks like bureaucracy, the failure it cites is the argument.

---

## 1. The gate

```bash
make check     # gofmt (all tracked .go) + go vet + go mod tidy -diff + go test
```

Run it before every commit. **Never commit with `make fmt-check` failing.**

A `pre-commit` hook is committed in `hooks/`. Activate it once per clone:

```bash
make hooks     # sets core.hooksPath=hooks
```

The hook is the fast subset (~0.2s): `gofmt` on **staged** `.go` files, plus
`go vet` and `go mod tidy -diff`. It deliberately does not run the test suite —
a gate slow enough to be annoying is a gate that gets bypassed. (`pkg/art` used
to be 17s of a 21s suite; Phase 1's curation brought the whole suite to ~2s, but
the rule holds as the suite grows back.) `make check` is the complete gate, over
every tracked `.go` file; run it yourself.

One caveat worth knowing:

- `gofmt -l` reads the **working tree**, not the index. A partially-staged file is
  judged by its on-disk contents.

**The repo is `gofmt`-clean as of Phase 1's first commit.** Phase 1 §1.7 originally
scheduled `gofmt -w` as the phase's *last* commit, to avoid reformatting 1,745 lines
that curation was about to delete. That was reversed, for two reasons: six of the
sixteen unformatted files are files Phase 1 must edit, and the pre-commit hook rejects
staged unformatted `.go` files — so every functional commit in the phase would have
been blocked, or pushed through with `--no-verify`. And the "wasted diff" premise
dissolved when curation was scoped to keep the GIF set rather than cut it to ~40:
`scenes_modular.go` loses ~120 lines, not ~1,800.

`make lint` (golangci-lint) is **not** part of `make check`, so `check` still
works on a machine without the linter installed.

---

## 2. Commits

**Subject: conventional-commit prefix, imperative, under 72 characters.**

Existing history is the counter-example. Real subjects run to 280 characters and
announce intent rather than describe the diff:

> `gnhf 16: Implemented 4 new handcrafted cozy procedural bit-art scenes (Kyoto Machiya Garden & Rain Chains, Cyber Lofi Diner, ...) with rich living micro-animations and expanded curated presets, palettes, and tests.`

Most of those scenes did not render. The subject asserted an outcome nobody had
checked. Concretely, this style means `git log --oneline` is unreadable,
`git log --grep` finds nothing, and no changelog can be generated.

Put the *what* in the subject and the *why* in the body. Claims about behavior
belong in a test, not a commit message.

---

## 3. Testing

### Never assert `len(output) != 0` as a rendering check

This is the single most expensive mistake in this repo's history.

`TestHandcraftedProceduralScenes` rendered all 36 scenes across 4 times of day —
144 frames — and asserted only that the output string was non-empty. When a scene
fails to load its asset it falls back to `c.Clear(PicoNavy)`, so **the test was
asserting that a solid navy rectangle is not zero-length.** It passed through 8
consecutive commits while 20 scenes were silently broken.

Assert something a broken render cannot satisfy: a distinct-RGB count, an exact
golden, a specific pixel. "It produced bytes" is not a test. Phase 1 replaced it
with a distinct-color floor over each scene's *own* layers — `renderSceneLayers`
in `pkg/art/scene_test.go` records why the composited frame cannot be measured
instead, and that reasoning is the transferable part.

The same applies to any `len(x) == 0` guard standing in for a correctness check.

### Test hygiene

- **`t.TempDir()`**, never a bare filename. `pkg/features/presets_test.go:11`
  writes `test_presets.json` into the package source directory and cleans up only
  on a non-panicking exit.
- **`t.Setenv`**, never `os.Setenv` with a manual `defer` restore.
  `pkg/art/terminal_adapter_test.go:9-38` does the latter; it is not
  parallel-safe and leaks on panic.
- **`t.Run` subtests** over monoliths. `TestModelViewModals`
  (`pkg/ui/view_test.go:37-321`) is 285 lines with ~40 unrelated assertions
  against shared mutating state — a failure in the middle hides everything after.
- **Guard slow tests** with `testing.Short()`. `TestVerifyDownloadedGifs`
  (`pkg/art/gif_loader_test.go`) decoded every GIF in `assets/` into the unbounded
  loader cache on each run — 15.75s of a 21s suite. Phase 1 guarded it and cut what
  it decodes from 363 files to 10.

---

## 4. Registries and enums

**Anything that registers by key needs a uniqueness test.** Two live bugs, one
root cause — nothing in this repo ever checks for a duplicate key:

- `RegisterScene` (`pkg/art/scene_interface.go`) ended in an unconditional
  `SceneRegistry[s.ID()] = s`. Of 380 registrations, **20 IDs were claimed twice**
  and the second silently won. That is the 20 dead scenes. Phase 1 made it panic;
  `TestRegisterSceneRejectsDuplicateID` and `TestNoDuplicateSceneRegistrations`
  (`pkg/art/scene_test.go`) hold the line.
- `GetCommandList` (`pkg/ui/command_palette.go`) is an append-only slice with no
  key at all. `Shortcut: "[ K ]"` appears on two different entries, while the `"K"`
  case in `pkg/ui/model.go` binds only the first — so the palette advertises a
  shortcut that does something else. Phase 13 owns the fix.

Prefer **panic in `init()`** over log-and-skip. Log-and-skip converts a
silent-overwrite bug into a silent-skip bug; a panic fails the process and every
test binary immediately and unmissably.

**`SceneType` is append-only until Phase 3.** Saved presets serialize the raw
`iota` ints, so reordering or removing a constant mid-block silently corrupts
every existing `user_presets.json`. Add at the end; delete only from the end.

---

## 5. Assets

Before Phase 1's curation, `assets/` held 367 files totalling 31MB. §1.8's
`git filter-repo` purge stripped the pre-curation asset blobs from every commit
in this range, so no commit before Phase 1's re-encode contains any file under
`assets/`. **§7 is the bar every asset must clear** — read it before adding or
cutting one.

- **No asset over ~300KB, over 64 frames, or above 640×360** — see §7 for why all
  three caps are needed and why on-disk size alone is not enough. On-disk size does
  not predict decoded cost: a 515KB/431-frame GIF is 142MB resident. Re-encode
  rather than ship an offender; drop it only if it fails the aesthetic bar too.
- **Never commit build artifacts.** `make build` writes `./gottage`, which is
  gitignored.
- **Batch asset churn.** Replacing an asset adds a permanent blob to history —
  it is not recoverable space without a history rewrite. Do not iterate on
  binaries in-repo; get the file right outside, then commit once.

### Do not introduce Git LFS

The instinct is reasonable and the answer is still no, on four counts:

1. **The remote does not create a need for it.** A private `origin`
   (`git@github.com:jnishwanth/gottage.git`) now exists, so the earlier form of
   this point — "`git remote -v` is empty" — is spent. It was never the
   load-bearing one: LFS pays off when large objects would otherwise bloat every
   clone and every fetch, and at the size in (2) they do not.
2. **The scale does not justify it.** Post-curation the asset set is 1-3MB. LFS
   starts earning its keep in the hundreds of megabytes.
3. **It would silently break Phase 2's `//go:embed`.** This is the decisive one.
   An LFS pointer file is ~130 bytes of text. If the smudge filter has not run — a
   clone without `git lfs pull`, LFS not installed, `GIT_LFS_SKIP_SMUDGE` set —
   then `go:embed` embeds *that text* instead of the GIF. The build succeeds,
   `assets.Open` succeeds, and every scene renders flat navy. That is exactly the
   silent-failure class this codebase already shipped once. Phase 2 has not landed
   yet — there is no `//go:embed` in the tree today — so this is a forward-looking
   constraint, not a description of the current build.
4. **It adds a required tool dependency** to every clone and every agent session,
   in exchange for nothing above.

**Revisit only if all three become true:**

1. a remote exists — **already true**, see (1) above;
2. assets exceed ~50MB — **not** true, 1-3MB post-curation;
3. *and* Phase 2's embed strategy has been replaced by runtime asset loading —
   **not** true. Today's CWD-relative `os.Open` loader is the *pre*-Phase-2 state
   that embedding is scheduled to replace, not a replacement for it. Do not read
   this condition as satisfied merely because no `//go:embed` exists yet.

What actually controls asset bloat here is the §7 caps and the churn discipline
above, not LFS.

---

## 6. The phase plan

`docs/plan/` holds a 19-phase improvement plan, numbered 0-17 plus **2b**.
**`docs/plan/plan-phases.md` is the index — read it first.**
`docs/plan/plan-phase-N.md` is the source of truth for what phase N does, and
every phase doc has a **"Do NOT touch"** section that exists to keep phases from
colliding. Respect it. `plan-phase-2b.md` is a deliberate exception to that
filename pattern — 2b was inserted between Phases 2 and 3 rather than renumbering
ten downstream docs — so treat it as a phase doc like any other.

The ordering rule the plan follows:

> A phase may replace a structure. It may not fix bugs inside a structure that a
> later phase replaces.

So some known bugs are deliberately left alone — unbound keybindings until Phase
13 regenerates them from one table, modal precedence until Phase 12 makes the
divergence inexpressible. "This is obviously broken, let me just fix it" is how
work gets done twice.

**Line anchors drift.** Re-derive them with `grep -n` rather than trusting numbers
cited in a doc, especially in `pkg/art/scenes_modular.go` (Phase 1 took it 7,909 →
6,035 lines, so every anchor into it in a Phase 4+ doc is wrong) and `pkg/ui/view.go`
(Phases 12-14 move it repeatedly).

---

## 7. Asset curation bar

An asset earns its place in `assets/animations/` only if it clears **both** bars
below. This is also the criterion Phase 8's reference frames get judged against, so
it is written down before any curation happens rather than reverse-engineered after.

### Aesthetic

- **Reads at 74×52 half-blocks.** Silhouette and value structure survive the
  downsample; fine line-art and small text do not. Check the actual rendered frame,
  not the source GIF in an image viewer.
- **A coherent cozy interior or view** — a room, a window, a street at night. Not an
  abstract loop, a logo, a character sheet, or a demo reel.
- **Visible micro-animation.** Rain, steam, flicker, drift, a passing train. A frame
  that only cross-fades is a still image with extra bytes.
- **A palette compatible with the app's warm tokens.** Phase 7 defines them; until
  then, judge against the existing handcrafted scenes.
- **Somewhere a person would want to sit.** This is the whole product thesis and it
  is the criterion that actually does the cutting.

### Hard caps

Mechanically checkable, all three required:

| Cap | Value | Why |
|---|---|---|
| On-disk size | **≤ 300 KB** | Phase 2 embeds every asset into the binary. |
| Frame count | **≤ 64** | Decoded cost is `frames × W × H × 24` bytes (`colorful.Color` is 3×float64) and the loader cache is unbounded until Phase 6. |
| Source dimensions | **≤ 640×360** | The renderer downsamples to canvas size. Anything larger is wasted bytes and a slower decode. |

`TestAssetsRespectHardCaps` (`pkg/art/asset_caps_test.go`) enforces all three
against every file in `assets/animations`, so an offending asset fails `make check`
rather than shipping.

**The frame cap is not redundant with the size cap.** Measured on the pre-curation set,
`waneella_street.gif` was 515 KB on disk — comfortably under 300 KB after a downscale —
but 431 frames, which is **142 MB resident** once decoded at 160×90. `cyber_bot.gif` was
152 KB and 221 frames (73 MB). On-disk size does not predict memory cost; only frame
count does.

**An oversized asset that clears the aesthetic bar gets re-encoded, not dropped.** A
large GIF is large because of source resolution or frame count, and neither survives the
downsample to 74×52. `ffmpeg` with `fps=…,scale=…,palettegen/paletteuse` handles both.
Prefer the `fps` filter over frame-dropping: it re-times the result, so the animation
still plays at its authored wall-clock speed once Phase 15 starts honoring `Delays`.

Re-encode **outside the repo and commit once** — see §5. Every intermediate version you
commit is a permanent blob.
