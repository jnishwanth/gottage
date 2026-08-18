package ui

import (
	"fmt"
	"strings"
	"time"

	"gottage/pkg/art"
	"gottage/pkg/audio"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) mainWidth() int {
	w := m.Width - 36
	if w < 54 {
		return 54
	}
	if w > 120 {
		return 120
	}
	return w
}

func (m Model) mainHeight() int {
	h := m.Height - 6
	if h < 22 {
		return 22
	}
	if h > 50 {
		return 50
	}
	return h
}

// View renders the main UI layout.
func (m Model) View() string {
	if m.ZenMode {
		return m.renderZenView()
	}

	var body string

	header := m.renderHeader()

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
	footerText := "Press [?] help • [V] Catalog • [T] Themes • [M] Spectrum • [Shift+F] Zen Mode • [1-4] Tabs • [space] Audio • [q] quit"
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

	// Theme Selector Modal
	if m.ShowThemePicker {
		return m.renderThemePickerModal(mainView)
	}

	// Scene Picker Catalog Modal
	if m.ShowScenePicker {
		return m.renderScenePickerModal(mainView)
	}

	// Help Overlay Popup
	if m.ShowHelp {
		return m.renderHelpModal(mainView)
	}

	return mainView
}

func (m Model) renderHeader() string {
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
	tabBar := strings.Join(tabStrings, " ")

	title := m.Styles.Title.Render("🏡 GOTTAGE") + " " + m.Styles.MutedText.Render("v1.0")
	proto := m.Styles.StatusBadge.Render(" " + art.DetectTerminalProtocol().String() + " ")
	themeBadge := m.Styles.PillBadge.Render(" " + m.Styles.Theme.Name + " ")

	return lipgloss.JoinHorizontal(lipgloss.Center, title, "  ", proto, "  ", themeBadge, "    ", tabBar)
}

func (m Model) renderArtView() string {
	w := m.mainWidth()
	h := m.mainHeight()

	var artStr string
	if m.UseGPURender && art.DetectTerminalProtocol() == art.ProtocolKitty {
		artStr = m.Scene.Render()
	} else {
		m.Scene.RenderFrame()
		artStr = strings.TrimRight(m.Scene.Canvas.Render(), "\r\n")
	}

	favBadge := ""
	if m.FavoriteScenes != nil && m.FavoriteScenes[m.Scene.CurrentScene] {
		favBadge = " ⭐"
	}

	metaHeader := fmt.Sprintf("🎨 %s%s  |  ☀️ %s  |  🌧️ %s  |  🎨 %s  |  ✨ %s",
		m.Styles.BoldAccent.Render(m.Scene.CurrentScene.String()),
		favBadge,
		m.Scene.EffectiveTimeOfDay().String(),
		m.Scene.Weather.String(),
		m.Scene.ColorFilter.String(),
		m.Scene.ParticleOverlay.String(),
	)

	content := metaHeader + "\n\n" + artStr
	return m.Styles.Box.Width(w).Height(h).Render(content)
}

func (m Model) renderTimerView() string {
	w := m.mainWidth()
	h := m.mainHeight()

	var sb strings.Builder
	sb.WriteString(m.Styles.Header.Render("⏱️ COZY FOCUS TIMER & POMODORO") + "\n\n")

	clockStr := m.renderBigClock(m.Timer.FormattedTime())
	sb.WriteString(clockStr + "\n\n")

	barWidth := w - 24
	if barWidth < 20 {
		barWidth = 20
	}
	completed := int(m.Timer.ProgressRatio() * float64(barWidth))
	if completed > barWidth {
		completed = barWidth
	}
	progressBar := strings.Repeat("█", completed) + strings.Repeat("░", barWidth-completed)

	sb.WriteString(fmt.Sprintf("Mode: %s    Status: %s\n",
		m.Styles.Title.Render(m.Timer.Mode.String()),
		m.timerStateStr()))
	sb.WriteString(fmt.Sprintf("Progress: [%s] %.0f%%\n\n", m.Styles.ProgressBar.Render(progressBar), m.Timer.ProgressRatio()*100))
	sb.WriteString(fmt.Sprintf("☕ Sessions Completed Today: %s\n\n", m.Styles.BoldAccent.Render(fmt.Sprintf("%d", m.Timer.CompletedSessions))))

	breath := GetZenBreathState(m.Scene.FrameCount)
	sb.WriteString(m.Styles.Title.Render("🧘 ZEN BREATHING RHYTHM") + "\n")
	sb.WriteString(fmt.Sprintf("  %s %s\n", m.Styles.BoldAccent.Render(breath.PhaseName), m.Styles.Header.Render(breath.RingVisual)))
	sb.WriteString(m.Styles.MutedText.Render("  "+breath.Instruction) + "\n\n")

	sb.WriteString(m.Styles.MutedText.Render("Quick Actions: ") + m.Styles.KeyCap.Render("[x]") + " Start/Pause  " + m.Styles.KeyCap.Render("[z]") + " Reset")

	return m.Styles.Box.Width(w).Height(h).Render(sb.String())
}

