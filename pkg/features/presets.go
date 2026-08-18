package features

import (
	"encoding/json"
	"os"

	"gottage/pkg/art"
)

// AtmospherePreset captures a complete user atmosphere configuration.
type AtmospherePreset struct {
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Scene           art.SceneType       `json:"scene"`
	Theme           string              `json:"theme"`
	ColorFilter     art.ColorFilter     `json:"color_filter"`
	ParticleOverlay art.ParticleOverlay `json:"particle_overlay"`
	ShaderEffect    art.ShaderEffect    `json:"shader_effect"`
	RainAudio       bool                `json:"rain_audio"`
	FireAudio       bool                `json:"fire_audio"`
	WindAudio       bool                `json:"wind_audio"`
	AnimSpeed       art.AnimSpeed       `json:"anim_speed"`
}

// PresetManager handles saving, loading, and listing atmosphere presets.
type PresetManager struct {
	Presets  []AtmospherePreset
	filePath string
}

// NewPresetManager initializes PresetManager with default built-in presets and attempts disk loading.
func NewPresetManager(filePath string) *PresetManager {
	pm := &PresetManager{
		filePath: filePath,
	}
	pm.loadBuiltinPresets()
	pm.LoadFromFile()
	return pm
}

func (pm *PresetManager) loadBuiltinPresets() {
	pm.Presets = []AtmospherePreset{
		{
			Name:            "☕ Rainy Coffee Nook",
			Description:     "Cozy rain artwork + rain audio + sepia filter + raindrop glass shader",
			Scene:           art.SceneRainyCity,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterSepia,
			ParticleOverlay: art.OverlaySnow,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "⚡ Cyberpunk Hackathon Deck",
			Description:     "Cyber matrix artwork + cool synth filter + matrix code rain + waveform shader",
			Scene:           art.SceneCyberMatrix,
			Theme:           "Cyberpunk Neon",
			ColorFilter:     art.FilterCoolSynth,
			ParticleOverlay: art.OverlayMatrix,
			ShaderEffect:    art.ShaderAudioWaveform,
			RainAudio:       false,
			AnimSpeed:       art.SpeedFast,
		},
		{
			Name:            "🌲 Deep Forest Meditation",
			Description:     "Enchanted forest artwork + wind audio + emerald theme + fireflies shader",
			Scene:           art.SceneEnchantedForest,
			Theme:           "Emerald Forest",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlaySakura,
			ShaderEffect:    art.ShaderFireflies,
			WindAudio:       true,
			AnimSpeed:       art.SpeedSlow,
		},
		{
			Name:            "🌌 Stargazing Observatory",
			Description:     "Cosmic observatory artwork + fireplace crackle + midnight filter + aurora waves",
			Scene:           art.SceneCosmicObservatory,
			Theme:           "Tokyo Night",
			ColorFilter:     art.FilterMidnight,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderAuroraWaves,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
	}
}

// SavePreset adds or updates a preset and persists to disk.
func (pm *PresetManager) SavePreset(preset AtmospherePreset) {
	found := false
	for i, p := range pm.Presets {
		if p.Name == preset.Name {
			pm.Presets[i] = preset
			found = true
			break
		}
	}
	if !found {
		pm.Presets = append(pm.Presets, preset)
	}
	_ = pm.SaveToFile()
}

// DeletePreset removes a preset by index.
func (pm *PresetManager) DeletePreset(index int) bool {
	if index < 0 || index >= len(pm.Presets) {
		return false
	}
	pm.Presets = append(pm.Presets[:index], pm.Presets[index+1:]...)
	_ = pm.SaveToFile()
	return true
}

func (pm *PresetManager) SaveToFile() error {
	if pm.filePath == "" {
		return nil
	}
	data, err := json.MarshalIndent(pm.Presets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(pm.filePath, data, 0644)
}

func (pm *PresetManager) LoadFromFile() error {
	if pm.filePath == "" {
		return nil
	}
	data, err := os.ReadFile(pm.filePath)
	if err != nil {
		return err
	}
	var loaded []AtmospherePreset
	if err := json.Unmarshal(data, &loaded); err != nil {
		return err
	}
	if len(loaded) > 0 {
		pm.Presets = loaded
	}
	return nil
}
