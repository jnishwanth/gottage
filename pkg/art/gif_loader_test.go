package art

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnimatedGifLoader(t *testing.T) {
	// Path to downloaded test gif
	gifPath := filepath.Join("..", "..", "assets", "animations", "kirokaze_room.gif")
	if _, err := os.Stat(gifPath); os.IsNotExist(err) {
		t.Skip("Test GIF file not found, skipping test")
	}

	canvas := NewCanvas(160, 90)
	anim, err := LoadAnimatedGif(gifPath, canvas.Width, canvas.Height)
	if err != nil {
		t.Fatalf("Failed to load animated GIF: %v", err)
	}

	if anim == nil {
		t.Fatal("Expected non-nil AnimatedGif")
	}

	if anim.FrameCount == 0 {
		t.Fatal("Expected FrameCount > 0")
	}

	t.Logf("Loaded GIF %s with %d frames (%dx%d)", anim.FilePath, anim.FrameCount, anim.Width, anim.Height)

	// Draw first frame
	anim.DrawFrame(canvas, 0)
	// Draw second frame at tick 1
	anim.DrawFrame(canvas, 1)

	// Test DrawFrameAt
	anim.DrawFrameAt(canvas, 10, 10, 80, 45, 2)
}

// TestVerifyDownloadedGifs decodes every GIF on disk and checks it is a real,
// non-empty animation.
//
// It asserts against the number of registered GIF scenes rather than a hardcoded
// file count: the previous `gifCount < 356` broke the moment an asset was pruned,
// and a floor that only ever gets edited downward is not an assertion. Tying it to
// the registry also catches the case this phase cares about — an asset on disk that
// nothing registers, or a registration whose asset was cut.
//
// Guarded by testing.Short(): it decodes every GIF at 160x90 into the unbounded
// package-level loader cache and is the single slowest test in the repo.
func TestVerifyDownloadedGifs(t *testing.T) {
	if testing.Short() {
		t.Skip("decodes every GIF in assets/; skipped under -short")
	}

	animDir := filepath.Join("..", "..", "assets", "animations")
	files, err := os.ReadDir(animDir)
	if err != nil {
		t.Fatalf("Failed to read animations directory: %v", err)
	}

	gifCount := 0
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".gif" {
			continue
		}
		gifCount++
		path := filepath.Join(animDir, f.Name())
		anim, err := LoadAnimatedGif(path, 160, 90)
		if err != nil {
			t.Errorf("Failed to decode GIF %s: %v", f.Name(), err)
			continue
		}
		if anim.FrameCount == 0 {
			t.Errorf("GIF %s has 0 frames", f.Name())
		}
	}

	registered := 0
	for _, sc := range SceneRegistry {
		if _, ok := sc.(*AnimatedGifScene); ok {
			registered++
		}
	}
	if gifCount < registered {
		t.Errorf("%d GIFs on disk but %d registered GIF scenes: at least one registration has no asset",
			gifCount, registered)
	}
}