func (m Model) timerStateStr() string {
	if m.Timer.Running {
		return m.Styles.BoldAccent.Render("RUNNING 📖")
	}
	return m.Styles.MutedText.Render("PAUSED ⏸️")
}

func (m Model) renderBigClock(timeStr string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(m.Styles.Theme.Title).
		Padding(1, 4).
		Bold(true).
		Render(timeStr)
}

func (m Model) renderTasksView() string {
	w := m.mainWidth()
	h := m.mainHeight()

	var sb strings.Builder
	completedCount := 0
	for _, item := range m.Tasks.Items {
		if item.Completed {
			completedCount++
		}
	}

	header := fmt.Sprintf("📝 COTTAGE TO-DO NOTEBOOK  (%d/%d Completed)", completedCount, len(m.Tasks.Items))
	sb.WriteString(m.Styles.Header.Render(header) + "\n\n")

	if m.AddingTask {
		sb.WriteString(m.Styles.Title.Render("Add New Focus Task:") + "\n\n")
		sb.WriteString(m.TextInput.View() + "\n\n")
		sb.WriteString(m.Styles.MutedText.Render("Press Enter to save • Esc to cancel"))
		return m.Styles.Box.Width(w).Height(h).Render(sb.String())
	}

	if len(m.Tasks.Items) == 0 {
		sb.WriteString(m.Styles.MutedText.Render("Notebook is empty! Press [a] to add your first cozy focus task."))
	} else {
		for i, item := range m.Tasks.Items {
			cursor := "  "
			if i == m.Tasks.Cursor {
				cursor = "👉"
			}

			box := "[ ]"
			titleStyle := lipgloss.NewStyle().Foreground(m.Styles.Theme.Title)
			if item.Completed {
				box = "[✓]"
				titleStyle = titleStyle.Strikethrough(true).Foreground(m.Styles.Theme.Subdued)
			}

			line := fmt.Sprintf("%s %s %s", cursor, m.Styles.BoldAccent.Render(box), titleStyle.Render(item.Title))
			sb.WriteString(line + "\n")
		}
	}

	sb.WriteString("\n" + strings.Repeat("─", w-6) + "\n")
	sb.WriteString(m.Styles.MutedText.Render("Actions: ") +
		m.Styles.KeyCap.Render("[a]") + " Add  " +
		m.Styles.KeyCap.Render("[Enter]") + " Toggle  " +
		m.Styles.KeyCap.Render("[d]") + " Delete  " +
		m.Styles.KeyCap.Render("[j/k]") + " Navigate")

	return m.Styles.Box.Width(w).Height(h).Render(sb.String())
}

