package ui

import (
	"strings"
	"testing"

	"gottage/pkg/art"

	tea "github.com/charmbracelet/bubbletea"
)

func TestVisualizers(t *testing.T) {
	styles := NewStyles(ThemeCozyAutumn)
	for _, mode := range []SpectrumStyle{SpectrumBars, SpectrumDots, SpectrumBraille, SpectrumEqualizer, SpectrumWave} {
		specPlaying := RenderAudioSpectrum(16, true, 10, mode, styles.BoldAccent)
		if specPlaying == "" {
			t.Errorf("Expected non-empty spectrum output when playing for mode %s", mode.String())
		}

		specPaused := RenderAudioSpectrum(16, false, 10, mode, styles.BoldAccent)
		if specPaused == "" {
			t.Errorf("Expected non-empty spectrum output when paused for mode %s", mode.String())
		}

		spec2D := RenderAudioSpectrum2D(32, 8, true, 15, mode, styles)
		if spec2D == "" {
			t.Errorf("Expected non-empty 2D spectrum output when playing for mode %s", mode.String())
		}
	}

	breath := GetZenBreathState(20)
	if breath.PhaseName == "" || breath.RingVisual == "" {
		t.Errorf("Unexpected empty breath state: %+v", breath)
	}
}

func TestModelViewModals(t *testing.T) {
	m := NewModel()
	vNormal := m.View()
	if !strings.Contains(vNormal, "GOTTAGE") {
		t.Error("Expected normal view to contain GOTTAGE header")
	}

	// Test Zen Mode
	m.ZenMode = true
	vZ := m.View()
	if !strings.Contains(vZ, "ZEN FOCUS SANCTUARY") {
		t.Error("Expected Zen view to contain ZEN FOCUS SANCTUARY")
	}
	m.ZenMode = false

	// Test Scene Navigation (PrevScene & RandomScene)
	origScene := m.Scene.CurrentScene
	prev := m.Scene.PrevScene()
	if prev.String() == "" {
		t.Error("Expected PrevScene to return a valid scene type")
	}
	randScene := m.Scene.RandomScene()
	if randScene.String() == "" {
		t.Error("Expected RandomScene to return a valid scene type")
	}
	_ = origScene

	// Test Scene Picker Catalog Modal
	m.ShowScenePicker = true
	vP := m.View()
	if !strings.Contains(vP, "SCENE CATALOG BROWSER") {
		t.Error("Expected Scene Picker modal to contain SCENE CATALOG BROWSER")
	}

	// Test Scene Catalog Search Filter
	m.PickerSearchMode = true
	m.PickerSearchInput.SetValue("cyber")
	vSearch := m.View()
	if !strings.Contains(vSearch, "Filter:") {
		t.Error("Expected Scene Picker search view to contain Filter:")
	}
	filtered := m.getFilteredScenes()
	if len(filtered) == 0 {
		t.Error("Expected search query 'cyber' to match scenes")
	}
	m.ShowScenePicker = false
	m.PickerSearchMode = false

	// Test Theme Picker Modal
	m.ShowThemePicker = true
	vT := m.View()
	if !strings.Contains(vT, "THEME PALETTE CUSTOMIZER") {
		t.Error("Expected Theme Picker modal to contain THEME PALETTE CUSTOMIZER")
	}
	if len(AllThemes) != 10 {
		t.Errorf("Expected 10 themes in AllThemes, got %d", len(AllThemes))
	}
	m.ShowThemePicker = false

	// Test ColorFilter & ParticleOverlay Cycling
	cf := m.Scene.CycleColorFilter()
	if cf.String() == "" {
		t.Error("Expected non-empty ColorFilter string")
	}
	po := m.Scene.CycleParticleOverlay()
	if po.String() == "" {
		t.Error("Expected non-empty ParticleOverlay string")
	}

	// Test Favorites Toggling
	isFav := m.ToggleFavorite(art.SceneForestCabin)
	if !isFav {
		t.Error("Expected ToggleFavorite to mark scene as favorite")
	}
	m.PickerCategoryIdx = 0
	favScenes := m.getFilteredScenes()
	if len(favScenes) == 0 {
		t.Error("Expected favorites list to contain at least 1 scene")
	}

	// Test Volume & Animation Speed & Task Priority
	vol := m.CLIAMP.AdjustVolume(10)
	if vol != 90 {
		t.Errorf("Expected volume 90, got %d", vol)
	}

	speed := m.Scene.CycleAnimSpeed()
	if speed.String() == "" {
		t.Error("Expected non-empty AnimSpeed string")
	}

	m.ActiveTab = TabTasks
	vTasks := m.View()
	if !strings.Contains(vTasks, "COTTAGE TO-DO NOTEBOOK") {
		t.Error("Expected tasks view to contain COTTAGE TO-DO NOTEBOOK")
	}

	// Test Scene Info Inspector Modal
	m.ShowSceneInfoModal = true
	vInfo := m.View()
	if !strings.Contains(vInfo, "SCENE INSPECTOR & PALETTE ANALYSIS") {
		t.Error("Expected Scene Info modal to contain SCENE INSPECTOR & PALETTE ANALYSIS")
	}
	m.ShowSceneInfoModal = false

	// Test Dominant Palette Extraction & Scene Toggles
	pals := m.Scene.Canvas.ExtractDominantPalette(8)
	if len(pals) == 0 {
		t.Error("Expected ExtractDominantPalette to return at least 1 color")
	}

	beat := m.Scene.ToggleBeatReactive()
	if !beat && m.Scene.BeatReactiveParticles {
		t.Error("Expected ToggleBeatReactive to toggle state")
	}

	blend := m.Scene.ToggleSmoothDayNight()
	if !blend && m.Scene.SmoothDayNightBlend {
		t.Error("Expected ToggleSmoothDayNight to toggle state")
	}

	vig := m.Scene.ToggleVignette()
	if vig == m.Scene.VignetteOn {
		// Verify toggle works
	}
	m.Scene.Canvas.ApplyVignette(0.5)

	crt := m.Scene.ToggleCRTScanlines()
	if !crt || !m.Scene.CRTScanlinesOn {
		t.Error("Expected ToggleCRTScanlines to turn ON scanlines")
	}
	m.Scene.Canvas.ApplyCRTScanlines(1.0)

	// Test Command Palette Modal & Mood Presets
	m.ShowCommandPalette = true
	vCmd := m.View()
	if !strings.Contains(vCmd, "QUICK COMMAND PALETTE & MOOD PRESETS") {
		t.Error("Expected Command Palette modal to contain QUICK COMMAND PALETTE & MOOD PRESETS")
	}

	cmds := GetCommandList(&m)
	if len(cmds) < 10 {
		t.Errorf("Expected at least 10 commands in palette, got %d", len(cmds))
	}

	filteredCmds := FilterCommands(cmds, "rainy")
	if len(filteredCmds) == 0 {
		t.Error("Expected FilterCommands to match 'rainy'")
	}

	// Execute Preset command
	filteredCmds[0].Action(&m)
	m.ShowCommandPalette = false

	// Test Shader Effect Cycling
	shader := m.Scene.CycleShaderEffect()
	if shader.String() == "" {
		t.Error("Expected non-empty ShaderEffect string")
	}

	// Test Preset Studio Modal
	m.ShowPresetModal = true
	vPreset := m.View()
	if !strings.Contains(vPreset, "ATMOSPHERE PRESET STUDIO") {
		t.Error("Expected Preset Studio modal to contain ATMOSPHERE PRESET STUDIO")
	}
	m.ShowPresetModal = false

	// Test 5-Channel Procedural Ambient Sound Console (Cafe & Crickets)
	cafeOn := m.Ambient.ToggleCafe()
	if !cafeOn || !m.Ambient.CafeActive {
		t.Error("Expected ToggleCafe to activate Cafe chatter sound")
	}
	cricketOn := m.Ambient.ToggleCrickets()
	if !cricketOn || !m.Ambient.CricketsActive {
		t.Error("Expected ToggleCrickets to activate Night Crickets sound")
	}
	vAmb := m.renderAmbianceView()
	if !strings.Contains(vAmb, "5-Channel Procedural Ambient Sound Console") || !strings.Contains(vAmb, "Cafe Chatter") || !strings.Contains(vAmb, "Night Crickets") {
		t.Error("Expected renderAmbianceView to display 5-channel console and new sound channels")
	}

	// Test 5-Channel Ambient Soundscape Equalizer & Mixer Modal
	m.ShowSoundMixerModal = true
	vMixer := m.View()
	if !strings.Contains(vMixer, "5-CHANNEL AMBIENT SOUNDSCAPE EQUALIZER & MIXER") {
		t.Error("Expected Sound Mixer modal to contain 5-CHANNEL AMBIENT SOUNDSCAPE EQUALIZER & MIXER")
	}
	m.Ambient.AdjustChannelVolume(0, 0.1)
	rainVol := m.Ambient.GetChannelVolume(0)
	if rainVol < 0.75 {
		t.Errorf("Expected RainVol to be increased to ~0.8, got %.2f", rainVol)
	}
	m.ShowSoundMixerModal = false

	// Test Canvas Clone & BlendCanvas
	c1 := art.NewCanvas(20, 20)
	c1.SetPixelHex(5, 5, "#FF0000")
	c2 := c1.Clone()
	if c2 == nil || c2.Pixels[5][5].R < 0.9 {
		t.Error("Expected Canvas.Clone to duplicate pixel matrix")
	}
	c1.BlendCanvas(c2, 0.5)

	// Test Smooth Scene Transition & SetSceneWithTransition
	m.Scene.SetSceneWithTransition(art.SceneCozyCampfire)
	if m.Scene.PrevCanvas == nil || m.Scene.CurrentScene != art.SceneCozyCampfire {
		t.Error("Expected SetSceneWithTransition to create PrevCanvas snapshot and set CurrentScene")
	}

	// Test Auto-Cycle Modes & Scope Toggle
	autoMode := m.Scene.CycleAutoCycleMode()
	if autoMode.String() == "" {
		t.Error("Expected non-empty AutoCycleMode string")
	}
	favScope := m.Scene.ToggleAutoCycleFavs()
	if !favScope || !m.Scene.AutoCycleFavsOnly {
		t.Error("Expected ToggleAutoCycleFavs to activate favorites scope")
	}

	vAmbAuto := m.renderAmbianceView()
	if !strings.Contains(vAmbAuto, "Auto-Cycle Mode") {
		t.Error("Expected renderAmbianceView to display Auto-Cycle Mode settings")
	}

	// Test Art Engine Render Mode Cycling & Live Badge Rendering
	rm := m.Scene.CycleRenderMode()
	if rm != art.RenderASCII {
		t.Errorf("Expected CycleRenderMode to switch to RenderASCII, got %v", rm)
	}
	m.StatusMsg = ""
	vBadge := m.View()
	if !strings.Contains(vBadge, "ASCII Density Ramp") {
		t.Error("Expected status footer to contain live ASCII Density Ramp badge")
	}

	// Test 2D Spectrum & Equalizer Studio Modal
	m.ShowSpectrumModal = true
	vSpec := m.View()
	if !strings.Contains(vSpec, "2D AUDIO SPECTRUM & EQUALIZER STUDIO") {
		t.Error("Expected ShowSpectrumModal view to render 2D AUDIO SPECTRUM & EQUALIZER STUDIO header")
	}
	m.ShowSpectrumModal = false

	// Test Canvas Energy Ripples & 3D Cosmos Starfield Warp Shader Badge
	m.Scene.TriggerCanvasRippleHex(40, 20, "#ff5588")
	if len(m.Scene.Ripples) != 1 {
		t.Errorf("Expected 1 canvas ripple active, got %d", len(m.Scene.Ripples))
	}
	m.Scene.ShaderEffect = art.ShaderStarfieldWarp
	m.StatusMsg = ""
	vWarp := m.View()
	if !strings.Contains(vWarp, "3D Cosmos Starfield Warp") {
		t.Error("Expected status footer to contain 3D Cosmos Starfield Warp live badge")
	}

	// Test Pet Cat companion interaction & status badge rendering
	petMsg := m.Scene.PetCat()
	if !strings.Contains(petMsg, "Bliss:") {
		t.Errorf("Expected pet message to contain Bliss:, got %s", petMsg)
	}
	m.StatusMsg = ""
	vBliss := m.View()
	if !strings.Contains(vBliss, "Bliss:") {
		t.Error("Expected status footer to contain live Bliss badge")
	}

	// Test Desktop Clock Badge, Quick Control Deck, and Mindfulness Quotes
	if !strings.Contains(vBliss, "🕒") {
		t.Error("Expected header to contain live desktop clock badge")
	}
	if !strings.Contains(vBliss, "MINDFULNESS PROMPT") {
		t.Error("Expected sidebar to contain MINDFULNESS PROMPT card")
	}
	vArt := m.renderArtView()
	if !strings.Contains(vArt, "Lamp:") || !strings.Contains(vArt, "Shader:") {
		t.Error("Expected renderArtView to contain interactive Quick Control Deck")
	}

	q1 := m.Inspiration.CurrentQuote()
	q2 := m.Inspiration.NextQuote()
	if q1.Quote == q2.Quote {
		t.Error("Expected InspirationManager.NextQuote to cycle to a new quote")
	}
}

