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
	Width             int
	Height            int
	Styles            Styles
	ActiveTab         ActiveTab
	ShowHelp          bool
	ShowScenePicker   bool
	PickerCategoryIdx int
	PickerSceneIdx    int
	PickerSearchMode  bool
	PickerSearchInput textinput.Model
	ShowThemePicker   bool
	PickerThemeIdx    int
	SpectrumMode      SpectrumStyle
	ZenMode           bool
	StatusMsg         string
	StatusMsgTime     time.Time
	UseGPURender      bool
	FavoriteScenes    map[art.SceneType]bool

	// Core Components
	Scene   *art.CottageScene
	CLIAMP  *audio.CLIAMPController
	Ambient *audio.AmbientManager
	Timer   *features.FocusTimer
	Tasks   *features.TaskList

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

	styles := NewStyles(ThemeCozyAutumn)
	scene := art.NewCottageScene(160, 90)

	favs := make(map[art.SceneType]bool)
	favs[art.SceneRainyCity] = true

	return Model{
		Width:             120,
		Height:            40,
		Styles:            styles,
		ActiveTab:         TabArt,
		ShowHelp:          false,
		ShowScenePicker:   false,
		PickerSearchInput: si,
		ShowThemePicker:   false,
		SpectrumMode:      SpectrumBars,
		UseGPURender:      art.DetectTerminalProtocol() == art.ProtocolKitty,
		FavoriteScenes:    favs,

		Scene:   scene,
		CLIAMP:  audio.NewCLIAMPController(),
		Ambient: audio.NewAmbientManager(),
		Timer:   features.NewFocusTimer(),
		Tasks:   features.NewTaskList(),

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
		cmds = append(cmds, m.animateTick())

	case cliampPollMsg:
		cmds = append(cmds, m.pollCLIAMP())

	case timerTickMsg:
		if m.Timer.Tick() {
			m.SetStatusMsg("🎉 Timer completed! Time for a cozy break!")
		}
		cmds = append(cmds, m.timerTick())

	case tea.KeyMsg:
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
			m.SpectrumMode = (m.SpectrumMode + 1) % 4
			m.SetStatusMsg("Visualizer Mode: " + m.SpectrumMode.String() + " 🎶")

		case "C":
			filter := m.Scene.CycleColorFilter()
			m.SetStatusMsg("Color Filter: " + filter.String() + " 🎨")

		case "P":
			overlay := m.Scene.CycleParticleOverlay()
			m.SetStatusMsg("Particle Overlay: " + overlay.String() + " ✨")

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

		case "f":
			m.Scene.FireOn = !m.Scene.FireOn
			if m.Scene.FireOn {
				m.SetStatusMsg("Fireplace ON 🔥")
			} else {
				m.SetStatusMsg("Fireplace OFF 💤")
			}

		case "s":
			m.Scene.FairyLightsOn = !m.Scene.FairyLightsOn
			if m.Scene.FairyLightsOn {
				m.SetStatusMsg("Fairy Lights ON ✨")
			} else {
				m.SetStatusMsg("Fairy Lights OFF 💤")
			}

		case "w":
			m.Scene.Weather = (m.Scene.Weather + 1) % 5
			m.SetStatusMsg("Weather: " + m.Scene.Weather.String())

		case "t":
			m.Scene.TimeOfDay = (m.Scene.TimeOfDay + 1) % 5
			m.SetStatusMsg("Time of Day: " + m.Scene.TimeOfDay.String())

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
			if m.CLIAMP != nil {
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
