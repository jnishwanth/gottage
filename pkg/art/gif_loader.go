package art

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"sync"

	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/draw"
)

// AnimatedGif holds decoded and pre-scaled pixel frames for multi-frame 8-bit / pixel-art GIF animations.
type AnimatedGif struct {
	FilePath   string
	Width      int
	Height     int
	Frames     [][][]colorful.Color // [frameIndex][y][x]
	Delays     []int                // Frame delays in centiseconds (10ms units)
	FrameCount int
}

// AnimatedGifLoader manages cached multi-frame GIF decoding and high-quality bilinear scaling.
type AnimatedGifLoader struct {
	mu    sync.Mutex
	cache map[string]*AnimatedGif
}

// Global default GIF loader instance.
var defaultGifLoader = &AnimatedGifLoader{
	cache: make(map[string]*AnimatedGif),
}

// LoadAnimatedGif loads or returns cached multi-frame AnimatedGif scaled to (targetWidth, targetHeight).
func LoadAnimatedGif(filePath string, targetWidth, targetHeight int) (*AnimatedGif, error) {
	return defaultGifLoader.Load(filePath, targetWidth, targetHeight)
}

// Load decodes an animated GIF file, scales all frames to (targetWidth, targetHeight), and caches the result.
func (loader *AnimatedGifLoader) Load(filePath string, targetWidth, targetHeight int) (*AnimatedGif, error) {
	loader.mu.Lock()
	defer loader.mu.Unlock()

	cacheKey := fmt.Sprintf("%s_%dx%d", filePath, targetWidth, targetHeight)
	if cached, ok := loader.cache[cacheKey]; ok {
		return cached, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open GIF file %s: %w", filePath, err)
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode GIF %s: %w", filePath, err)
	}

	if len(g.Image) == 0 {
		return nil, fmt.Errorf("GIF %s has 0 frames", filePath)
	}

	// Determine original GIF dimensions
	origW, origH := g.Config.Width, g.Config.Height
	if origW == 0 || origH == 0 {
		b := g.Image[0].Bounds()
		origW, origH = b.Dx(), b.Dy()
	}
	if origW == 0 {
		origW = 160
	}
	if origH == 0 {
		origH = 90
	}

	// Prepare compositing canvas for disposal handling
	composite := image.NewRGBA(image.Rect(0, 0, origW, origH))
	anim := &AnimatedGif{
		FilePath:   filePath,
		Width:      targetWidth,
		Height:     targetHeight,
		Frames:     make([][][]colorful.Color, len(g.Image)),
		Delays:     make([]int, len(g.Image)),
		FrameCount: len(g.Image),
	}

	for i, img := range g.Image {
		// Set delay (default to 10 centiseconds if zero)
		delay := g.Delay[i]
		if delay <= 0 {
			delay = 10
		}
		anim.Delays[i] = delay

		// Draw frame onto compositing buffer
		draw.Draw(composite, img.Bounds(), img, img.Bounds().Min, draw.Over)

		// Create target scaled RGBA image
		dstRGBA := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
		draw.BiLinear.Scale(dstRGBA, dstRGBA.Bounds(), composite, composite.Bounds(), draw.Over, nil)

		// Convert scaled RGBA into 2D colorful.Color matrix
		matrix := make([][]colorful.Color, targetHeight)
		for y := 0; y < targetHeight; y++ {
			matrix[y] = make([]colorful.Color, targetWidth)
			for x := 0; x < targetWidth; x++ {
				r, g, b, _ := dstRGBA.At(x, y).RGBA()
				matrix[y][x] = colorful.Color{
					R: float64(r) / 65535.0,
					G: float64(g) / 65535.0,
					B: float64(b) / 65535.0,
				}
			}
		}
		anim.Frames[i] = matrix

		// Handle GIF frame disposal
		var disposal byte = gif.DisposalNone
		if i < len(g.Disposal) {
			disposal = g.Disposal[i]
		}
		if disposal == gif.DisposalBackground {
			draw.Draw(composite, img.Bounds(), image.Transparent, image.Point{}, draw.Src)
		}
	}

	loader.cache[cacheKey] = anim
	return anim, nil
}

// GetFrame returns the 2D colorful.Color pixel matrix for a given animation tick / frame count.
func (anim *AnimatedGif) GetFrame(tick uint64) [][]colorful.Color {
	if anim.FrameCount == 0 {
		return nil
	}
	idx := int(tick) % anim.FrameCount
	return anim.Frames[idx]
}

// DrawFrame renders the GIF frame corresponding to tick directly onto canvas at (0, 0).
func (anim *AnimatedGif) DrawFrame(canvas *Canvas, tick uint64) {
	matrix := anim.GetFrame(tick)
	if matrix == nil {
		return
	}
	h := minInt(canvas.Height, len(matrix))
	for y := 0; y < h; y++ {
		w := minInt(canvas.Width, len(matrix[y]))
		copy(canvas.Pixels[y][:w], matrix[y][:w])
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
