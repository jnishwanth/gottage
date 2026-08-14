package art

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// SceneType identifies active scene visual layout.
type SceneType int

const (
	SceneRainyCity SceneType = iota
	SceneForestCabin
	SceneCottageClassic
	SceneObservatory
	SceneKirokazeRoom
	SceneBalconyLofi
	SceneRainSetup
	SceneKirokazeCity
	SceneHomeSweetHome
	SceneWaneellaStreet
	SceneSunsetCity
	SceneSnowCabin
	SceneCyberpunkCity
	SceneCozyLibrary
	ScenePixelTrain
	SceneOceanWaves
	SceneStargazing
	SceneCoffeeShop
	SceneAutumnPark
	SceneMidnightOwl
	SceneRiverBridge
	SceneNeonRunner
	SceneMagicSanctuary
	SceneCyberBot
	SceneSakuraShrine
	SceneCozyCampfire
	SceneSpaceStation
	SceneRamenShop
	SceneDeepSea
	SceneCyberArcade
	SceneNeonBistro
	SceneSunsetStation
	SceneMountainView
	ScenePixelBakery
	SceneEnchantedForest
	SceneSkyIsland
	SceneBambooGrove
	SceneLighthouseCliff
	SceneVolcanoForge
	SceneCrystalCave
	SceneDesertOasis
	SceneSteampunkLab
	SceneAlienPlanet
	SceneHauntedManor
	SceneCloudTemple
	SceneCyberBar
	SceneCoralReef
	SceneFrostPeak
	SceneCyberDojo
	SceneMushroomForest
	SceneRetroArcade
	ScenePirateCove
	SceneCozyGreenhouse
	SceneCosmicObservatory
	SceneCyberRooftop
	SceneCozyIgloo
	SceneAtlantisTemple
	SceneCyberMatrix
	SceneSunsetBeach
	SceneFairytaleWindmill
	SceneCyberExpress
	SceneBiolumSwamp
	SceneClockworkTower
	SceneStarlightCamp
	SceneSpaceNebula
	SceneCyberMarket
	SceneCyberDiner
	SceneVolcanoView
	SceneCrystalLagoon
	SceneSteampunkAirship
	SceneCozyAttic
	SceneRetroArcadeHall
)

// Rect represents a rectangular boundary for scene-aware masking.
type Rect struct {
	X, Y, W, H int
}

func (s SceneType) String() string {
	if sc, ok := SceneRegistry[s]; ok {
		return sc.Name()
	}
	switch s {
	case SceneRainyCity:
		return "8-Bit Rainy City Cafe 🏙️"
	case SceneForestCabin:
		return "8-Bit Forest Cabin 🌲"
	case SceneCottageClassic:
		return "8-Bit Classic Cottage 🏡"
	case SceneObservatory:
		return "8-Bit Starlight Observatory 🌌"
	default:
		return "8-Bit Pixel Animation 🎨"
	}
}

// WindowRect returns the exact window glass boundaries for scene-aware particle masking!
func (cs *CottageScene) WindowRect() Rect {
	return GetScene(cs.CurrentScene).WindowBounds()
}

