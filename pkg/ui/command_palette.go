package ui

import (
	"fmt"
	"strings"
)

// CommandItem defines a single quick action inside the Command Palette.
type CommandItem struct {
	Title       string
	Category    string
	Shortcut    string
	Description string
	Action      func(m *Model)
}

// GetCommandList returns all available actions and mood presets.
func GetCommandList(m *Model) []CommandItem {
	var items []CommandItem

	// Atmosphere Presets (Builtin & User Saved)
	if m.Presets != nil {
		for i, p := range m.Presets.Presets {
			presetCopy := p
			items = append(items, CommandItem{
				Title:       fmt.Sprintf("✨ Preset: %s", presetCopy.Name),
				Category:    "Atmospheric Mood",
				Shortcut:    fmt.Sprintf("Preset %d", i+1),
				Description: presetCopy.Description,
				Action: func(m *Model) {
					m.ApplyPreset(presetCopy)
				},
			})
		}
	}

	// Dialogs & Studio Modals
	items = append(items, CommandItem{
		Title:       "💾 Studio: Preset Studio & Custom Manager",
		Category:    "Dialogs",
		Shortcut:    "[ Shift+S ]",
		Description: "Open preset manager modal to apply, save, or delete custom atmosphere presets",
		Action: func(m *Model) {
			m.ShowPresetModal = true
			m.PresetIdx = 0
			m.PresetSaving = false
		},
	})

	items = append(items, CommandItem{
		Title:       "✨ Shader: Cycle Procedural Canvas Shaders",
		Category:    "Aesthetics",
		Shortcut:    "[ H ]",
		Description: "Cycle canvas animation shaders (Rain-on-Glass, Fireflies, Aurora, Audio Waveform)",
		Action: func(m *Model) {
			shader := m.Scene.CycleShaderEffect()
			m.SetStatusMsg("Canvas Shader: " + shader.String() + " ✨")
		},
	})

	// Views & Navigation
	items = append(items,
		CommandItem{
			Title:       "🏡 View: Cottage Art Main Canvas",
			Category:    "Navigation",
			Shortcut:    "[ 1 ]",
			Description: "Switch to full-screen 8-bit pixel art canvas view",
			Action:      func(m *Model) { m.ActiveTab = TabArt },
		},
		CommandItem{
			Title:       "⏱️ View: Focus Timer & Pomodoro",
			Category:    "Navigation",
			Shortcut:    "[ 2 ]",
			Description: "Switch to Focus Timer & Pomodoro control view",
			Action:      func(m *Model) { m.ActiveTab = TabTimer },
		},
		CommandItem{
			Title:       "📝 View: Cottage To-Do Notebook",
			Category:    "Navigation",
			Shortcut:    "[ 3 ]",
			Description: "Switch to task notebook & priority checklist view",
			Action:      func(m *Model) { m.ActiveTab = TabTasks },
		},
		CommandItem{
			Title:       "🌿 View: Ambiance & Audio Soundboard",
			Category:    "Navigation",
			Shortcut:    "[ 4 ]",
			Description: "Switch to ambient sound mixer and audio visualizer view",
			Action:      func(m *Model) { m.ActiveTab = TabAmbiance },
		},
		CommandItem{
			Title:       "🧘 View: Zen Minimal Focus Mode",
			Category:    "Navigation",
			Shortcut:    "[ Shift+F ]",
			Description: "Toggle zero-clutter focus view with floating timer overlay",
			Action:      func(m *Model) { m.ZenMode = !m.ZenMode },
		},

		// Dialog Modals
		CommandItem{
			Title:       "🖼️ Catalog: Scene Browser Modal",
			Category:    "Dialogs",
			Shortcut:    "[ Shift+V ]",
			Description: "Open categorized scene catalog with 370+ pixel-art presets",
			Action: func(m *Model) {
				m.ShowScenePicker = true
				m.PickerSceneIdx = 0
			},
		},
		CommandItem{
			Title:       "🎨 Theme: Palette Customizer Modal",
			Category:    "Dialogs",
			Shortcut:    "[ Shift+T ]",
			Description: "Open interactive theme selector with live color swatches",
			Action: func(m *Model) {
				m.ShowThemePicker = true
				m.PickerThemeIdx = 0
			},
		},
		CommandItem{
			Title:       "ℹ️ Info: Scene Inspector & Palette Analysis",
			Category:    "Dialogs",
			Shortcut:    "[ Shift+I ]",
			Description: "Inspect resolution, animation metrics, and 8 dominant scene colors",
			Action:      func(m *Model) { m.ShowSceneInfoModal = true },
		},

		// Aesthetics & Controls
		CommandItem{
			Title:       "🎲 Scene: Surprise Random Jump",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+R ]",
			Description: "Instantly jump to a random artwork preset across all themes",
			Action: func(m *Model) {
				ns := m.Scene.RandomScene()
				m.SetStatusMsg("🎲 Random Scene: " + ns.String() + " ✨")
			},
		},
		CommandItem{
			Title:       "👈 Scene: Previous Scene History",
			Category:    "Aesthetics",
			Shortcut:    "[ b ]",
			Description: "Return to the previously viewed background artwork",
			Action: func(m *Model) {
				ns := m.Scene.PrevScene()
				m.SetStatusMsg("Switched Scene: " + ns.String())
			},
		},
		CommandItem{
			Title:       "🎨 Filter: Cycle Color Grading Filter",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+C ]",
			Description: "Cycle RGB post-processing filters (Amber, Synth, Sepia, Phosphor)",
			Action: func(m *Model) {
				cf := m.Scene.CycleColorFilter()
				m.SetStatusMsg("Color Filter: " + cf.String() + " 🎨")
			},
		},
		CommandItem{
			Title:       "✨ Particles: Cycle Screen Particle Overlay",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+P ]",
			Description: "Cycle ambient physics overlays (Sakura, Embers, Rain, Snow, Matrix)",
			Action: func(m *Model) {
				po := m.Scene.CycleParticleOverlay()
				m.SetStatusMsg("Particle Overlay: " + po.String() + " ✨")
			},
		},
		CommandItem{
			Title:       "🎶 Visualizer: Cycle Audio Spectrum Style",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+M ]",
			Description: "Cycle spectrum styles (Bars, Dots, Braille, Equalizer, Waveform)",
			Action: func(m *Model) {
				m.SpectrumMode = (m.SpectrumMode + 1) % 5
				m.SetStatusMsg("Visualizer Mode: " + m.SpectrumMode.String() + " 🎶")
			},
		},
		CommandItem{
			Title:       "🎵 Audio Pulse: Toggle Beat-Reactive Particles",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+B ]",
			Description: "Toggle particle movement pulsing with Lofi Radio audio beat",
			Action: func(m *Model) {
				b := m.Scene.ToggleBeatReactive()
				if b {
					m.SetStatusMsg("Audio-Reactive Particle Pulsing: ON 🎵")
				} else {
					m.SetStatusMsg("Audio-Reactive Particle Pulsing: OFF")
				}
			},
		},
		CommandItem{
			Title:       "🌅 Sky Blend: Toggle Day/Night Smooth Blending",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+D ]",
			Description: "Toggle real-time smooth day/night sky color gradient blending",
			Action: func(m *Model) {
				d := m.Scene.ToggleSmoothDayNight()
				if d {
					m.SetStatusMsg("Smooth Day/Night Sky Blending: ON 🌅")
				} else {
					m.SetStatusMsg("Smooth Day/Night Sky Blending: OFF")
				}
			},
		},

		// Sound & Volume Controls
		CommandItem{
			Title:       "🌧️ Audio: Toggle Procedural Rain Sound",
			Category:    "Audio",
			Shortcut:    "[ r ]",
			Description: "Toggle soft low-pass ambient rain audio loop",
			Action: func(m *Model) {
				r := m.Ambient.ToggleRain()
				if r {
					m.SetStatusMsg("Procedural Rain Audio: ON 🌧️")
				} else {
					m.SetStatusMsg("Procedural Rain Audio: OFF")
				}
			},
		},
		CommandItem{
			Title:       "🔥 Audio: Toggle Fireplace Crackle Sound",
			Category:    "Audio",
			Shortcut:    "[ f ]",
			Description: "Toggle warm fireplace hearth rumble audio loop",
			Action: func(m *Model) {
				f := m.Ambient.ToggleFire()
				if f {
					m.SetStatusMsg("Fireplace Audio: ON 🔥")
				} else {
					m.SetStatusMsg("Fireplace Audio: OFF")
				}
			},
		},
		CommandItem{
			Title:       "🍃 Audio: Toggle Wind Breeze Sound",
			Category:    "Audio",
			Shortcut:    "[ W ]",
			Description: "Toggle ambient wind breeze audio loop",
			Action: func(m *Model) {
				w := m.Ambient.ToggleWind()
				if w {
					m.SetStatusMsg("Wind Breeze Audio: ON 🍃")
				} else {
					m.SetStatusMsg("Wind Breeze Audio: OFF")
				}
			},
		},
		CommandItem{
			Title:       "☕ Audio: Toggle Cafe Chatter Ambience",
			Category:    "Audio",
			Shortcut:    "[ K ]",
			Description: "Toggle ambient coffee shop chatter audio loop",
			Action: func(m *Model) {
				c := m.Ambient.ToggleCafe()
				if c {
					m.SetStatusMsg("Cafe Ambience: ON ☕")
				} else {
					m.SetStatusMsg("Cafe Ambience: OFF")
				}
			},
		},
		CommandItem{
			Title:       "🦗 Audio: Toggle Night Crickets Ambience",
			Category:    "Audio",
			Shortcut:    "[ e ]",
			Description: "Toggle ambient night crickets chirping audio loop",
			Action: func(m *Model) {
				cr := m.Ambient.ToggleCrickets()
				if cr {
					m.SetStatusMsg("Night Crickets Audio: ON 🦗")
				} else {
					m.SetStatusMsg("Night Crickets Audio: OFF")
				}
			},
		},
		CommandItem{
			Title:       "🔊 Audio: Volume Up (+10%)",
			Category:    "Audio",
			Shortcut:    "[ + ]",
			Description: "Increase master audio output volume by 10%",
			Action: func(m *Model) {
				v := m.CLIAMP.AdjustVolume(10)
				m.SetStatusMsg(fmt.Sprintf("Audio Volume: 🔊 %d%%", v))
			},
		},
		CommandItem{
			Title:       "🔉 Audio: Volume Down (-10%)",
			Category:    "Audio",
			Shortcut:    "[ - ]",
			Description: "Decrease master audio output volume by 10%",
			Action: func(m *Model) {
				v := m.CLIAMP.AdjustVolume(-10)
				m.SetStatusMsg(fmt.Sprintf("Audio Volume: 🔉 %d%%", v))
			},
		},
	)

	return items
}

// FilterCommands returns commands matching query string.
func FilterCommands(commands []CommandItem, query string) []CommandItem {
	query = strings.TrimSpace(strings.ToLower(query))
	if query == "" {
		return commands
	}

	var results []CommandItem
	for _, cmd := range commands {
		if strings.Contains(strings.ToLower(cmd.Title), query) ||
			strings.Contains(strings.ToLower(cmd.Category), query) ||
			strings.Contains(strings.ToLower(cmd.Description), query) ||
			strings.Contains(strings.ToLower(cmd.Shortcut), query) {
			results = append(results, cmd)
		}
	}
	return results
}
