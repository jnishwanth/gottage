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
	if w < 78 {
		return 78
	}
	if w > 120 {
		return 120
	}
	return w
}

func (m Model) mainHeight() int {
	h := m.Height - 6
	if h < 32 {
		return 32
	}
	if h > 52 {
		return 52
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
	footerText := ""
	if m.StatusMsg != "" && time.Since(m.StatusMsgTime) < 4*time.Second {
		footerText = m.Styles.StatusBadge.Render(" " + m.StatusMsg + " ")
	} else {
		var liveBadges []string
		st := m.CLIAMP.GetStatus()
		if st.Playing {
			liveBadges = append(liveBadges, "🎵 Playing")
		}
		if m.Ambient.RainActive {
			liveBadges = append(liveBadges, "🌧️ Rain")
		}
		if m.Ambient.FireActive {
			liveBadges = append(liveBadges, "🔥 Fire")
		}
		if m.Ambient.WindActive {
			liveBadges = append(liveBadges, "🌬️ Wind")
		}
		if m.Ambient.CafeActive {
			liveBadges = append(liveBadges, "☕ Cafe")
		}
		if m.Ambient.CricketsActive {
			liveBadges = append(liveBadges, "🦗 Crickets")
		}
		if m.Timer.Running {
			liveBadges = append(liveBadges, "⏱️ "+m.Timer.FormattedTime())
		}
		if m.Scene.ColorFilter != art.FilterNormal {
			liveBadges = append(liveBadges, "🎨 "+m.Scene.ColorFilter.String())
		}
		if m.Scene.ParticleOverlay != art.OverlayNone {
			liveBadges = append(liveBadges, "✨ "+m.Scene.ParticleOverlay.String())
		}
		if m.Scene.ShaderEffect != art.ShaderNone {
			liveBadges = append(liveBadges, "🌌 "+m.Scene.ShaderEffect.String())
		}
		if m.Scene.VignetteOn {
			liveBadges = append(liveBadges, "🖼️ Vignette")
		}
		if m.Scene.CRTScanlinesOn {
			liveBadges = append(liveBadges, "📺 CRT")
		}
		if m.Scene.RenderMode != art.RenderHalfBlock {
			liveBadges = append(liveBadges, "🎨 "+m.Scene.RenderMode.String())
		}
		if m.Scene.AutoCycleMode != art.AutoCycleOff {
			liveBadges = append(liveBadges, "🎬 "+m.Scene.AutoCycleMode.String())
		}
		if m.Scene.CatBliss > 0 {
			liveBadges = append(liveBadges, fmt.Sprintf("🐾 Bliss: %d%%", m.Scene.CatBliss))
		}

		prefix := ""
		if len(liveBadges) > 0 {
			prefix = strings.Join(liveBadges, " • ") + "  |  "
		}
		footerText = prefix + "Press [?] Help • [K] Pet Cat 🐾 • [Q] Quote ✨ • [F] Zen 🌿 • [:/Ctrl+P] Palette • [Shift+X] Mixer • [M] Spectrum • [q] quit"
	}
	footer := m.Styles.Footer.Render(footerText)

	mainView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		content,
		"",
		footer,
	)

	// Command Palette Modal
	if m.ShowCommandPalette {
		return m.renderCommandPaletteModal(mainView)
	}

	// 2D Audio Visualizer & Spectrum Studio Modal
	if m.ShowSpectrumModal {
		return m.renderSpectrumModal(mainView)
	}

	// 5-Channel Ambient Soundscape Mixer Modal
	if m.ShowSoundMixerModal {
		return m.renderSoundMixerModal(mainView)
	}

	// Preset Studio & Custom Preset Manager Modal
	if m.ShowPresetModal {
		return m.renderPresetModal(mainView)
	}

	// Scene Info & Palette Inspector Modal
	if m.ShowSceneInfoModal {
		return m.renderSceneInfoModal(mainView)
	}

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
	tabs := []struct {
		tab   ActiveTab
		short string
	}{
		{TabArt, "1 🏡 Room"},
		{TabTimer, "2 ⏱️ Timer"},
		{TabTasks, "3 📝 Tasks"},
		{TabAmbiance, "4 🌿 Ambiance"},
	}

	var tabStrings []string
	for _, item := range tabs {
		idxStr := fmt.Sprintf("[%s]", item.short)
		if m.ActiveTab == item.tab {
			tabStrings = append(tabStrings, m.Styles.ActiveTab.Render(idxStr))
		} else {
			tabStrings = append(tabStrings, m.Styles.InactiveTab.Render(idxStr))
		}
	}
	tabBar := strings.Join(tabStrings, " ")

	title := m.Styles.Title.Render("🏡 GOTTAGE") + " " + m.Styles.MutedText.Render("• Virtual Cottage")
	clockBadge := m.Styles.PillBadge.Render("🕒 " + time.Now().Format("15:04:05"))
	themeBadge := m.Styles.StatusBadge.Render(" " + m.Styles.Theme.Name + " ")

	var headerBadges []string
	headerBadges = append(headerBadges, clockBadge)
	if m.Timer.Running {
		headerBadges = append(headerBadges, m.Styles.BoldAccent.Render(" ⏱️ "+m.Timer.FormattedTime()+" "))
	}
	st := m.CLIAMP.GetStatus()
	if st.Playing {
		miniSpec := RenderAudioSpectrum(6, true, m.Scene.FrameCount, SpectrumBars, m.Styles.BoldAccent)
		headerBadges = append(headerBadges, m.Styles.StatusBadge.Render(" 🎵 "+miniSpec+" "))
	}
	if m.Beverage != nil {
		headerBadges = append(headerBadges, m.Styles.StatusBadge.Render(" "+m.Beverage.StatusBadge()+" "))
	}
	headerBadges = append(headerBadges, m.Styles.MutedText.Render(m.catBlissBadge()))
	headerBadges = append(headerBadges, themeBadge)

	badgeBar := strings.Join(headerBadges, "  ")

	return lipgloss.JoinHorizontal(lipgloss.Center, title, "   ", badgeBar, "     ", tabBar)
}

