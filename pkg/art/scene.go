package art

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	Canvas        *Canvas
	CurrentScene  SceneType
	Buddy         *FocusBuddy
	FrameCount    uint64
	TimeOfDay     TimeOfDay
	Weather       WeatherMode
	CatBreed      CatBreed
	LampOn        bool
	FireOn        bool
	FairyLightsOn bool
	Lightning     bool

	RainParticles  []Particle
	SnowParticles  []Particle
	SteamParticles []Particle
	Stars          []struct{ X, Y int; Phase float64 }

	rng *rand.Rand
}

// NewCottageScene creates a new 8-bit retro cottage scene.
func NewCottageScene(width, height int) *CottageScene {
	canvas := NewCanvas(width, height)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	cs := &CottageScene{
		Canvas:        canvas,
		CurrentScene:  SceneRainyCity,
		Buddy:         NewFocusBuddy(),
		TimeOfDay:     TimeAuto,
		Weather:       WeatherRain,
		CatBreed:      CatOrange,
		LampOn:        true,
		FireOn:        true,
		FairyLightsOn: false,
		rng:           rng,
	}

	cs.initStars()
	cs.initParticles()
	return cs
}

// CycleScene cycles to the next 8-bit retro scene preset.
func (cs *CottageScene) CycleScene() SceneType {
	cs.CurrentScene = (cs.CurrentScene + 1) % 4
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
}

// RenderFrame renders active 8-bit retro scene with scene-aware particle masking.
func (cs *CottageScene) RenderFrame() {
	c := cs.Canvas

	switch cs.CurrentScene {
	case SceneRainyCity:
		cs.RenderRainyCity(c)
	case SceneForestCabin:
		cs.RenderForestCabin(c)
	case SceneCottageClassic:
		cs.RenderCottageClassic(c)
	case SceneObservatory:
		cs.RenderObservatory(c)
	default:
		cs.RenderRainyCity(c)
	}

	// Render Scene-Aware Window Particles (Clipped strictly to window glass!)
	cs.renderSceneWindowParticles(c)
}

func (cs *CottageScene) renderSceneWindowParticles(c *Canvas) {
	wRect := cs.WindowRect()

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

	// 3. Desk Lamp Cone Light Glow
	if cs.LampOn {
		c.ApplyLightCone(10, 36, 18, 14, PicoYellow, 0.35)
	}

	// 4. Lightning Flash (Thunderstorm mode)
	if cs.Lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// Render returns ANSI character string of the rendered frame.
func (cs *CottageScene) Render() string {
	cs.RenderFrame()
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
