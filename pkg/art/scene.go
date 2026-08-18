package art

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

// TimeOfDay specifies the ambient sky/lighting mode.
type TimeOfDay int

const (
	TimeAuto TimeOfDay = iota
	TimeDawn
	TimeDay
	TimeSunset
	TimeNight
)

func (t TimeOfDay) String() string {
	switch t {
	case TimeDawn:
		return "Dawn 🌅"
	case TimeDay:
		return "Day ☀️"
	case TimeSunset:
		return "Sunset 🌇"
	case TimeNight:
		return "Night 🌙"
	default:
		return "Auto (Real Time) ⏰"
	}
}

// WeatherMode specifies the window weather effect.
type WeatherMode int

const (
	WeatherClear WeatherMode = iota
	WeatherRain
	WeatherSnow
	WeatherStorm
	WeatherFog
)

func (w WeatherMode) String() string {
	switch w {
	case WeatherClear:
		return "Clear 🌤️"
	case WeatherRain:
		return "Rain 🌧️"
	case WeatherSnow:
		return "Snow ❄️"
	case WeatherStorm:
		return "Thunderstorm 🌩️"
	case WeatherFog:
		return "Cozy Fog 🌫️"
	default:
		return "Clear 🌤️"
	}
}

// ColorFilter specifies visual color grading overlay filter.
type ColorFilter int

const (
	FilterNormal ColorFilter = iota
	FilterWarmAmber
	FilterCoolSynth
	FilterPhosphor
	FilterSepia
	FilterMidnight
)

func (f ColorFilter) String() string {
	switch f {
	case FilterWarmAmber:
		return "Warm Amber 🌅"
	case FilterCoolSynth:
		return "Cool Synth 🌌"
	case FilterPhosphor:
		return "CRT Phosphor 📟"
	case FilterSepia:
		return "Vintage Sepia 📜"
	case FilterMidnight:
		return "Midnight Blue 🌙"
	default:
		return "Natural 🎨"
	}
}

// ParticleOverlay specifies organic screen-wide particle animation overlays.
type ParticleOverlay int

const (
	OverlayNone ParticleOverlay = iota
	OverlaySakura
	OverlayEmbers
	OverlaySnow
	OverlayMatrix
	OverlayStars
)

func (p ParticleOverlay) String() string {
	switch p {
	case OverlaySakura:
		return "Sakura Petals 🌸"
	case OverlayEmbers:
		return "Hearth Embers 🔥"
	case OverlaySnow:
		return "Gentle Snow ❄️"
	case OverlayMatrix:
		return "Matrix Rain 📟"
	case OverlayStars:
		return "Starlight 🌟"
	default:
		return "None ❌"
	}
}

// AnimSpeed specifies scene animation playback speed.
type AnimSpeed int

const (
	SpeedNormal AnimSpeed = iota
	SpeedSlow
	SpeedFast
	SpeedPaused
)

func (s AnimSpeed) String() string {
	switch s {
	case SpeedSlow:
		return "Relaxing (0.5x) 🐢"
	case SpeedFast:
		return "Fast Focus (2.0x) ⚡"
	case SpeedPaused:
		return "Paused ⏸️"
	default:
		return "Normal (1.0x) 🎬"
	}
}

// CatBreed specifies cat color theme.
type CatBreed int

const (
	CatOrange CatBreed = iota
	CatTuxedo
	CatCalico
	CatBlack
)

func (c CatBreed) String() string {
	switch c {
	case CatOrange:
		return "Ginger Tabby 🐱"
	case CatTuxedo:
		return "Tuxedo 🐈"
	case CatCalico:
		return "Calico 🐱"
	case CatBlack:
		return "Midnight Black 🐈‍⬛"
	default:
		return "Ginger Tabby 🐱"
	}
}

// Particle represents weather drops or steam.
type Particle struct {
	X, Y    float64
	Vx, Vy  float64
	Life    float64
	MaxLife float64
	Char    rune
}

// CottageScene encapsulates dynamic state for the 8-bit retro scene engine.
type CottageScene struct {
	Canvas          *Canvas
	CurrentScene    SceneType
	Buddy           *FocusBuddy
	FrameCount      uint64
	TimeOfDay       TimeOfDay
	Weather         WeatherMode
	CatBreed        CatBreed
	ColorFilter     ColorFilter
	ParticleOverlay ParticleOverlay
	AnimSpeed       AnimSpeed
	LampOn          bool
	FireOn          bool
	FairyLightsOn         bool
	Lightning             bool
	BeatReactiveParticles bool
	SmoothDayNightBlend   bool

	RainParticles    []Particle
	SnowParticles    []Particle
	SteamParticles   []Particle
	OverlayParticles []Particle
	Stars            []struct{ X, Y int; Phase float64 }

	rng *rand.Rand
}