func (m Model) catBlissBadge() string {
	b := m.Scene.CatBliss
	switch {
	case b >= 90:
		return fmt.Sprintf("😻 %d%% Euphoric", b)
	case b >= 70:
		return fmt.Sprintf("😸 %d%% Blissful", b)
	case b >= 40:
		return fmt.Sprintf("😺 %d%% Content", b)
	default:
		return fmt.Sprintf("🐱 %d%% Sleepy", b)
	}
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

	lampIcon := "🌙"
	if m.Scene.LampOn {
		lampIcon = "💡"
	}
	fireIcon := "💤"
	if m.Scene.FireOn {
		fireIcon = "🔥"
	}

	metaHeader := fmt.Sprintf(" %s%s  │  %s  │  %s  │  Lamp: %s  │  Fire: %s  │  Shader: %s",
		m.Styles.BoldAccent.Render(m.Scene.CurrentScene.String()),
		favBadge,
		m.Scene.EffectiveTimeOfDay().String(),
		m.Scene.Weather.String(),
		onOffStr(m.Scene.LampOn),
		onOffStr(m.Scene.FireOn),
		m.Scene.ShaderEffect.String(),
	)

	bevStr := "☕"
	if m.Beverage != nil {
		bevStr = m.Beverage.CurrentDrink().Icon
	}

	quickDeck := fmt.Sprintf(" %s Scene  %s Weather  %s Time  %s Lamp %s  %s Fire %s  %s Brew %s  %s Sip ♨️  %s Pet 🐾",
		m.Styles.KeyCap.Render("v"),
		m.Styles.KeyCap.Render("w"),
		m.Styles.KeyCap.Render("t"),
		m.Styles.KeyCap.Render("l"), lampIcon,
		m.Styles.KeyCap.Render("f"), fireIcon,
		m.Styles.KeyCap.Render("J"), bevStr,
		m.Styles.KeyCap.Render("d"),
		m.Styles.KeyCap.Render("K"),
	)

	content := metaHeader + "\n\n" + artStr + "\n\n" + m.Styles.MutedText.Render(quickDeck)
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
	items := m.Tasks.Items
	filtered := m.Tasks.FilteredItems()

	completedCount := 0
	for _, item := range items {
		if item.Completed {
			completedCount++
		}
	}

	total := len(items)
	ratio := 0.0
	if total > 0 {
		ratio = float64(completedCount) / float64(total)
	}

	barWidth := w - 42
	if barWidth < 15 {
		barWidth = 15
	}
	filled := int(ratio * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	pBar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	header := fmt.Sprintf("📝 COTTAGE TO-DO NOTEBOOK  [%s] %.0f%% (%d/%d Completed)",
		m.Styles.ProgressBar.Render(pBar),
		ratio*100,
		completedCount,
		total,
	)
	sb.WriteString(m.Styles.Header.Render(header) + "\n\n")

	filterTabs := fmt.Sprintf("Filter Mode: %s    (Press [f] to cycle)", m.Styles.BoldAccent.Render(m.Tasks.Filter.String()))
	sb.WriteString(filterTabs + "\n\n")

	if m.AddingTask {
		sb.WriteString(m.Styles.Title.Render("Add New Focus Task:") + "\n\n")
		sb.WriteString(m.TextInput.View() + "\n\n")
		sb.WriteString(m.Styles.MutedText.Render("Press Enter to save • Esc to cancel"))
		return m.Styles.Box.Width(w).Height(h).Render(sb.String())
	}

	if len(filtered) == 0 {
		sb.WriteString(m.Styles.MutedText.Render("No tasks in current filter! Press [f] to change filter or [a] to add a task."))
	} else {
		for i, item := range filtered {
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

			priBadge := fmt.Sprintf("[%s]", item.Priority.String())
			line := fmt.Sprintf("%s %s %s %s", cursor, m.Styles.BoldAccent.Render(box), m.Styles.MutedText.Render(priBadge), titleStyle.Render(item.Title))
			sb.WriteString(line + "\n")
		}
	}

	sb.WriteString("\n" + strings.Repeat("─", w-6) + "\n")
	sb.WriteString(m.Styles.MutedText.Render("Actions: ") +
		m.Styles.KeyCap.Render("[a]") + " Add  " +
		m.Styles.KeyCap.Render("[Enter]") + " Toggle  " +
		m.Styles.KeyCap.Render("[p]") + " Priority  " +
		m.Styles.KeyCap.Render("[f]") + " Filter  " +
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
	favScopeStr := "All 370+ Scenes 🎨"
	if m.Scene.AutoCycleFavsOnly {
		favScopeStr = "Favorites Only ⭐"
	}
	sb.WriteString(fmt.Sprintf("  %s Scene Preset:   %s    %s Random Scene: %s\n",
		m.Styles.KeyCap.Render("[v/b]"), m.Styles.BoldAccent.Render(m.Scene.CurrentScene.String()),
		m.Styles.KeyCap.Render("[R]"), m.Styles.MutedText.Render("Surprise Me")))
	sb.WriteString(fmt.Sprintf("  %s Auto-Cycle Mode:%s    %s Auto-Cycle Scope: %s\n",
		m.Styles.KeyCap.Render("[Y]"), m.Scene.AutoCycleMode.String(),
		m.Styles.KeyCap.Render("[U]"), favScopeStr))
	sb.WriteString(fmt.Sprintf("  %s Color Filter:   %s    %s Animation Speed: %s\n",
		m.Styles.KeyCap.Render("[C]"), m.Scene.ColorFilter.String(),
		m.Styles.KeyCap.Render("[S]"), m.Scene.AnimSpeed.String()))
	sb.WriteString(fmt.Sprintf("  %s Screen Overlay: %s    %s Vignette Edge: %s\n",
		m.Styles.KeyCap.Render("[P]"), m.Scene.ParticleOverlay.String(),
		m.Styles.KeyCap.Render("[G]"), onOffStr(m.Scene.VignetteOn)))
	sb.WriteString(fmt.Sprintf("  %s CRT Scanlines:  %s    %s Time of Day:   %s\n",
		m.Styles.KeyCap.Render("[O]"), onOffStr(m.Scene.CRTScanlinesOn),
		m.Styles.KeyCap.Render("[t]"), m.Scene.EffectiveTimeOfDay().String()))
	sb.WriteString(fmt.Sprintf("  %s Weather Mode:   %s\n", m.Styles.KeyCap.Render("[w]"), m.Scene.Weather.String()))
	sb.WriteString(fmt.Sprintf("  %s Cat Companion:  %s\n", m.Styles.KeyCap.Render("[c]"), m.Scene.CatBreed.String()))
	sb.WriteString(fmt.Sprintf("  %s Desk Lamp:      %s   %s Fireplace: %s   %s Fairy Lights: %s\n\n",
		m.Styles.KeyCap.Render("[l]"), onOffStr(m.Scene.LampOn),
		m.Styles.KeyCap.Render("[f]"), onOffStr(m.Scene.FireOn),
		m.Styles.KeyCap.Render("[s]"), onOffStr(m.Scene.FairyLightsOn)))

	sb.WriteString(m.Styles.Title.Render("🌧️ 5-Channel Procedural Ambient Sound Console") + "  " + m.Styles.MutedText.Render("(Press [Shift+X] for Equalizer Mixer)") + "\n")
	sb.WriteString(fmt.Sprintf("  %s Rain: %s (%.0f%%)    %s Fireplace: %s (%.0f%%)    %s Wind Breeze: %s (%.0f%%)\n",
		m.Styles.KeyCap.Render("[r]"), onOffStr(m.Ambient.RainActive), m.Ambient.RainVol*100,
		m.Styles.KeyCap.Render("[m]"), onOffStr(m.Ambient.FireActive), m.Ambient.FireVol*100,
		m.Styles.KeyCap.Render("[W]"), onOffStr(m.Ambient.WindActive), m.Ambient.WindVol*100))
	sb.WriteString(fmt.Sprintf("  %s Cafe Chatter: %s (%.0f%%)    %s Night Crickets: %s (%.0f%%)\n",
		m.Styles.KeyCap.Render("[K]"), onOffStr(m.Ambient.CafeActive), m.Ambient.CafeVol*100,
		m.Styles.KeyCap.Render("[e]"), onOffStr(m.Ambient.CricketsActive), m.Ambient.CricketsVol*100))
	sb.WriteString(fmt.Sprintf("  %s Weather Audio Auto-Sync: %s\n\n",
		m.Styles.KeyCap.Render("[A]"), onOffStr(m.AutoSyncAudio)))

	if m.Beverage != nil {
		drink := m.Beverage.CurrentDrink()
		wCompleted := m.Beverage.Warmth / 10
		warmthBar := strings.Repeat("█", wCompleted) + strings.Repeat("░", 10-wCompleted)
		sb.WriteString(m.Styles.Title.Render("☕ Artisan Hot Beverage & Tea Bar") + "  " + m.Styles.MutedText.Render("(Press [J] to brew next, [d] or [Ctrl+B] to sip)") + "\n")
		sb.WriteString(fmt.Sprintf("  Active Drink: %s    Warmth: [%s] %d%%    Sips Taken: %d\n",
			m.Styles.BoldAccent.Render(drink.Name),
			m.Styles.ProgressBar.Render(warmthBar),
			m.Beverage.Warmth,
			m.Beverage.SipCount))
		sb.WriteString(fmt.Sprintf("  Aroma & Notes: %s\n\n", m.Styles.MutedText.Render(drink.FlavorNotes)))
	}

	sb.WriteString(m.Styles.Title.Render("🎵 Curated Lofi Streams & Volume (Press n/p to cycle, +/- volume)") + "\n")
	volStr := "80%"
	if m.CLIAMP != nil {
		volStr = m.CLIAMP.GetStatus().Volume
		if volStr == "" {
			volStr = fmt.Sprintf("%d%%", m.CLIAMP.VolumePercent)
		}
	}
	sb.WriteString(fmt.Sprintf("  Volume: %s %s\n", m.Styles.BoldAccent.Render("🔊 "+volStr), m.Styles.MutedText.Render("([+] Up / [-] Down)")))
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

	// 1. Music & Soundscape Card
	st := m.CLIAMP.GetStatus()
	trackName := audio.CuratedLofiPresets[m.CLIAMP.ActivePreset].Name
	if st.Playing && st.TrackTitle != "" {
		trackName = st.TrackTitle
	}
	stateIcon := "⏸️ Paused"
	if st.Playing {
		stateIcon = "▶️ Playing"
	}
	volStr := "80%"
	if st.Volume != "" {
		volStr = st.Volume
	} else if m.CLIAMP != nil {
		volStr = fmt.Sprintf("%d%%", m.CLIAMP.VolumePercent)
	}

	sb.WriteString(m.Styles.Title.Render("🎧 24/7 LOFI RADIO") + "\n")
	sb.WriteString(m.Styles.BoldAccent.Render(truncate(trackName, 22)) + "\n")
	spec := RenderAudioSpectrum(18, st.Playing, m.Scene.FrameCount, m.SpectrumMode, m.Styles.BoldAccent)
	sb.WriteString("Wave: " + spec + "\n")
	sb.WriteString(fmt.Sprintf("%s • 🔊 %s\n", stateIcon, volStr))

	var activeSounds []string
	if m.Ambient.RainActive {
		activeSounds = append(activeSounds, "🌧️")
	}
	if m.Ambient.FireActive {
		activeSounds = append(activeSounds, "🔥")
	}
	if m.Ambient.WindActive {
		activeSounds = append(activeSounds, "🌬️")
	}
	if m.Ambient.CafeActive {
		activeSounds = append(activeSounds, "☕")
	}
	if m.Ambient.CricketsActive {
		activeSounds = append(activeSounds, "🦗")
	}
	if len(activeSounds) > 0 {
		sb.WriteString("Sound: " + strings.Join(activeSounds, " ") + "\n")
	}
	sb.WriteString(m.Styles.MutedText.Render("[Space] Play • [n/p] • [X] EQ") + "\n\n")

	sb.WriteString(strings.Repeat("─", 24) + "\n\n")

	// 2. Focus Pomodoro Card
	sb.WriteString(m.Styles.Title.Render("⏱️ FOCUS TIMER") + "\n")
	sb.WriteString(fmt.Sprintf("Time: %s\n", m.Styles.Header.Render(m.Timer.FormattedTime())))
	sb.WriteString(fmt.Sprintf("Mode: %s\n", m.Timer.Mode.String()))

	barWidth := 18
	completed := int(m.Timer.ProgressRatio() * float64(barWidth))
	if completed > barWidth {
		completed = barWidth
	}
	progressBar := strings.Repeat("█", completed) + strings.Repeat("░", barWidth-completed)
	sb.WriteString(fmt.Sprintf("[%s]\n", m.Styles.ProgressBar.Render(progressBar)))
	sb.WriteString(m.Styles.MutedText.Render("[x] Start/Pause • [z] Reset") + "\n\n")

	sb.WriteString(strings.Repeat("─", 24) + "\n\n")

	// 3. Cozy Warm Beverage Card
	if m.Beverage != nil {
		drink := m.Beverage.CurrentDrink()
		sb.WriteString(m.Styles.Title.Render("☕ WARM BEVERAGE") + "\n")
		sb.WriteString(m.Styles.BoldAccent.Render(truncate(drink.Name, 22)) + "\n")
		wCompleted := m.Beverage.Warmth / 10
		wBar := strings.Repeat("█", wCompleted) + strings.Repeat("░", 10-wCompleted)
		sb.WriteString(fmt.Sprintf("Warmth: [%s] %d%%\n", m.Styles.ProgressBar.Render(wBar), m.Beverage.Warmth))
		sb.WriteString(m.Styles.MutedText.Render(truncate(drink.Aroma, 22)) + "\n")
		sb.WriteString(m.Styles.MutedText.Render("[d] Sip ♨️ • [J] Brew Next") + "\n\n")
		sb.WriteString(strings.Repeat("─", 24) + "\n\n")
	}

	// 4. Mindfulness Prompt Card
	q := m.Inspiration.CurrentQuote()
	sb.WriteString(m.Styles.Title.Render("✨ MINDFULNESS PROMPT") + "\n")
	sb.WriteString(m.Styles.BoldAccent.Render("\""+truncate(q.Quote, 22)+"\"") + "\n")
	sb.WriteString(m.Styles.MutedText.Render("— "+q.Author) + " " + m.Styles.KeyCap.Render("[Q]") + "\n\n")

	sb.WriteString(strings.Repeat("─", 24) + "\n\n")

	// 5. Quick Companion Controls Card
	sb.WriteString(m.Styles.Title.Render("🌿 COZY COMPANION") + "\n")
	catName := "Sparky"
	switch m.Scene.CatBreed {
	case art.CatTuxedo:
		catName = "Sylvester"
	case art.CatCalico:
		catName = "Mochi"
	case art.CatBlack:
		catName = "Luna"
	}
	sb.WriteString(fmt.Sprintf("🐾 %s: %s\n", catName, m.Styles.Header.Render(fmt.Sprintf("%d%% Bliss 💖", m.Scene.CatBliss))))
	sb.WriteString(m.Styles.MutedText.Render("Press [K] to pet cat companion!") + "\n\n")

	sb.WriteString(fmt.Sprintf("%s Catalog  %s Presets\n", m.Styles.KeyCap.Render("[V]"), m.Styles.KeyCap.Render("[Shift+S]")))
	sb.WriteString(fmt.Sprintf("%s Themes   %s Zen Focus\n", m.Styles.KeyCap.Render("[T]"), m.Styles.KeyCap.Render("[Shift+F]")))
	sb.WriteString(fmt.Sprintf("%s Palette  %s Shortcuts\n", m.Styles.KeyCap.Render("[:]"), m.Styles.KeyCap.Render("[?]")))

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
  [Shift+S] Open Preset Studio & Custom Manager Modal
  [Shift+X] Open 5-Channel Soundscape Equalizer & Mixer Modal
  [H]      Cycle Canvas Animation Shader (Rain-on-Glass, Fireflies, Aurora, Waveform)
  [V]      Open Categorized Scene Catalog Modal (370+ Scenes)
  [/]      Search / Filter Scenes by Keyword (in Catalog)
  [*]      Bookmark / Toggle Favorite Scene Preset
  [R]      Jump to Random Scene Preset (Surprise Me)
  [Y]      Cycle Auto-Cycle Scene Rotation Mode (1m, 3m, 5m, 10m, Timer-Linked)
  [U]      Toggle Auto-Cycle Scope (Favorites Only / All 370+ Scenes)
  [v/b]    Cycle Next / Previous Scene Preset

  COZY BEVERAGE & COMPANION:
  [J]      Brew Next Warm Drink (Coffee, Matcha, Earl Grey, Chai, Cocoa, Chamomile)
  [d]      Take a Soothing Warm Sip from Hot Mug ([Ctrl+B])
  [K]      Pet Cozy Cat Companion (Heart Particles & Purrs)
  [Q]      Cycle Mindfulness & Focus Inspiration Quotes

  [C]      Cycle Color Filter (Natural, Warm Amber, Cool Synth, Phosphor, Sepia, Midnight)
  [P]      Cycle Screen Particle Overlay (Sakura, Embers, Snow, Matrix, Stars)
  [S]      Open Preset Studio Modal (Shift+S) / Cycle Speed
  [A]      Toggle Weather-Audio Auto-Sync (ON/OFF)
  [T]      Open Interactive Theme Palette Selector Modal (10 Themes)
  [M]      Cycle Audio Spectrum Visualizer Style (Bars, Dots, Braille, EQ, Wave)
  [Shift+F] Toggle Zen Minimal Focus Mode (Zero Clutter)
  [l]      Toggle Desk Lamp (ON/OFF)
  [f]      Toggle Fireplace (ON/OFF) (in Art View)
  [s]      Toggle Fairy Lights (ON/OFF)
  [w]      Cycle Weather (Clear, Rain, Snow, Storm, Fog)
  [t]      Cycle Time of Day (Dawn, Day, Sunset, Night, Auto)
  [c]      Cycle Cat Breed (Tabby, Tuxedo, Calico, Black)
  [g]      Toggle Ghostty GPU / Half-Block Canvas

  AUDIO & LOFI CONTROLS:
  [space]  Play/Pause CLIAMP / Audio
  [n/p]    Next / Previous Lofi Track Preset
  [+/-]    Adjust Lofi Audio Volume (Up / Down)
  [r]      Toggle Rain Ambient Sound
  [m]      Toggle Fireplace Ambient Sound
  [W]      Toggle Wind Ambient Sound
  [K]      Toggle Cafe Chatter Ambient Sound
  [e]      Toggle Night Crickets Ambient Sound

  TIMER & TASKS:
  [x]      Start / Pause Pomodoro Timer
  [z]      Reset Pomodoro Timer
  [a]      Add new Task (in Tasks View)
  [p]      Cycle Task Priority (High, Med, Low) (in Tasks View)
  [f]      Cycle Task Filter (All, Active, Completed) (in Tasks View)
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
	quote := m.Inspiration.CurrentQuote()

	headerBar := lipgloss.JoinHorizontal(lipgloss.Center,
		m.Styles.Title.Render("🌿 ZEN FOCUS SANCTUARY"),
		"  •  ",
		m.Styles.BoldAccent.Render(m.Scene.CurrentScene.String()),
		"  •  ",
		m.Styles.PillBadge.Render("🕒 "+time.Now().Format("15:04:05")),
	)

	statusLine := fmt.Sprintf("⏱️ %s (%s)  │  %s  │  🧘 %s %s  │  🎵 %s",
		m.Styles.BoldAccent.Render(m.Timer.FormattedTime()),
		m.timerStateStr(),
		m.Styles.BoldAccent.Render(m.catBlissBadge()),
		m.Styles.Title.Render(breath.PhaseName),
		m.Styles.Header.Render(breath.RingVisual),
		spec,
	)

	quoteLine := fmt.Sprintf("✨ %s %s",
		m.Styles.BoldAccent.Render("\""+quote.Quote+"\""),
		m.Styles.MutedText.Render("— "+quote.Author),
	)

	zenBox := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Title).
		Padding(0, 2).
		Render(lipgloss.JoinVertical(lipgloss.Center,
			statusLine,
			quoteLine,
		))

	footer := m.Styles.Footer.Render("Press [Esc] or [F] to exit Zen Mode • [K] Pet Cat 🐾 • [Q] Quote ✨ • [Space] Audio • [x] Timer • [v] Cycle Scene")

	content := lipgloss.JoinVertical(lipgloss.Center,
		headerBar,
		"",
		artCanvas,
		"",
		zenBox,
		"",
		footer,
	)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) renderPresetModal(bg string) string {
	dialogWidth := 76
	title := m.Styles.Title.Render(" 💾 ATMOSPHERE PRESET STUDIO & CUSTOM MANAGER ")

	var content string

	if m.PresetSaving {
		content = fmt.Sprintf(
			"%s\n\n%s\n\n%s  %s",
			m.Styles.BoldAccent.Render("Save Current Ambiance as Custom Preset:"),
			m.PresetInputText.View(),
			m.Styles.KeyCap.Render(" Enter ")+" Save Preset",
			m.Styles.KeyCap.Render(" Esc ")+" Cancel",
		)
	} else {
		var presetRows []string
		presets := m.Presets.Presets

		for i, p := range presets {
			cursor := "  "
			rowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#A3A6B4"))
			if i == m.PresetIdx {
				cursor = "▶ "
				rowStyle = m.Styles.BoldAccent
			}

			details := fmt.Sprintf("%s • %s • Shader: %s", p.Scene.String(), p.Theme, p.ShaderEffect.String())
			row := fmt.Sprintf("%s%s\n   %s", cursor, rowStyle.Render(p.Name), m.Styles.MutedText.Render(details))
			presetRows = append(presetRows, row)
		}

		selectedDesc := ""
		if len(presets) > 0 && m.PresetIdx < len(presets) {
			p := presets[m.PresetIdx]
			selectedDesc = fmt.Sprintf("\n%s\n  %s",
				m.Styles.Header.Render("Preset Description:"),
				m.Styles.MutedText.Render(p.Description),
			)
		}

		controls := fmt.Sprintf(
			"%s Apply  %s Save Current Ambiance  %s Delete  %s Close",
			m.Styles.KeyCap.Render(" Enter "),
			m.Styles.KeyCap.Render(" a "),
			m.Styles.KeyCap.Render(" d "),
			m.Styles.KeyCap.Render(" Esc "),
		)

		content = strings.Join(presetRows, "\n") + selectedDesc + "\n\n" + controls
	}

	dialogBox := m.Styles.Box.Width(dialogWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title,
			"",
			content,
		),
	)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, dialogBox)
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
	endIdx := startIdx + 11
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

	leftCol := fmt.Sprintf("%s\n\n%s\n\n%s",
		tabsRow,
		searchHeader,
		strings.Join(listLines, "\n"),
	)

	// Right Column: Metadata Preview Card
	previewCard := ""
	if len(scenes) > 0 && m.PickerSceneIdx < len(scenes) {
		selScene := scenes[m.PickerSceneIdx]
		meta := art.GetSceneMeta(selScene)

		favStatus := "Not Bookmarked"
		if m.FavoriteScenes != nil && m.FavoriteScenes[selScene] {
			favStatus = "⭐ Bookmarked"
		}

		previewContent := fmt.Sprintf("%s\n\n%s\n\n%s %s\n\n%s\n  %s\n\n%s\n  %s\n\n%s %s",
			m.Styles.Header.Render("👁️ SCENE PREVIEW"),
			m.Styles.Title.Render(truncate(selScene.String(), 28)),
			m.Styles.MutedText.Render("Mood Tag:"), m.Styles.BoldAccent.Render(meta.MoodTag),
			m.Styles.MutedText.Render("Weather Match:"), meta.RecommendedWeather,
			m.Styles.MutedText.Render("Atmosphere:"), m.Styles.MutedText.Render(truncate(meta.Atmosphere, 32)),
			m.Styles.MutedText.Render("Favorites:"), m.Styles.Header.Render(favStatus),
		)

		previewCard = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(m.Styles.Theme.Title).
			Padding(0, 1).
			Width(34).
			Render(previewContent)
	}

	mainBody := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, "  ", previewCard)

	modalText := fmt.Sprintf("🎨 SCENE CATALOG BROWSER (%d Total Scenes)\n\n%s\n\n%s",
		len(art.SceneOrder),
		mainBody,
		m.Styles.MutedText.Render("Nav: [h/l] Category • [/] Search • [f/*] Favorite • [j/k] Select • [Enter] Apply • [Esc/V] Close"),
	)

	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Width(96).
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

