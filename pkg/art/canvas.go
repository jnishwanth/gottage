package art

import (
	"fmt"
	"math"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

// CanvasRenderMode specifies terminal art rendering aesthetic style.
type CanvasRenderMode int

const (
	RenderHalfBlock CanvasRenderMode = iota // ▀ 24-Bit TrueColor Half-Blocks (Default pixel art)
	RenderASCII                             // ASCII Luminance Density Ramp ( .:-=+*#%@#) with 24-bit TrueColor
	RenderBraille                           // Unicode Braille Matrix (⠀-⣿) with pixel subgrid density
	RenderPhosphor                          // Retro Monochromatic CRT Glow (Amber / Emerald Phosphor)
)

func (r CanvasRenderMode) String() string {
	switch r {
	case RenderASCII:
		return "ASCII Density Ramp ( .:-=+*#%@)"
	case RenderBraille:
		return "Unicode Braille Matrix (⠀-⣿)"
	case RenderPhosphor:
		return "Retro CRT Phosphor Glow (📟 Amber/Green)"
	default:
		return "24-Bit TrueColor Half-Blocks (▀)"
	}
}

// Canvas represents a high-density 24-bit TrueColor pixel canvas.
// Standard character cells have a 1:2 width:height aspect ratio.
// Using half-block '▀', each character cell renders 2 vertical pixels (top & bottom).
type Canvas struct {
	Width       int // width in pixel columns (also character columns)
	Height      int // height in pixel rows (character rows = Height / 2)
	Pixels      [][]colorful.Color
	Glyphs      [][]rune
	GlyphColors [][]colorful.Color
	RenderMode  CanvasRenderMode
}

// NewCanvas creates a new Canvas with the specified width (in characters/pixels)
// and height (in pixel rows, should be even).
func NewCanvas(width, height int) *Canvas {
	if height%2 != 0 {
		height++
	}

	pixels := make([][]colorful.Color, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]colorful.Color, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = colorful.Color{R: 0, G: 0, B: 0}
		}
	}

	charHeight := height / 2
	glyphs := make([][]rune, charHeight)
	glyphColors := make([][]colorful.Color, charHeight)
	for y := 0; y < charHeight; y++ {
		glyphs[y] = make([]rune, width)
		glyphColors[y] = make([]colorful.Color, width)
	}

	return &Canvas{
		Width:       width,
		Height:      height,
		Pixels:      pixels,
		Glyphs:      glyphs,
		GlyphColors: glyphColors,
		RenderMode:  RenderHalfBlock,
	}
}

// Clear fills the entire canvas with a background color and removes all glyphs.
func (c *Canvas) Clear(bg colorful.Color) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.Pixels[y][x] = bg
		}
	}
	charHeight := c.Height / 2
	for y := 0; y < charHeight; y++ {
		for x := 0; x < c.Width; x++ {
			c.Glyphs[y][x] = 0
		}
	}
}

// SetPixel sets a single pixel at (x, y) where y is in pixel coordinates [0, Height-1].
func (c *Canvas) SetPixel(x, y int, color colorful.Color) {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return
	}
	c.Pixels[y][x] = color
}

// SetPixelHex sets a pixel using a hex string like "#ff5588".
func (c *Canvas) SetPixelHex(x, y int, hex string) {
	col, err := colorful.Hex(hex)
	if err == nil {
		c.SetPixel(x, y, col)
	}
}

// BlendPixel alpha-blends a new color over existing pixel at (x, y).
func (c *Canvas) BlendPixel(x, y int, color colorful.Color, alpha float64) {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return
	}
	if alpha <= 0 {
		return
	}
	if alpha >= 1.0 {
		c.Pixels[y][x] = color
		return
	}
	current := c.Pixels[y][x]
	c.Pixels[y][x] = current.BlendRgb(color, alpha)
}

// DrawGlyph overlay an ASCII/Unicode character at character cell (x, charY).
func (c *Canvas) DrawGlyph(x, charY int, r rune, fg colorful.Color) {
	if x < 0 || x >= c.Width || charY < 0 || charY >= (c.Height/2) {
		return
	}
	c.Glyphs[charY][x] = r
	c.GlyphColors[charY][x] = fg
}

