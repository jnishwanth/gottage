# Phase 17 — Docs & Final Polish

> **Session size:** Small · **Depends on:** Phase 16 · **Touches:** `README.md`, `AGENTS.md`, `CONTRIBUTING.md`

---

## Goal

Make the documentation true, and make it *stay* true by generating what can be generated and testing the rest.

## Why this is last

It documents the final state. Written earlier, it would describe something that no longer exists — which is precisely the failure mode below.

---

## How wrong the README currently is

Nearly every specific claim in it is false:

| README claims | Reality |
|---|---|
| **"4 Science-Backed Research Scenes"** cycled with `v` | **46 registered scenes** since Phase 1's curation — 36 handcrafted + 10 GIFs. (All 360 `SceneType` constants remain, deliberately unpruned; 314 are simply unregistered.) The README *understates* by ~11×, while the app's own strings *overstate* by ~8×: eight surviving `370+` literals across `model.go`, `command_palette.go` and `view.go` (`"Type keyword to filter 370+ scenes..."`, `"All 370+ Scenes 🎨"`, …). Fix both against `len(art.SceneOrder)`, not against each other |
| Architecture diagram lists 10 files | 34 Go files. **Omits `scenes_modular.go` — 45% of all code** — plus `scenes.go`, `palette.go`, `gif_loader.go`, `command_palette.go`, `presets.go`, and more |
| Ambient sounds: "Rain, Fireplace Crackle, Wind" | Five channels — and **none of them made sound** (Phase 11) |
| Keybinding table | Missing ~25 bindings that exist |
| "Go 1.20+" | `go.mod` says `go 1.26.5` |
| "MIT License © 2026" + MIT badge | No LICENSE file existed (Phase 0) |
| Badge `img.shields.golang.org` | Was not a real host; corrected to `img.shields.io` (Phase 0) |
| `https://github.com/cliamp` | Bare user URL, not a repo |
| `go build -o gottage main.go` | The documented command that regenerated the committed 7MB binary in place |

The in-app help was independently wrong too (Phase 13 fixed it by generating it).

**The lesson to encode:** documentation that is hand-maintained alongside a table drifts. Generate what can be generated; test the rest.

---

## Workstreams

### 17.1 — Rewrite the README

Accurate, and structured so drift is detectable:

- What it is; a screenshot or recorded demo (this is a visual app — Phases 8-10 produced real art, show it)
- Install and run, including `GOTTAGE_ASSETS` (Phase 2)
- Feature list matching what exists
- **Keybindings generated from Phase 13's command table**, not hand-written
- Scene count as a **tested number**
- Architecture reflecting the post-Phase-10 layout
- Config file location and schema (Phase 3)
- Known limitations — explicitly, including anything Phase 16 deferred
- Fix the cliamp link

### 17.2 — Make docs self-checking

- `TestSceneCountMatchesReadme` — parse the number out of the README, compare to `len(SceneOrder)`.
- `TestKeybindingsDocMatchesTable` — or better, **generate** the keybinding section via `go generate` and have `make check` fail on drift.
- `TestArchitectureDocListsAllPackages`.

A doc that can go stale silently will. These are cheap.

### 17.3 — Finish `AGENTS.md`

Phase 0 created it as process control. Update it with everything learned:

- The migration pattern for adding a scene (Phase 9/10)
- Tokens only, no inline hex (Phase 7)
- Components before hand-rolled art (Phase 8)
- Golden regeneration is a **deliberate, separately-committed, visually-reviewed** act
- Never assert `len(output) != 0` as a rendering test — **and say why**: that assertion let 20 broken scenes ship across 8 commits
- Never introduce a UI toggle without a test that the engine reads it — the `FairyLightsOn` / `SmoothDayNightBlend` / `SpeedFast` / silent-ambient class
- The lint set and why it widened

### 17.4 — `CONTRIBUTING.md`

Short: build, test, `make check`, commit conventions, where phase docs live.

### 17.5 — Retire the backup mirror

No history work is needed here. **Phase 1 §1.8 already purged the binary and the dropped assets in one `git-filter-repo` pass**, taking `.git` from 88MB to ~8MB.

What remains is bookkeeping:

```bash
du -sh .git                                    # ~8M, confirming Phase 1 held
git count-objects -vH | grep size-pack
```

The backup mirror at `/Users/nish/proj/gottage-backup.git` has been the only copy of the pre-rewrite history since Phase 1. Sixteen phases later, if nothing has needed it, delete it deliberately — or keep it if the `gnhf N:` provenance is worth preserving as an archive. Either is fine; **decide, rather than leaving it lying around unlabelled.**

### 17.6 — Final walkthrough

Run the app for a while. Every tab, every modal, several terminal sizes, several scenes, audio on.

Ask the question the whole plan was aimed at: **is this cozy, and would you actually leave it on screen while working?** If not, the honest answer is another art iteration — not a shipped README claiming otherwise.

---

## Do NOT touch in this phase

- Code, beyond doc-generation plumbing and anything the walkthrough proves broken.
- Features.

---

## Verification

```bash
go test ./... -count=1
make check
go generate ./... && git diff --exit-code          # generated docs are current

# every README claim, checked
grep -c 'img.shields.golang.org' README.md         # 0
grep -o '[0-9]\+ scenes' README.md                 # matches len(SceneOrder)
du -sh .git                                        # ~8M  (purged back in Phase 1)

# runs clean from anywhere
go build -o /tmp/gottage . && cd /tmp && ./gottage
```

## Exit criteria

- [ ] README accurate; every claim verified or tested
- [ ] Keybindings and scene count generated or test-guarded
- [ ] `AGENTS.md` carries the lessons; `CONTRIBUTING.md` exists
- [ ] Backup mirror deleted or deliberately archived
- [ ] Full manual walkthrough passed
- [ ] `make check` green

---

## Closing the loop

Worth recording at the end, because it is the point of the whole exercise. The starting state was:

- **20 of 36 handcrafted scenes never rendered** — 8 commits of work, invisible
- **Ambient audio never played a sound** — the engine was never called
- **All 34 built-in presets silently failed** to change the theme
- **The app was unusable on an 80×24 terminal**
- **Browsing scenes could consume ~1GB of RAM**
- **The renderer allocated ~64,000 strings per second**
- **The README understated the scene count by ~90×**

Every one of those passed a green test suite, because the tests asserted that output was non-empty.

The durable fix was never any individual patch. It was: **`RegisterScene` panics. Goldens are byte-exact. Help is generated. Layout is tested at every size. Errors surface.** Make the machine fail loudly, and the bugs stop being invisible.
