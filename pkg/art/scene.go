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

// ShaderEffect specifies dynamic procedural canvas animation shaders.
type ShaderEffect int

const (
	ShaderNone ShaderEffect = iota
	ShaderRainGlass
	ShaderFireflies
	ShaderAuroraWaves
	ShaderAudioWaveform
	ShaderStarfieldWarp
	ShaderFluidWaves
	ShaderFireFlames
	ShaderMatrixDepth
)

func (s ShaderEffect) String() string {
	switch s {
	case ShaderRainGlass:
		return "Rain-on-Glass Window Drops 🌧️"
	case ShaderFireflies:
		return "Bioluminescent Firefly Swarm 🌌"
	case ShaderAuroraWaves:
		return "Atmospheric Aurora Waves 🌅"
	case ShaderAudioWaveform:
		return "8-Bit Audio Waveform Spectrum 🌊"
	case ShaderStarfieldWarp:
		return "3D Cosmos Starfield Warp 🌌"
	case ShaderFluidWaves:
		return "Bioluminescent Fluid Waves 🌊"
	case ShaderFireFlames:
		return "Procedural Flame Engine 🔥"
	case ShaderMatrixDepth:
		return "3D Layered Matrix Code Rain 📟"
	default:
		return "None ❌"
	}
}

// AutoCycleMode specifies background scene auto-rotation interval.
type AutoCycleMode int

const (
	AutoCycleOff AutoCycleMode = iota
	AutoCycle1Min
	AutoCycle3Min
	AutoCycle5Min
	AutoCycle10Min
	AutoCycleTimerLinked
)

func (a AutoCycleMode) String() string {
	switch a {
	case AutoCycle1Min:
		return "1 Minute ⏱️"
	case AutoCycle3Min:
		return "3 Minutes ⏱️"
	case AutoCycle5Min:
		return "5 Minutes ⏱️"
	case AutoCycle10Min:
		return "10 Minutes ⏱️"
	case AutoCycleTimerLinked:
		return "Timer Linked ⏳"
	default:
		return "Off ❌"
	}
}