// DrawRectOutline draws a rectangular border in pixel coordinates.
func (c *Canvas) DrawRectOutline(x, y, w, h int, color colorful.Color) {
	for px := x; px < x+w; px++ {
		c.SetPixel(px, y, color)
		c.SetPixel(px, y+h-1, color)
	}
	for py := y; py < y+h; py++ {
		c.SetPixel(x, py, color)
		c.SetPixel(x+w-1, py, color)
	}
}

// FillRect fills a rectangle in pixel coordinates.
func (c *Canvas) FillRect(x, y, w, h int, color colorful.Color) {
	for py := y; py < y+h; py++ {
		for px := x; px < x+w; px++ {
			c.SetPixel(px, py, color)
		}
	}
}

// BlendRect fills a rectangle with alpha blending.
func (c *Canvas) BlendRect(x, y, w, h int, color colorful.Color, alpha float64) {
	for py := y; py < y+h; py++ {
		for px := x; px < x+w; px++ {
			c.BlendPixel(px, py, color, alpha)
		}
	}
}

// DrawVerticalGradient fills a rectangle with a top-to-bottom color gradient.
func (c *Canvas) DrawVerticalGradient(x, y, w, h int, topCol, botCol colorful.Color) {
	if h <= 0 {
		return
	}
	for py := 0; py < h; py++ {
		t := float64(py) / float64(h-1)
		if h == 1 {
			t = 0
		}
		col := topCol.BlendRgb(botCol, t)
		for px := x; px < x+w; px++ {
			c.SetPixel(px, y+py, col)
		}
	}
}

// DrawHorizontalGradient fills a rectangle with a left-to-right color gradient.
func (c *Canvas) DrawHorizontalGradient(x, y, w, h int, leftCol, rightCol colorful.Color) {
	if w <= 0 {
		return
	}
	for px := 0; px < w; px++ {
		t := float64(px) / float64(w-1)
		if w == 1 {
			t = 0
		}
		col := leftCol.BlendRgb(rightCol, t)
		for py := y; py < y+h; py++ {
			c.SetPixel(x+px, py, col)
		}
	}
}

