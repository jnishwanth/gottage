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

	RegisterScene(&AnimatedGifScene{
		id:        SceneKirokazeRoom,
		name:      "Kirokaze Cozy Room 🎨",
		gifPath:   "assets/animations/kirokaze_room.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBalconyLofi,
		name:      "Lofi Balcony View 🏙️",
		gifPath:   "assets/animations/balcony_lofi.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRainSetup,
		name:      "Pixel Rain Setup 🌧️",
		gifPath:   "assets/animations/rain_setup.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneKirokazeCity,
		name:      "Kirokaze Retro City 🌃",
		gifPath:   "assets/animations/kirokaze_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneHomeSweetHome,
		name:      "Home Sweet Home 🏡",
		gifPath:   "assets/animations/home_sweet_home.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneWaneellaStreet,
		name:      "Waneella Pixel Street 🏮",
		gifPath:   "assets/animations/waneella_street.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSunsetCity,
		name:      "Sunset City Skyline 🌇",
		gifPath:   "assets/animations/sunset_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSnowCabin,
		name:      "Winter Cabin Snowfall ❄️",
		gifPath:   "assets/animations/snow_cabin.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
}

// ---------------------------------------------------------
// AnimatedGifScene (Generic Decoded GIF Animation Scene)
// ---------------------------------------------------------
type AnimatedGifScene struct {
	id        SceneType
	name      string
	gifPath   string
	winBounds Rect
	showBuddy bool
	buddyX    int
	buddyY    int
}

func (s *AnimatedGifScene) ID() SceneType       { return s.id }
func (s *AnimatedGifScene) Name() string     { return s.name }
func (s *AnimatedGifScene) WindowBounds() Rect { return s.winBounds }

func (s *AnimatedGifScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	anim, err := LoadAnimatedGif(s.gifPath, c.Width, c.Height)
	if err == nil && anim != nil {
		anim.DrawFrame(c, frameCount)
	} else {
		c.Clear(PicoNavy)
	}
}

func (s *AnimatedGifScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	if s.showBuddy && buddy != nil {
		buddy.Draw(c, s.buddyX, s.buddyY, frameCount)
	}
}

func (s *AnimatedGifScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.50)
	}
}

// ---------------------------------------------------------
// 1. CityCafeScene (16-Bit Rainy City Cafe)
// ---------------------------------------------------------
type CityCafeScene struct{}

func (s *CityCafeScene) ID() SceneType { return SceneRainyCity }
func (s *CityCafeScene) Name() string { return "16-Bit Rainy City Cafe 🏙️" }
func (s *CityCafeScene) WindowBounds() Rect { return Rect{X: 20, Y: 8, W: 120, H: 58} }

func (s *CityCafeScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	c.DrawVerticalGradient(0, 0, c.Width, 66, PicoNavy, PicoBlack)

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

	// 16-Bit Skyscrapers Silhouettes
	bldgColors := []colorful.Color{PicoBlack, PicoNavy, PicoDarkPurple}
	neonColors := []colorful.Color{PicoPink, PicoCyan, PicoYellow, PicoLime}

	for b := 0; b < 10; b++ {
		bx := wx + b*12 + 2
		bw := 10
		bh := 24 + (b*13)%28
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		if tod == TimeNight || tod == TimeSunset {
			neonCol := neonColors[(b+int(frameCount/15))%4]
			for wyPos := wy + wh - bh + 4; wyPos < wy+wh-4; wyPos += 6 {
				for wxPos := bx + 2; wxPos < bx+bw-2; wxPos += 3 {
					if (wxPos+wyPos)%2 == 0 {
						c.SetPixel(wxPos, wyPos, neonCol)
					}
				}
			}
		}
	}

	// 16-Bit Neon Sign
	if tod == TimeNight || tod == TimeSunset {
		neonSignCol := neonColors[int(frameCount/10)%4]
		c.DrawGlyph(wx+8, wy+10, 'L', neonSignCol)
		c.DrawGlyph(wx+10, wy+10, 'O', neonSignCol)
		c.DrawGlyph(wx+12, wy+10, '-', neonSignCol)
		c.DrawGlyph(wx+14, wy+10, 'F', neonSignCol)
		c.DrawGlyph(wx+16, wy+10, 'I', neonSignCol)
	}
}