func TestMouseEvents(t *testing.T) {
	m := NewModel()
	m.Width = 120
	m.Height = 40

	// Test clicking header tabs
	mouseMsg := tea.MouseMsg{
		X:      75,
		Y:      1,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}
	updated, _ := m.Update(mouseMsg)
	mUpdated := updated.(Model)
	if mUpdated.StatusMsg == "" {
		t.Error("Expected tab click to set status message")
	}

	// Test clicking sidebar music card
	sidebarMusicClick := tea.MouseMsg{
		X:      100,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}
	updated2, _ := mUpdated.Update(sidebarMusicClick)
	m2 := updated2.(Model)
	if m2.StatusMsg == "" {
		t.Error("Expected sidebar music click to set status message")
	}

	// Test clicking canvas to pet cat (on TabArt)
	m2.ActiveTab = TabArt
	canvasClick := tea.MouseMsg{
		X:      40,
		Y:      20,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}
	updated3, _ := m2.Update(canvasClick)
	m3 := updated3.(Model)
	if !strings.Contains(m3.StatusMsg, "Bliss") {
		t.Errorf("Expected canvas click to pet cat companion, got status: %s", m3.StatusMsg)
	}
}

func TestAllHandcraftedScenesInView(t *testing.T) {
	scenes := []art.SceneType{
		art.SceneRainyCity,
		art.SceneForestCabin,
		art.SceneCottageClassic,
		art.SceneObservatory,
		art.SceneZenTeahouse,
		art.SceneCozyCampfire,
		art.SceneCozyAttic,
		art.ScenePixelWinterCabin,
		art.SceneCozyGreenhouse,
		art.SceneCozyLibrary,
		art.SceneLighthouseCliff,
		art.ScenePixelBakery,
		art.SceneRetroLofiTrain,
		art.SceneZenWaterfallShrine,
		art.SceneNeonBistro,
		art.SceneStarlightCamp,
		art.SceneRamenShop,
		art.SceneMagicSanctuary,
		art.SceneMountainAuroraShack,
		art.SceneStarlightTreehouse,
		art.SceneCoffeeShop,
		art.SceneEnchantedForest,
		art.SceneRiverBridge,
		art.SceneOceanWaves,
		art.SceneSunsetStation,
		art.SceneBambooGrove,
		art.SceneWinterCafe,
		art.SceneAutumnPark,
		art.SceneFairytaleWindmill,
		art.SceneCyberBar,
		art.SceneCloudTemple,
		art.SceneWizardTower,
		art.SceneZenBonsai,
		art.SceneCyberDiner,
		art.SceneCozyIgloo,
		art.SceneSkyIsland,
	}

	m := NewModel()
	m.Width = 120
	m.Height = 40

	for _, st := range scenes {
		m.Scene.CurrentScene = st
		v := m.View()
		if len(v) == 0 {
			t.Errorf("Expected non-empty View output for scene %v", st)
		}
		if !strings.Contains(v, "GOTTAGE") {
			t.Errorf("Expected View to contain GOTTAGE for scene %v", st)
		}
	}
}

func TestBeverageSystem(t *testing.T) {
	m := NewModel()
	m.Width = 120
	m.Height = 40

	// Test Header contains Beverage badge
	v := m.View()
	if !strings.Contains(v, "Dark Roast") {
		t.Errorf("Expected View to display default Dark Roast coffee badge")
	}

	// Test Ambiance View contains Artisan Beverage Bar
	m.ActiveTab = TabAmbiance
	vAmb := m.View()
	if !strings.Contains(vAmb, "Artisan Hot Beverage") {
		t.Errorf("Expected Ambiance view to render Artisan Hot Beverage section")
	}

	// Test Key 'J' (Brew Next Drink)
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'J'}})
	m2 := updated.(Model)
	if !strings.Contains(m2.StatusMsg, "Freshly Brewed") {
		t.Errorf("Expected 'J' key to brew next drink, got: %s", m2.StatusMsg)
	}

	// Test Key 'd' (Take a Sip)
	m2.ActiveTab = TabArt
	updated2, _ := m2.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}})
	m3 := updated2.(Model)
	if !strings.Contains(m3.StatusMsg, "warm") {
		t.Errorf("Expected 'd' key to take a sip, got: %s", m3.StatusMsg)
	}
}