// DrawLine draws a 1-pixel line from (x0, y0) to (x1, y1) using Bresenham's algorithm.
func (c *Canvas) DrawLine(x0, y0, x1, y1 int, color colorful.Color) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := -int(math.Abs(float64(y1 - y0)))
	sx := 1
	if x0 >= x1 {
		sx = -1
	}
	sy := 1
	if y0 >= y1 {
		sy = -1
	}
	err := dx + dy

	for {
		c.SetPixel(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

// DrawCircle draws a 1-pixel circle outline centered at (cx, cy) with radius r.
func (c *Canvas) DrawCircle(cx, cy, r int, color colorful.Color) {
	if r <= 0 {
		c.SetPixel(cx, cy, color)
		return
	}
	x := 0
	y := r
	d := 3 - 2*r
	c.drawCirclePoints(cx, cy, x, y, color)
	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		c.drawCirclePoints(cx, cy, x, y, color)
	}
}

func (c *Canvas) drawCirclePoints(cx, cy, x, y int, color colorful.Color) {
	c.SetPixel(cx+x, cy+y, color)
	c.SetPixel(cx-x, cy+y, color)
	c.SetPixel(cx+x, cy-y, color)
	c.SetPixel(cx-x, cy-y, color)
	c.SetPixel(cx+y, cy+x, color)
	c.SetPixel(cx-y, cy+x, color)
	c.SetPixel(cx+y, cy-x, color)
	c.SetPixel(cx-y, cy-x, color)
}

// FillCircle draws a filled circle centered at (cx, cy) with radius r.
func (c *Canvas) FillCircle(cx, cy, r int, color colorful.Color) {
	if r <= 0 {
		c.SetPixel(cx, cy, color)
		return
	}
	for dy := -r; dy <= r; dy++ {
		span := int(math.Sqrt(float64(r*r - dy*dy)))
		for dx := -span; dx <= span; dx++ {
			c.SetPixel(cx+dx, cy+dy, color)
		}
	}
}

// DrawEllipse draws an ellipse outline with horizontal radius rx and vertical radius ry.
func (c *Canvas) DrawEllipse(cx, cy, rx, ry int, color colorful.Color) {
	if rx <= 0 || ry <= 0 {
		c.SetPixel(cx, cy, color)
		return
	}
	for angle := 0.0; angle < 2*math.Pi; angle += 0.05 {
		px := cx + int(math.Round(float64(rx)*math.Cos(angle)))
		py := cy + int(math.Round(float64(ry)*math.Sin(angle)))
		c.SetPixel(px, py, color)
	}
}

// FillEllipse fills an ellipse with horizontal radius rx and vertical radius ry.
func (c *Canvas) FillEllipse(cx, cy, rx, ry int, color colorful.Color) {
	if rx <= 0 || ry <= 0 {
		c.SetPixel(cx, cy, color)
		return
	}
	for dy := -ry; dy <= ry; dy++ {
		t := float64(dy) / float64(ry)
		if math.Abs(t) <= 1.0 {
			span := int(math.Round(float64(rx) * math.Sqrt(1.0-t*t)))
			for dx := -span; dx <= span; dx++ {
				c.SetPixel(cx+dx, cy+dy, color)
			}
		}
	}
}

// DrawArc draws a circular arc from startAngle to endAngle (in radians).
func (c *Canvas) DrawArc(cx, cy, r int, startAngle, endAngle float64, color colorful.Color) {
	if r <= 0 {
		c.SetPixel(cx, cy, color)
		return
	}
	step := 1.0 / float64(r*2)
	for angle := startAngle; angle <= endAngle; angle += step {
		px := cx + int(math.Round(float64(r)*math.Cos(angle)))
		py := cy + int(math.Round(float64(r)*math.Sin(angle)))
		c.SetPixel(px, py, color)
	}
}


// ApplyRadialLight applies a warm radial light source centered at (cx, cy) in pixel coords.
func (c *Canvas) ApplyRadialLight(cx, cy int, radius float64, lightColor colorful.Color, maxIntensity float64) {
	minY := int(math.Max(0, float64(cy)-radius))
	maxY := int(math.Min(float64(c.Height-1), float64(cy)+radius))
	minX := int(math.Max(0, float64(cx)-radius*2)) // adjust aspect ratio
	maxX := int(math.Min(float64(c.Width-1), float64(cx)+radius*2))

	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			// Account for cell aspect ratio (pixels are roughly twice as tall as wide in character mode)
			dx := float64(px-cx) / 1.8
			dy := float64(py - cy)
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < radius {
				// Quadratic falloff for smooth lighting
				factor := 1.0 - (dist / radius)
				intensity := factor * factor * maxIntensity
				c.BlendPixel(px, py, lightColor, intensity)
			}
		}
	}
}

// ApplyLightCone applies a spot/cone light (like a desk lamp pointing down).
func (c *Canvas) ApplyLightCone(cx, cy int, height, angleWidth int, lightColor colorful.Color, intensity float64) {
	for dy := 0; dy < height; dy++ {
		py := cy + dy
		if py < 0 || py >= c.Height {
			continue
		}
		spread := int(float64(dy) * float64(angleWidth) / float64(height))
		minX := cx - spread
		maxX := cx + spread
		tY := 1.0 - float64(dy)/float64(height)
		layerIntensity := intensity * (0.3 + 0.7*tY)

		for px := minX; px <= maxX; px++ {
			if px < 0 || px >= c.Width {
				continue
			}
			distCenter := math.Abs(float64(px-cx)) / float64(spread+1)
			edgeFade := 1.0 - distCenter*distCenter
			if edgeFade > 0 {
				c.BlendPixel(px, py, lightColor, layerIntensity*edgeFade)
			}
		}
	}
}

// Render converts the Canvas into an ANSI TrueColor string based on c.RenderMode.
func (c *Canvas) Render() string {
	switch c.RenderMode {
	case RenderASCII:
		return c.renderASCII()
	case RenderBraille:
		return c.renderBraille()
	case RenderPhosphor:
		return c.renderPhosphor()
	default:
		return c.renderHalfBlock()
	}
}