func (m Model) renderSceneInfoModal(bg string) string {
	st := m.Scene.CurrentScene
	meta := art.GetSceneMeta(st)

	favStatus := "Not Bookmarked"
	if m.FavoriteScenes != nil && m.FavoriteScenes[st] {
		favStatus = "⭐ Favorites Bookmarked"
	}

	pals := m.Scene.Canvas.ExtractDominantPalette(8)
	var swatchElems []string
	for _, col := range pals {
		hexStr := col.Hex()
		block := lipgloss.NewStyle().
			Background(lipgloss.Color(hexStr)).
			Foreground(lipgloss.Color("#FFFFFF")).
			Render(" " + hexStr + " ")
		swatchElems = append(swatchElems, block)
	}
	paletteRow := strings.Join(swatchElems, " ")
	if paletteRow == "" {
		paletteRow = m.Styles.MutedText.Render("N/A")
	}

	catName := art.SceneCategories[art.CategorizeScene(st)]

	col1 := fmt.Sprintf("%s\n  Preset:       %s\n  Category:     %s\n  Favorites:    %s\n\n%s\n  Mood Tag:     %s\n  Weather Match: %s\n  Atmosphere:   %s",
		m.Styles.Header.Render("🎨 SCENE OVERVIEW"),
		m.Styles.BoldAccent.Render(st.String()),
		m.Styles.Title.Render(catName),
		m.Styles.Header.Render(favStatus),
		m.Styles.Header.Render("🌿 ATMOSPHERIC METADATA"),
		m.Styles.BoldAccent.Render(meta.MoodTag),
		meta.RecommendedWeather,
		m.Styles.MutedText.Render(truncate(meta.Atmosphere, 40)),
	)

	col2 := fmt.Sprintf("%s\n  Resolution:   %dx%d px\n  Frames:       %d\n  Render Mode:  %s\n  Anim Speed:   %s\n  GPU Mode:     %s\n  Color Filter: %s\n  Particle:     %s\n  Vignette:     %s\n  CRT Scanlines: %s\n  Beat Pulse:   %s\n  Sky Blend:    %s",
		m.Styles.Header.Render("⚡ ENGINE METRICS"),
		m.Scene.Canvas.Width, m.Scene.Canvas.Height,
		m.Scene.FrameCount,
		m.Scene.Canvas.RenderMode.String(),
		m.Scene.AnimSpeed.String(),
		onOffStr(m.UseGPURender),
		m.Scene.ColorFilter.String(),
		m.Scene.ParticleOverlay.String(),
		onOffStr(m.Scene.VignetteOn),
		onOffStr(m.Scene.CRTScanlinesOn),
		onOffStr(m.Scene.BeatReactiveParticles),
		onOffStr(m.Scene.SmoothDayNightBlend),
	)

	topCols := lipgloss.JoinHorizontal(lipgloss.Top, col1, "    ", col2)

	modalText := fmt.Sprintf("🎨 SCENE INSPECTOR & PALETTE ANALYSIS\n\n%s\n\n%s\n  %s\n\n%s",
		topCols,
		m.Styles.Header.Render("🖌️ EXTRACTED 8-COLOR SCENE PALETTE"),
		paletteRow,
		m.Styles.MutedText.Render("Quick Toggles: [Shift+E] Render Mode • [B] Beat Pulse • [D] Day/Night Blend • [C] Filter • [P] Particle • [Esc/i] Close"),
	)

	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Width(86).
		Render(modalText)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}

