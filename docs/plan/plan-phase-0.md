# Phase 0 — Repo & Toolchain Foundation

> **Session size:** Small · **Depends on:** nothing · **Touches:** no application code

> **Status: completed.** Corrections found while executing this phase, recorded
> so §0.8 below is not read as current:
>
> - **§0.8's branch table was stale.** `just-making-the-repo-better-and-less-of-a-mess`
>   did not exist and there was no second worktree. The actual third branch was
>   `gnhf/plan-research-and-fi-75ef78` at `08be2ed`, 10 commits *behind* HEAD and
>   0 ahead. It was deleted with `git branch -d`, which machine-confirmed it held
>   no unique commits. `main` was fast-forwarded to HEAD.
> - **§0.6's errcheck estimate was low.** 47 actionable, not 42 (19 in
>   `ambient.go` as stated). More importantly, `errcheck` at defaults reports
>   **122**, because 111 are `strings.Builder`/`bytes.Buffer` writes whose error
>   is always nil — `.golangci.yml` excludes those by function.
> - **§0.7:** `golang.org/x/image` is imported by *two* files —
>   `gif_loader.go:11` and `image_loader.go:13`, the latter deleted by Phase 1.
> - **§0.5's hook was changed to gate staged files only.** A hook running
>   repo-wide `make check` would have blocked every Phase 0 and Phase 1 commit
>   until §1.7, training the `--no-verify` habit. `make check` stays repo-wide.
> - **The Goal's `.git` shrink claim was wrong for *this* phase.** Phase 0 does
>   not shrink `.git`; it stays at 88MB, deliberately (§0.1). The purge belongs to
>   Phase 1 §1.8, which rewrites binary *and* dropped-asset blobs in one
>   `git-filter-repo` pass. The Goal below is scoped accordingly.
> - **§0.5's fmt gate was narrowed to tracked files.** `fmt-check` runs `gofmt`
>   over `git ls-files '*.go'`, not `gofmt -l .`, so it no longer walks
>   `.gnhf/` or other untracked scratch. §0.3's and §0.5's `gofmt -l .` wording
>   below predates that; `make fmt-check` is the current command.
> - Phase-1 handoff notes were written into `plan-phase-1.md`.

---

## Goal

Make the repository a place where mistakes fail loudly, before changing a single line of app logic. Establish the conventions file, the license, and the local build/lint/test gate, and untrack the committed binary so that **Phase 1 §1.8** can shrink `.git` from 88MB to ~35MB in a single history rewrite. This phase leaves `.git` at 88MB — see §0.1.

## Why this phase is first

Every later phase commits code. This phase decides what "a good commit" means and installs the gate that enforces it. It touches no `.go` file bodies, so it cannot conflict with anything and nothing later can obsolete it.

The history purge specifically **must** be first: it rewrites all 95 SHAs. Doing it later would invalidate every commit reference written into phase docs 1–17.

### The context you need

95 commits, 93% of them AI-agent `gnhf N:` commits, produced over 3 days with zero written conventions. The direct, mechanical consequences visible today:

- `gofmt -l .` reports **16 of 34 Go files** unformatted (1,745 diff lines). Nobody has ever run it.
- A **7MB compiled binary is tracked**, touched by **39 separate commits** → `.git` is 88MB for a 17k-line project.
- README displays an MIT badge and claims "MIT License © 2026". **There is no LICENSE file.**
- Eight consecutive commits shipped scenes that never render (Phase 1), and every one of them claimed "100% passing tests" — because the tests asserted `len(rendered) != 0`.

A conventions file plus a gate is what stops commit #96 from repeating this.

---

## Workstreams

### 0.1 — Backup, and stop the bleeding

**The history rewrite itself moved to Phase 1** — see the note below. This phase takes the backup and untracks the binary so it stops accumulating.

```bash
# 1. Backup mirror, kept outside the repo. Do this before ANY of the 18 phases.
git clone --mirror /Users/nish/proj/gottage /Users/nish/proj/gottage-backup.git
git -C /Users/nish/proj/gottage-backup.git log --oneline | wc -l    # -> 95

# 2. Confirm the damage
git ls-files | grep -x gottage              # -> gottage
git log --oneline -- gottage | wc -l        # -> 39 commits carrying a ~7MB blob
du -sh .git                                 # -> 88M

# 3. Untrack it (with .gitignore from 0.2, it stops coming back)
git rm --cached gottage
```

`.git` stays at 88MB until Phase 1. That is deliberate.

> **Why the purge waits for Phase 1.**
> Confirmed with the user: the binary **and** the assets Phase 1 drops (~304 GIFs + 4 JPGs, ~25MB) get purged together in **one** `git-filter-repo` pass, at the end of Phase 1 — see `plan-phase-1.md` §1.8.
>
> It has to be that way round: you cannot purge assets from history before curation decides which ones survive. Splitting it into two rewrites would mean rewriting all SHAs twice for no benefit.
>
> Consequence to be aware of: **Phase 1 rewrites this phase's commit SHAs too.** That is harmless — no phase doc references a commit SHA, only `file:line` anchors, which the rewrite does not touch. But do not build anything on a Phase 0 SHA in the meantime.