func (c *Canvas) renderHalfBlock() string {
	var sb strings.Builder
	charHeight := c.Height / 2

	var lastFgR, lastFgG, lastFgB int = -1, -1, -1
	var lastBgR, lastBgG, lastBgB int = -1, -1, -1

	for charY := 0; charY < charHeight; charY++ {
		topY := charY * 2
		botY := charY*2 + 1

		for x := 0; x < c.Width; x++ {
			topCol := c.Pixels[topY][x]
			botCol := c.Pixels[botY][x]

			topR, topG, topB := int(topCol.R*255), int(topCol.G*255), int(topCol.B*255)
			botR, botG, botB := int(botCol.R*255), int(botCol.G*255), int(botCol.B*255)

			r := c.Glyphs[charY][x]
			if r != 0 {
				fgCol := c.GlyphColors[charY][x]
				fgR, fgG, fgB := int(fgCol.R*255), int(fgCol.G*255), int(fgCol.B*255)

				if fgR != lastFgR || fgG != lastFgG || fgB != lastFgB {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", fgR, fgG, fgB))
					lastFgR, lastFgG, lastFgB = fgR, fgG, fgB
				}
				if botR != lastBgR || botG != lastBgG || botB != lastBgB {
					sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", botR, botG, botB))
					lastBgR, lastBgG, lastBgB = botR, botG, botB
				}
				sb.WriteRune(r)
			} else {
				if topR != lastFgR || topG != lastFgG || topB != lastFgB {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", topR, topG, topB))
					lastFgR, lastFgG, lastFgB = topR, topG, topB
				}
				if botR != lastBgR || botG != lastBgG || botB != lastBgB {
					sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", botR, botG, botB))
					lastBgR, lastBgG, lastBgB = botR, botG, botB
				}
				sb.WriteString("▀")
			}
		}
		if charY < charHeight-1 {
			sb.WriteString("\x1b[0m\n")
		} else {
			sb.WriteString("\x1b[0m")
		}
		lastFgR, lastFgG, lastFgB = -1, -1, -1
		lastBgR, lastBgG, lastBgB = -1, -1, -1
	}

	return sb.String()
}

func (c *Canvas) renderASCII() string {
	var sb strings.Builder
	charHeight := c.Height / 2
	ramp := []rune{' ', '.', ':', '-', '=', '+', '*', '#', '%', '@'}

	var lastFgR, lastFgG, lastFgB int = -1, -1, -1

	for charY := 0; charY < charHeight; charY++ {
		topY := charY * 2
		botY := charY*2 + 1

		for x := 0; x < c.Width; x++ {
			topCol := c.Pixels[topY][x]
			botCol := c.Pixels[botY][x]
			avgCol := topCol.BlendRgb(botCol, 0.5)

			r := c.Glyphs[charY][x]
			if r == 0 {
				lum := 0.299*avgCol.R + 0.587*avgCol.G + 0.114*avgCol.B
				idx := int(lum * float64(len(ramp)-1))
				if idx < 0 {
					idx = 0
				}
				if idx >= len(ramp) {
					idx = len(ramp) - 1
				}
				r = ramp[idx]
			} else {
				avgCol = c.GlyphColors[charY][x]
			}

			fgR, fgG, fgB := int(avgCol.R*255), int(avgCol.G*255), int(avgCol.B*255)
			if fgR != lastFgR || fgG != lastFgG || fgB != lastFgB {
				sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", fgR, fgG, fgB))
				lastFgR, lastFgG, lastFgB = fgR, fgG, fgB
			}
			sb.WriteRune(r)
		}
		if charY < charHeight-1 {
			sb.WriteString("\x1b[0m\n")
		} else {
			sb.WriteString("\x1b[0m")
		}
		lastFgR, lastFgG, lastFgB = -1, -1, -1
	}

	return sb.String()
}