func (m Model) renderAmbianceView() string {
	w := m.mainWidth()
	h := m.mainHeight()

	var sb strings.Builder
	sb.WriteString(m.Styles.Header.Render("🌿 AMBIANCE, ENVIRONMENT & MUSIC CONTROL") + "\n\n")

	sb.WriteString(m.Styles.Title.Render("🎨 Visual Environment Settings") + "\n")
	sb.WriteString(fmt.Sprintf("  %s Scene Preset:   %s\n", m.Styles.KeyCap.Render("[v]"), m.Styles.BoldAccent.Render(m.Scene.CurrentScene.String())))
	sb.WriteString(fmt.Sprintf("  %s Color Filter:   %s\n", m.Styles.KeyCap.Render("[C]"), m.Scene.ColorFilter.String()))
	sb.WriteString(fmt.Sprintf("  %s Screen Overlay: %s\n", m.Styles.KeyCap.Render("[P]"), m.Scene.ParticleOverlay.String()))
	sb.WriteString(fmt.Sprintf("  %s Time of Day:    %s\n", m.Styles.KeyCap.Render("[t]"), m.Scene.EffectiveTimeOfDay().String()))
	sb.WriteString(fmt.Sprintf("  %s Weather Mode:   %s\n", m.Styles.KeyCap.Render("[w]"), m.Scene.Weather.String()))
	sb.WriteString(fmt.Sprintf("  %s Cat Companion:  %s\n", m.Styles.KeyCap.Render("[c]"), m.Scene.CatBreed.String()))
	sb.WriteString(fmt.Sprintf("  %s Desk Lamp:      %s   %s Fireplace: %s   %s Fairy Lights: %s\n\n",
		m.Styles.KeyCap.Render("[l]"), onOffStr(m.Scene.LampOn),
		m.Styles.KeyCap.Render("[f]"), onOffStr(m.Scene.FireOn),
		m.Styles.KeyCap.Render("[s]"), onOffStr(m.Scene.FairyLightsOn)))

	sb.WriteString(m.Styles.Title.Render("🌧️ Procedural Audio Channels") + "\n")
	sb.WriteString(fmt.Sprintf("  %s Rain Sound: %s    %s Fire Sound: %s    %s Wind Sound: %s\n\n",
		m.Styles.KeyCap.Render("[r]"), onOffStr(m.Ambient.RainActive),
		m.Styles.KeyCap.Render("[m]"), onOffStr(m.Ambient.FireActive),
		m.Styles.KeyCap.Render("[w]"), onOffStr(m.Ambient.WindActive)))

	sb.WriteString(m.Styles.Title.Render("🎵 Curated Lofi Streams (Press n/p to cycle)") + "\n")
	for i, preset := range audio.CuratedLofiPresets {
		prefix := "  "
		style := m.Styles.MutedText
		if i == m.CLIAMP.ActivePreset {
			prefix = "▶ "
			style = m.Styles.BoldAccent
		}
		sb.WriteString(fmt.Sprintf("%s%s (%s)\n", prefix, style.Render(preset.Name), preset.Tag))
	}

	return m.Styles.Box.Width(w).Height(h).Render(sb.String())
}

