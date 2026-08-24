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
`pkg/art` alone takes 17s, and a gate slow enough to be annoying is a gate that
gets bypassed. `make check` is the complete gate, over every tracked `.go`
file; run it yourself.

Two caveats worth knowing:

- `gofmt -l` reads the **working tree**, not the index. A partially-staged file is
  judged by its on-disk contents.
- `make check` is expected to **fail on `gofmt` until Phase 1 §1.7**, which runs
  `gofmt -w` as its final commit. Reformatting 1,745 lines that Phase 1 is about
  to delete would be wasted diff. Do not "fix" this early.

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

`TestHandcraftedProceduralScenes` renders all 36 scenes across 4 times of day —
144 frames — and asserts only that the output string is non-empty. When a scene
fails to load its asset it falls back to `c.Clear(PicoNavy)`, so **the test was
asserting that a solid navy rectangle is not zero-length.** It passed through 8
consecutive commits while 20 scenes were silently broken.

Assert something a broken render cannot satisfy: a distinct-RGB count, an exact
golden, a specific pixel. "It produced bytes" is not a test.

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
  (`pkg/art/gif_loader_test.go:41`) decoded every GIF in `assets/` on each run.
  Since the §1.8 purge its `../../assets/animations` path no longer resolves, so
  it now fails outright — the known cause of the red suite at this commit.

---

## 4. Registries and enums

**Anything that registers by key needs a uniqueness test.** Two live bugs, one
root cause — nothing in this repo ever checks for a duplicate key:

- `RegisterScene` (`pkg/art/scene_interface.go:30`) ends in an unconditional
  `SceneRegistry[s.ID()] = s`. Of 380 registrations, **20 IDs are claimed twice**;
  the second silently wins. That is the 20 dead scenes. Phase 1 makes it panic.
- `GetCommandList` (`pkg/ui/command_palette.go`) is an append-only slice with no
  key at all. `Shortcut: "[ K ]"` appears at both `:78` and `:510`, while
  `pkg/ui/model.go:966` binds `K` to only the first — so the palette advertises a
  shortcut that does something else. Phase 13 owns the fix.

Prefer **panic in `init()`** over log-and-skip. Log-and-skip converts a
silent-overwrite bug into a silent-skip bug; a panic fails the process and every
test binary immediately and unmissably.

**`SceneType` is append-only until Phase 3.** Saved presets serialize the raw
`iota` ints, so reordering or removing a constant mid-block silently corrupts
every existing `user_presets.json`. Add at the end; delete only from the end.

---

## 5. Assets

Before Phase 1's curation, `assets/` held 367 files totalling 31MB; Phase 1
curates that set to ~40 GIFs (`plan-phase-1.md` §1.3). §1.8's `git filter-repo`
purge stripped the pre-curation asset blobs from every commit in this range, so
no commit here contains any file under `assets/`.

- **No asset over ~300KB** without discussion. That distribution was heavily
  skewed — mean 76KB, but the top 5 files alone were 9MB. A large GIF is
  usually large because of frame count or palette width, neither of which
  survives downsampling to 74×52 half-blocks. Re-encode rather than ship it.
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

What actually controls asset bloat here is the size cap and the churn discipline
above, not LFS.

---

## 6. The phase plan

`docs/plan/` holds an 18-phase improvement plan. **`docs/plan/plan-phases.md` is
the index — read it first.** `docs/plan/plan-phase-N.md` is the source of truth
for what phase N does, and every phase doc has a **"Do NOT touch"** section that
exists to keep phases from colliding. Respect it.

The ordering rule the plan follows:

> A phase may replace a structure. It may not fix bugs inside a structure that a
> later phase replaces.

So some known bugs are deliberately left alone — unbound keybindings until Phase
13 regenerates them from one table, modal precedence until Phase 12 makes the
divergence inexpressible. "This is obviously broken, let me just fix it" is how
work gets done twice.

**Line anchors drift.** Re-derive them with `grep -n` rather than trusting numbers
cited in a doc, especially in `pkg/art/scenes_modular.go` (Phase 1 halves it) and
`pkg/ui/view.go` (Phases 12-14 move it repeatedly).

---

## 7. Asset curation bar

<!-- PLACEHOLDER — Phase 1 §1.3 fills this in as its first act, before cutting
     any asset. It is the criterion Phase 8's reference frames get judged
     against, so it has to be written down before curation starts, not after.

     Suggested starting point from plan-phase-1.md §1.3: reads clearly at 74×52
     half-blocks; a coherent cozy interior or view rather than an abstract loop;
     visible micro-animation; a palette compatible with the app's warm tokens;
     under ~300KB. -->

_Not yet written — see Phase 1 §1.3._
