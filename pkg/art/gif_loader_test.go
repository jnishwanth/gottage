package art

import (
	"path"
	"testing"

	"gottage/assets"
)

func TestAnimatedGifLoader(t *testing.T) {
	// The registration key, resolved through package assets -- not a path
	// relative to this package directory.
	//
	// The previous version stat'd the source tree and called t.Skip when the file
	// was absent. A vanished asset has to fail the suite: a test that skips itself
	// out of existence is the same silent-failure class as the navy fallback it is
	// here to catch.
	gifPath := "assets/animations/kirokaze_room.gif"

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

}

// TestVerifyDownloadedGifs decodes every embedded GIF and checks each one is a
// real, non-empty animation.
//
// It deliberately asserts nothing about registry/asset agreement: TestEverySceneHasAsset
// already asserts that in both directions at path granularity, which is strictly
// stronger than any count comparison here could be.
//
// Guarded by testing.Short(): it decodes every GIF at 160x90 into the unbounded
// package-level loader cache and is the single slowest test in the repo.
func TestVerifyDownloadedGifs(t *testing.T) {
	if testing.Short() {
		t.Skip("decodes every embedded GIF; skipped under -short")
	}

	names, err := assets.List()
	if err != nil {
		t.Fatalf("assets.List: %v", err)
	}
	if len(names) == 0 {
		t.Fatal("no embedded assets: this test would pass vacuously")
	}

	for _, name := range names {
		anim, err := LoadAnimatedGif(name, 160, 90)
		if err != nil {
			t.Errorf("Failed to decode GIF %s: %v", path.Base(name), err)
			continue
		}
		if anim.FrameCount == 0 {
			t.Errorf("GIF %s has 0 frames", path.Base(name))
		}
	}
}
