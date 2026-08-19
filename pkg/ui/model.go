package ui

import (
	"fmt"
	"strings"
	"time"

	"gottage/pkg/art"
	"gottage/pkg/audio"
	"gottage/pkg/features"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// ActiveTab specifies the main focus view tab.
type ActiveTab int

const (
	TabArt ActiveTab = iota
	TabTimer
	TabTasks
	TabAmbiance
)

func (t ActiveTab) String() string {
	switch t {
	case TabArt:
		return "🏡 Cottage Art"
	case TabTimer:
		return "⏱️ Focus Timer"
	case TabTasks:
		return "📝 Task List"
	case TabAmbiance:
		return "🌿 Ambiance & Audio"
	default:
		return "Cottage"
	}
}

// Tick Messages
type tickMsg time.Time
type cliampPollMsg audio.CLIAMPStatus
type timerTickMsg time.Time

// Model is the main Bubbletea model.
type Model struct {
	Width              int
	Height             int
	Styles             Styles
	ActiveTab          ActiveTab
	ShowHelp           bool
	ShowScenePicker    bool
	PickerCategoryIdx  int
	PickerSceneIdx     int
	PickerSearchMode   bool
	PickerSearchInput  textinput.Model
	ShowThemePicker    bool
	PickerThemeIdx     int
	ShowSceneInfoModal bool
	ShowCommandPalette bool
	CommandInput       textinput.Model
	CommandIdx         int
	ShowPresetModal    bool
	PresetIdx          int
	PresetSaving       bool
	PresetInputText    textinput.Model
	ShowSoundMixerModal bool
	MixerChannelIdx    int
	ShowSpectrumModal  bool
	SpectrumMode       SpectrumStyle
	ZenMode            bool
	AutoSyncAudio      bool
	StatusMsg          string
	StatusMsgTime      time.Time
	UseGPURender       bool
	FavoriteScenes     map[art.SceneType]bool

	// Core Components
	Scene       *art.CottageScene
	CLIAMP      *audio.CLIAMPController
	Ambient     *audio.AmbientManager
	Timer       *features.FocusTimer
	Tasks       *features.TaskList
	Presets     *features.PresetManager
	Inspiration *features.InspirationManager

	// Sub-view State
	AddingTask bool
	TextInput  textinput.Model
}

// NewModel creates initial Bubbletea model.
func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter new focus task..."
	ti.CharLimit = 60
	ti.Width = 40

	si := textinput.New()
	si.Placeholder = "Type keyword to filter 370+ scenes..."
	si.CharLimit = 30
	si.Width = 36

	ci := textinput.New()
	ci.Placeholder = "Type command, tab, theme, filter, or mood..."
	ci.CharLimit = 40
	ci.Width = 44

	pi := textinput.New()
	pi.Placeholder = "Enter preset name (e.g. My Study Haven)..."
	pi.CharLimit = 32
	pi.Width = 36

	styles := NewStyles(ThemeCozyAutumn)
	scene := art.NewCottageScene(74, 52)

	favs := make(map[art.SceneType]bool)
	favs[art.SceneRainyCity] = true

	return Model{
		Width:              120,
		Height:             40,
		Styles:             styles,
		ActiveTab:          TabArt,
		ShowHelp:           false,
		ShowScenePicker:    false,
		PickerSearchInput:  si,
		ShowCommandPalette: false,
		CommandInput:       ci,
		ShowThemePicker:    false,
		ShowPresetModal:    false,
		PresetInputText:    pi,
		SpectrumMode:       SpectrumBars,
		AutoSyncAudio:      true,
		UseGPURender:       art.DetectTerminalProtocol() == art.ProtocolKitty,
		FavoriteScenes:     favs,

		Scene:       scene,
		CLIAMP:      audio.NewCLIAMPController(),
		Ambient:     audio.NewAmbientManager(),
		Timer:       features.NewFocusTimer(),
		Tasks:       features.NewTaskList(),
		Presets:     features.NewPresetManager("user_presets.json"),
		Inspiration: features.NewInspirationManager(),

		AddingTask: false,
		TextInput:  ti,
	}
}

// ToggleFavorite toggles bookmark status for a scene.
func (m *Model) ToggleFavorite(st art.SceneType) bool {
	if m.FavoriteScenes == nil {
		m.FavoriteScenes = make(map[art.SceneType]bool)
	}
	m.FavoriteScenes[st] = !m.FavoriteScenes[st]
	return m.FavoriteScenes[st]
}