// NewCottageScene creates a new 8-bit retro cottage scene.
func NewCottageScene(width, height int) *CottageScene {
	canvas := NewCanvas(width, height)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	cs := &CottageScene{
		Canvas:                canvas,
		CurrentScene:          SceneRainyCity,
		Buddy:                 NewFocusBuddy(),
		TimeOfDay:             TimeAuto,
		Weather:               WeatherRain,
		CatBreed:              CatOrange,
		ColorFilter:           FilterNormal,
		ParticleOverlay:       OverlayNone,
		AnimSpeed:             SpeedNormal,
		LampOn:                true,
		FireOn:                true,
		FairyLightsOn:         false,
		BeatReactiveParticles: true,
		SmoothDayNightBlend:   true,
		rng:                   rng,
	}

	cs.initStars()
	cs.initParticles()
	return cs
}

// CycleAnimSpeed cycles animation playback speed (Normal -> Slow -> Fast -> Paused -> Normal).
func (cs *CottageScene) CycleAnimSpeed() AnimSpeed {
	cs.AnimSpeed = (cs.AnimSpeed + 1) % 4
	return cs.AnimSpeed
}

// CycleColorFilter cycles to the next color grading filter preset.
func (cs *CottageScene) CycleColorFilter() ColorFilter {
	cs.ColorFilter = (cs.ColorFilter + 1) % 6
	return cs.ColorFilter
}

// CycleParticleOverlay cycles to the next organic screen particle overlay mode.
func (cs *CottageScene) CycleParticleOverlay() ParticleOverlay {
	cs.ParticleOverlay = (cs.ParticleOverlay + 1) % 6
	cs.initOverlayParticles()
	return cs.ParticleOverlay
}

// ToggleBeatReactive toggles audio-reactive particle pulsing.
func (cs *CottageScene) ToggleBeatReactive() bool {
	cs.BeatReactiveParticles = !cs.BeatReactiveParticles
	return cs.BeatReactiveParticles
}

// ToggleSmoothDayNight toggles real-time smooth day/night sky blending.
func (cs *CottageScene) ToggleSmoothDayNight() bool {
	cs.SmoothDayNightBlend = !cs.SmoothDayNightBlend
	return cs.SmoothDayNightBlend
}

// CycleScene cycles to the next 8-bit retro scene preset.
func (cs *CottageScene) CycleScene() SceneType {
	if len(SceneOrder) == 0 {
		return cs.CurrentScene
	}
	idx := 0
	for i, id := range SceneOrder {
		if id == cs.CurrentScene {
			idx = i
			break
		}
	}
	nextIdx := (idx + 1) % len(SceneOrder)
	cs.CurrentScene = SceneOrder[nextIdx]
	return cs.CurrentScene
}

// PrevScene cycles to the previous 8-bit retro scene preset.
func (cs *CottageScene) PrevScene() SceneType {
	if len(SceneOrder) == 0 {
		return cs.CurrentScene
	}
	idx := 0
	for i, id := range SceneOrder {
		if id == cs.CurrentScene {
			idx = i
			break
		}
	}
	prevIdx := (idx + len(SceneOrder) - 1) % len(SceneOrder)
	cs.CurrentScene = SceneOrder[prevIdx]
	return cs.CurrentScene
}

// RandomScene selects a random 8-bit retro scene preset.
func (cs *CottageScene) RandomScene() SceneType {
	if len(SceneOrder) == 0 {
		return cs.CurrentScene
	}
	nextIdx := cs.rng.Intn(len(SceneOrder))
	cs.CurrentScene = SceneOrder[nextIdx]
	return cs.CurrentScene
}

func (cs *CottageScene) initStars() {
	cs.Stars = nil
	for i := 0; i < 20; i++ {
		cs.Stars = append(cs.Stars, struct{ X, Y int; Phase float64 }{
			X:     cs.rng.Intn(60),
			Y:     cs.rng.Intn(30),
			Phase: cs.rng.Float64() * math.Pi * 2,
		})
	}
}