func (c *Canvas) renderBraille() string {
	var sb strings.Builder
	charHeight := c.Height / 2

	var lastFgR, lastFgG, lastFgB int = -1, -1, -1

	for charY := 0; charY < charHeight; charY++ {
		topY := charY * 2
		botY := charY*2 + 1

		for x := 0; x < c.Width; x++ {
			topCol := c.Pixels[topY][x]
			botCol := c.Pixels[botY][x]
			avgCol := topCol.BlendRgb(botCol, 0.5)

			r := c.Glyphs[charY][x]
			if r == 0 {
				lumTop := 0.299*topCol.R + 0.587*topCol.G + 0.114*topCol.B
				lumBot := 0.299*botCol.R + 0.587*botCol.G + 0.114*botCol.B

				var mask int
				if lumTop > 0.05 {
					mask |= 0x1
				}
				if lumTop > 0.25 {
					mask |= 0x8
				}
				if lumTop > 0.50 {
					mask |= 0x2
				}
				if lumTop > 0.75 {
					mask |= 0x10
				}
				if lumBot > 0.05 {
					mask |= 0x4
				}
				if lumBot > 0.25 {
					mask |= 0x20
				}
				if lumBot > 0.50 {
					mask |= 0x40
				}
				if lumBot > 0.75 {
					mask |= 0x80
				}

				if mask == 0 {
					r = ' '
				} else {
					r = rune(0x2800 | mask)
				}
			} else {
				avgCol = c.GlyphColors[charY][x]
			}

			fgR, fgG, fgB := int(avgCol.R*255), int(avgCol.G*255), int(avgCol.B*255)
			if fgR != lastFgR || fgG != lastFgG || fgB != lastFgB {
				sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", fgR, fgG, fgB))
				lastFgR, lastFgG, lastFgB = fgR, fgG, fgB
			}
			sb.WriteRune(r)
		}
		if charY < charHeight-1 {
			sb.WriteString("\x1b[0m\n")
		} else {
			sb.WriteString("\x1b[0m")
		}
		lastFgR, lastFgG, lastFgB = -1, -1, -1
	}

	return sb.String()
}

func (c *Canvas) renderPhosphor() string {
	var sb strings.Builder
	charHeight := c.Height / 2

	var lastFgR, lastFgG, lastFgB int = -1, -1, -1
	var lastBgR, lastBgG, lastBgB int = -1, -1, -1

	for charY := 0; charY < charHeight; charY++ {
		topY := charY * 2
		botY := charY*2 + 1

		for x := 0; x < c.Width; x++ {
			topCol := c.Pixels[topY][x]
			botCol := c.Pixels[botY][x]

			topLum := 0.299*topCol.R + 0.587*topCol.G + 0.114*topCol.B
			botLum := 0.299*botCol.R + 0.587*botCol.G + 0.114*botCol.B

			topPhos := colorful.Color{R: topLum, G: topLum * 0.72, B: topLum * 0.05}
			botPhos := colorful.Color{R: botLum, G: botLum * 0.72, B: botLum * 0.05}

			topR, topG, topB := int(topPhos.R*255), int(topPhos.G*255), int(topPhos.B*255)
			botR, botG, botB := int(botPhos.R*255), int(botPhos.G*255), int(botPhos.B*255)

			r := c.Glyphs[charY][x]
			if r != 0 {
				if topR != lastFgR || topG != lastFgG || topB != lastFgB {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", topR, topG, topB))
					lastFgR, lastFgG, lastFgB = topR, topG, topB
				}
				if botR != lastBgR || botG != lastBgG || botB != lastBgB {
					sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", botR, botG, botB))
					lastBgR, lastBgG, lastBgB = botR, botG, botB
				}
				sb.WriteRune(r)
			} else {
				if topR != lastFgR || topG != lastFgG || topB != lastFgB {
					sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", topR, topG, topB))
					lastFgR, lastFgG, lastFgB = topR, topG, topB
				}
				if botR != lastBgR || botG != lastBgG || botB != lastBgB {
					sb.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", botR, botG, botB))
					lastBgR, lastBgG, lastBgB = botR, botG, botB
				}
				sb.WriteString("▀")
			}
		}
		if charY < charHeight-1 {
			sb.WriteString("\x1b[0m\n")
		} else {
			sb.WriteString("\x1b[0m")
		}
		lastFgR, lastFgG, lastFgB = -1, -1, -1
		lastBgR, lastBgG, lastBgB = -1, -1, -1
	}

	return sb.String()
}

