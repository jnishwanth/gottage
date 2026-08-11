package ui

import (
	"fmt"
	"strings"
	"time"

	"gottage/pkg/art"
	"gottage/pkg/audio"

	"github.com/charmbracelet/lipgloss"
)

// View renders the main UI layout.
func (m Model) View() string {
	var body string

	// Build Tab Navigation Bar
	tabs := []ActiveTab{TabArt, TabTimer, TabTasks, TabAmbiance}
	var tabStrings []string
	for i, tab := range tabs {
		idxStr := fmt.Sprintf("[%d] %s", i+1, tab.String())
		if m.ActiveTab == tab {
			tabStrings = append(tabStrings, m.Styles.ActiveTab.Render(idxStr))
		} else {
			tabStrings = append(tabStrings, m.Styles.InactiveTab.Render(idxStr))
		}
	}
	tabBar := strings.Join(tabStrings, "  ")

	headerTitle := m.Styles.Title.Render("🏡 GOTTAGE") + " " + m.Styles.Header.Render("v1.0")
	protoBadge := m.Styles.StatusBadge.Render(" " + art.DetectTerminalProtocol().String() + " ")
	header := lipgloss.JoinHorizontal(lipgloss.Center, headerTitle, "  ", protoBadge, "   ", tabBar)

	// Render Main Tab View
	switch m.ActiveTab {
	case TabArt:
		body = m.renderArtView()
	case TabTimer:
		body = m.renderTimerView()
	case TabTasks:
		body = m.renderTasksView()
	case TabAmbiance:
		body = m.renderAmbianceView()
	}

	// If active tab is NOT TabArt, clear any leftover Kitty GPU image overlay from Ghostty!
	if m.ActiveTab != TabArt {
		body = art.ClearKittyGPU() + body
	}

	// Render Right Sidebar
	sidebar := m.renderSidebar()

	// Combine Main Body and Sidebar
	content := lipgloss.JoinHorizontal(lipgloss.Top, body, "  ", sidebar)

	// Toast / Status Footer
	footerText := "Press [?] for help • [1-4] switch tabs • [space] audio • [g] toggle GPU/Text • [q] quit"
	if m.StatusMsg != "" && time.Since(m.StatusMsgTime) < 4*time.Second {
		footerText = m.Styles.StatusBadge.Render(" " + m.StatusMsg + " ")
	}
	footer := m.Styles.Footer.Render(footerText)

	mainView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		content,
		"",
		footer,
	)

	// Help Overlay Popup
	if m.ShowHelp {
		return m.renderHelpModal(mainView)
	}

	return mainView
}

func (m Model) renderArtView() string {
	// If GPU mode is enabled and terminal is Kitty/Ghostty, return Kitty GPU graphics string
	if m.UseGPURender && art.DetectTerminalProtocol() == art.ProtocolKitty {
		artStr := m.Scene.Render()
		return m.Styles.Box.Width(76).Height(35).Render(artStr)
	}

	// Fallback to crisp PICO-8 half-block text canvas ('▀')
	m.Scene.RenderFrame()
	artStr := strings.TrimRight(m.Scene.Canvas.Render(), "\r\n")
	return m.Styles.Box.Width(76).Height(35).Render(artStr)
}

func (m Model) renderTimerView() string {
	var sb strings.Builder

	sb.WriteString(m.Styles.Header.Render("⏱️ COZY FOCUS TIMER") + "\n\n")
	sb.WriteString(fmt.Sprintf("Mode: %s\n\n", m.Styles.Title.Render(m.Timer.Mode.String())))

	// Giant ASCII Digital Clock
	clockStr := m.renderBigClock(m.Timer.FormattedTime())
	sb.WriteString(clockStr + "\n\n")

	// Progress Bar
	barWidth := 36
	completed := int(m.Timer.ProgressRatio() * float64(barWidth))
	if completed > barWidth {
		completed = barWidth
	}
	progressBar := strings.Repeat("█", completed) + strings.Repeat("░", barWidth-completed)
	sb.WriteString(fmt.Sprintf("Progress: [%s] %.0f%%\n\n", m.Styles.ProgressBar.Render(progressBar), m.Timer.ProgressRatio()*100))

	stateStr := "PAUSED ⏸️"
	if m.Timer.Running {
		stateStr = "RUNNING 📖"
	}
	sb.WriteString(fmt.Sprintf("Status: %s\n", stateStr))
	sb.WriteString(fmt.Sprintf("Completed Sessions Today: %d ☕\n\n", m.Timer.CompletedSessions))

	sb.WriteString(m.Styles.Footer.Render("Press [x] to Start/Pause • [z] to Reset"))

	return m.Styles.Box.Width(76).Height(35).Render(sb.String())
}

