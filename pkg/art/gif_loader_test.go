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