// RenderRainyCity renders a native 8-bit retro lofi city cafe with animated body-double companion.
func (cs *CottageScene) RenderRainyCity(c *Canvas) {
	// Dark 8-bit interior wall (PICO-8 Navy & Black)
	c.DrawVerticalGradient(0, 0, c.Width, 46, PicoNavy, PicoBlack)

	// Sleek 8-bit bar counter table (Rows 46 to Height-1)
	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoPeach)

	// High-Rise City Window (X: 10..64, Y: 4..42)
	r := cs.WindowRect()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	tod := cs.EffectiveTimeOfDay()
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

	if cs.Lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// 8-Bit Skyscrapers Silhouette
	bldgColors := []colorful.Color{PicoBlack, PicoNavy, PicoDarkPurple}
	neonColors := []colorful.Color{PicoPink, PicoCyan, PicoYellow, PicoLime}

	for b := 0; b < 9; b++ {
		bx := wx + b*6 + 1
		bw := 5
		bh := 14 + (b*7)%16
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		// Animated 8-Bit Neon Window Lights
		if tod == TimeNight || tod == TimeSunset {
			neonCol := neonColors[(b+int(cs.FrameCount/15))%4]
			for wyPos := wy + wh - bh + 2; wyPos < wy+wh-2; wyPos += 3 {
				for wxPos := bx + 1; wxPos < bx+bw-1; wxPos += 2 {
					if (wxPos+wyPos)%2 == 0 {
						c.SetPixel(wxPos, wyPos, neonCol)
					}
				}
			}
		}
	}

	// Animated 8-Bit Neon Signs (LO-FI)
	if tod == TimeNight || tod == TimeSunset {
		neonSignCol := neonColors[int(cs.FrameCount/10)%4]
		c.DrawGlyph(wx+4, wy+6, 'L', neonSignCol)
		c.DrawGlyph(wx+5, wy+6, 'O', neonSignCol)
		c.DrawGlyph(wx+6, wy+6, '-', neonSignCol)
		c.DrawGlyph(wx+7, wy+6, 'F', neonSignCol)
		c.DrawGlyph(wx+8, wy+6, 'I', neonSignCol)
	}

	// Crisp 8-Bit Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoDarkGrey)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoDarkGrey)

	// 8-Bit Body-Double Companion at Desk (X: 30, Y: 34)
	cs.Buddy.Draw(c, 30, 34, cs.FrameCount)

	// Glowing 8-Bit Laptop on Bar Counter (X: 42, Y: 42)
	c.FillRect(42, 40, 14, 6, PicoDarkGrey)
	c.FillRect(44, 41, 10, 4, PicoCyan)
	for l := 0; l < 3; l++ {
		c.FillRect(45, 42+l, 3+l*2, 1, PicoLime)
	}
	c.FillRect(40, 46, 18, 2, PicoLightGrey)

	// Coffee Mug on Bar Counter (X: 40, Y: 43)
	c.FillRect(40, 43, 3, 3, PicoRed)
}

// RenderForestCabin renders a native 8-bit retro cabin with body-double companion.
func (cs *CottageScene) RenderForestCabin(c *Canvas) {
	c.DrawVerticalGradient(0, 0, c.Width, 46, PicoBrown, PicoDarkGrey)
	for y := 0; y < 46; y += 6 {
		c.FillRect(0, y, c.Width, 1, PicoBlack)
	}

	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoDarkGrey)

	r := cs.WindowRect()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	tod := cs.EffectiveTimeOfDay()
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

	if cs.Lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Pine Trees
	for i := 0; i < 7; i++ {
		tx := wx + i*7 + 2
		th := 14 + (i*5)%10
		c.FillRect(tx, wy+wh-th, 5, th, PicoBlack)
		c.FillRect(tx+1, wy+wh-th-4, 3, 4, PicoGreen)
	}

	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoBrown)
	c.FillRect(wx+ww/3, wy, 2, wh, PicoBrown)
	c.FillRect(wx+(ww*2)/3, wy, 2, wh, PicoBrown)

	// 8-Bit Body-Double Companion at Bench Desk (X: 18, Y: 34)
	cs.Buddy.Draw(c, 18, 34, cs.FrameCount)

	// 8-Bit Stone Fireplace (Right side: X: 52..68)
	cs.render8BitFireplace(c, 52, 24, 18, 22)

	// 8-Bit Sleeping Cat on Rug (X: 38..48)
	cs.render8BitCat(c, 38, 50)
}