func (m Model) renderBigClock(timeStr string) string {
	// Simple clean big digit representation
	return lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(m.Styles.Theme.Title).
		Padding(1, 4).
		Bold(true).
		Render(timeStr)
}

func (m Model) renderTasksView() string {
	var sb strings.Builder

	sb.WriteString(m.Styles.Header.Render("📝 COTTAGE TO-DO NOTEBOOK") + "\n\n")

	if m.AddingTask {
		sb.WriteString("Add New Task:\n")
		sb.WriteString(m.TextInput.View() + "\n\n")
		sb.WriteString(m.Styles.Footer.Render("Press Enter to add, Esc to cancel"))
		return m.Styles.Box.Width(76).Height(35).Render(sb.String())
	}

	if len(m.Tasks.Items) == 0 {
		sb.WriteString(m.Styles.Footer.Render("No tasks yet. Press [a] to add a task!"))
	} else {
		for i, item := range m.Tasks.Items {
			cursor := "  "
			if i == m.Tasks.Cursor {
				cursor = "👉"
			}

			box := "[ ]"
			titleStyle := lipgloss.NewStyle()
			if item.Completed {
				box = "[✓]"
				titleStyle = titleStyle.Strikethrough(true).Foreground(m.Styles.Theme.Subdued)
			}

			line := fmt.Sprintf("%s %s %s", cursor, box, titleStyle.Render(item.Title))
			sb.WriteString(line + "\n")
		}
	}

	sb.WriteString("\n" + strings.Repeat("─", 40) + "\n")
	sb.WriteString(m.Styles.Footer.Render("[a] Add Task • [Enter] Toggle Done • [d] Delete • [j/k] Navigate"))

	return m.Styles.Box.Width(76).Height(35).Render(sb.String())
}

func (m Model) renderAmbianceView() string {
	var sb strings.Builder

	sb.WriteString(m.Styles.Header.Render("🌿 AMBIANCE & MUSIC CONTROL CENTER") + "\n\n")

	// Scene Visual Controls
	sb.WriteString(m.Styles.Title.Render("Bit-Art Environment Controls:") + "\n")
	sb.WriteString(fmt.Sprintf(" [v] Scene Preset:  %s\n", m.Styles.Header.Render(m.Scene.CurrentScene.String())))
	sb.WriteString(fmt.Sprintf(" [t] Time of Day:   %s\n", m.Scene.EffectiveTimeOfDay().String()))
	sb.WriteString(fmt.Sprintf(" [w] Weather Mode:  %s\n", m.Scene.Weather.String()))
	sb.WriteString(fmt.Sprintf(" [c] Cat Breed:     %s\n", m.Scene.CatBreed.String()))
	sb.WriteString(fmt.Sprintf(" [l] Desk Lamp:     %s\n", onOffStr(m.Scene.LampOn)))
	sb.WriteString(fmt.Sprintf(" [f] Fireplace:     %s\n", onOffStr(m.Scene.FireOn)))
	sb.WriteString(fmt.Sprintf(" [s] Fairy Lights:  %s\n\n", onOffStr(m.Scene.FairyLightsOn)))

	// Ambient Audio Controls
	sb.WriteString(m.Styles.Title.Render("Native Ambient Sound Engine:") + "\n")
	sb.WriteString(fmt.Sprintf(" [r] Rain Audio:    %s\n", onOffStr(m.Ambient.RainActive)))
	sb.WriteString(fmt.Sprintf(" [m] Fire Audio:    %s\n", onOffStr(m.Ambient.FireActive)))
	sb.WriteString(fmt.Sprintf("     Wind Audio:    %s\n\n", onOffStr(m.Ambient.WindActive)))

	// CLIAMP Details
	sb.WriteString(m.Styles.Title.Render("CLIAMP Music Player Status:") + "\n")
	st := m.CLIAMP.GetStatus()
	if st.Installed {
		sb.WriteString(fmt.Sprintf(" CLIAMP Installed: YES (/opt/homebrew/bin/cliamp)\n"))
		sb.WriteString(fmt.Sprintf(" Daemon Status:   %s\n", daemonStr(st.DaemonRunning)))
		if st.TrackTitle != "" {
			sb.WriteString(fmt.Sprintf(" Current Track:   %s\n", m.Styles.Header.Render(st.TrackTitle)))
		}
	} else {
		sb.WriteString(" CLIAMP Installed: NO (Running standalone procedural audio engine)\n")
	}

	sb.WriteString("\n" + m.Styles.Title.Render("Curated Lofi Radio Streams (Press n/p to cycle):") + "\n")
	for i, preset := range audio.CuratedLofiPresets {
		prefix := "  "
		if i == m.CLIAMP.ActivePreset {
			prefix = "▶ "
		}
		sb.WriteString(fmt.Sprintf("%s%s (%s)\n", prefix, preset.Name, preset.Tag))
	}

	return m.Styles.Box.Width(76).Height(35).Render(sb.String())
}