func (m Model) renderSidebar() string {
	h := m.mainHeight()
	var sb strings.Builder

	// 1. Music Widget Card
	st := m.CLIAMP.GetStatus()
	trackName := audio.CuratedLofiPresets[m.CLIAMP.ActivePreset].Name
	if st.Playing && st.TrackTitle != "" {
		trackName = st.TrackTitle
	}
	stateIcon := "⏸️ Paused"
	if st.Playing {
		stateIcon = "▶️ Playing"
	}

	sb.WriteString(m.Styles.Title.Render("🎵 LOFI RADIO") + "\n")
	sb.WriteString(m.Styles.BoldAccent.Render(truncate(trackName, 22)) + "\n")
	spec := RenderAudioSpectrum(18, st.Playing, m.Scene.FrameCount, m.SpectrumMode, m.Styles.BoldAccent)
	sb.WriteString("Spectrum: " + spec + "\n")
	sb.WriteString("State: " + stateIcon + "\n")
	sb.WriteString(m.Styles.MutedText.Render("[Space] Play/Pause • [n/p]") + "\n\n")

	sb.WriteString(strings.Repeat("─", 24) + "\n\n")

	// 2. Focus Timer Card
	sb.WriteString(m.Styles.Title.Render("⏱️ FOCUS TIMER") + "\n")
	sb.WriteString(fmt.Sprintf("Time: %s\n", m.Styles.Header.Render(m.Timer.FormattedTime())))
	sb.WriteString(fmt.Sprintf("Mode: %s\n", m.Timer.Mode.String()))
	sb.WriteString(m.Styles.MutedText.Render("[x] Start/Pause • [z] Reset") + "\n\n")

	sb.WriteString(strings.Repeat("─", 24) + "\n\n")

	// 3. Quick Toggles Card
	sb.WriteString(m.Styles.Title.Render("✨ QUICK CONTROLS") + "\n")
	sb.WriteString(fmt.Sprintf("%s Catalog: %s\n", m.Styles.KeyCap.Render("[V]"), m.Styles.MutedText.Render("Browse 370+")))
	sb.WriteString(fmt.Sprintf("%s Theme: %s\n", m.Styles.KeyCap.Render("[T]"), truncate(m.Styles.Theme.Name, 10)))
	sb.WriteString(fmt.Sprintf("%s Filter: %s\n", m.Styles.KeyCap.Render("[C]"), truncate(m.Scene.ColorFilter.String(), 10)))
	sb.WriteString(fmt.Sprintf("%s Overlay: %s\n", m.Styles.KeyCap.Render("[P]"), truncate(m.Scene.ParticleOverlay.String(), 10)))
	sb.WriteString(fmt.Sprintf("%s Spectrum: %s\n", m.Styles.KeyCap.Render("[M]"), truncate(m.SpectrumMode.String(), 10)))
	sb.WriteString(fmt.Sprintf("%s Zen Mode: %s\n", m.Styles.KeyCap.Render("[Shift+F]"), m.Styles.MutedText.Render("Minimal")))
	sb.WriteString(fmt.Sprintf("%s Favorite: %s\n", m.Styles.KeyCap.Render("[*]"), m.Styles.MutedText.Render("Toggle")))

	return m.Styles.Box.Width(28).Height(h).Render(sb.String())
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
  [V]      Open Categorized Scene Catalog Modal (370+ Scenes)
  [/]      Search / Filter Scenes by Keyword (in Catalog)
  [*]      Bookmark / Toggle Favorite Scene Preset
  [C]      Cycle Color Filter (Natural, Warm Amber, Cool Synth, Phosphor, Sepia)
  [P]      Cycle Screen Particle Overlay (Sakura, Embers, Snow, Matrix)
  [T]      Open Interactive Theme Palette Selector Modal (10 Themes)
  [M]      Cycle Audio Spectrum Visualizer Style (Bars, Dots, Braille, EQ)
  [Shift+F] Toggle Zen Minimal Focus Mode (Zero Clutter)
  [v]      Cycle Scene Preset (Rainy City, Forest Cabin, Classic, Observatory)
  [l]      Toggle Desk Lamp (ON/OFF)
  [f]      Toggle Fireplace (ON/OFF)
  [s]      Toggle Fairy Lights (ON/OFF)
  [w]      Cycle Weather (Clear, Rain, Snow, Storm, Fog)
  [t]      Cycle Time of Day (Dawn, Day, Sunset, Night, Auto)
  [c]      Cycle Cat Breed (Tabby, Tuxedo, Calico, Black)
  [g]      Toggle Ghostty GPU / Half-Block Canvas

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

func (m Model) renderZenView() string {
	m.Scene.RenderFrame()
	artCanvas := strings.TrimRight(m.Scene.Canvas.Render(), "\r\n")

	st := m.CLIAMP.GetStatus()
	spec := RenderAudioSpectrum(16, st.Playing, m.Scene.FrameCount, m.SpectrumMode, m.Styles.BoldAccent)
	breath := GetZenBreathState(m.Scene.FrameCount)

	cardContent := fmt.Sprintf("%s\n\n⏱️ %s   |   %s\n\n%s  %s\n%s\n\n%s %s",
		m.Styles.Title.Render("🌿 ZEN FOCUS MODE"),
		m.Styles.BoldAccent.Render(m.Timer.FormattedTime()),
		m.timerStateStr(),
		m.Styles.Title.Render(breath.PhaseName),
		m.Styles.Header.Render(breath.RingVisual),
		m.Styles.MutedText.Render("  "+breath.Instruction),
		m.Styles.MutedText.Render("Spectrum: "),
		spec,
	)

	floatingCard := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Title).
		Padding(1, 3).
		Render(cardContent)

	overlay := lipgloss.Place(m.Width, m.Height-2, lipgloss.Center, lipgloss.Center, floatingCard)
	footer := m.Styles.Footer.Render("Press [Esc] or [Shift+F] to exit Zen Mode • [Space] Audio • [x] Timer • [v] Cycle Scene")

	return artCanvas + "\n" + overlay + "\n" + footer
}

