package art

import (
	"bytes"
	"image/gif"
	"path"
	"testing"

	"gottage/assets"
)

// Hard caps from AGENTS.md §7. All three are required of every asset.
const (
	maxAssetBytes  = 300 * 1024
	maxAssetFrames = 64
	maxAssetWidth  = 640
	maxAssetHeight = 360
)

// TestAssetsRespectHardCaps enforces AGENTS.md §7's three asset caps against every
// GIF actually on disk.
//
// The frame cap is not redundant with the size cap: decoded cost is
// frames x W x H x 24 bytes against a loader cache that stays unbounded until
// Phase 6, so on-disk size does not predict memory cost. Pre-curation,
// waneella_street.gif was 515KB -- under the size cap after a downscale -- but 431
// frames, which is 142MB resident at 160x90.
//
// Source dimensions are read from the GIF's own logical screen descriptor rather
// than from LoadAnimatedGif, which rescales to whatever target size it is handed
// and therefore can never observe a source-dimension violation.
//
// It measures the embedded copies rather than the source tree, because those are
// the bytes that ship in the binary -- and reading them through package assets is
// what lets this test run from any directory.
func TestAssetsRespectHardCaps(t *testing.T) {
	names, err := assets.List()
	if err != nil {
		t.Fatalf("assets.List: %v", err)
	}

	checked := 0
	for _, full := range names {
		if path.Ext(full) != ".gif" {
			continue
		}
		name := path.Base(full)
		checked++

		b, err := assets.ReadFile(full)
		if err != nil {
			t.Errorf("%s: read: %v", name, err)
			continue
		}
		if int64(len(b)) > maxAssetBytes {
			t.Errorf("%s: %d bytes embedded, cap is %d", name, len(b), maxAssetBytes)
		}

		g, err := gif.DecodeAll(bytes.NewReader(b))
		if err != nil {
			t.Errorf("%s: decode: %v", name, err)
			continue
		}

		if n := len(g.Image); n > maxAssetFrames {
			t.Errorf("%s: %d frames, cap is %d", name, n, maxAssetFrames)
		}

		w, h := g.Config.Width, g.Config.Height
		if w <= 0 || h <= 0 {
			t.Errorf("%s: logical screen reports %dx%d, so the dimension cap cannot be checked", name, w, h)
			continue
		}
		if w > maxAssetWidth || h > maxAssetHeight {
			t.Errorf("%s: source is %dx%d, cap is %dx%d", name, w, h, maxAssetWidth, maxAssetHeight)
		}
	}

	if checked == 0 {
		t.Fatal("no embedded .gif files: the caps would pass vacuously")
	}
}