### 0.2 — `.gitignore`

There is no `.gitignore` anywhere. The only ignore mechanism is a hand-edited `.git/info/exclude` containing `.gnhf/runs/`, which is **not tracked and not shared** — any reclone picks `.gnhf/` back up (113 files, 2.3MB).

Create `.gitignore` covering at minimum: the `gottage` binary, `.gnhf/`, `user_presets.json` (Phase 3 relocates it, but it litters CWD until then), `*.test`, `coverage.out`, `dist/`, editor dirs.

### 0.2b — `.gitattributes`, and the Git LFS decision

There is **no `.gitattributes` anywhere** in the repo. Add one:

```gitattributes
*.gif binary
*.jpg binary
*.png binary
*.wav binary
*.go  text eol=lf
```

Marking the binaries stops git attempting diffs/merges on them and prevents any CRLF mangling. Cheap, and it should have existed from the start.

#### Should the assets be Git LFS objects? **No — and the reason is worth recording.**

The instinct is right (27MB of binary GIFs in a 17k-line repo), but LFS is the wrong tool *here*, on four counts:

1. **There is no remote.** `git remote -v` is empty; the repo has never been pushed. LFS's entire value is that large objects live on a separate server and are fetched on demand. With no remote, LFS objects sit in `.git/lfs/objects/` — **the same bytes on the same disk, plus a layer of indirection.** It buys literally nothing today.
2. **The scale does not justify it.** Assets are 27MB now, but Phase 1 curates to ~40 GIFs under a size cap. The distribution is heavily skewed — mean 76KB, but 12 files over 500KB and the top 5 alone accounting for 9MB — so a size-aware curation lands the whole set at **1-3MB**. LFS starts earning its keep in the hundreds of MB.
3. **It would silently break Phase 2's `//go:embed`.** This is the decisive one — see `plan-phase-2.md`, which documents it in full. Short version: an LFS pointer file is ~130 bytes of text, and `go:embed` would embed *that* rather than the GIF. The build succeeds and every scene renders garbage.
4. **It adds a required tool dependency** for every clone and every future agent session, in exchange for nothing.

**Revisit only if all of these become true:** a remote exists, assets exceed ~50MB, and Phase 2's embed strategy has been replaced by runtime asset loading. Note that trio in `AGENTS.md` so a future session does not "helpfully" add LFS and break the build.

What actually controls asset bloat here is not LFS — it is **Phase 1's size cap** and the churn discipline in `AGENTS.md` (§0.3).

### 0.3 — `AGENTS.md`

**This is the highest-leverage artifact in the phase.** Given 93% agent authorship, written conventions are process control, not documentation polish. It must state at minimum:

- Run `make check` before every commit. Never commit with `gofmt -l .` non-empty.
- Commit subjects: conventional-commit style, imperative, under 72 chars. (Current subjects are up-to-280-char feature-announcement paragraphs describing *intent*, not the diff — `git log --grep` is useless and no changelog can be generated.)
- **Never assert `len(output) != 0` as a rendering test.** State why: that exact assertion let 20 broken scenes ship across 8 commits.
- Never commit build artifacts. Never add an asset over ~300KB without discussion — see the size cap in `plan-phase-1.md` §1.3.
- **Do not introduce Git LFS.** It would silently break `//go:embed` (see §0.2b). The revisit conditions are listed there.
- Replacing an asset adds a permanent blob to history. Batch asset churn; do not iterate on binaries in-repo.
- Registry/enum additions must come with a uniqueness test.
- Where the phase docs live and that `plan-phase-N.md` is the source of truth for sequencing.

### 0.4 — `LICENSE`

MIT, matching the README's existing claim. Also fix the README's broken badge host `img.shields.golang.org` (not a real host) while here — it is a one-line change and Phase 17 does the full rewrite.

### 0.5 — `Makefile` + pre-commit hook

The user chose **personal tool, no CI**. So the gate is local:

```make
check: fmt-check vet test
fmt-check:  # fails if gofmt -l . is non-empty
vet:        # go vet ./...
test:       # go test ./... -count=1
test-race:  # go test -race ./...   (note: pkg/art is 17s -> 222s under -race)
lint:       # golangci-lint run
build:      # go build -o gottage .
run:        # go run .
```

Install a `.git/hooks/pre-commit` running `make check`, and commit a `hooks/` copy so it survives reclone.

### 0.6 — `.golangci.yml`, deliberately narrow

Enable **only** `gofmt`, `govet`, `errcheck`, `ineffassign`.