func (m Model) renderCommandPaletteModal(bg string) string {
	allCmds := GetCommandList(&m)
	filtered := FilterCommands(allCmds, m.CommandInput.Value())

	searchHeader := fmt.Sprintf("🔍 QUICK COMMAND PALETTE & MOOD PRESETS   (Press Esc to exit)\n\n  > %s", m.CommandInput.View())

	var listLines []string
	startIdx := 0
	if m.CommandIdx > 5 {
		startIdx = m.CommandIdx - 5
	}
	endIdx := startIdx + 7
	if endIdx > len(filtered) {
		endIdx = len(filtered)
	}

	if len(filtered) == 0 {
		listLines = append(listLines, m.Styles.MutedText.Render("  No matching commands or presets found!"))
	} else {
		for i := startIdx; i < endIdx; i++ {
			cmd := filtered[i]
			cursor := "  "
			titleStyle := m.Styles.MutedText
			if i == m.CommandIdx {
				cursor = "👉"
				titleStyle = m.Styles.BoldAccent
			}

			catBadge := m.Styles.Card.Render(cmd.Category)
			scBadge := m.Styles.KeyCap.Render(cmd.Shortcut)

			line := fmt.Sprintf("%s %s %s %s\n     %s",
				cursor,
				catBadge,
				titleStyle.Render(cmd.Title),
				scBadge,
				m.Styles.MutedText.Render(cmd.Description),
			)
			listLines = append(listLines, line)
		}
	}

	modalText := fmt.Sprintf("%s\n\n%s\n\n%s",
		searchHeader,
		strings.Join(listLines, "\n\n"),
		m.Styles.MutedText.Render(fmt.Sprintf("Nav: [↑/↓] Select (%d/%d matches) • [Enter] Execute • [Esc] Close", len(filtered), len(allCmds))),
	)

	modal := m.Styles.Box.
		BorderForeground(m.Styles.Theme.Accent).
		Width(88).
		Render(modalText)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}

