package art

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	RegisterScene(&CityCafeScene{})
	RegisterScene(&ForestCabinScene{})
	RegisterScene(&ClassicCottageScene{})
	RegisterScene(&ObservatoryScene{})
}

// ---------------------------------------------------------
// 1. CityCafeScene (8-Bit Rainy City Cafe)
// ---------------------------------------------------------
type CityCafeScene struct{}

func (s *CityCafeScene) ID() SceneType { return SceneRainyCity }
func (s *CityCafeScene) Name() string { return "8-Bit Rainy City Cafe 🏙️" }
func (s *CityCafeScene) WindowBounds() Rect { return Rect{X: 10, Y: 4, W: 54, H: 38} }

func (s *CityCafeScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	// Dark 8-bit interior wall background
	c.DrawVerticalGradient(0, 0, c.Width, 46, PicoNavy, PicoBlack)

	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoDarkPurple, PicoOrange
	case TimeDay:
		skyTop, skyBot = PicoNavy, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoRed
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// 8-Bit Skyscrapers Silhouette outside window
	bldgColors := []colorful.Color{PicoBlack, PicoNavy, PicoDarkPurple}
	neonColors := []colorful.Color{PicoPink, PicoCyan, PicoYellow, PicoLime}

	for b := 0; b < 9; b++ {
		bx := wx + b*6 + 1
		bw := 5
		bh := 14 + (b*7)%16
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		if tod == TimeNight || tod == TimeSunset {
			neonCol := neonColors[(b+int(frameCount/15))%4]
			for wyPos := wy + wh - bh + 2; wyPos < wy+wh-2; wyPos += 3 {
				for wxPos := bx + 1; wxPos < bx+bw-1; wxPos += 2 {
					if (wxPos+wyPos)%2 == 0 {
						c.SetPixel(wxPos, wyPos, neonCol)
					}
				}
			}
		}
	}

	// 8-Bit Neon Signs
	if tod == TimeNight || tod == TimeSunset {
		neonSignCol := neonColors[int(frameCount/10)%4]
		c.DrawGlyph(wx+4, wy+6, 'L', neonSignCol)
		c.DrawGlyph(wx+5, wy+6, 'O', neonSignCol)
		c.DrawGlyph(wx+6, wy+6, '-', neonSignCol)
		c.DrawGlyph(wx+7, wy+6, 'F', neonSignCol)
		c.DrawGlyph(wx+8, wy+6, 'I', neonSignCol)
	}
}

func (s *CityCafeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Crisp 8-Bit Window Frame (Renders ON TOP of weather particles!)
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoDarkGrey)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoDarkGrey)

	// Sleek 8-bit bar counter table (Renders ON TOP of weather particles!)
	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoPeach)

	// 8-Bit Body-Double Companion (Renders ON TOP of weather particles!)
	if buddy != nil {
		buddy.Draw(c, 30, 34, frameCount)
	}

	// Glowing 8-Bit Laptop on Bar Counter
	c.FillRect(42, 40, 14, 6, PicoDarkGrey)
	c.FillRect(44, 41, 10, 4, PicoCyan)
	for l := 0; l < 3; l++ {
		c.FillRect(45, 42+l, 3+l*2, 1, PicoLime)
	}
	c.FillRect(40, 46, 18, 2, PicoLightGrey)

	// Coffee Mug on Bar Counter
	c.FillRect(40, 43, 3, 3, PicoRed)
}