func (m Model) renderThemePickerModal(bg string) string {
	var listLines []string
	for i, theme := range AllThemes {
		cursor := "  "
		style := m.Styles.MutedText
		if i == m.PickerThemeIdx {
			cursor = "👉"
			style = m.Styles.BoldAccent
		}

		activeBadge := ""
		if theme.Name == m.Styles.Theme.Name {
			activeBadge = " " + m.Styles.StatusBadge.Render("ACTIVE")
		}

		// Color preview block
		bBlock := lipgloss.NewStyle().Background(theme.Border).Render("  ")
		tBlock := lipgloss.NewStyle().Background(theme.Title).Render("  ")
		aBlock := lipgloss.NewStyle().Background(theme.Accent).Render("  ")
		hBlock := lipgloss.NewStyle().Background(theme.Highlight).Render("  ")

		swatch := fmt.Sprintf("[%s%s%s%s]", bBlock, tBlock, aBlock, hBlock)
		listLines = append(listLines, fmt.Sprintf("%s %s %s%s", cursor, swatch, style.Render(theme.Name), activeBadge))
	}

	modalText := fmt.Sprintf("🎨 THEME PALETTE CUSTOMIZER\n\nSelect color palette theme for Gottage TUI:\n\n%s\n\n%s",
		strings.Join(listLines, "\n\n"),
		m.Styles.MutedText.Render("Nav: [j/k] Select Theme • [Enter] Apply • [Esc/T] Close"),
	)

	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Width(65).
		Render(modalText)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}

func (m Model) renderScenePickerModal(bg string) string {
	var scenes []art.SceneType
	if m.PickerSearchMode || strings.TrimSpace(m.PickerSearchInput.Value()) != "" {
		scenes = m.getFilteredScenes()
	} else {
		scenes = art.GetScenesByCategory(m.PickerCategoryIdx)
	}

	var categoryTabs []string
	for i, cName := range art.SceneCategories {
		if i == m.PickerCategoryIdx && !m.PickerSearchMode && strings.TrimSpace(m.PickerSearchInput.Value()) == "" {
			categoryTabs = append(categoryTabs, m.Styles.ActiveTab.Render(cName))
		} else {
			categoryTabs = append(categoryTabs, m.Styles.InactiveTab.Render(cName))
		}
	}
	tabsRow := strings.Join(categoryTabs, " ")

	searchHeader := ""
	if m.PickerSearchMode || strings.TrimSpace(m.PickerSearchInput.Value()) != "" {
		searchHeader = fmt.Sprintf("🔍 Filter: %s (%d matches)", m.PickerSearchInput.View(), len(scenes))
	} else {
		searchHeader = fmt.Sprintf("Category: %s (%d scenes)   (Press [/] to Search)", m.Styles.Title.Render(art.SceneCategories[m.PickerCategoryIdx]), len(scenes))
	}

	var listLines []string
	startIdx := 0
	if m.PickerSceneIdx > 8 {
		startIdx = m.PickerSceneIdx - 8
	}
	endIdx := startIdx + 12
	if endIdx > len(scenes) {
		endIdx = len(scenes)
	}

	if len(scenes) == 0 {
		listLines = append(listLines, m.Styles.MutedText.Render("  No matching scenes found!"))
	} else {
		for i := startIdx; i < endIdx; i++ {
			st := scenes[i]
			cursor := "  "
			style := m.Styles.MutedText
			if i == m.PickerSceneIdx {
				cursor = "👉"
				style = m.Styles.BoldAccent
			}
			activeBadge := ""
			if st == m.Scene.CurrentScene {
				activeBadge = " " + m.Styles.StatusBadge.Render("ACTIVE")
			}
			favBadge := ""
			if m.FavoriteScenes != nil && m.FavoriteScenes[st] {
				favBadge = " ⭐"
			}
			listLines = append(listLines, fmt.Sprintf("%s %s%s%s", cursor, style.Render(st.String()), favBadge, activeBadge))
		}
	}

	modalText := fmt.Sprintf("🎨 SCENE CATALOG BROWSER (%d Total Scenes)\n\n%s\n\n%s\n\n%s\n\n%s",
		len(art.SceneOrder),
		tabsRow,
		searchHeader,
		strings.Join(listLines, "\n"),
		m.Styles.MutedText.Render("Nav: [h/l] Category • [/] Search • [f/*] Toggle Favorite • [j/k] Select • [Enter] Apply • [Esc/V] Close"),
	)

	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Width(80).
		Render(modalText)

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