func (cs *CottageScene) initParticles() {
	// Rain drops
	for i := 0; i < 24; i++ {
		cs.RainParticles = append(cs.RainParticles, Particle{
			X:    float64(10 + cs.rng.Intn(50)),
			Y:    float64(4 + cs.rng.Intn(30)),
			Vy:   1.2 + cs.rng.Float64()*0.4,
			Vx:   -0.2,
			Char: '\\',
		})
	}

	// Snow flakes
	for i := 0; i < 18; i++ {
		cs.SnowParticles = append(cs.SnowParticles, Particle{
			X:       float64(10 + cs.rng.Intn(50)),
			Y:       float64(4 + cs.rng.Intn(30)),
			Vy:      0.25 + cs.rng.Float64()*0.2,
			Vx:      -0.1 + cs.rng.Float64()*0.2,
			Life:    cs.rng.Float64() * 10,
			MaxLife: 10,
		})
	}
}

// EffectiveTimeOfDay returns resolved TimeOfDay.
func (cs *CottageScene) EffectiveTimeOfDay() TimeOfDay {
	if cs.TimeOfDay != TimeAuto {
		return cs.TimeOfDay
	}
	hour := time.Now().Hour()
	if hour >= 5 && hour < 8 {
		return TimeDawn
	} else if hour >= 8 && hour < 17 {
		return TimeDay
	} else if hour >= 17 && hour < 20 {
		return TimeSunset
	}
	return TimeNight
}

// Update advances the animation frame and updates scene-aware particle physics.
func (cs *CottageScene) Update() {
	if cs.AnimSpeed == SpeedPaused {
		return
	}
	if cs.AnimSpeed == SpeedSlow && cs.FrameCount%2 != 0 {
		cs.FrameCount++
		return
	}

	cs.FrameCount++

	if cs.Weather == WeatherStorm {
		if !cs.Lightning && cs.rng.Float64() < 0.02 {
			cs.Lightning = true
		} else if cs.Lightning && cs.rng.Float64() < 0.30 {
			cs.Lightning = false
		}
	} else {
		cs.Lightning = false
	}

	// Fetch current active scene window bounds
	wRect := cs.WindowRect()

	// Rain physics - strictly reset inside active window glass area
	for i := range cs.RainParticles {
		p := &cs.RainParticles[i]
		p.Y += p.Vy
		p.X += p.Vx
		if p.Y >= float64(wRect.Y+wRect.H) || p.X < float64(wRect.X) {
			p.Y = float64(wRect.Y) + cs.rng.Float64()*2.0
			p.X = float64(wRect.X) + float64(cs.rng.Intn(wRect.W))
		}
	}

	// Snow physics - strictly reset inside active window glass area
	for i := range cs.SnowParticles {
		p := &cs.SnowParticles[i]
		p.Y += p.Vy
		p.Life += 0.1
		p.X += math.Sin(p.Life)*0.3 + p.Vx
		if p.Y >= float64(wRect.Y+wRect.H) || p.X < float64(wRect.X) || p.X >= float64(wRect.X+wRect.W) {
			p.Y = float64(wRect.Y) + cs.rng.Float64()*2.0
			p.X = float64(wRect.X) + float64(cs.rng.Intn(wRect.W))
		}
	}

	// Mug Steam Physics
	if cs.rng.Float64() < 0.25 {
		cs.SteamParticles = append(cs.SteamParticles, Particle{
			X:       19.0 + (cs.rng.Float64()-0.5)*0.8,
			Y:       42.0,
			Vx:      (cs.rng.Float64() - 0.4) * 0.15,
			Vy:      -0.2 - cs.rng.Float64()*0.15,
			Life:    0,
			MaxLife: 10 + cs.rng.Float64()*6,
		})
	}

	for i := len(cs.SteamParticles) - 1; i >= 0; i-- {
		st := &cs.SteamParticles[i]
		st.Life++
		st.X += st.Vx
		st.Y += st.Vy
		if st.Life >= st.MaxLife || st.Y < 34 {
			cs.SteamParticles = append(cs.SteamParticles[:i], cs.SteamParticles[i+1:]...)
		}
	}

	// Overlay Particle physics
	if cs.ParticleOverlay != OverlayNone {
		for i := range cs.OverlayParticles {
			p := &cs.OverlayParticles[i]
			p.Life += 0.1
			switch cs.ParticleOverlay {
			case OverlaySakura:
				p.Y += 0.35 + cs.rng.Float64()*0.2
				p.X += math.Sin(p.Life)*0.4 + 0.1
			case OverlayEmbers:
				p.Y -= 0.35 + cs.rng.Float64()*0.2
				p.X += math.Sin(p.Life) * 0.3
			case OverlaySnow:
				p.Y += 0.2 + cs.rng.Float64()*0.15
				p.X += math.Sin(p.Life*0.5) * 0.2
			case OverlayMatrix:
				p.Y += 1.2 + cs.rng.Float64()*0.5
			case OverlayStars:
				p.Life += 0.05
			}
			if p.Y >= float64(cs.Canvas.Height) || p.Y < 0 || p.X >= float64(cs.Canvas.Width) || p.X < 0 {
				p.X = float64(cs.rng.Intn(cs.Canvas.Width))
				if cs.ParticleOverlay == OverlayEmbers {
					p.Y = float64(cs.Canvas.Height - 1)
				} else {
					p.Y = 0
				}
			}
		}
	}
}

