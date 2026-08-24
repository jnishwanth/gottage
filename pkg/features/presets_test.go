package features

import (
	"os"
	"testing"

	"gottage/pkg/art"
)

func TestPresetManager(t *testing.T) {
	tmpFile := "test_presets.json"
	defer os.Remove(tmpFile)

	pm := NewPresetManager(tmpFile)
	if len(pm.Presets) < 10 {
		t.Fatalf("expected at least 10 builtin presets, got %d", len(pm.Presets))
	}

	custom := AtmospherePreset{
		Name:            "Test Custom Preset",
		Description:     "Unit test custom atmosphere",
		Scene:           art.SceneCyberDiner,
		Theme:           "Cyberpunk Neon",
		ColorFilter:     art.FilterCoolSynth,
		ParticleOverlay: art.OverlayMatrix,
		ShaderEffect:    art.ShaderRainGlass,
		RainAudio:       true,
	}

	pm.SavePreset(custom)

	if len(pm.Presets) < 5 {
		t.Fatalf("expected 5 presets after saving custom, got %d", len(pm.Presets))
	}

	pm2 := NewPresetManager(tmpFile)
	if len(pm2.Presets) != len(pm.Presets) {
		t.Fatalf("expected %d presets after reload, got %d", len(pm.Presets), len(pm2.Presets))
	}

	deleted := pm2.DeletePreset(len(pm2.Presets) - 1)
	if !deleted {
		t.Fatalf("expected successful preset deletion")
	}
}

// TestBuiltinPresetScenesAreRegistered guards against a preset pointing at a
// SceneType that nothing registers.
//
// art.GetScene falls back to SceneRainyCity for an unregistered type, so such a
// preset silently loads the wrong scene -- no error, no log. Phase 1 pruned the
// GIF scene set from 344 to 10, and two built-in presets (Cyberpunk Hackathon Deck
// and Stargazing Observatory) referenced scenes that went with it. This test is
// what makes the next prune fail loudly instead.
func TestBuiltinPresetScenesAreRegistered(t *testing.T) {
	pm := NewPresetManager("")
	for _, p := range pm.Presets {
		if _, ok := art.SceneRegistry[p.Scene]; !ok {
			t.Errorf("preset %q references SceneType %d, which no scene registers "+
				"(GetScene would silently fall back to SceneRainyCity)", p.Name, p.Scene)
		}
	}
}