// ExtractDominantPalette samples the canvas pixels and returns up to n dominant colors.
func (c *Canvas) ExtractDominantPalette(n int) []colorful.Color {
	if n <= 0 {
		return nil
	}
	colorCounts := make(map[string]int)
	colorMap := make(map[string]colorful.Color)

	// Sample grid every 4th pixel for speed and efficiency
	for y := 0; y < c.Height; y += 4 {
		for x := 0; x < c.Width; x += 4 {
			col := c.Pixels[y][x]
			r := uint8(math.Round(col.R*15) * 17)
			g := uint8(math.Round(col.G*15) * 17)
			b := uint8(math.Round(col.B*15) * 17)
			qCol := colorful.Color{R: float64(r) / 255.0, G: float64(g) / 255.0, B: float64(b) / 255.0}
			hex := qCol.Hex()
			colorCounts[hex]++
			colorMap[hex] = qCol
		}
	}

	type colCount struct {
		hex   string
		count int
	}
	var sorted []colCount
	for h, cnt := range colorCounts {
		sorted = append(sorted, colCount{hex: h, count: cnt})
	}

	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j].count > sorted[i].count {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	var result []colorful.Color
	for i := 0; i < len(sorted) && i < n; i++ {
		result = append(result, colorMap[sorted[i].hex])
	}
	return result
}

// ApplyVignette applies a subtle radial edge darkening to the canvas.
func (c *Canvas) ApplyVignette(intensity float64) {
	if intensity <= 0 || c.Width == 0 || c.Height == 0 {
		return
	}
	halfW := float64(c.Width) / 2.0
	halfH := float64(c.Height) / 2.0
	black := colorful.Color{R: 0, G: 0, B: 0}

	for y := 0; y < c.Height; y++ {
		dy := (float64(y) - halfH) / halfH
		for x := 0; x < c.Width; x++ {
			dx := (float64(x) - halfW) / halfW
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist > 0.4 {
				vignetteFactor := math.Pow((dist-0.4)/0.6, 1.6) * intensity
				if vignetteFactor > 0.75 {
					vignetteFactor = 0.75
				}
				c.Pixels[y][x] = c.Pixels[y][x].BlendRgb(black, vignetteFactor)
			}
		}
	}
}

// Clone returns a deep copy of the Canvas.
func (c *Canvas) Clone() *Canvas {
	if c == nil {
		return nil
	}
	clone := NewCanvas(c.Width, c.Height)
	clone.RenderMode = c.RenderMode
	for y := 0; y < c.Height; y++ {
		copy(clone.Pixels[y], c.Pixels[y])
	}
	charHeight := c.Height / 2
	for y := 0; y < charHeight; y++ {
		copy(clone.Glyphs[y], c.Glyphs[y])
		copy(clone.GlyphColors[y], c.GlyphColors[y])
	}
	return clone
}

// BlendCanvas blends pixels from another canvas into this canvas with a given alpha weight.
func (c *Canvas) BlendCanvas(other *Canvas, alpha float64) {
	if other == nil || alpha <= 0 || c.Width != other.Width || c.Height != other.Height {
		return
	}
	if alpha >= 1.0 {
		for y := 0; y < c.Height; y++ {
			copy(c.Pixels[y], other.Pixels[y])
		}
		return
	}
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			c.Pixels[y][x] = c.Pixels[y][x].BlendRgb(other.Pixels[y][x], alpha)
		}
	}
}

// ApplyCRTScanlines applies alternating horizontal scanline dimming for retro CRT monitor aesthetics.
func (c *Canvas) ApplyCRTScanlines(depth float64) {
	if depth <= 0 {
		return
	}
	mult := 1.0 - (0.22 * depth)
	if mult < 0.5 {
		mult = 0.5
	}
	for y := 1; y < c.Height; y += 2 {
		for x := 0; x < c.Width; x++ {
			col := c.Pixels[y][x]
			c.Pixels[y][x] = colorful.Color{
				R: col.R * mult,
				G: col.G * mult,
				B: col.B * mult,
			}
		}
	}
}