func (cs *CottageScene) initOverlayParticles() {
	cs.OverlayParticles = nil
	if cs.ParticleOverlay == OverlayNone {
		return
	}
	count := 32
	for i := 0; i < count; i++ {
		cs.OverlayParticles = append(cs.OverlayParticles, Particle{
			X:       float64(cs.rng.Intn(cs.Canvas.Width)),
			Y:       float64(cs.rng.Intn(cs.Canvas.Height)),
			Vx:      (cs.rng.Float64() - 0.5) * 0.3,
			Vy:      0.2 + cs.rng.Float64()*0.4,
			Life:    cs.rng.Float64() * 20,
			MaxLife: 20,
		})
	}
}

// RenderFrame executes the modular 6-step layered rendering pipeline:
// Step 1: Render Sky, Stars, Moon, Mountains/Buildings outside window (Layer 1)
// Step 2: Render Weather Particles (Rain/Snow) clipped strictly to WindowBounds (Layer 2)
// Step 3: Render Window Frames, Room Interior, Desk, Objects, and Focus Buddy (Layers 3-5)
// Step 4: Apply Dynamic Lighting Cones, Hearth Glows, and Lightning Flashes (Layer 6)
// Step 5: Render Organic Screen Particle Overlays (Sakura, Embers, Matrix Rain)
// Step 6: Apply Visual Color Filter Post-Processing Tint
func (cs *CottageScene) RenderFrame() {
	c := cs.Canvas
	s := GetScene(cs.CurrentScene)

	// Step 1: Sky & Background (Layer 1)
	s.RenderSky(c, cs.EffectiveTimeOfDay(), cs.Lightning, cs.FrameCount, cs.Stars)

	// Step 2: Weather & Rain / Snow Particles (Layer 2 - Rendered BEHIND window frames, desk, and person!)
	cs.renderSceneWindowParticles(c, s.WindowBounds())

	// Step 3: Window Frames, Room Interior, Objects, and Focus Buddy Companion (Layers 3-5 - Rendered ON TOP of rain!)
	s.RenderForeground(c, cs.FrameCount, cs.Buddy, cs.LampOn, cs.FireOn, cs.CatBreed)

	// Step 4: Dynamic Lighting Cones, Hearth Glows, and Lightning Flashes (Layer 6)
	s.ApplyLighting(c, cs.LampOn, cs.FireOn, cs.Lightning)

	// Step 5: Organic Screen Particle Overlay
	cs.renderOverlayParticles(c)

	// Step 6: Visual Color Filter Post-Processing
	if cs.ColorFilter != FilterNormal {
		cs.applyColorFilter(c)
	}
}

func (cs *CottageScene) renderOverlayParticles(c *Canvas) {
	if cs.ParticleOverlay == OverlayNone {
		return
	}
	for _, p := range cs.OverlayParticles {
		px, py := int(p.X), int(p.Y)
		if px < 0 || px >= c.Width || py < 0 || py >= c.Height {
			continue
		}
		switch cs.ParticleOverlay {
		case OverlaySakura:
			c.BlendPixel(px, py, PicoPink, 0.85)
			c.DrawGlyph(px, py/2, '🌸', PicoPink)
		case OverlayEmbers:
			c.BlendPixel(px, py, PicoOrange, 0.9)
			c.DrawGlyph(px, py/2, '•', PicoOrange)
		case OverlaySnow:
			c.BlendPixel(px, py, PicoWhite, 0.8)
			c.DrawGlyph(px, py/2, '•', PicoWhite)
		case OverlayMatrix:
			c.BlendPixel(px, py, PicoGreen, 0.9)
			r := rune('0' + cs.rng.Intn(10))
			c.DrawGlyph(px, py/2, r, PicoGreen)
		case OverlayStars:
			phase := math.Sin(p.Life)
			if phase > 0.3 {
				c.BlendPixel(px, py, PicoYellow, phase*0.8)
				c.DrawGlyph(px, py/2, '✦', PicoYellow)
			}
		}
	}
}

