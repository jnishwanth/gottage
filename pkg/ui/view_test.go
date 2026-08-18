package ui

import (
	"strings"
	"testing"

	"gottage/pkg/art"
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
	if !strings.Contains(vZ, "ZEN FOCUS MODE") {
		t.Error("Expected Zen view to contain ZEN FOCUS MODE")
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

	// Test Command Palette Modal & Mood Presets
	m.ShowCommandPalette = true
	vCmd := m.View()
	if !strings.Contains(vCmd, "QUICK COMMAND PALETTE & MOOD PRESETS") {
		t.Error("Expected Command Palette modal to contain QUICK COMMAND PALETTE & MOOD PRESETS")
	}

	cmds := GetCommandList(&m)
	if len(cmds) < 15 {
		t.Errorf("Expected at least 15 commands in palette, got %d", len(cmds))
	}

	filteredCmds := FilterCommands(cmds, "rainy")
	if len(filteredCmds) == 0 {
		t.Error("Expected FilterCommands to match 'rainy'")
	}

	// Execute Mood Preset 1 (Rainy Cafe Nook)
	filteredCmds[0].Action(&m)
	if m.Scene.CurrentScene != art.SceneRainyCity {
		t.Errorf("Expected CurrentScene to be SceneRainyCity, got %s", m.Scene.CurrentScene.String())
	}
	m.ShowCommandPalette = false
}