// RenderCottageClassic renders classic 8-bit cottage with body-double companion.
func (cs *CottageScene) RenderCottageClassic(c *Canvas) {
	c.DrawVerticalGradient(0, 0, c.Width, 46, PicoDarkPurple, PicoNavy)
	c.FillRect(0, 46, c.Width, c.Height-46, PicoBrown)
	c.FillRect(0, 46, c.Width, 1, PicoBlack)

	r := cs.WindowRect()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	tod := cs.EffectiveTimeOfDay()
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

	if cs.Lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	if tod == TimeNight && !cs.Lightning {
		c.FillRect(wx+24, wy+3, 4, 3, PicoWhite)
		c.FillRect(wx+23, wy+3, 2, 3, skyTop)
	} else if tod == TimeDay {
		c.FillRect(wx+24, wy+3, 4, 3, PicoYellow)
	}

	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoBrown)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoBrown)

	// 8-Bit Desk Surface (Left)
	cs.render8BitDesk(c, 6, 42, 28)

	// 8-Bit Body-Double Companion at Cottage Desk (X: 10, Y: 32)
	cs.Buddy.Draw(c, 10, 32, cs.FrameCount)

	// 8-Bit Fireplace (Right)
	cs.render8BitFireplace(c, 52, 26, 18, 20)

	// 8-Bit Cat on Cushion
	cs.render8BitCat(c, 36, 50)
}

// RenderObservatory renders native 8-bit celestial observatory with body-double companion.
func (cs *CottageScene) RenderObservatory(c *Canvas) {
	c.Clear(PicoBlack)

	r := cs.WindowRect()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, PicoBlack, PicoNavy)

	c.FillRect(wx+36, wy+4, 6, 5, PicoWhite)
	c.FillRect(wx+34, wy+4, 4, 5, PicoBlack)

	for i, st := range cs.Stars {
		px := wx + (st.X % ww)
		py := wy + (st.Y % wh)
		twinkle := math.Sin(float64(cs.FrameCount)*0.1 + st.Phase)
		if twinkle > 0.2 {
			c.SetPixel(px, py, PicoWhite)
			if i%3 == 0 {
				c.DrawGlyph(px, py/2, '✦', PicoCyan)
			}
		}
	}

	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, PicoDarkGrey)
	c.FillRect(wx+ww/2-1, wy, 2, wh, PicoDarkGrey)
	c.FillRect(wx, wy+wh/2-1, ww, 2, PicoDarkGrey)

	// 8-Bit Body-Double Companion at Telescope Desk (X: 20, Y: 34)
	cs.Buddy.Draw(c, 20, 34, cs.FrameCount)

	// Telescope
	c.FillRect(32, 48, 12, 2, PicoPeach)
	c.FillRect(36, 50, 4, 10, PicoDarkGrey)
}

func (cs *CottageScene) render8BitDesk(c *Canvas, dx, dy, dw int) {
	c.FillRect(dx, dy, dw, 3, PicoBrown)
	c.FillRect(dx, dy, dw, 1, PicoPeach)
	c.FillRect(dx+2, dy+3, 3, 15, PicoBrown)

	if cs.LampOn {
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

func (cs *CottageScene) render8BitFireplace(c *Canvas, fx, fy, fw, fh int) {
	c.FillRect(fx, fy, fw, fh, PicoBrown)
	c.FillRect(fx+2, fy+4, fw-4, fh-4, PicoBlack)

	if !cs.FireOn {
		return
	}

	fireCenter := fx + fw/2
	fireTick := (cs.FrameCount / 3) % 4

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

func (cs *CottageScene) render8BitCat(c *Canvas, cx, cy int) {
	c.FillRect(cx-2, cy+4, 12, 3, PicoDarkPurple)

	var furMain colorful.Color
	switch cs.CatBreed {
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
	if (cs.FrameCount/12)%2 == 0 {
		breathOffset = -1
	}

	c.FillRect(cx, cy+breathOffset, 8, 4, furMain)
	c.FillRect(cx-1, cy+1+breathOffset, 3, 3, furMain)

	if (cs.FrameCount/10)%3 == 1 {
		c.DrawGlyph(cx-2, (cy-2)/2, 'z', PicoCyan)
	}
}
