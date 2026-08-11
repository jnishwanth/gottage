package art

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"

	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/image/draw"
)

// ImageLoader handles decoding and high-quality bilinear scaling of image assets onto the Canvas.
type ImageLoader struct {
	mu    sync.Mutex
	cache map[string][][]colorful.Color
}

// Global Image Loader instance
var defaultLoader = &ImageLoader{
	cache: make(map[string][][]colorful.Color),
}

// LoadImageToCanvas decodes an image file, resizes it using high-quality bilinear interpolation
// via golang.org/x/image/draw, and writes the RGB pixel colors directly into canvas.
func LoadImageToCanvas(filePath string, canvas *Canvas) error {
	return defaultLoader.Load(filePath, canvas)
}

// Load loads or fetches cached scaled pixel matrix for canvas.
func (il *ImageLoader) Load(filePath string, canvas *Canvas) error {
	il.mu.Lock()
	defer il.mu.Unlock()

	cacheKey := fmt.Sprintf("%s_%dx%d", filePath, canvas.Width, canvas.Height)
	if cachedPixels, ok := il.cache[cacheKey]; ok {
		// Fast copy from cache
		for y := 0; y < canvas.Height; y++ {
			copy(canvas.Pixels[y], cachedPixels[y])
		}
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open image file %s: %w", filePath, err)
	}
	defer file.Close()

	srcImg, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image %s: %w", filePath, err)
	}

	// Create target RGBA buffer matching canvas dimensions
	dstRGBA := image.NewRGBA(image.Rect(0, 0, canvas.Width, canvas.Height))

	// High-quality bilinear scaling using standard extended library golang.org/x/image/draw
	draw.BiLinear.Scale(dstRGBA, dstRGBA.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)

	// Copy into colorful.Color 2D matrix
	cachedMatrix := make([][]colorful.Color, canvas.Height)
	for y := 0; y < canvas.Height; y++ {
		cachedMatrix[y] = make([]colorful.Color, canvas.Width)
		for x := 0; x < canvas.Width; x++ {
			r, g, b, _ := dstRGBA.At(x, y).RGBA()
			// RGBA() returns 16-bit values [0..65535]
			col := colorful.Color{
				R: float64(r) / 65535.0,
				G: float64(g) / 65535.0,
				B: float64(b) / 65535.0,
			}
			cachedMatrix[y][x] = col
			canvas.Pixels[y][x] = col
		}
	}

	il.cache[cacheKey] = cachedMatrix
	return nil
}