func (cs *CottageScene) applyColorFilter(c *Canvas) {
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			col := c.Pixels[y][x]
			r, g, b := col.R, col.G, col.B
			switch cs.ColorFilter {
			case FilterWarmAmber:
				r = math.Min(1.0, r*1.2)
				g = math.Min(1.0, g*0.95)
				b = math.Min(1.0, b*0.7)
			case FilterCoolSynth:
				r = math.Min(1.0, r*0.85+b*0.25)
				g = math.Min(1.0, g*0.8)
				b = math.Min(1.0, b*1.25)
			case FilterPhosphor:
				lum := 0.299*r + 0.587*g + 0.114*b
				r = lum * 0.15
				g = lum * 1.15
				b = lum * 0.15
			case FilterSepia:
				lum := 0.299*r + 0.587*g + 0.114*b
				r = math.Min(1.0, lum*1.2)
				g = math.Min(1.0, lum*1.0)
				b = math.Min(1.0, lum*0.75)
			case FilterMidnight:
				r = math.Min(1.0, r*0.45)
				g = math.Min(1.0, g*0.65)
				b = math.Min(1.0, b*1.3)
			}
			c.Pixels[y][x] = colorful.Color{
				R: math.Max(0, math.Min(1.0, r)),
				G: math.Max(0, math.Min(1.0, g)),
				B: math.Max(0, math.Min(1.0, b)),
			}
		}
	}
}

func (cs *CottageScene) renderSceneWindowParticles(c *Canvas, wRect Rect) {
	// 1. Rain Drops - Masked strictly inside window glass rect
	if cs.Weather == WeatherRain || cs.Weather == WeatherStorm {
		for _, p := range cs.RainParticles {
			px, py := int(p.X), int(p.Y)
			if px >= wRect.X && px < wRect.X+wRect.W && py >= wRect.Y && py < wRect.Y+wRect.H {
				c.BlendPixel(px, py, PicoCyan, 0.75)
				if cs.FrameCount%3 == 0 {
					c.DrawGlyph(px, py/2, '\\', PicoCyan)
				}
			}
		}
	} else if cs.Weather == WeatherSnow {
		for _, p := range cs.SnowParticles {
			px, py := int(p.X), int(p.Y)
			if px >= wRect.X && px < wRect.X+wRect.W && py >= wRect.Y && py < wRect.Y+wRect.H {
				c.BlendPixel(px, py, PicoWhite, 0.85)
				c.DrawGlyph(px, py/2, '•', PicoWhite)
			}
		}
	} else if cs.Weather == WeatherFog {
		c.BlendRect(wRect.X, wRect.Y+wRect.H-12, wRect.W, 12, PicoLightGrey, 0.25)
	}

	// 2. Mug Steam Particles
	for _, st := range cs.SteamParticles {
		px, py := int(st.X), int(st.Y)
		if px >= 0 && px < c.Width && py >= 0 && py < c.Height {
			alpha := (1.0 - st.Life/st.MaxLife) * 0.65
			c.BlendPixel(px, py, PicoWhite, alpha)
			if st.Life < 6 {
				c.DrawGlyph(px, py/2, '~', PicoWhite)
			}
		}
	}
}

// Render returns ANSI string or Kitty GPU graphics payload based on terminal capabilities.
func (cs *CottageScene) Render() string {
	cs.RenderFrame()

	proto := DetectTerminalProtocol()
	if proto == ProtocolKitty {
		return EncodeKittyGPU(cs.Canvas, 74, 34)
	}

	return cs.Canvas.Render()
}

// StatusSummary returns art status summary.
func (cs *CottageScene) StatusSummary() string {
	return fmt.Sprintf("%s | %s | Cat: %s | Lamp: %s | Fire: %s",
		cs.EffectiveTimeOfDay().String(),
		cs.Weather.String(),
		cs.CatBreed.String(),
		onOff(cs.LampOn),
		onOff(cs.FireOn),
	)
}

func onOff(v bool) string {
	if v {
		return "ON ✨"
	}
	return "OFF 💤"
}