func (m Model) renderSoundMixerModal(bg string) string {
	dialogWidth := 74
	title := m.Styles.Title.Render(" 🎛️ 5-CHANNEL AMBIENT SOUNDSCAPE EQUALIZER & MIXER ")

	channels := []struct {
		name   string
		icon   string
		active bool
		vol    float64
		key    string
	}{
		{name: "Rain Sound", icon: "🌧️", active: m.Ambient.RainActive, vol: m.Ambient.RainVol, key: "[r]"},
		{name: "Fireplace Crackle", icon: "🔥", active: m.Ambient.FireActive, vol: m.Ambient.FireVol, key: "[m]"},
		{name: "Wind Breeze", icon: "🌬️", active: m.Ambient.WindActive, vol: m.Ambient.WindVol, key: "[W]"},
		{name: "Cafe Chatter", icon: "☕", active: m.Ambient.CafeActive, vol: m.Ambient.CafeVol, key: "[K]"},
		{name: "Night Crickets", icon: "🦗", active: m.Ambient.CricketsActive, vol: m.Ambient.CricketsVol, key: "[e]"},
	}

	var rows []string
	for i, ch := range channels {
		cursor := "  "
		style := m.Styles.MutedText
		if i == m.MixerChannelIdx {
			cursor = "👉"
			style = m.Styles.BoldAccent
		}

		barWidth := 20
		filled := int(ch.vol * float64(barWidth))
		if filled > barWidth {
			filled = barWidth
		}
		sliderBar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

		statusStr := onOffStr(ch.active)
		volPct := fmt.Sprintf("%3.0f%%", ch.vol*100)

		row := fmt.Sprintf("%s %s %-18s %s [%s] %s  %s",
			cursor,
			ch.icon,
			style.Render(ch.name),
			statusStr,
			m.Styles.ProgressBar.Render(sliderBar),
			m.Styles.BoldAccent.Render(volPct),
			m.Styles.KeyCap.Render(ch.key),
		)
		rows = append(rows, row)
	}

	controls := fmt.Sprintf("%s Select  %s Vol Up/Down  %s Toggle  %s Close",
		m.Styles.KeyCap.Render(" j/k "),
		m.Styles.KeyCap.Render(" h/l / +/- "),
		m.Styles.KeyCap.Render(" Enter "),
		m.Styles.KeyCap.Render(" Esc/Shift+X "),
	)

	content := strings.Join(rows, "\n\n") + "\n\n" + strings.Repeat("─", dialogWidth-8) + "\n" + controls

	dialogBox := m.Styles.Box.Width(dialogWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title,
			"",
			content,
		),
	)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, dialogBox)
}