// ApplyPreset applies a complete atmosphere preset to Model state.
func (m *Model) ApplyPreset(p features.AtmospherePreset) {
	m.Scene.CurrentScene = p.Scene
	m.Scene.ColorFilter = p.ColorFilter
	m.Scene.ParticleOverlay = p.ParticleOverlay
	m.Scene.ShaderEffect = p.ShaderEffect
	m.Scene.AnimSpeed = p.AnimSpeed

	for _, theme := range AllThemes {
		if strings.EqualFold(theme.Name, p.Theme) {
			m.Styles = NewStyles(theme)
			break
		}
	}

	if p.RainAudio && !m.Ambient.RainActive {
		m.Ambient.ToggleRain()
	} else if !p.RainAudio && m.Ambient.RainActive {
		m.Ambient.ToggleRain()
	}

	if p.FireAudio && !m.Ambient.FireActive {
		m.Ambient.ToggleFire()
	} else if !p.FireAudio && m.Ambient.FireActive {
		m.Ambient.ToggleFire()
	}

	if p.WindAudio && !m.Ambient.WindActive {
		m.Ambient.ToggleWind()
	} else if !p.WindAudio && m.Ambient.WindActive {
		m.Ambient.ToggleWind()
	}

	m.SetStatusMsg(fmt.Sprintf("Activated Preset: %s ✨", p.Name))
}

// Init initializes Bubbletea commands (animation ticks & audio polling).
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.animateTick(),
		m.pollCLIAMP(),
		m.timerTick(),
	)
}

func (m Model) animateTick() tea.Cmd {
	return tea.Tick(60*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) pollCLIAMP() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		if m.CLIAMP != nil {
			return cliampPollMsg(m.CLIAMP.PollStatus())
		}
		return cliampPollMsg{}
	})
}

func (m Model) timerTick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return timerTickMsg(t)
	})
}

// SetStatusMsg sets a temporary toast message.
func (m *Model) SetStatusMsg(msg string) {
	m.StatusMsg = msg
	m.StatusMsgTime = time.Now()
}