func (m Model) renderSidebar() string {
	var sb strings.Builder

	// 1. Music Player Widget
	sb.WriteString(m.Styles.Title.Render("🎵 LOFI PLAYER") + "\n")
	st := m.CLIAMP.GetStatus()
	trackName := audio.CuratedLofiPresets[m.CLIAMP.ActivePreset].Name
	if st.Playing && st.TrackTitle != "" {
		trackName = st.TrackTitle
	}

	stateIcon := "⏸️ Paused"
	if st.Playing {
		stateIcon = "▶️ Playing"
	}

	sb.WriteString(fmt.Sprintf("%s\n", m.Styles.Header.Render(truncate(trackName, 26))))
	sb.WriteString(fmt.Sprintf("State: %s\n", stateIcon))
	sb.WriteString(m.Styles.Footer.Render("[space] Play/Pause • [n/p] Prev/Next") + "\n\n")

	sb.WriteString(strings.Repeat("─", 28) + "\n\n")

	// 2. Timer Widget
	sb.WriteString(m.Styles.Title.Render("⏱️ FOCUS TIMER") + "\n")
	sb.WriteString(fmt.Sprintf("Time: %s\n", m.Styles.Header.Render(m.Timer.FormattedTime())))
	sb.WriteString(fmt.Sprintf("Mode: %s\n", m.Timer.Mode.String()))
	sb.WriteString(m.Styles.Footer.Render("[x] Start/Pause • [z] Reset") + "\n\n")

	sb.WriteString(strings.Repeat("─", 28) + "\n\n")

	// 3. Quick Toggles & System Info
	sb.WriteString(m.Styles.Title.Render("✨ QUICK TOGGLES & THEME") + "\n")
	sb.WriteString(fmt.Sprintf("[v] Scene: %s\n", truncate(m.Scene.CurrentScene.String(), 18)))
	sb.WriteString(fmt.Sprintf("[Shift+T] Theme: %s\n", truncate(m.Styles.Theme.Name, 16)))
	sb.WriteString(fmt.Sprintf("[l] Lamp:  %s\n", onOffStr(m.Scene.LampOn)))
	sb.WriteString(fmt.Sprintf("[f] Fire:  %s\n", onOffStr(m.Scene.FireOn)))
	sb.WriteString(fmt.Sprintf("[w] Weather: %s\n", m.Scene.Weather.String()))

	return m.Styles.Box.Width(32).Height(35).Render(sb.String())
}

func (m Model) renderHelpModal(bg string) string {
	helpText := `
  🏡 GOTTAGE — KEYBOARD SHORTCUTS & HELP
  ---------------------------------------
  [1]      Switch to Cottage Art View
  [2]      Switch to Focus Timer
  [3]      Switch to Task Notebook
  [4]      Switch to Ambiance & Audio
  [Tab]    Cycle Views

  DYNAMIC BIT-ART SCENES & CONTROLS:
  [v]      Cycle Scene Preset (Rainy City, Forest Cabin, Classic, Observatory)
  [l]      Toggle Desk Lamp (ON/OFF)
  [f]      Toggle Fireplace (ON/OFF)
  [s]      Toggle Fairy Lights (ON/OFF)
  [w]      Cycle Weather (Clear, Rain, Snow, Storm, Fog)
  [t]      Cycle Time of Day (Dawn, Day, Sunset, Night, Auto)
  [c]      Cycle Cat Breed (Tabby, Tuxedo, Calico, Black)

  AUDIO & LOFI CONTROLS:
  [space]  Play/Pause CLIAMP / Audio
  [n]      Next Lofi Track Preset
  [p]      Previous Lofi Track Preset
  [r]      Toggle Rain Ambient Sound
  [m]      Toggle Fireplace Ambient Sound

  TIMER & TASKS:
  [x]      Start / Pause Pomodoro Timer
  [z]      Reset Pomodoro Timer
  [a]      Add new Task (in Tasks View)
  [Enter]  Toggle Task Done
  [d]      Delete Task

  [?]      Close Help Overlay
  [q]      Quit Gottage
`
	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Render(helpText)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}

func onOffStr(v bool) string {
	if v {
		return "ON ✨"
	}
	return "OFF 💤"
}

func daemonStr(v bool) string {
	if v {
		return "ACTIVE 🟢"
	}
	return "STOPPED 🔴"
}

func truncate(str string, maxLen int) string {
	if len(str) <= maxLen {
		return str
	}
	return str[:maxLen-3] + "..."
}
