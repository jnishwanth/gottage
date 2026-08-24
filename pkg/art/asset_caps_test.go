package art

import (
	"image/gif"
	"os"
	"path/filepath"
	"testing"
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
func TestAssetsRespectHardCaps(t *testing.T) {
	animDir := filepath.Join("..", "..", "assets", "animations")
	entries, err := os.ReadDir(animDir)
	if err != nil {
		t.Fatalf("read %s: %v", animDir, err)
	}

	checked := 0
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".gif" {
			continue
		}
		name := e.Name()
		path := filepath.Join(animDir, name)
		checked++

		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("%s: stat: %v", name, err)
			continue
		}
		if info.Size() > maxAssetBytes {
			t.Errorf("%s: %d bytes on disk, cap is %d", name, info.Size(), maxAssetBytes)
		}

		f, err := os.Open(path)
		if err != nil {
			t.Errorf("%s: open: %v", name, err)
			continue
		}
		g, err := gif.DecodeAll(f)
		f.Close()
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
		t.Fatalf("no .gif files in %s: the caps would pass vacuously", animDir)
	}
}