func (m Model) renderSpectrumModal(bg string) string {
	dialogWidth := 76
	title := m.Styles.Title.Render(" 🎶 2D AUDIO SPECTRUM & EQUALIZER STUDIO ")

	st := m.CLIAMP.GetStatus()
	spec2D := RenderAudioSpectrum2D(64, 10, st.Playing, m.Scene.FrameCount, m.SpectrumMode, m.Styles)

	modeHeader := fmt.Sprintf("Visualizer Mode: %s  |  Style: %s",
		m.Styles.BoldAccent.Render(m.SpectrumMode.String()),
		m.Styles.Header.Render(fmt.Sprintf("%d/5", int(m.SpectrumMode)+1)),
	)

	trackInfo := "Track: " + m.Styles.Header.Render("24/7 Lofi Focus Radio Stream")
	if st.TrackTitle != "" {
		trackInfo = "Track: " + m.Styles.Header.Render(st.TrackTitle)
	}

	stateBadge := m.Styles.SwitchOff.Render(" PAUSED ⏸️ ")
	if st.Playing {
		stateBadge = m.Styles.SwitchOn.Render(" PLAYING ▶️ ")
	}

	activeAmbients := []string{}
	if m.Ambient.RainActive {
		activeAmbients = append(activeAmbients, "🌧️ Rain")
	}
	if m.Ambient.FireActive {
		activeAmbients = append(activeAmbients, "🔥 Fire")
	}
	if m.Ambient.WindActive {
		activeAmbients = append(activeAmbients, "🌬️ Wind")
	}
	if m.Ambient.CafeActive {
		activeAmbients = append(activeAmbients, "☕ Cafe")
	}
	if m.Ambient.CricketsActive {
		activeAmbients = append(activeAmbients, "🦗 Crickets")
	}

	ambientStr := "None"
	if len(activeAmbients) > 0 {
		ambientStr = strings.Join(activeAmbients, " • ")
	}

	audioCard := fmt.Sprintf("%s  |  %s\nAmbient Layers: %s",
		trackInfo,
		stateBadge,
		m.Styles.BoldAccent.Render(ambientStr),
	)

	controls := fmt.Sprintf("%s Cycle Style  %s Play/Pause  %s Radio 1-5  %s Close",
		m.Styles.KeyCap.Render(" Tab/Space "),
		m.Styles.KeyCap.Render(" p "),
		m.Styles.KeyCap.Render(" 1-5 "),
		m.Styles.KeyCap.Render(" Esc/M "),
	)

	divider := strings.Repeat("─", dialogWidth-8)

	content := lipgloss.JoinVertical(lipgloss.Left,
		modeHeader,
		"",
		spec2D,
		"",
		divider,
		audioCard,
		divider,
		controls,
	)

	dialogBox := m.Styles.Box.Width(dialogWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title,
			"",
			content,
		),
	)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, dialogBox)
}