func (a AutoCycleMode) Duration() time.Duration {
	switch a {
	case AutoCycle1Min:
		return 1 * time.Minute
	case AutoCycle3Min:
		return 3 * time.Minute
	case AutoCycle5Min:
		return 5 * time.Minute
	case AutoCycle10Min:
		return 10 * time.Minute
	default:
		return 0
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

// CanvasRipple represents an expanding energy ring dropped onto the canvas.
type CanvasRipple struct {
	X, Y      float64
	Radius    float64
	MaxRadius float64
	Color     colorful.Color
	Life      float64
	MaxLife   float64
}

// CottageScene encapsulates dynamic state for the 8-bit retro scene engine.
type CottageScene struct {
	Canvas                *Canvas
	PrevCanvas            *Canvas
	TransitionStep        int
	TransitionMax         int
	CurrentScene          SceneType
	Buddy                 *FocusBuddy
	FrameCount            uint64
	TimeOfDay             TimeOfDay
	Weather               WeatherMode
	CatBreed              CatBreed
	ColorFilter           ColorFilter
	ParticleOverlay       ParticleOverlay
	ShaderEffect          ShaderEffect
	RenderMode            CanvasRenderMode
	AnimSpeed             AnimSpeed
	AutoCycleMode         AutoCycleMode
	AutoCycleFavsOnly     bool
	LastAutoCycleTime     time.Time
	LampOn                bool
	FireOn                bool
	FairyLightsOn         bool
	Lightning             bool
	BeatReactiveParticles bool
	SmoothDayNightBlend   bool
	VignetteOn            bool
	CRTScanlinesOn        bool
	CatBliss              int
	LastPetTime           time.Time

	RainParticles    []Particle
	SnowParticles    []Particle
	SteamParticles   []Particle
	OverlayParticles []Particle
	ShaderParticles  []Particle
	Ripples          []CanvasRipple
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
		ShaderEffect:          ShaderNone,
		RenderMode:            RenderHalfBlock,
		AnimSpeed:             SpeedNormal,
		AutoCycleMode:         AutoCycleOff,
		AutoCycleFavsOnly:     false,
		LastAutoCycleTime:     time.Now(),
		LampOn:                true,
		FireOn:                true,
		FairyLightsOn:         false,
		BeatReactiveParticles: true,
		SmoothDayNightBlend:   true,
		VignetteOn:            true,
		CRTScanlinesOn:        false,
		CatBliss:              85,
		LastPetTime:           time.Now(),
		rng:                   rng,
	}

	cs.initStars()
	cs.initParticles()
	cs.initShaderParticles()
	return cs
}

// SetSceneWithTransition switches scene with a smooth multi-frame dissolve crossfade.
func (cs *CottageScene) SetSceneWithTransition(newScene SceneType) {
	if cs.CurrentScene == newScene {
		return
	}
	if cs.Canvas != nil {
		cs.PrevCanvas = cs.Canvas.Clone()
		cs.TransitionStep = 0
		cs.TransitionMax = 6
	}
	cs.CurrentScene = newScene
}

// CycleAutoCycleMode cycles auto-rotation mode (Off -> 1m -> 3m -> 5m -> 10m -> TimerLinked -> Off).
func (cs *CottageScene) CycleAutoCycleMode() AutoCycleMode {
	cs.AutoCycleMode = (cs.AutoCycleMode + 1) % 6
	cs.LastAutoCycleTime = time.Now()
	return cs.AutoCycleMode
}

// ToggleAutoCycleFavs toggles whether auto-cycle selects only from favorites.
func (cs *CottageScene) ToggleAutoCycleFavs() bool {
	cs.AutoCycleFavsOnly = !cs.AutoCycleFavsOnly
	return cs.AutoCycleFavsOnly
}

// TriggerCanvasRipple drops a serene expanding energy ring on the canvas.
func (cs *CottageScene) TriggerCanvasRipple(x, y int, col colorful.Color) {
	if cs.Canvas == nil {
		return
	}
	if x < 0 || x >= cs.Canvas.Width || y < 0 || y >= cs.Canvas.Height {
		x = cs.Canvas.Width / 2
		y = cs.Canvas.Height / 2
	}
	cs.Ripples = append(cs.Ripples, CanvasRipple{
		X:         float64(x),
		Y:         float64(y),
		Radius:    1.0,
		MaxRadius: 14.0 + cs.rng.Float64()*8.0,
		Color:     col,
		Life:      0,
		MaxLife:   15,
	})
}

// TriggerCanvasRippleHex drops a serene expanding energy ring using a hex color string.
func (cs *CottageScene) TriggerCanvasRippleHex(x, y int, hex string) {
	col, err := colorful.Hex(hex)
	if err != nil {
		col = PicoCyan
	}
	cs.TriggerCanvasRipple(x, y, col)
}

// PetCat pets the cozy cat companion, boosting bliss and launching heart particles.
func (cs *CottageScene) PetCat() string {
	cs.CatBliss += 15
	if cs.CatBliss > 100 {
		cs.CatBliss = 100
	}
	cs.LastPetTime = time.Now()

	cx := 48
	cy := 44
	if cs.Canvas != nil {
		if cx >= cs.Canvas.Width {
			cx = cs.Canvas.Width - 8
		}
		if cy >= cs.Canvas.Height {
			cy = cs.Canvas.Height - 8
		}
	}

	cs.TriggerCanvasRippleHex(cx, cy, "#ff7597")

	heartRunes := []rune{'♥', '💖', '✨', '🐾', '💕'}
	for i := 0; i < 8; i++ {
		angle := cs.rng.Float64() * 2 * math.Pi
		speed := 0.5 + cs.rng.Float64()*1.2
		cs.OverlayParticles = append(cs.OverlayParticles, Particle{
			X:       float64(cx) + cs.rng.Float64()*10 - 5,
			Y:       float64(cy) + cs.rng.Float64()*4 - 2,
			Vx:      math.Cos(angle) * speed,
			Vy:      -0.6 - cs.rng.Float64()*0.8,
			Life:    0,
			MaxLife: 20 + cs.rng.Float64()*15,
			Char:    heartRunes[cs.rng.Intn(len(heartRunes))],
		})
	}

	catName := "Sparky"
	switch cs.CatBreed {
	case CatTuxedo:
		catName = "Sylvester"
	case CatCalico:
		catName = "Mochi"
	case CatBlack:
		catName = "Luna"
	}

	return fmt.Sprintf("🐾 %s purrs happily! Bliss: %d%% 💖", catName, cs.CatBliss)
}

// CycleShaderEffect cycles to the next procedural canvas shader animation.
func (cs *CottageScene) CycleShaderEffect() ShaderEffect {
	cs.ShaderEffect = (cs.ShaderEffect + 1) % 9
	cs.initShaderParticles()
	return cs.ShaderEffect
}

// CycleRenderMode cycles terminal art rendering style (HalfBlock -> ASCII -> Braille -> Phosphor -> HalfBlock).
func (cs *CottageScene) CycleRenderMode() CanvasRenderMode {
	cs.RenderMode = (cs.RenderMode + 1) % 4
	if cs.Canvas != nil {
		cs.Canvas.RenderMode = cs.RenderMode
	}
	return cs.RenderMode
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

// ToggleVignette toggles soft radial vignette canvas edge shading.
func (cs *CottageScene) ToggleVignette() bool {
	cs.VignetteOn = !cs.VignetteOn
	return cs.VignetteOn
}

// ToggleCRTScanlines toggles retro CRT monitor scanline dimming overlay.
func (cs *CottageScene) ToggleCRTScanlines() bool {
	cs.CRTScanlinesOn = !cs.CRTScanlinesOn
	return cs.CRTScanlinesOn
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
	cs.SetSceneWithTransition(SceneOrder[nextIdx])
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
	cs.SetSceneWithTransition(SceneOrder[prevIdx])
	return cs.CurrentScene
}

// RandomScene selects a random 8-bit retro scene preset.
func (cs *CottageScene) RandomScene() SceneType {
	if len(SceneOrder) == 0 {
		return cs.CurrentScene
	}
	nextIdx := cs.rng.Intn(len(SceneOrder))
	cs.SetSceneWithTransition(SceneOrder[nextIdx])
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

	if cs.FrameCount%300 == 0 && cs.CatBliss > 40 {
		cs.CatBliss--
	}

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
		mugX := 22.0
		if cs.CurrentScene == SceneRainyCity {
			mugX = 30.0
		} else if cs.CurrentScene == SceneForestCabin {
			mugX = 28.0
		}
		cs.SteamParticles = append(cs.SteamParticles, Particle{
			X:       mugX + (cs.rng.Float64()-0.5)*0.8,
			Y:       30.0,
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
		if st.Life >= st.MaxLife || st.Y < 22 {
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

	// Ripple Energy physics
	for i := len(cs.Ripples) - 1; i >= 0; i-- {
		rp := &cs.Ripples[i]
		rp.Life++
		rp.Radius += (rp.MaxRadius - rp.Radius) * 0.25
		if rp.Life >= rp.MaxLife {
			cs.Ripples = append(cs.Ripples[:i], cs.Ripples[i+1:]...)
		}
	}

	// Shader Particle physics
	if cs.ShaderEffect != ShaderNone {
		for i := range cs.ShaderParticles {
			p := &cs.ShaderParticles[i]
			p.Life += 0.1
			switch cs.ShaderEffect {
			case ShaderRainGlass:
				p.Y += 0.8 + cs.rng.Float64()*0.4
			case ShaderFireflies:
				p.Y += math.Sin(p.Life*0.4) * 0.3
				p.X += math.Cos(p.Life*0.4) * 0.3
			case ShaderMatrixDepth:
				p.Y += 1.2 + cs.rng.Float64()*0.6
			case ShaderStarfieldWarp:
				p.Life += 0.15
			}
			if p.Y >= float64(cs.Canvas.Height) || p.Y < 0 || p.X >= float64(cs.Canvas.Width) || p.X < 0 {
				p.X = float64(cs.rng.Intn(cs.Canvas.Width))
				p.Y = float64(cs.rng.Intn(cs.Canvas.Height))
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

func (cs *CottageScene) initShaderParticles() {
	cs.ShaderParticles = nil
	if cs.ShaderEffect == ShaderNone {
		return
	}
	count := 30
	if cs.ShaderEffect == ShaderRainGlass {
		count = 25
	} else if cs.ShaderEffect == ShaderFireflies {
		count = 20
	} else if cs.ShaderEffect == ShaderStarfieldWarp {
		count = 45
	} else if cs.ShaderEffect == ShaderMatrixDepth {
		count = 35
	}
	for i := 0; i < count; i++ {
		cs.ShaderParticles = append(cs.ShaderParticles, Particle{
			X:       float64(cs.rng.Intn(cs.Canvas.Width)),
			Y:       float64(cs.rng.Intn(cs.Canvas.Height)),
			Vx:      (cs.rng.Float64() - 0.5) * 0.2,
			Vy:      0.3 + cs.rng.Float64()*0.5,
			Life:    cs.rng.Float64() * 30,
			MaxLife: 30,
		})
	}
}

// RenderFrame executes the modular 6-step layered rendering pipeline:
// Step 1: Render Sky, Stars, Moon, Mountains/Buildings outside window (Layer 1)
// Step 2: Render Weather Particles (Rain/Snow) clipped strictly to WindowBounds (Layer 2)
// Step 3: Render Window Frames, Room Interior, Desk, Objects, and Focus Buddy (Layers 3-5)
// Step 4: Apply Dynamic Lighting Cones, Hearth Glows, and Lightning Flashes (Layer 6)
// Step 5: Render Organic Screen Particle Overlays (Sakura, Embers, Matrix Rain)
// Step 5.5: Render Procedural Canvas Shader Effects (Rain-on-Glass, Fireflies, Aurora, Starfield Warp, Fluid Waves, Flames, Matrix Depth)
// Step 5.6: Render Interactive Canvas Energy Ripples
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

	// Step 5.5: Procedural Canvas Shader Effects
	if cs.ShaderEffect != ShaderNone {
		cs.renderShaderEffect(c)
	}

	// Step 5.6: Interactive Canvas Energy Ripples
	if len(cs.Ripples) > 0 {
		cs.renderCanvasRipples(c)
	}

	// Step 5.7: Apply Smooth Multi-Frame Scene Dissolve Crossfade
	if cs.PrevCanvas != nil && cs.TransitionStep < cs.TransitionMax {
		alpha := 1.0 - (float64(cs.TransitionStep) / float64(cs.TransitionMax))
		c.BlendCanvas(cs.PrevCanvas, alpha)
		cs.TransitionStep++
		if cs.TransitionStep >= cs.TransitionMax {
			cs.PrevCanvas = nil
		}
	}

	// Step 6: Visual Color Filter Post-Processing
	if cs.ColorFilter != FilterNormal {
		cs.applyColorFilter(c)
	}

	// Step 7: Cinematic Radial Vignette Edge Shading
	if cs.VignetteOn {
		c.ApplyVignette(0.55)
	}

	// Step 8: Retro CRT Monitor Scanline Overlay
	if cs.CRTScanlinesOn {
		c.ApplyCRTScanlines(1.0)
	}
}

func (cs *CottageScene) renderCanvasRipples(c *Canvas) {
	for _, rp := range cs.Ripples {
		alpha := (1.0 - rp.Life/rp.MaxLife) * 0.85
		cx, cy := int(rp.X), int(rp.Y)
		r := rp.Radius
		minY := int(math.Max(0, float64(cy)-r))
		maxY := int(math.Min(float64(c.Height-1), float64(cy)+r))
		minX := int(math.Max(0, float64(cx)-r*2))
		maxX := int(math.Min(float64(c.Width-1), float64(cx)+r*2))

		for py := minY; py <= maxY; py++ {
			for px := minX; px <= maxX; px++ {
				dx := float64(px-cx) / 2.0
				dy := float64(py - cy)
				dist := math.Sqrt(dx*dx + dy*dy)
				if math.Abs(dist-r) < 1.2 {
					c.BlendPixel(px, py, rp.Color, alpha)
				}
			}
		}
	}
}

func (cs *CottageScene) renderShaderEffect(c *Canvas) {
	if cs.ShaderEffect == ShaderNone {
		return
	}
	switch cs.ShaderEffect {
	case ShaderRainGlass:
		for _, p := range cs.ShaderParticles {
			px, py := int(p.X), int(p.Y)
			if px >= 0 && px < c.Width && py >= 0 && py < c.Height {
				c.BlendPixel(px, py, PicoCyan, 0.75)
				if py > 0 && py < c.Height {
					c.BlendPixel(px, py-1, PicoNavy, 0.4)
				}
				c.DrawGlyph(px, py/2, '💧', PicoCyan)
			}
		}

	case ShaderFireflies:
		for i, p := range cs.ShaderParticles {
			px, py := int(p.X), int(p.Y)
			if px >= 0 && px < c.Width && py >= 0 && py < c.Height {
				phase := math.Sin(float64(cs.FrameCount)*0.1 + float64(i)*0.8)
				if phase > 0 {
					glowColor := PicoYellow
					if i%2 == 0 {
						glowColor = PicoGreen
					}
					c.BlendPixel(px, py, glowColor, phase*0.85)
					if phase > 0.5 {
						c.DrawGlyph(px, py/2, '✦', glowColor)
					}
				}
			}
		}

	case ShaderAuroraWaves:
		t := float64(cs.FrameCount) * 0.05
		for x := 0; x < c.Width; x++ {
			waveY := int(8.0 + math.Sin(t+float64(x)*0.08)*4.0 + math.Cos(t*0.5+float64(x)*0.04)*2.0)
			for dy := -2; dy <= 2; dy++ {
				py := waveY + dy
				if py >= 0 && py < c.Height/2 {
					col := PicoGreen
					if dy == -2 || dy == 2 {
						col = PicoPink
					} else if dy == -1 || dy == 1 {
						col = PicoCyan
					}
					c.BlendPixel(x, py, col, 0.35-math.Abs(float64(dy))*0.08)
				}
			}
		}

	case ShaderAudioWaveform:
		t := float64(cs.FrameCount) * 0.15
		baseY := c.Height - 6
		for x := 0; x < c.Width; x++ {
			h := int(math.Abs(math.Sin(t+float64(x)*0.15)*4.0) + math.Abs(math.Cos(t*0.8+float64(x)*0.25)*3.0))
			for dy := 0; dy <= h; dy++ {
				py := baseY - dy
				if py >= 0 && py < c.Height {
					col := PicoCyan
					if dy > 2 {
						col = PicoPink
					}
					c.BlendPixel(x, py, col, 0.7)
				}
			}
		}

	case ShaderStarfieldWarp:
		cx, cy := float64(c.Width)/2.0, float64(c.Height)/2.0
		for i := range cs.ShaderParticles {
			t := float64(cs.FrameCount)*0.08 + float64(i)*1.5
			dist := math.Mod(t*12.0, math.Max(cx, cy))
			angle := float64(i) * (math.Pi * 2 / float64(len(cs.ShaderParticles)))
			px := cx + math.Cos(angle)*dist*2.0
			py := cy + math.Sin(angle)*dist
			if px >= 0 && px < float64(c.Width) && py >= 0 && py < float64(c.Height) {
				alpha := math.Min(1.0, dist/15.0)
				col := PicoCyan
				if i%3 == 0 {
					col = PicoPink
				} else if i%3 == 1 {
					col = PicoYellow
				}
				c.BlendPixel(int(px), int(py), col, alpha*0.9)
				if dist > 15 {
					c.DrawGlyph(int(px), int(py)/2, '✦', col)
				}
			}
		}

	case ShaderFluidWaves:
		t := float64(cs.FrameCount) * 0.08
		baseY := c.Height - 14
		for x := 0; x < c.Width; x++ {
			wave1 := math.Sin(t+float64(x)*0.07) * 4.0
			wave2 := math.Cos(t*0.7+float64(x)*0.12) * 2.5
			yPeak := baseY + int(wave1+wave2)
			for py := yPeak; py < c.Height; py++ {
				if py >= 0 && py < c.Height {
					depth := float64(py-yPeak) / 14.0
					col := PicoCyan.BlendRgb(PicoNavy, depth)
					c.BlendPixel(x, py, col, 0.45*(1.0-depth*0.5))
				}
			}
		}

	case ShaderFireFlames:
		t := float64(cs.FrameCount) * 0.12
		baseY := c.Height - 1
		for x := 0; x < c.Width; x++ {
			flameH := int(math.Abs(math.Sin(t+float64(x)*0.2)*6.0) + math.Abs(math.Cos(t*1.3+float64(x)*0.35)*4.0))
			for dy := 0; dy <= flameH; dy++ {
				py := baseY - dy
				if py >= 0 && py < c.Height {
					ratio := float64(dy) / float64(flameH+1)
					col := PicoYellow
					if ratio > 0.6 {
						col = PicoRed
					} else if ratio > 0.3 {
						col = PicoOrange
					}
					c.BlendPixel(x, py, col, 0.75*(1.0-ratio*0.4))
				}
			}
		}

	case ShaderMatrixDepth:
		for _, p := range cs.ShaderParticles {
			px, py := int(p.X), int(p.Y)
			for trail := 0; trail < 6; trail++ {
				ty := py - trail
				if px >= 0 && px < c.Width && ty >= 0 && ty < c.Height {
					alpha := (1.0 - float64(trail)/6.0) * 0.85
					col := PicoGreen
					if trail == 0 {
						col = PicoWhite
					}
					c.BlendPixel(px, ty, col, alpha)
					if ty%2 == 0 {
						r := rune('0' + (px+ty+cs.rng.Intn(10))%10)
						c.DrawGlyph(px, ty/2, r, col)
					}
				}
			}
		}
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
	if cs.Canvas != nil {
		cs.Canvas.RenderMode = cs.RenderMode
	}

	proto := DetectTerminalProtocol()
	if proto == ProtocolKitty && cs.RenderMode == RenderHalfBlock {
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
