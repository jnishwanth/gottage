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
	if len(pm.Presets) < 4 {
		t.Fatalf("expected at least 4 builtin presets, got %d", len(pm.Presets))
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