func (s *CityCafeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// 16-Bit Window Frame (Renders ON TOP of rain drops!)
	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoDarkGrey)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoDarkGrey)

	// Sleek 16-bit Bar Counter Table (Rows 66 to Height-1)
	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoPeach)

	// 16-Bit Focus Companion at Desk (Renders ON TOP of rain drops!)
	if buddy != nil {
		buddy.Draw(c, 60, 40, frameCount)
	}

	// Glowing 16-Bit Laptop on Bar Counter
	c.FillRect(88, 52, 28, 12, PicoDarkGrey)
	c.FillRect(92, 54, 20, 8, PicoCyan)
	for l := 0; l < 4; l++ {
		c.FillRect(94, 56+l*2, 6+l*3, 1, PicoLime)
	}
	c.FillRect(84, 64, 36, 4, PicoLightGrey)

	// Coffee Mug on Bar Counter
	c.FillRect(80, 58, 6, 6, PicoRed)
}

func (s *CityCafeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(92, 52, 32, 24, PicoYellow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 2. ForestCabinScene (16-Bit Forest Sanctuary Cabin)
// ---------------------------------------------------------
type ForestCabinScene struct{}

func (s *ForestCabinScene) ID() SceneType { return SceneForestCabin }
func (s *ForestCabinScene) Name() string { return "16-Bit Forest Cabin 🌲" }
func (s *ForestCabinScene) WindowBounds() Rect { return Rect{X: 24, Y: 10, W: 112, H: 54} }

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

	// Pine Trees
	for i := 0; i < 8; i++ {
		tx := wx + i*14 + 4
		th := 20 + (i*9)%18
		c.FillRect(tx, wy+wh-th, 10, th, PicoBlack)
		c.FillRect(tx+2, wy+wh-th-8, 6, 8, PicoGreen)
	}
}

func (s *ForestCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(0, 0, wx-2, 66, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(0, 0, c.Width, wy-2, PicoBrown, PicoDarkGrey)
	for y := 0; y < 66; y += 10 {
		c.FillRect(0, y, c.Width, 2, PicoBlack)
	}

	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoDarkGrey)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoBrown)
	c.FillRect(wx+ww/3, wy, 4, wh, PicoBrown)
	c.FillRect(wx+(ww*2)/3, wy, 4, wh, PicoBrown)

	if buddy != nil {
		buddy.Draw(c, 36, 40, frameCount)
	}

	renderFireplaceHelper(c, 108, 32, 36, 34, fireOn, frameCount)
	renderCatHelper(c, 76, 70, catBreed, frameCount)
}

func (s *ForestCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(126, 48, 28.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 3. ClassicCottageScene (16-Bit Classic Cottage)
// ---------------------------------------------------------
type ClassicCottageScene struct{}

func (s *ClassicCottageScene) ID() SceneType { return SceneCottageClassic }
func (s *ClassicCottageScene) Name() string { return "16-Bit Classic Cottage 🏡" }
func (s *ClassicCottageScene) WindowBounds() Rect { return Rect{X: 40, Y: 10, W: 80, H: 46} }

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
		c.FillRect(wx+48, wy+6, 8, 6, PicoWhite)
		c.FillRect(wx+46, wy+6, 4, 6, skyTop)
	} else if tod == TimeDay {
		c.FillRect(wx+48, wy+6, 8, 6, PicoYellow)
	}
}

func (s *ClassicCottageScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(0, 0, wx-2, 66, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(0, 0, c.Width, wy-2, PicoDarkPurple, PicoNavy)

	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoBlack)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoBrown)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoBrown)

	renderDeskHelper(c, 12, 58, 48, lampOn)

	if buddy != nil {
		buddy.Draw(c, 20, 38, frameCount)
	}

	renderFireplaceHelper(c, 108, 34, 36, 32, fireOn, frameCount)
	renderCatHelper(c, 74, 70, catBreed, frameCount)
}