Do **not** enable `funlen` or `gocyclo` yet — they will fail loudly against `Update()` (794 lines, `pkg/ui/model.go:236-1029`) and `init()` (2,117 lines, `pkg/art/scenes_modular.go:9`), neither of which is fixable until Phases 13 and 1 respectively. A lint config that cannot pass gets disabled and then ignored. **Phase 16 widens it.**

Note `errcheck` will flag 42 ignored errors repo-wide (19 in `pkg/audio/ambient.go` alone). Use a baseline/`//nolint` with a `TODO(phase-N)` reference rather than fixing them here — the files that own most of them get rewritten in Phases 3, 11, and 16.

### 0.7 — `go mod tidy`

`golang.org/x/image v0.44.0` is marked `// indirect` in `go.mod` but is **directly imported** by `pkg/art/gif_loader.go`. Run `go mod tidy` and add `go mod tidy -diff` to `make check`.

Leave the 17 outdated dependencies alone — none is security-relevant, and a dependency bump in the same phase as a history rewrite makes bisection miserable.

### 0.8 — Branch reconciliation

```
main                                        2a747b9   (7 commits)
gnhf/make-the-tui-design-d3afb2             65ba4e3   (HEAD, +88)
just-making-the-repo-better-and-less-of-a-mess  2a747b9   (stale, in a worktree)
```

`git rev-list --left-right --count main...HEAD` → `0  88`. Clean fast-forward, no divergence. Fast-forward `main` to HEAD and delete the stale third branch, or the next 17 phases stack on a tree whose default branch is 88 commits stale. There is no remote, so this is purely local bookkeeping.

---

## Do NOT touch in this phase

- Any `.go` file body. `go mod tidy` may edit `go.mod`/`go.sum`; nothing else.
- **Do not run `gofmt -w`.** It belongs at the *end of Phase 1*, after ~3,500 lines are deleted — reformatting lines you are about to delete is wasted diff and guarantees conflicts.
- Do not delete assets. Phase 1 owns curation.
- Do not add `funlen`/`gocyclo`.

---

## Verification

```bash
git ls-files | grep -c '^gottage$'              # 0  (untracked)
du -sh .git                                     # still ~88M — Phase 1 purges
test -f /Users/nish/proj/gottage-backup.git/HEAD   # backup exists
test -f LICENSE && test -f AGENTS.md && test -f .gitignore && test -f .gitattributes && test -f Makefile
go mod tidy -diff                               # exit 0
go build ./... && go vet ./...                  # clean
go test ./... -count=1                          # passes (unchanged from baseline)
git rev-list --left-right --count main...HEAD   # 0  0
make check                                      # fails ONLY on gofmt (expected; Phase 1 fixes)
```

`make check` failing on `gofmt` at the end of this phase is the **correct** outcome — it proves the gate works and hands Phase 1 a concrete task.

## Exit criteria

Ticked only where the outcome was machine-checked, per §2 of `AGENTS.md`.

- [x] `gottage` untracked and gitignored (history purge is Phase 1 §1.8) — `git ls-files | grep -x gottage` is empty; `/gottage` is in `.gitignore`
- [x] `.gitignore`, `.gitattributes`, `LICENSE`, `AGENTS.md`, `Makefile`, `.golangci.yml`, `hooks/pre-commit` all committed — all seven resolve under `git ls-files`
- [x] LFS decision recorded in `AGENTS.md` with its revisit conditions
- [x] `go mod tidy -diff` exits 0
- [x] Test suite unchanged in count — 22 test functions, counted
- [ ] Backup mirror exists outside the repo, and its `git log --oneline | wc -l` is 95 — **unverified here**: the mirror is outside the repo, so it cannot be checked from a clone. Confirm on the machine that took it.
- [ ] `main` fast-forwarded to HEAD; stale branch deleted — carried out and recorded in the correction block at the top of this file, but **not re-verifiable** from an arbitrary checkout.
- [ ] Test suite still green — **not true on this tree.** Phase 1 §1.8's history rewrite retroactively removed the GIFs that `pkg/art` loads, so this commit's tree has no `assets/`. The suite is exercised against a tree that has them in Phase 1.

`make check` is likewise expected to fail here, on `fmt-check` only — see the note just above **Exit criteria**.

## Re-validation checklist for Phase 1

Before trusting `plan-phase-1.md`, confirm:

1. **Line anchors moved?** This phase should not have shifted any `.go` line numbers. Verify one anchor: `sed -n '29,34p' pkg/art/scene_interface.go` still shows `RegisterScene`.
2. **Collision count still 20?** Re-run the extraction in Phase 1 §1.1 — it should still report 20. If the user hand-edited scenes in between, the list changes.
3. **Backup verified?** Phase 1 ends with an irreversible rewrite of all 95 commits. Confirm the mirror clone exists and is complete *before* starting Phase 1, not during it.
4. **Is `errcheck` noise blocking `make check`?** If the baseline approach was not taken, Phase 1 inherits 42 findings — decide whether to fix or baseline before starting.