func (s *CityCafeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(44, 38, 16, 14, PicoYellow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 2. ForestCabinScene (8-Bit Forest Sanctuary Cabin)
// ---------------------------------------------------------
type ForestCabinScene struct{}

func (s *ForestCabinScene) ID() SceneType { return SceneForestCabin }
func (s *ForestCabinScene) Name() string { return "8-Bit Forest Cabin 🌲" }
func (s *ForestCabinScene) WindowBounds() Rect { return Rect{X: 12, Y: 6, W: 50, H: 34} }

func (s *ForestCabinScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoDarkPurple, PicoOrange
	case TimeDay:
		skyTop, skyBot = PicoGreen, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoPeach
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Pine Trees outside window
	for i := 0; i < 7; i++ {
		tx := wx + i*7 + 2
		th := 14 + (i*5)%10
		c.FillRect(tx, wy+wh-th, 5, th, PicoBlack)
		c.FillRect(tx+1, wy+wh-th-4, 3, 4, PicoGreen)
	}
}

func (s *ForestCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Log Wall (Renders ON TOP of weather outside window!)
	c.DrawVerticalGradient(0, 0, wx-1, 46, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 46, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(0, 0, c.Width, wy-1, PicoBrown, PicoDarkGrey)
	for y := 0; y < 46; y += 6 {
		c.FillRect(0, y, c.Width, 1, PicoBlack)
	}

	// Floor
	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoDarkGrey)

	// Timber Window Frame (Renders ON TOP of weather!)
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoBrown)
	c.FillRect(wx+ww/3, wy, 2, wh, PicoBrown)
	c.FillRect(wx+(ww*2)/3, wy, 2, wh, PicoBrown)

	// 8-Bit Body-Double Companion at Bench Desk (Renders ON TOP of weather!)
	if buddy != nil {
		buddy.Draw(c, 18, 34, frameCount)
	}

	// 8-Bit Stone Fireplace
	renderFireplaceHelper(c, 52, 24, 18, 22, fireOn, frameCount)

	// 8-Bit Sleeping Cat
	renderCatHelper(c, 38, 50, catBreed, frameCount)
}

func (s *ForestCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(61, 35, 16.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 3. ClassicCottageScene (8-Bit Classic Cottage)
// ---------------------------------------------------------
type ClassicCottageScene struct{}

func (s *ClassicCottageScene) ID() SceneType { return SceneCottageClassic }
func (s *ClassicCottageScene) Name() string { return "8-Bit Classic Cottage 🏡" }
func (s *ClassicCottageScene) WindowBounds() Rect { return Rect{X: 18, Y: 6, W: 36, H: 28} }

func (s *ClassicCottageScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoLavender, PicoPeach
	case TimeDay:
		skyTop, skyBot = PicoNavy, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoRed
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	if tod == TimeNight && !lightning {
		c.FillRect(wx+24, wy+3, 4, 3, PicoWhite)
		c.FillRect(wx+23, wy+3, 2, 3, skyTop)
	} else if tod == TimeDay {
		c.FillRect(wx+24, wy+3, 4, 3, PicoYellow)
	}
}

func (s *ClassicCottageScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Wallpaper & Floor (Renders ON TOP of weather outside window!)
	c.DrawVerticalGradient(0, 0, wx-1, 46, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 46, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(0, 0, c.Width, wy-1, PicoDarkPurple, PicoNavy)

	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoBlack)

	// Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoBrown)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoBrown)

	// Desk
	renderDeskHelper(c, 6, 42, 28, lampOn)

	// 8-Bit Body-Double Companion at Cottage Desk (Renders ON TOP of weather!)
	if buddy != nil {
		buddy.Draw(c, 10, 32, frameCount)
	}

	// Fireplace
	renderFireplaceHelper(c, 52, 26, 18, 20, fireOn, frameCount)

	// Cat
	renderCatHelper(c, 36, 50, catBreed, frameCount)
}

func (s *ClassicCottageScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(10, 36, 18, 14, PicoYellow, 0.35)
	}
	if fireOn {
		c.ApplyRadialLight(61, 36, 16.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 4. ObservatoryScene (8-Bit Starlight Observatory)
// ---------------------------------------------------------
type ObservatoryScene struct{}

func (s *ObservatoryScene) ID() SceneType { return SceneObservatory }
func (s *ObservatoryScene) Name() string { return "8-Bit Starlight Observatory 🌌" }
func (s *ObservatoryScene) WindowBounds() Rect { return Rect{X: 10, Y: 4, W: 54, H: 36} }

func (s *ObservatoryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, PicoBlack, PicoNavy)

	c.FillRect(wx+36, wy+4, 6, 5, PicoWhite)
	c.FillRect(wx+34, wy+4, 4, 5, PicoBlack)

	for i, st := range stars {
		px := wx + (st.X % ww)
		py := wy + (st.Y % wh)
		twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
		if twinkle > 0.2 {
			c.SetPixel(px, py, PicoWhite)
			if i%3 == 0 {
				c.DrawGlyph(px, py/2, '✦', PicoCyan)
			}
		}
	}
}

func (s *ObservatoryScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Wall interior around window
	c.FillRect(0, 0, wx-1, 46, PicoBlack)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 46, PicoBlack)
	c.FillRect(0, 0, c.Width, wy-1, PicoBlack)
	c.FillRect(0, wy+wh+1, c.Width, c.Height-(wy+wh+1), PicoBlack)

	// Dome Frame (Renders ON TOP of space stars!)
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoDarkGrey)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoDarkGrey)
	c.FillRect(wx, wy+wh/2-1, ww, 2, PicoDarkGrey)

	// 8-Bit Body-Double Companion at Telescope Desk (Renders ON TOP of space stars!)
	if buddy != nil {
		buddy.Draw(c, 20, 34, frameCount)
	}

	// Telescope
	c.FillRect(32, 48, 12, 2, PicoPeach)
	c.FillRect(36, 50, 4, 10, PicoDarkGrey)
}

func (s *ObservatoryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// Helper drawing functions
func renderDeskHelper(c *Canvas, dx, dy, dw int, lampOn bool) {
	c.FillRect(dx, dy, dw, 3, PicoBrown)
	c.FillRect(dx, dy, dw, 1, PicoPeach)
	c.FillRect(dx+2, dy+3, 3, 15, PicoBrown)

	if lampOn {
		c.FillRect(dx+3, dy-1, 3, 1, PicoYellow)
		c.FillRect(dx+4, dy-5, 1, 4, PicoYellow)
		c.FillRect(dx+2, dy-8, 5, 3, PicoGreen)
	} else {
		c.FillRect(dx+3, dy-1, 3, 1, PicoDarkGrey)
		c.FillRect(dx+4, dy-5, 1, 4, PicoDarkGrey)
		c.FillRect(dx+2, dy-8, 5, 3, PicoDarkGrey)
	}

	c.FillRect(dx+12, dy-3, 3, 3, PicoRed)
}

func renderFireplaceHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	c.FillRect(fx, fy, fw, fh, PicoBrown)
	c.FillRect(fx+2, fy+4, fw-4, fh-4, PicoBlack)

	if !fireOn {
		return
	}

	fireCenter := fx + fw/2
	fireTick := (frameCount / 3) % 4

	for i := -3; i <= 3; i++ {
		px := fireCenter + i
		flameH := int(4.0 + 2.0*math.Sin(float64(i)+float64(fireTick)*1.5))

		for fyPos := 0; fyPos < flameH; fyPos++ {
			py := fy + fh - 2 - fyPos
			if py < fy {
				continue
			}
			if fyPos < 2 {
				c.SetPixel(px, py, PicoYellow)
			} else if fyPos < 4 {
				c.SetPixel(px, py, PicoOrange)
			} else {
				c.SetPixel(px, py, PicoRed)
			}
		}
	}
}

func renderCatHelper(c *Canvas, cx, cy int, breed CatBreed, frameCount uint64) {
	c.FillRect(cx-2, cy+4, 12, 3, PicoDarkPurple)

	var furMain colorful.Color
	switch breed {
	case CatOrange:
		furMain = PicoOrange
	case CatTuxedo:
		furMain = PicoNavy
	case CatCalico:
		furMain = PicoPeach
	case CatBlack:
		furMain = PicoBlack
	}

	breathOffset := 0
	if (frameCount/12)%2 == 0 {
		breathOffset = -1
	}

	c.FillRect(cx, cy+breathOffset, 8, 4, furMain)
	c.FillRect(cx-1, cy+1+breathOffset, 3, 3, furMain)

	if (frameCount/10)%3 == 1 {
		c.DrawGlyph(cx-2, (cy-2)/2, 'z', PicoCyan)
	}
}