// Update handles messages and keybindings.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tickMsg:
		// Advance bit-art scene animation frame
		m.Scene.Update()

		// Auto-Cycle Scene Timer Evaluation
		if m.Scene.AutoCycleMode >= art.AutoCycle1Min && m.Scene.AutoCycleMode <= art.AutoCycle10Min {
			dur := m.Scene.AutoCycleMode.Duration()
			if dur > 0 && time.Since(m.Scene.LastAutoCycleTime) >= dur {
				m.Scene.LastAutoCycleTime = time.Now()
				if m.Scene.AutoCycleFavsOnly && len(m.FavoriteScenes) > 0 {
					var favs []art.SceneType
					for st, ok := range m.FavoriteScenes {
						if ok {
							favs = append(favs, st)
						}
					}
					if len(favs) > 0 {
						nextFav := favs[int(time.Now().Unix())%len(favs)]
						m.Scene.SetSceneWithTransition(nextFav)
					} else {
						m.Scene.RandomScene()
					}
				} else {
					m.Scene.RandomScene()
				}
				m.SetStatusMsg("🎬 Auto-Cycle: " + m.Scene.CurrentScene.String() + " ✨")
			}
		}

		cmds = append(cmds, m.animateTick())

	case cliampPollMsg:
		cmds = append(cmds, m.pollCLIAMP())

	case timerTickMsg:
		if m.Timer.Tick() {
			if m.Scene.AutoCycleMode == art.AutoCycleTimerLinked {
				if m.Timer.Mode == features.ModeShortBreak || m.Timer.Mode == features.ModeLongBreak {
					m.Scene.SetSceneWithTransition(art.SceneCozyCampfire)
					if m.AutoSyncAudio {
						m.Ambient.RainActive = true
						m.Ambient.FireActive = true
					}
					m.SetStatusMsg("☕ Break Time! Switched to Cozy Hearth Campfire ☕")
				} else {
					m.Scene.SetSceneWithTransition(art.SceneCozyLibrary)
					m.SetStatusMsg("📖 Focus Session Started! Switched to Cozy Library 📖")
				}
			} else {
				m.SetStatusMsg("🎉 Timer completed! Time for a cozy break!")
			}
		}
		cmds = append(cmds, m.timerTick())

	case tea.MouseMsg:
		if msg.Button == tea.MouseButtonLeft && msg.Action == tea.MouseActionPress {
			if msg.Y <= 2 {
				// Header tabs clicked
				if msg.X > m.Width/2 {
					tabSpan := 12
					tabIdx := (msg.X - (m.Width / 2)) / tabSpan
					if tabIdx >= 0 && tabIdx < 4 {
						m.ActiveTab = ActiveTab(tabIdx)
						m.SetStatusMsg("Switched Tab: " + m.ActiveTab.String())
					}
				}
			} else if msg.X > m.mainWidth() {
				// Sidebar cards clicked
				if msg.Y >= 3 && msg.Y <= 10 {
					if m.CLIAMP != nil {
						_ = m.CLIAMP.TogglePlayPause()
						st := m.CLIAMP.GetStatus()
						if st.Playing {
							m.SetStatusMsg(fmt.Sprintf("Playing Lofi: %s 🎵", st.TrackTitle))
						} else {
							m.SetStatusMsg("Lofi Radio Paused ⏸️")
						}
					}
				} else if msg.Y >= 12 && msg.Y <= 18 {
					m.Timer.Toggle()
					if m.Timer.Running {
						m.SetStatusMsg("Timer Started ⏱️")
					} else {
						m.SetStatusMsg("Timer Paused ⏸️")
					}
				} else if msg.Y >= 20 {
					petMsg := m.Scene.PetCat()
					m.SetStatusMsg(petMsg)
				}
			} else if m.ActiveTab == TabArt {
				// Canvas click -> Pet cat companion
				petMsg := m.Scene.PetCat()
				m.SetStatusMsg(petMsg)
			}
		}

	case tea.KeyMsg:
		// Modal: 2D Audio Visualizer & Spectrum Studio
		if m.ShowSpectrumModal {
			switch msg.String() {
			case "esc", "q", "M":
				m.ShowSpectrumModal = false
			case "tab", "space", "l", "right":
				m.SpectrumMode = (m.SpectrumMode + 1) % 5
				m.SetStatusMsg("Visualizer Mode: " + m.SpectrumMode.String() + " 🎶")
			case "h", "left":
				m.SpectrumMode = (m.SpectrumMode + 4) % 5
				m.SetStatusMsg("Visualizer Mode: " + m.SpectrumMode.String() + " 🎶")
			case "p":
				if m.CLIAMP != nil {
					_ = m.CLIAMP.TogglePlayPause()
					st := m.CLIAMP.GetStatus()
					if st.Playing {
						m.SetStatusMsg("CLIAMP Playback Resumed ▶️")
					} else {
						m.SetStatusMsg("CLIAMP Playback Paused ⏸️")
					}
				}
			case "1", "2", "3", "4", "5":
				idx := int(msg.String()[0] - '1')
				presets := audio.CuratedLofiPresets
				if idx >= 0 && idx < len(presets) && m.CLIAMP != nil {
					_ = m.CLIAMP.PlayPreset(idx)
					m.SetStatusMsg("Radio Stream: " + presets[idx].Name + " 📻")
				}
			}
			return m, nil
		}

		// Modal: Scene Info & Palette Inspector
		if m.ShowSceneInfoModal {
			switch msg.String() {
			case "esc", "q", "I", "i":
				m.ShowSceneInfoModal = false
			}
			return m, nil
		}

		// Modal: Quick Command Palette
		if m.ShowCommandPalette {
			allCmds := GetCommandList(&m)
			filtered := FilterCommands(allCmds, m.CommandInput.Value())

			switch msg.String() {
			case "esc":
				m.ShowCommandPalette = false
				m.CommandInput.Blur()
			case "j", "down", "ctrl+n":
				if m.CommandIdx < len(filtered)-1 {
					m.CommandIdx++
				}
			case "k", "up", "ctrl+p":
				if m.CommandIdx > 0 {
					m.CommandIdx--
				}
			case "enter":
				if len(filtered) > 0 && m.CommandIdx < len(filtered) {
					targetCmd := filtered[m.CommandIdx]
					targetCmd.Action(&m)
				}
				m.ShowCommandPalette = false
				m.CommandInput.Blur()
			default:
				var cmd tea.Cmd
				m.CommandInput, cmd = m.CommandInput.Update(msg)
				m.CommandIdx = 0
				return m, cmd
			}
			return m, nil
		}

		// Modal: Interactive Soundscape Mixer Console
		if m.ShowSoundMixerModal {
			switch msg.String() {
			case "esc", "q", "X":
				m.ShowSoundMixerModal = false
			case "j", "down":
				m.MixerChannelIdx = (m.MixerChannelIdx + 1) % 5
			case "k", "up":
				m.MixerChannelIdx = (m.MixerChannelIdx + 4) % 5
			case "l", "right", "+", "=":
				st := audio.SoundType(m.MixerChannelIdx)
				if st == audio.SoundClock { // skip invalid if mapped, order is Rain(0), Fire(1), Wind(2), Cafe(4), Crickets(5)
					st = audio.SoundRain
				}
				channelTypes := []audio.SoundType{audio.SoundRain, audio.SoundFire, audio.SoundWind, audio.SoundCafe, audio.SoundCrickets}
				if m.MixerChannelIdx < len(channelTypes) {
					st = channelTypes[m.MixerChannelIdx]
				}
				vol := m.Ambient.AdjustChannelVolume(st, 0.1)
				m.SetStatusMsg(fmt.Sprintf("Channel Volume adjusted to %.0f%% 🔊", vol*100))
			case "h", "left", "-", "_":
				channelTypes := []audio.SoundType{audio.SoundRain, audio.SoundFire, audio.SoundWind, audio.SoundCafe, audio.SoundCrickets}
				st := audio.SoundRain
				if m.MixerChannelIdx < len(channelTypes) {
					st = channelTypes[m.MixerChannelIdx]
				}
				vol := m.Ambient.AdjustChannelVolume(st, -0.1)
				m.SetStatusMsg(fmt.Sprintf("Channel Volume adjusted to %.0f%% 🔉", vol*100))
			case "enter", "space":
				switch m.MixerChannelIdx {
				case 0:
					m.Ambient.ToggleRain()
				case 1:
					m.Ambient.ToggleFire()
				case 2:
					m.Ambient.ToggleWind()
				case 3:
					m.Ambient.ToggleCafe()
				case 4:
					m.Ambient.ToggleCrickets()
				}
			}
			return m, nil
		}

		// Modal: Theme Picker
		if m.ShowThemePicker {
			switch msg.String() {
			case "esc", "q", "T":
				m.ShowThemePicker = false
			case "j", "down":
				m.PickerThemeIdx = (m.PickerThemeIdx + 1) % len(AllThemes)
			case "k", "up":
				m.PickerThemeIdx = (m.PickerThemeIdx + len(AllThemes) - 1) % len(AllThemes)
			case "enter":
				selectedTheme := AllThemes[m.PickerThemeIdx]
				m.Styles = NewStyles(selectedTheme)
				m.SetStatusMsg("Applied Theme: " + selectedTheme.Name + " ✨")
				m.ShowThemePicker = false
			}
			return m, nil
		}

		// Modal: Atmosphere Preset Studio & Custom Preset Manager
		if m.ShowPresetModal {
			if m.PresetSaving {
				switch msg.String() {
				case "esc":
					m.PresetSaving = false
					m.PresetInputText.Blur()
				case "enter":
					name := strings.TrimSpace(m.PresetInputText.Value())
					if name != "" {
						themeName := m.Styles.Theme.Name
						preset := features.AtmospherePreset{
							Name:            name,
							Description:     fmt.Sprintf("Custom user preset created in %s", themeName),
							Scene:           m.Scene.CurrentScene,
							Theme:           themeName,
							ColorFilter:     m.Scene.ColorFilter,
							ParticleOverlay: m.Scene.ParticleOverlay,
							ShaderEffect:    m.Scene.ShaderEffect,
							RainAudio:       m.Ambient.RainActive,
							FireAudio:       m.Ambient.FireActive,
							WindAudio:       m.Ambient.WindActive,
							AnimSpeed:       m.Scene.AnimSpeed,
						}
						m.Presets.SavePreset(preset)
						m.SetStatusMsg(fmt.Sprintf("Saved Preset: %s 💾", name))
					}
					m.PresetSaving = false
					m.PresetInputText.Blur()
				default:
					var cmd tea.Cmd
					m.PresetInputText, cmd = m.PresetInputText.Update(msg)
					return m, cmd
				}
				return m, nil
			}

			switch msg.String() {
			case "esc", "q", "S":
				m.ShowPresetModal = false
			case "j", "down":
				if m.PresetIdx < len(m.Presets.Presets)-1 {
					m.PresetIdx++
				}
			case "k", "up":
				if m.PresetIdx > 0 {
					m.PresetIdx--
				}
			case "enter":
				if len(m.Presets.Presets) > 0 && m.PresetIdx < len(m.Presets.Presets) {
					p := m.Presets.Presets[m.PresetIdx]
					m.ApplyPreset(p)
					m.ShowPresetModal = false
				}
			case "a":
				m.PresetSaving = true
				m.PresetInputText.Reset()
				m.PresetInputText.Focus()
				return m, textinput.Blink
			case "d":
				if len(m.Presets.Presets) > 0 && m.PresetIdx < len(m.Presets.Presets) {
					delName := m.Presets.Presets[m.PresetIdx].Name
					m.Presets.DeletePreset(m.PresetIdx)
					if m.PresetIdx >= len(m.Presets.Presets) && m.PresetIdx > 0 {
						m.PresetIdx--
					}
					m.SetStatusMsg(fmt.Sprintf("Deleted Preset: %s 🗑️", delName))
				}
			}
			return m, nil
		}

		// Modal: Interactive Scene Picker
		if m.ShowScenePicker {
			if m.PickerSearchMode {
				switch msg.String() {
				case "esc":
					if m.PickerSearchInput.Value() != "" {
						m.PickerSearchInput.Reset()
					} else {
						m.PickerSearchMode = false
						m.PickerSearchInput.Blur()
					}
				case "down", "j":
					filtered := m.getFilteredScenes()
					if m.PickerSceneIdx < len(filtered)-1 {
						m.PickerSceneIdx++
					}
				case "up", "k":
					if m.PickerSceneIdx > 0 {
						m.PickerSceneIdx--
					}
				case "enter":
					filtered := m.getFilteredScenes()
					if len(filtered) > 0 && m.PickerSceneIdx < len(filtered) {
						m.Scene.CurrentScene = filtered[m.PickerSceneIdx]
						m.SetStatusMsg("Switched Scene: " + m.Scene.CurrentScene.String() + " ✨")
					}
					m.ShowScenePicker = false
					m.PickerSearchMode = false
					m.PickerSearchInput.Blur()
				default:
					var cmd tea.Cmd
					m.PickerSearchInput, cmd = m.PickerSearchInput.Update(msg)
					m.PickerSceneIdx = 0
					return m, cmd
				}
				return m, nil
			}

			catScenes := m.getFilteredScenes()
			switch msg.String() {
			case "esc", "q", "V":
				m.ShowScenePicker = false
			case "/":
				m.PickerSearchMode = true
				m.PickerSearchInput.Focus()
				return m, textinput.Blink
			case "f", "*":
				if len(catScenes) > 0 && m.PickerSceneIdx < len(catScenes) {
					target := catScenes[m.PickerSceneIdx]
					fav := m.ToggleFavorite(target)
					if fav {
						m.SetStatusMsg("Added to Favorites: " + target.String() + " ⭐")
					} else {
						m.SetStatusMsg("Removed from Favorites: " + target.String())
					}
				}
			case "h", "left":
				m.PickerCategoryIdx = (m.PickerCategoryIdx + len(art.SceneCategories) - 1) % len(art.SceneCategories)
				m.PickerSceneIdx = 0
			case "l", "right":
				m.PickerCategoryIdx = (m.PickerCategoryIdx + 1) % len(art.SceneCategories)
				m.PickerSceneIdx = 0
			case "j", "down":
				if m.PickerSceneIdx < len(catScenes)-1 {
					m.PickerSceneIdx++
				}
			case "k", "up":
				if m.PickerSceneIdx > 0 {
					m.PickerSceneIdx--
				}
			case "enter":
				if len(catScenes) > 0 && m.PickerSceneIdx < len(catScenes) {
					m.Scene.CurrentScene = catScenes[m.PickerSceneIdx]
					m.SetStatusMsg("Switched Scene: " + m.Scene.CurrentScene.String() + " ✨")
				}
				m.ShowScenePicker = false
			}
			return m, nil
		}

		// Modal: Zen Focus Mode
		if m.ZenMode {
			switch msg.String() {
			case "esc", "F":
				m.ZenMode = false
				m.SetStatusMsg("Exited Zen Focus Mode 🌿")
				return m, nil
			case "space", " ":
				if m.CLIAMP != nil {
					_ = m.CLIAMP.TogglePlayPause()
				}
			case "x":
				m.Timer.Toggle()
			case "z":
				m.Timer.Reset()
			case "v":
				newScene := m.Scene.CycleScene()
				m.SetStatusMsg("Switched Scene: " + newScene.String())
			}
			return m, nil
		}

		// Task text input capture
		if m.AddingTask {
			switch msg.String() {
			case "enter":
				val := m.TextInput.Value()
				if val != "" {
					m.Tasks.Add(val)
					m.TextInput.Reset()
					m.SetStatusMsg(fmt.Sprintf("Added task: %s ✨", val))
				}
				m.AddingTask = false
				m.TextInput.Blur()
			case "esc":
				m.AddingTask = false
				m.TextInput.Blur()
			default:
				var cmd tea.Cmd
				m.TextInput, cmd = m.TextInput.Update(msg)
				return m, cmd
			}
			return m, nil
		}

		// Main App Keybindings
		switch msg.String() {
		case "q", "ctrl+c":
			m.Ambient.Cleanup()
			return m, tea.Quit

		case "ctrl+p", ":":
			m.ShowCommandPalette = !m.ShowCommandPalette
			m.CommandIdx = 0
			m.CommandInput.Reset()
			if m.ShowCommandPalette {
				m.CommandInput.Focus()
				return m, textinput.Blink
			}

		case "?":
			m.ShowHelp = !m.ShowHelp

		case "V":
			m.ShowScenePicker = !m.ShowScenePicker
			m.PickerSceneIdx = 0
			m.PickerSearchMode = false

		case "F":
			m.ZenMode = !m.ZenMode
			if m.ZenMode {
				m.SetStatusMsg("Zen Minimal Focus Mode Active 🌿 (Press Esc to exit)")
			} else {
				m.SetStatusMsg("Exited Zen Focus Mode 🌿")
			}

		case "M":
			m.ShowSpectrumModal = !m.ShowSpectrumModal

		case "R":
			newScene := m.Scene.RandomScene()
			m.SetStatusMsg("🎲 Random Scene: " + newScene.String() + " ✨")

		case "Y":
			mode := m.Scene.CycleAutoCycleMode()
			m.SetStatusMsg("🎬 Auto-Cycle Mode: " + mode.String())

		case "U":
			favOnly := m.Scene.ToggleAutoCycleFavs()
			if favOnly {
				m.SetStatusMsg("🎬 Auto-Cycle Scope: Favorites Only ⭐")
			} else {
				m.SetStatusMsg("🎬 Auto-Cycle Scope: All 370+ Scenes 🎨")
			}

		case "Q":
			q := m.Inspiration.NextQuote()
			m.SetStatusMsg(fmt.Sprintf("✨ Quote: \"%s\" — %s", q.Quote, q.Author))

		case "b":
			newScene := m.Scene.PrevScene()
			m.SetStatusMsg("Switched Scene: " + newScene.String())


		case "C":
			filter := m.Scene.CycleColorFilter()
			m.SetStatusMsg("Color Filter: " + filter.String() + " 🎨")

		case "I", "i":
			m.ShowSceneInfoModal = !m.ShowSceneInfoModal

		case "B":
			beat := m.Scene.ToggleBeatReactive()
			if beat {
				m.SetStatusMsg("Audio-Reactive Particle Pulsing: ON 🎵")
			} else {
				m.SetStatusMsg("Audio-Reactive Particle Pulsing: OFF")
			}

		case "D":
			blend := m.Scene.ToggleSmoothDayNight()
			if blend {
				m.SetStatusMsg("Smooth Day/Night Sky Blending: ON 🌅")
			} else {
				m.SetStatusMsg("Smooth Day/Night Sky Blending: OFF")
			}

		case "G":
			vig := m.Scene.ToggleVignette()
			if vig {
				m.SetStatusMsg("Cinematic Vignette Shading: ON 🖼️")
			} else {
				m.SetStatusMsg("Cinematic Vignette Shading: OFF")
			}

		case "O":
			crt := m.Scene.ToggleCRTScanlines()
			if crt {
				m.SetStatusMsg("Retro CRT Scanlines Overlay: ON 📺")
			} else {
				m.SetStatusMsg("Retro CRT Scanlines Overlay: OFF")
			}

		case "P":
			overlay := m.Scene.CycleParticleOverlay()
			m.SetStatusMsg("Particle Overlay: " + overlay.String() + " ✨")

		case "H", "h":
			shader := m.Scene.CycleShaderEffect()
			m.SetStatusMsg("Canvas Shader: " + shader.String() + " ✨")

		case "E":
			rm := m.Scene.CycleRenderMode()
			m.SetStatusMsg("Art Engine Render Mode: " + rm.String() + " 🎨")

		case "S":
			m.ShowPresetModal = !m.ShowPresetModal
			m.PresetIdx = 0
			m.PresetSaving = false

		case "*":
			fav := m.ToggleFavorite(m.Scene.CurrentScene)
			if fav {
				m.SetStatusMsg("Added to Favorites: " + m.Scene.CurrentScene.String() + " ⭐")
			} else {
				m.SetStatusMsg("Removed from Favorites: " + m.Scene.CurrentScene.String())
			}

		case "tab":
			m.ActiveTab = (m.ActiveTab + 1) % 4

		case "1":
			m.ActiveTab = TabArt
		case "2":
			m.ActiveTab = TabTimer
		case "3":
			m.ActiveTab = TabTasks
		case "4":
			m.ActiveTab = TabAmbiance

		// Interactive Bit-Art Toggles
		case "g":
			m.UseGPURender = !m.UseGPURender
			if m.UseGPURender {
				m.SetStatusMsg("Graphics Mode: Ghostty GPU (4x Crisp Native) ⚡")
			} else {
				m.SetStatusMsg("Graphics Mode: 8-Bit Half-Block Text Canvas ▀")
			}

		case "T":
			m.ShowThemePicker = !m.ShowThemePicker
			m.PickerThemeIdx = 0
			for i, t := range AllThemes {
				if t.Name == m.Styles.Theme.Name {
					m.PickerThemeIdx = i
					break
				}
			}

		case "v":
			newScene := m.Scene.CycleScene()
			m.SetStatusMsg("Switched Scene: " + newScene.String())

		case "l":
			m.Scene.LampOn = !m.Scene.LampOn
			if m.Scene.LampOn {
				m.SetStatusMsg("Desk Lamp ON 💡")
			} else {
				m.SetStatusMsg("Desk Lamp OFF 🌙")
			}

		case "A":
			m.AutoSyncAudio = !m.AutoSyncAudio
			if m.AutoSyncAudio {
				m.SetStatusMsg("Atmospheric Audio Auto-Sync: ON 🌿")
			} else {
				m.SetStatusMsg("Atmospheric Audio Auto-Sync: OFF")
			}

		case "+", "=":
			if m.CLIAMP != nil {
				vol := m.CLIAMP.AdjustVolume(10)
				m.SetStatusMsg(fmt.Sprintf("Lofi Volume: %d%% 🔊", vol))
			}

		case "-", "_":
			if m.CLIAMP != nil {
				vol := m.CLIAMP.AdjustVolume(-10)
				m.SetStatusMsg(fmt.Sprintf("Lofi Volume: %d%% 🔉", vol))
			}

		case "f":
			if m.ActiveTab == TabTasks {
				flt := m.Tasks.CycleFilter()
				m.SetStatusMsg("Task Filter: " + flt.String())
			} else {
				m.Scene.FireOn = !m.Scene.FireOn
				if m.AutoSyncAudio {
					m.Ambient.FireActive = m.Scene.FireOn
				}
				if m.Scene.FireOn {
					m.SetStatusMsg("Fireplace ON 🔥")
				} else {
					m.SetStatusMsg("Fireplace OFF 💤")
				}
			}

		case "w":
			m.Scene.Weather = (m.Scene.Weather + 1) % 5
			if m.AutoSyncAudio {
				if m.Scene.Weather == art.WeatherRain || m.Scene.Weather == art.WeatherStorm {
					m.Ambient.RainActive = true
				} else {
					m.Ambient.RainActive = false
				}
			}
			m.SetStatusMsg("Weather: " + m.Scene.Weather.String())

		case "c":
			m.Scene.CatBreed = (m.Scene.CatBreed + 1) % 4
			m.SetStatusMsg("Cat Breed: " + m.Scene.CatBreed.String())

		// Audio & CLIAMP / Native Lofi Radio Controls
		case "space", " ":
			if m.CLIAMP != nil {
				_ = m.CLIAMP.TogglePlayPause()
				st := m.CLIAMP.GetStatus()
				if st.Playing {
					m.SetStatusMsg(fmt.Sprintf("Playing Lofi: %s 🎵", st.TrackTitle))
				} else {
					m.SetStatusMsg("Lofi Radio Paused ⏸️")
				}
			}

		case "n":
			if m.CLIAMP != nil {
				_ = m.CLIAMP.NextTrack()
				st := m.CLIAMP.GetStatus()
				m.SetStatusMsg(fmt.Sprintf("Playing Lofi: %s 🎶", st.TrackTitle))
			}

		case "p":
			if m.ActiveTab == TabTasks {
				pri := m.Tasks.CyclePriorityCurrent()
				m.SetStatusMsg("Task Priority: " + pri.String())
			} else if m.CLIAMP != nil {
				_ = m.CLIAMP.PrevTrack()
				st := m.CLIAMP.GetStatus()
				m.SetStatusMsg(fmt.Sprintf("Playing Lofi: %s 🎶", st.TrackTitle))
			}

		case "r":
			active := m.Ambient.ToggleRain()
			if active {
				m.SetStatusMsg("Rain Audio ON 🌧️")
			} else {
				m.SetStatusMsg("Rain Audio OFF")
			}

		case "m":
			active := m.Ambient.ToggleFire()
			if active {
				m.SetStatusMsg("Fire Audio ON 🔥")
			} else {
				m.SetStatusMsg("Fire Audio OFF")
			}

		case "W":
			active := m.Ambient.ToggleWind()
			if active {
				m.SetStatusMsg("Wind Audio ON 🌬️")
			} else {
				m.SetStatusMsg("Wind Audio OFF")
			}

		case "K":
			msg := m.Scene.PetCat()
			m.SetStatusMsg(msg)

		case "X":
			m.ShowSoundMixerModal = !m.ShowSoundMixerModal
			m.MixerChannelIdx = 0

		case "e":
			active := m.Ambient.ToggleCrickets()
			if active {
				m.SetStatusMsg("Night Crickets Audio ON 🦗")
			} else {
				m.SetStatusMsg("Night Crickets Audio OFF")
			}

		// Timer Controls
		case "x":
			m.Timer.Toggle()
			if m.Timer.Running {
				m.SetStatusMsg("Timer Started ⏱️")
			} else {
				m.SetStatusMsg("Timer Paused ⏸️")
			}

		case "z":
			m.Timer.Reset()
			m.SetStatusMsg("Timer Reset 🔄")

		// Task List Controls
		case "j", "down":
			if m.ActiveTab == TabTasks {
				m.Tasks.MoveDown()
			}
		case "k", "up":
			if m.ActiveTab == TabTasks {
				m.Tasks.MoveUp()
			} else {
				msg := m.Scene.PetCat()
				m.SetStatusMsg(msg)
			}
		case "enter":
			if m.ActiveTab == TabTasks {
				m.Tasks.ToggleCurrent()
			}
		case "a":
			if m.ActiveTab == TabTasks {
				m.AddingTask = true
				m.TextInput.Focus()
				return m, textinput.Blink
			}
		case "d":
			if m.ActiveTab == TabTasks {
				m.Tasks.DeleteCurrent()
				m.SetStatusMsg("Task deleted 🗑️")
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) getFilteredScenes() []art.SceneType {
	query := strings.TrimSpace(strings.ToLower(m.PickerSearchInput.Value()))
	if query == "" {
		return art.GetScenesByCategoryAndFavorites(m.PickerCategoryIdx, m.FavoriteScenes)
	}

	var matched []art.SceneType
	for _, scene := range art.SceneOrder {
		if strings.Contains(strings.ToLower(scene.String()), query) {
			matched = append(matched, scene)
		}
	}
	return matched
}
