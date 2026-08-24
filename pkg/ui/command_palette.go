package ui

import (
	"fmt"
	"strings"

	"gottage/pkg/art"
	"gottage/pkg/features"
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
		Title:       "🎛️ Mixer: 5-Channel Soundscape Equalizer & Mixer",
		Category:    "Dialogs",
		Shortcut:    "[ Shift+X ]",
		Description: "Open 5-channel ambient audio mixer modal to adjust individual sound channel volumes",
		Action: func(m *Model) {
			m.ShowSoundMixerModal = true
			m.MixerChannelIdx = 0
		},
	})

	items = append(items, CommandItem{
		Title:       "✨ Shader: Cycle Procedural Canvas Shaders",
		Category:    "Aesthetics",
		Shortcut:    "[ H ]",
		Description: "Cycle canvas animation shaders (Rain-on-Glass, Fireflies, Aurora, Audio Waveform, Starfield Warp, Fluid Waves, Flames, Matrix Depth)",
		Action: func(m *Model) {
			shader := m.Scene.CycleShaderEffect()
			m.SetStatusMsg("Canvas Shader: " + shader.String() + " ✨")
		},
	})

	items = append(items, CommandItem{
		Title:       "🐾 Companion: Pet Cozy Cat (Purr & Heart Sparks)",
		Category:    "Interactive Companion",
		Shortcut:    "[ K ]",
		Description: "Pet your cozy companion cat to unleash heart particles, purrs, and boost focus bliss",
		Action: func(m *Model) {
			msg := m.Scene.PetCat()
			m.SetStatusMsg(msg)
		},
	})

	items = append(items, CommandItem{
		Title:       "✨ Quote: Cycle Focus Mindfulness Quote",
		Category:    "Mindfulness",
		Shortcut:    "[ Q ]",
		Description: "Cycle uplifting focus and mindfulness inspiration quotes",
		Action: func(m *Model) {
			q := m.Inspiration.NextQuote()
			m.SetStatusMsg(fmt.Sprintf("✨ Quote: \"%s\" — %s", q.Quote, q.Author))
		},
	})

	items = append(items, CommandItem{
		Title:       "☕ Beverage: Take a Warm Sip",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Ctrl+B / d ]",
		Description: "Take a soothing warm sip from your hot mug for clarity and comfort",
		Action: func(m *Model) {
			if m.Beverage != nil {
				msg, warmth := m.Beverage.TakeSip()
				m.SetStatusMsg(fmt.Sprintf("%s (%d%% ♨️)", msg, warmth))
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🫖 Beverage: Cycle Next Hot Drink",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Shift+J ]",
		Description: "Freshly brew the next warm beverage from the artisan cottage menu",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.NextDrink()
				m.SetStatusMsg("Freshly Brewed: " + drink.Name + " ✨ (" + drink.Aroma + ")")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "☕ Brew: Pour-Over Dark Roast Coffee",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Coffee ]",
		Description: "Slow-dripped Ethiopian dark roast with velvety crema and roasted hazelnut notes",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkDripCoffee))
				m.SetStatusMsg("Freshly Brewed: " + drink.Name + " ☕")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🍵 Brew: Ceremonial Matcha Latte",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Matcha ]",
		Description: "Uji ceremonial matcha whisked with warm oat foam for mindful focus",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkMatchaLatte))
				m.SetStatusMsg("Freshly Whisked: " + drink.Name + " 🍵")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🫖 Brew: Artisan Earl Grey Tea",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Earl Grey ]",
		Description: "Single-origin Ceylon black tea steeped with cold-pressed bergamot peel",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkEarlGrey))
				m.SetStatusMsg("Steeped Fresh: " + drink.Name + " 🫖")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🍫 Brew: Marshmallow Hot Chocolate",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Cocoa ]",
		Description: "Melted Valrhona cocoa simmered with sweet vanilla and fluffy marshmallows",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkHotCocoa))
				m.SetStatusMsg("Simmered Fresh: " + drink.Name + " 🍫")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🍯 Brew: Honey Cardamom Spiced Chai",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Chai ]",
		Description: "Simmered Assam tea with cracked spices, fresh ginger, and clover honey",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkSpicedChai))
				m.SetStatusMsg("Simmered Fresh: " + drink.Name + " 🍯")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "🌼 Brew: Calming Chamomile Blossom",
		Category:    "Cozy Beverage",
		Shortcut:    "[ Brew Chamomile ]",
		Description: "Caffeine-free whole flower chamomile infusion for deep tranquil calm",
		Action: func(m *Model) {
			if m.Beverage != nil {
				drink := m.Beverage.BrewDrink(int(features.DrinkChamomile))
				m.SetStatusMsg("Steeped Fresh: " + drink.Name + " 🌼")
			}
		},
	})

	items = append(items, CommandItem{
		Title:       "💧 Ripple: Drop Canvas Energy Ripple",
		Category:    "Aesthetics",
		Shortcut:    "[ Ripple ]",
		Description: "Drop a serene expanding energy ripple ring onto the terminal canvas",
		Action: func(m *Model) {
			m.Scene.TriggerCanvasRippleHex(m.mainWidth()/2, m.mainHeight(), string(m.Styles.Theme.Accent))
			m.SetStatusMsg("Dropped Serene Canvas Energy Ripple 💧")
		},
	})

	items = append(items, CommandItem{
		Title:       "🎨 Art Engine: Cycle Terminal Render Modes",
		Category:    "Aesthetics",
		Shortcut:    "[ Shift+E ]",
		Description: "Cycle terminal rendering style (24-Bit Half-Blocks, ASCII Ramp, Braille Matrix, CRT Phosphor Glow)",
		Action: func(m *Model) {
			rm := m.Scene.CycleRenderMode()
			m.SetStatusMsg("Art Engine Render Mode: " + rm.String() + " 🎨")
		},
	})

	items = append(items, CommandItem{
		Title:       "🎬 Auto-Cycle: Cycle Auto-Rotation Mode",
		Category:    "Aesthetics",
		Shortcut:    "[ Y ]",
		Description: "Cycle background scene auto-rotation (Off -> 1m -> 3m -> 5m -> 10m -> Timer-Linked)",
		Action: func(m *Model) {
			mode := m.Scene.CycleAutoCycleMode()
			m.SetStatusMsg("🎬 Auto-Cycle Mode: " + mode.String())
		},
	})

	items = append(items, CommandItem{
		Title:       "⭐ Auto-Cycle: Toggle Favorites Scope",
		Category:    "Aesthetics",
		Shortcut:    "[ U ]",
		Description: "Toggle whether auto-cycling selects only bookmarked favorite scenes or all 370+ scenes",
		Action: func(m *Model) {
			favOnly := m.Scene.ToggleAutoCycleFavs()
			if favOnly {
				m.SetStatusMsg("🎬 Auto-Cycle Scope: Favorites Only ⭐")
			} else {
				m.SetStatusMsg("🎬 Auto-Cycle Scope: All 370+ Scenes 🎨")
			}
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
			Title:       "🎧 Scene: Lofi Study Desk & Rainy Window",
			Category:    "Aesthetics",
			Shortcut:    "[ Lofi Desk ]",
			Description: "Switch to handcrafted Lofi Study Desk with rain-streaked window and sleeping cat",
			Action: func(m *Model) {
				m.Scene.SetSceneWithTransition(art.SceneSunsetStation)
				m.SetStatusMsg("Switched Scene: Lofi Study Desk & Rainy Window 🎧🌧️")
			},
		},
		CommandItem{
			Title:       "🎋 Scene: Zen Rain Garden & Bamboo Water Clacker",
			Category:    "Aesthetics",
			Shortcut:    "[ Zen Garden ]",
			Description: "Switch to Japanese tatami room with animated shishi-odoshi bamboo water fountain",
			Action: func(m *Model) {
				m.Scene.SetSceneWithTransition(art.SceneBambooGrove)
				m.SetStatusMsg("Switched Scene: Zen Rain Garden & Bamboo Water Clacker 🎋🌧️")
			},
		},
		CommandItem{
			Title:       "🐈 Scene: Cozy Cat Cafe & Espresso Nook",
			Category:    "Aesthetics",
			Shortcut:    "[ Cat Cafe ]",
			Description: "Switch to warm wooden Cat Cafe with climbing tree, chrome espresso machine, and sleeping cats",
			Action: func(m *Model) {
				m.Scene.SetSceneWithTransition(art.SceneWinterCafe)
				m.SetStatusMsg("Switched Scene: Cozy Cat Cafe & Espresso Nook 🐈☕")
			},
		},
		CommandItem{
			Title:       "🍂 Scene: Autumn Lake Dock & Sunrise Mist",
			Category:    "Aesthetics",
			Shortcut:    "[ Autumn Dock ]",
			Description: "Switch to peaceful autumn mountain lake dock with morning mist and plaid blanket",
			Action: func(m *Model) {
				m.Scene.SetSceneWithTransition(art.SceneAutumnPark)
				m.SetStatusMsg("Switched Scene: Autumn Lake Dock & Sunrise Mist 🍂🌅")
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
			Title:       "🎶 Visualizer: Open 2D Spectrum Studio Modal",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+M ]",
			Description: "Open 2D Graphic Equalizer & Audio Visualizer Studio",
			Action: func(m *Model) {
				m.ShowSpectrumModal = true
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
			Title:       "🖼️ Vignette: Toggle Radial Edge Shading",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+G ]",
			Description: "Toggle soft radial vignette canvas edge shading for focus depth",
			Action: func(m *Model) {
				v := m.Scene.ToggleVignette()
				if v {
					m.SetStatusMsg("Cinematic Vignette Shading: ON 🖼️")
				} else {
					m.SetStatusMsg("Cinematic Vignette Shading: OFF")
				}
			},
		},
		CommandItem{
			Title:       "📺 CRT: Toggle CRT Monitor Scanlines Overlay",
			Category:    "Aesthetics",
			Shortcut:    "[ Shift+O ]",
			Description: "Toggle retro CRT monitor scanline dimming overlay",
			Action: func(m *Model) {
				c := m.Scene.ToggleCRTScanlines()
				if c {
					m.SetStatusMsg("Retro CRT Scanlines Overlay: ON 📺")
				} else {
					m.SetStatusMsg("Retro CRT Scanlines Overlay: OFF")
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
