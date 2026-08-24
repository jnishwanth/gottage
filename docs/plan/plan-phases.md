# Gottage Improvement Plan — Phase Index

18 phases, each sized for **one fresh session**. Read this file first, then `plan-phase-N.md`.

## The ordering rule

> **A phase may replace a structure. It may not fix bugs inside a structure that a later phase replaces.**

Everything is ordered most-fundamental-first, so no phase's work is made obsolete by a later one. Three places where this inverted the obvious ordering:

- **Keybindings.** `[t]`/`[s]` unbound, `[K]` mislabeled — known since the audit, deliberately **not** fixed until Phase 13, because Phase 13's command table regenerates all of it from one source.
- **Modal precedence.** Update and View disagree today. Not patched — Phase 12's focus stack makes the divergence inexpressible.
- **Palette.** Tokens (7) before components (8) before scene migration (9/10), so components are never built on ad-hoc colors and then re-touched.

## How to run a phase

1. Read `plan-phase-N.md` fully, including **"Do NOT touch"** — those boundaries are what keep phases from colliding.
2. Run its **prerequisites / re-validation checklist** (written at the end of phase N-1). Line anchors drift; re-derive rather than trusting cited numbers, especially in `pkg/art/scenes_modular.go` (Phase 1 halves it) and `pkg/ui/view.go` (Phases 12-14 move it repeatedly).
3. Do the work. Commit in reviewable slices — **golden regeneration always gets its own labelled commit.**
4. Run the phase's verification block. Every exit criterion is mechanically checkable.
5. Write the re-validation checklist findings into the next phase doc if anything changed.

## The phases

| # | Phase | Size | Needs | Headline outcome |
|---|---|---|---|---|
| 0 | Repo & Toolchain Foundation | S | — | Backup, untrack binary; `AGENTS.md`, LICENSE, `make check` |
| 1 | Content Triage | L | 0 | **20 dead scenes start rendering**; assets 31MB → ~6MB; **`.git` 88MB → ~8MB** |
| 2 | Asset Pipeline & Embed | S | 1 | Runs from any directory |
| 3 | Config & Persistence | M | 2 | The one on-disk format break; state survives restart |
| 4 | Canvas Core | L | 0 | ~3,800 → <10 allocs/frame; `RenderTo` |
| 5 | Determinism & Goldens | M | 4 | Rendering becomes testable at all |
| 6 | `SceneInterface` v2 | L | 5 | Lifecycle hook; **~1GB cache bounded** |
| 7 | Palette → Tokens | M | 6 | 0 inline hex; 0 per-frame parsing |
| 8 | Reference + Components | M | 7 | **Signed-off art direction**; shared vocabulary |
| 9 | Scene Migration A | L | 8 | ~18 scenes rebuilt, resolution-independent |
| 10 | Scene Migration B + split | L | 9 | Remainder; `scenes_modular.go` gone |
| 11 | Audio on `oto` | M | 3 | **Ambient audio actually plays** |
| 12 | `Model` + Focus Stack | L | 11 | 40 fields decomposed; precedence bug impossible |
| 13 | Command Table | L | 12 | 794-line `Update()` → lookup; help generated |
| 14 | Responsive Layout | M | 13, 4 | **80×24 works** |
| 15 | Runtime | M | 13 | Idle CPU ~0; GIFs play at authored speed |
| 16 | Bug Sweep & Errors | M | 15 | Errors surface; wide lint |
| 17 | Docs & Polish | S | 16 | README true and self-checking |

Phases 9 and 10 split further if a session runs long — per-scene work cuts cleanly anywhere.

## Confirmed decisions

- **Procedural first, curate GIFs** — handcrafted wins every collision; ~36 + ~40 curated.
- **Ambient audio rewritten on a real Go audio library** — in-process PCM, not subprocesses.
- **History purged in one `git-filter-repo` pass at the end of Phase 1** — binary *and* dropped assets together, since curation must decide the asset list first. `.git` 88MB → ~8MB. No remote, so it is nearly free. **Back up before Phase 0.**
- **Personal tool, done well** — module path stays `gottage`; no CI/goreleaser; local `Makefile`; assets embedded.
- **No Git LFS.** With no remote it stores the same bytes on the same disk, the curated set is only 1-3MB, and an LFS pointer file would be silently embedded by `//go:embed` in place of the GIF — build green, every scene garbage. `.gitattributes` marks binaries instead. Revisit conditions in `plan-phase-0.md` §0.2b.

## What this is actually fixing

Every one of these passed a green test suite, because the tests asserted output was non-empty:

- 20 of 36 handcrafted scenes never rendered — 8 commits of work, invisible
- Ambient audio never played a sound — the engine was never called
- All 34 built-in presets silently failed to change the theme
- Unusable on an 80×24 terminal
- Browsing scenes could reach ~1GB of RAM
- The renderer allocated ~64,000 strings/second
- The README understated the scene count by ~90×

The durable fix is not any individual patch. It is: **`RegisterScene` panics. Goldens are byte-exact. Help is generated. Layout is tested at every size. Errors surface.**