func (s *ClassicCottageScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(20, 52, 32, 24, PicoYellow, 0.35)
	}
	if fireOn {
		c.ApplyRadialLight(126, 50, 28.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 4. ObservatoryScene (16-Bit Starlight Observatory)
// ---------------------------------------------------------
type ObservatoryScene struct{}

func (s *ObservatoryScene) ID() SceneType { return SceneObservatory }
func (s *ObservatoryScene) Name() string { return "16-Bit Starlight Observatory 🌌" }
func (s *ObservatoryScene) WindowBounds() Rect { return Rect{X: 20, Y: 8, W: 120, H: 54} }

func (s *ObservatoryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, PicoBlack, PicoNavy)

	c.FillRect(wx+72, wy+8, 12, 10, PicoWhite)
	c.FillRect(wx+68, wy+8, 8, 10, PicoBlack)

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

	c.FillRect(0, 0, wx-2, 66, PicoBlack)
	c.FillRect(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoBlack)
	c.FillRect(0, 0, c.Width, wy-2, PicoBlack)
	c.FillRect(0, wy+wh+2, c.Width, c.Height-(wy+wh+2), PicoBlack)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoDarkGrey)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoDarkGrey)
	c.FillRect(wx, wy+wh/2-2, ww, 4, PicoDarkGrey)

	if buddy != nil {
		buddy.Draw(c, 40, 40, frameCount)
	}

	c.FillRect(64, 68, 24, 4, PicoPeach)
	c.FillRect(72, 72, 8, 18, PicoDarkGrey)
}

func (s *ObservatoryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// Helper drawing functions
func renderDeskHelper(c *Canvas, dx, dy, dw int, lampOn bool) {
	c.FillRect(dx, dy, dw, 6, PicoBrown)
	c.FillRect(dx, dy, dw, 2, PicoPeach)
	c.FillRect(dx+4, dy+6, 6, 24, PicoBrown)

	if lampOn {
		c.FillRect(dx+6, dy-2, 6, 2, PicoYellow)
		c.FillRect(dx+8, dy-10, 2, 8, PicoYellow)
		c.FillRect(dx+4, dy-16, 10, 6, PicoGreen)
	} else {
		c.FillRect(dx+6, dy-2, 6, 2, PicoDarkGrey)
		c.FillRect(dx+8, dy-10, 2, 8, PicoDarkGrey)
		c.FillRect(dx+4, dy-16, 10, 6, PicoDarkGrey)
	}

	c.FillRect(dx+24, dy-6, 6, 6, PicoRed)
}

func renderFireplaceHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	c.FillRect(fx, fy, fw, fh, PicoBrown)
	c.FillRect(fx+4, fy+8, fw-8, fh-8, PicoBlack)

	if !fireOn {
		return
	}

	fireCenter := fx + fw/2
	fireTick := (frameCount / 3) % 4

	for i := -6; i <= 6; i++ {
		px := fireCenter + i
		flameH := int(8.0 + 4.0*math.Sin(float64(i)+float64(fireTick)*1.5))

		for fyPos := 0; fyPos < flameH; fyPos++ {
			py := fy + fh - 4 - fyPos
			if py < fy {
				continue
			}
			if fyPos < 3 {
				c.SetPixel(px, py, PicoYellow)
			} else if fyPos < 6 {
				c.SetPixel(px, py, PicoOrange)
			} else {
				c.SetPixel(px, py, PicoRed)
			}
		}
	}
}

func renderCatHelper(c *Canvas, cx, cy int, breed CatBreed, frameCount uint64) {
	c.FillRect(cx-4, cy+8, 24, 6, PicoDarkPurple)

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
		breathOffset = -2
	}

	c.FillRect(cx, cy+breathOffset, 16, 8, furMain)
	c.FillRect(cx-2, cy+2+breathOffset, 6, 6, furMain)

	if (frameCount/10)%3 == 1 {
		c.DrawGlyph(cx-4, (cy-4)/2, 'z', PicoCyan)
	}
}
