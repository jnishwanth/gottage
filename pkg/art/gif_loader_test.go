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

func TestVerifyDownloadedGifs(t *testing.T) {
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
			continue
		}
		t.Logf("✓ Verified GIF %-24s | Frames: %3d | Scaled: %dx%d", f.Name(), anim.FrameCount, anim.Width, anim.Height)
	}

	if gifCount < 38 {
		t.Errorf("Expected at least 38 GIF animations, found %d", gifCount)
	}
}

