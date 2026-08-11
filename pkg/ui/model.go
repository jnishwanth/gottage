package ui

import (
	"fmt"
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
	Width         int
	Height        int
	Styles        Styles
	ActiveTab     ActiveTab
	ShowHelp      bool
	StatusMsg     string
	StatusMsgTime time.Time

	// Core Components
	Scene     *art.CottageScene
	CLIAMP    *audio.CLIAMPController
	Ambient   *audio.AmbientManager
	Timer     *features.FocusTimer
	Tasks     *features.TaskList
	TextInput textinput.Model
	AddingTask bool
}

// NewModel constructs the initial Gottage app model.
func NewModel() Model {
	// Initialize Canvas width & height (74 wide x 68 pixel height = 34 character rows)
	scene := art.NewCottageScene(74, 68)
	cliampCtrl := audio.NewCLIAMPController()
	ambientMgr := audio.NewAmbientManager()
	ambientMgr.StartAmbientLoop()

	ti := textinput.New()
	ti.Placeholder = "Enter new task..."
	ti.CharLimit = 60

	return Model{
		Width:     100,
		Height:    40,
		Styles:    NewStyles(ThemeCozyAutumn),
		ActiveTab: TabArt,
		ShowHelp:  false,
		Scene:     scene,
		CLIAMP:    cliampCtrl,
		Ambient:   ambientMgr,
		Timer:     features.NewFocusTimer(),
		Tasks:     features.NewTaskList(),
		TextInput: ti,
	}
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
		case "T":
			newTheme := CycleTheme(m.Styles.Theme)
			m.Styles = NewStyles(newTheme)
			m.SetStatusMsg("Theme Switched: " + newTheme.Name)

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

		// Audio & CLIAMP Controls
		case "space":
			if m.CLIAMP != nil && m.CLIAMP.GetStatus().Installed {
				_ = m.CLIAMP.TogglePlayPause()
				m.SetStatusMsg("CLIAMP Toggled 🎵")
			} else {
				m.SetStatusMsg("CLIAMP not installed, running procedural ambient 🎧")
			}

		case "n":
			if m.CLIAMP != nil {
				_ = m.CLIAMP.NextTrack()
				m.SetStatusMsg(fmt.Sprintf("Preset: %s 🎶", audio.CuratedLofiPresets[m.CLIAMP.ActivePreset].Name))
			}

		case "p":
			if m.CLIAMP != nil {
				_ = m.CLIAMP.PrevTrack()
				m.SetStatusMsg(fmt.Sprintf("Preset: %s 🎶", audio.CuratedLofiPresets[m.CLIAMP.ActivePreset].Name))
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
