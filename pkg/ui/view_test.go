package ui

import (
	"strings"
	"testing"
)

func TestVisualizers(t *testing.T) {
	styles := NewStyles(ThemeCozyAutumn)
	specPlaying := RenderAudioSpectrum(16, true, 10, styles.BoldAccent)
	if specPlaying == "" {
		t.Error("Expected non-empty spectrum output when playing")
	}

	specPaused := RenderAudioSpectrum(16, false, 10, styles.BoldAccent)
	if specPaused == "" {
		t.Error("Expected non-empty spectrum output when paused")
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

	// Test Scene Picker Catalog Modal
	m.ShowScenePicker = true
	vP := m.View()
	if !strings.Contains(vP, "SCENE CATALOG BROWSER") {
		t.Error("Expected Scene Picker modal to contain SCENE CATALOG BROWSER")
	}
}
