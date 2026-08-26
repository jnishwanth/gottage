package art

import (
	"testing"

	"gottage/assets"
)

// Counts Phase 1's curation landed on. Asserted here rather than left implicit
// because the pre-curation registry claimed 380 registrations for 360 distinct
// IDs, and nothing in the repo noticed.
const (
	wantRegistrations = 46
	wantGifScenes     = 10
)

// TestEveryRegisteredAssetResolves pins the registry's shape and checks every
// asset-backed scene can reach its bytes.
//
// It overlaps two neighbours on purpose and adds what neither has:
// TestEverySceneHasAsset asserts the registry and the embedded set are a
// bijection, which stays true if both shrink together, and
// assets.TestEmbeddedAssetsAreRealGifs counts assets without knowing how many
// scenes there are. The counts below are the part that catches a scene
// disappearing.
func TestEveryRegisteredAssetResolves(t *testing.T) {
	if got := len(SceneRegistry); got != wantRegistrations {
		t.Errorf("%d scenes registered, want %d", got, wantRegistrations)
	}

	gifScenes := 0
	for _, id := range SceneOrder {
		gs, ok := SceneRegistry[id].(*AnimatedGifScene)
		if !ok {
			continue
		}
		gifScenes++
		f, err := assets.Open(gs.gifPath)
		if err != nil {
			t.Errorf("SceneType %d (%q): %v", id, gs.name, err)
			continue
		}
		f.Close()
	}
	if gifScenes != wantGifScenes {
		t.Errorf("%d GIF-backed scenes, want %d", gifScenes, wantGifScenes)
	}
}

// cwdCanvas is a size no other test renders at.
//
// The loader cache is package-global and keyed "path_WxH", so at a shared size
// an earlier test could hand this one a decode it performed while the working
// directory was still the package directory -- and the test would pass without
// ever resolving an asset from the temp directory it chdir'd into. Phase 6 owns
// the cache; keeping the key unique is this test's own defence.
const (
	cwdCanvasW = 73
	cwdCanvasH = 51
)

// minGifSceneColors is measured, not guessed. Across the ten curated GIFs at
// 73x51 the counts run 1481 (rain_setup) to 2421 (balcony_lofi); the failure mode
// this has to catch scores exactly 1, because c.Clear(PicoNavy) paints every
// pixel the same navy.
//
// The floor sits well under the measured minimum rather than just under it. A GIF
// carries at most 256 palette entries per frame and only exceeds that here
// because draw.BiLinear interpolates; if Phase 4 or 15 ever swaps in a
// palette-preserving scaler these counts drop by an order of magnitude without
// anything being wrong. 64 still fails a flat fill by 64x.
const minGifSceneColors = 64

// TestScenesRenderFromArbitraryCWD is the regression test for this phase.
//
// Every GIF path used to be a CWD-relative string literal opened with os.Open, so
// the scenes rendered as flat navy anywhere but the repo root -- and pkg/art's
// own tests run from pkg/art/, meaning they only ever exercised that failure
// path. Every later phase's rendering assertion depends on this not being true.
func TestScenesRenderFromArbitraryCWD(t *testing.T) {
	t.Chdir(t.TempDir())
	ClearAssetErrors()
	t.Cleanup(ClearAssetErrors)

	rendered := 0
	for _, id := range SceneOrder {
		gs, ok := SceneRegistry[id].(*AnimatedGifScene)
		if !ok {
			continue
		}
		rendered++
		t.Run(gs.name, func(t *testing.T) {
			c := NewCanvas(cwdCanvasW, cwdCanvasH)
			gs.RenderSky(c, TimeNight, false, 0, nil)
			got := distinctColors(c)
			t.Logf("%s: %d distinct colors", gs.gifPath, got)
			if got < minGifSceneColors {
				t.Errorf("%d distinct colors, want at least %d -- this is what the navy fallback looks like",
					got, minGifSceneColors)
			}
		})
	}

	if rendered == 0 {
		t.Fatal("no GIF-backed scenes rendered: this test would pass vacuously")
	}
	if errs := AssetErrors(); len(errs) != 0 {
		t.Errorf("assets failed to resolve from a temp directory: %v", keysOf(errs))
	}
}
