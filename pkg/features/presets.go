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
			Name:            "🌿 Glasshouse Solarium Sanctuary",
			Description:     "Sunlit botanical glasshouse + rain audio + emerald theme + drifting mist",
			Scene:           art.SceneCozyGreenhouse,
			Theme:           "Emerald Forest",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlaySakura,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "📚 Midnight Stained Glass Library",
			Description:     "Grand arched library + ticking pendulum clock + warm fire + starlight overlay",
			Scene:           art.SceneCozyLibrary,
			Theme:           "Tokyo Night",
			ColorFilter:     art.FilterMidnight,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireflies,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌊 Coastal Ocean Lighthouse Study",
			Description:     "Ocean bay window + rotating beacon + ocean wind + potbelly stove crackle",
			Scene:           art.SceneLighthouseCliff,
			Theme:           "Nordic Frost",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderFluidWaves,
			WindAudio:       true,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🥐 Hearthside Autumn Bakery",
			Description:     "Wood-fired stone hearth + freshly baked bread + warm amber + falling leaves",
			Scene:           art.ScenePixelBakery,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayEmbers,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🚂 Vintage Lofi Train Carriage",
			Description:     "Panoramic rain window + passing mountain ridge + cassette player + warm reading light",
			Scene:           art.SceneRetroLofiTrain,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "♨️ Moonlit Mountain Onsen",
			Description:     "Hinoki cedar onsen + rising thermal steam + Mount Fuji view + night crickets & wind",
			Scene:           art.SceneZenWaterfallShrine,
			Theme:           "Tokyo Night",
			ColorFilter:     art.FilterMidnight,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireflies,
			WindAudio:       true,
			AnimSpeed:       art.SpeedSlow,
		},
		{
			Name:            "📻 Rainy Midnight Vinyl Lounge",
			Description:     "Spinning vinyl turntable + glowing tube amplifier + acoustic guitar + rain bokeh",
			Scene:           art.SceneNeonBistro,
			Theme:           "Gruvbox Dark",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🏕️ Sunken Hearth & Blanket Fort",
			Description:     "Sunken conversation pit + fairy light rafters + crackling hearth + tea kettle",
			Scene:           art.SceneStarlightCamp,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayEmbers,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🍜 Midnight Tokyo Ramen Bar",
			Description:     "Simmering tonkotsu broth + glowing red chochin lantern + Tokyo alley rain + rain audio",
			Scene:           art.SceneRamenShop,
			Theme:           "Tokyo Night",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌿 Botanical Herbal Apothecary",
			Description:     "Hanging dried herbs + glowing amethyst & emerald elixirs + bubbling brass cauldron + starlight",
			Scene:           art.SceneMagicSanctuary,
			Theme:           "Emerald Forest",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireflies,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌌 Nordic Fjord Pier & Aurora Sauna",
			Description:     "Undulating emerald & violet Aurora Borealis + mirror fjord lake + cedar hot tub + arctic wind",
			Scene:           art.SceneMountainAuroraShack,
			Theme:           "Nordic Frost",
			ColorFilter:     art.FilterCoolSynth,
			ParticleOverlay: art.OverlaySnow,
			ShaderEffect:    art.ShaderAuroraWaves,
			WindAudio:       true,
			FireAudio:       true,
			AnimSpeed:       art.SpeedSlow,
		},
		{
			Name:            "🌟 Starlit Urban Rooftop Sanctuary",
			Description:     "Fairy light pergola + telescope stargazing + rooftop record player + city skyline bokeh",
			Scene:           art.SceneStarlightTreehouse,
			Theme:           "Sunset Pastel",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireflies,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "📚 Antique Corner Bookshop",
			Description:     "Towering mahogany bookcases + rolling brass ladder + velvet armchair + rain streaks + Earl Grey tea",
			Scene:           art.SceneCoffeeShop,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌿 Misty Potting Shed & Greenhouse",
			Description:     "Glasshouse roof + copper watering can + hanging macrame ferns + terrarium moss + garden rain",
			Scene:           art.SceneEnchantedForest,
			Theme:           "Emerald Forest",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlaySakura,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🏕️ Midnight Riverside Camp",
			Description:     "Glowing canvas tent + crackling campfire kettle + undulating moonlit river + starlight & embers",
			Scene:           art.SceneRiverBridge,
			Theme:           "Nordic Frost",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌊 Seaside Cliffside Balcony",
			Description:     "Nautical teak deck + rotating lighthouse beam + rolling ocean waves + coastal lavender breeze",
			Scene:           art.SceneOceanWaves,
			Theme:           "Catppuccin Mocha",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderFluidWaves,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🎧 Lofi Study Desk & Rain",
			Description:     "Rain-streaked window + corkboard notes + warm desk lamp + over-ear headphones + lofi rain audio",
			Scene:           art.SceneSunsetStation,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🎋 Zen Rain Garden & Water Clacker",
			Description:     "Japanese tatami engawa + tipping bamboo water fountain + stone tsukubai basin + fresh matcha tea",
			Scene:           art.SceneBambooGrove,
			Theme:           "Catppuccin Mocha",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlaySakura,
			ShaderEffect:    art.ShaderFireflies,
			RainAudio:       true,
			AnimSpeed:       art.SpeedSlow,
		},
		{
			Name:            "🐈 Cozy Cat Cafe & Espresso Nook",
			Description:     "Arched window cafe + sisal cat tree + chrome espresso machine + fresh croissants + cafe chatter",
			Scene:           art.SceneWinterCafe,
			Theme:           "Gruvbox Dark",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🍂 Autumn Lake Dock & Morning Mist",
			Description:     "Alpine lake sunrise + swirling morning mist + crimson maples + plaid wool blanket + tin campsite lantern",
			Scene:           art.SceneAutumnPark,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayEmbers,
			ShaderEffect:    art.ShaderFluidWaves,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌿 Shire Hill Hobbit Burrow",
			Description:     "Round timber door + rolling green wildflower hills + crackling fireplace + whistling copper kettle & seed cake",
			Scene:           art.SceneFairytaleWindmill,
			Theme:           "Cozy Autumn",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayEmbers,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🎷 Rainy Japanese Jazz Kissa",
			Description:     "Dark walnut counter + glass siphon coffee dripper + glowing vacuum tube amp & blue VU meters + vinyl jazz",
			Scene:           art.SceneCyberBar,
			Theme:           "Gruvbox Dark",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayNone,
			ShaderEffect:    art.ShaderRainGlass,
			RainAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🌌 Starlit A-Frame Loft Bed",
			Description:     "A-frame loft skylight + cosmic nebula & shooting stars + fairy light rafters + fluffy down duvet comforter",
			Scene:           art.SceneCloudTemple,
			Theme:           "Tokyo Night",
			ColorFilter:     art.FilterMidnight,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderStarfieldWarp,
			WindAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🔮 Alchemist's Starlit Tower",
			Description:     "Gothic arched granite window + rotating brass astrolabe + glowing potion elixirs + star scroll & candle sconces",
			Scene:           art.SceneWizardTower,
			Theme:           "Catppuccin Mocha",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlayStars,
			ShaderEffect:    art.ShaderFireflies,
			FireAudio:       true,
			AnimSpeed:       art.SpeedNormal,
		},
		{
			Name:            "🍵 Zen Engawa Tea Sanctuary",
			Description:     "Tatami porch overlooking koi pond & Mount Fuji + gentle crickets + matcha tea",
			Scene:           art.SceneZenTeahouse,
			Theme:           "Catppuccin Mocha",
			ColorFilter:     art.FilterNormal,
			ParticleOverlay: art.OverlaySakura,
			ShaderEffect:    art.ShaderFireflies,
			AnimSpeed:       art.SpeedSlow,
		},
		{
			Name:            "🏕️ Starlit Alpine Campfire",
			Description:     "Campfire by pine lake + crackling hearth audio + starlight stars overlay",
			Scene:           art.SceneCozyCampfire,
			Theme:           "Gruvbox Dark",
			ColorFilter:     art.FilterWarmAmber,
			ParticleOverlay: art.OverlayEmbers,
			ShaderEffect:    art.ShaderFireFlames,
			FireAudio:       true,
			WindAudio:       true,
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
