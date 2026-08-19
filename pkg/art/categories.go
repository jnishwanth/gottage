package art

import (
	"strings"
)

// SceneCategories lists display names for categorized scene browser.
var SceneCategories = []string{
	"⭐ Favorites",
	"🏡 Cozy & Cottage",
	"⚡ Cyberpunk & Neon",
	"🌿 Zen & Nature",
	"🌌 Sci-Fi & Cosmic",
	"⚙️ Steampunk & Fantasy",
	"🌅 Synthwave & Retro",
}

// CategorizeScene determines category index (1..6) for a given SceneType.
func CategorizeScene(st SceneType) int {
	name := strings.ToLower(st.String())
	switch {
	case strings.Contains(name, "cyber") || strings.Contains(name, "matrix") || strings.Contains(name, "neon") || strings.Contains(name, "kirokaze"):
		return 2
	case strings.Contains(name, "zen") || strings.Contains(name, "bamboo") || strings.Contains(name, "shrine") || strings.Contains(name, "torii") || strings.Contains(name, "waterfall") || strings.Contains(name, "cherry") || strings.Contains(name, "bonsai") || strings.Contains(name, "grove") || strings.Contains(name, "lotus") || strings.Contains(name, "garden") || strings.Contains(name, "onsen"):
		return 3
	case strings.Contains(name, "observatory") || strings.Contains(name, "space") || strings.Contains(name, "nebula") || strings.Contains(name, "astral") || strings.Contains(name, "zodiac") || strings.Contains(name, "alien") || strings.Contains(name, "stargate") || strings.Contains(name, "cosmos") || strings.Contains(name, "stargazing") || strings.Contains(name, "abyssal") || strings.Contains(name, "deep sea"):
		return 4
	case strings.Contains(name, "steampunk") || strings.Contains(name, "clockwork") || strings.Contains(name, "citadel") || strings.Contains(name, "airship") || strings.Contains(name, "castle") || strings.Contains(name, "magic") || strings.Contains(name, "wizard") || strings.Contains(name, "haunted") || strings.Contains(name, "volcano") || strings.Contains(name, "forge") || strings.Contains(name, "pirate"):
		return 5
	case strings.Contains(name, "synthwave") || strings.Contains(name, "highway") || strings.Contains(name, "arcade") || strings.Contains(name, "retro") || strings.Contains(name, "grid") || strings.Contains(name, "racer") || strings.Contains(name, "coastal"):
		return 6
	default:
		return 1 // Cozy & Cottage
	}
}

// GetScenesByCategory returns slice of SceneType in given category index.
func GetScenesByCategory(catIdx int) []SceneType {
	return GetScenesByCategoryAndFavorites(catIdx, nil)
}

// GetScenesByCategoryAndFavorites returns slice of SceneType for given category or favorites map.
func GetScenesByCategoryAndFavorites(catIdx int, favorites map[SceneType]bool) []SceneType {
	var list []SceneType
	if catIdx == 0 {
		for _, st := range SceneOrder {
			if favorites != nil && favorites[st] {
				list = append(list, st)
			}
		}
		if len(list) == 0 { // fallback to cozy if no favorites yet
			return GetScenesByCategoryAndFavorites(1, nil)
		}
		return list
	}
	for _, st := range SceneOrder {
		if CategorizeScene(st) == catIdx {
			list = append(list, st)
		}
	}
	if len(list) == 0 {
		return SceneOrder
	}
	return list
}

// SceneMeta provides curated atmospheric metadata for catalog preview.
type SceneMeta struct {
	MoodTag            string
	RecommendedWeather string
	Atmosphere         string
}

// GetSceneMeta returns atmospheric details for a scene.
func GetSceneMeta(st SceneType) SceneMeta {
	catIdx := CategorizeScene(st)
	switch catIdx {
	case 1:
		return SceneMeta{MoodTag: "☕ Cozy & Calm", RecommendedWeather: "🌧️ Rain / Snow", Atmosphere: "Warm indoor ambience with soft light and relaxing focus vibes"}
	case 2:
		return SceneMeta{MoodTag: "⚡ High-Tech Cyber", RecommendedWeather: "🌧️ Rain", Atmosphere: "Futuristic neon skyline with rain reflections and digital streams"}
	case 3:
		return SceneMeta{MoodTag: "🌿 Serene & Zen", RecommendedWeather: "🌸 Clear / Fog", Atmosphere: "Peaceful Japanese garden with flowing water and koi pond physics"}
	case 4:
		return SceneMeta{MoodTag: "🌌 Cosmic & Mystical", RecommendedWeather: "⭐ Starlight Night", Atmosphere: "Stargazing celestial observatory and deep space bioluminescent vistas"}
	case 5:
		return SceneMeta{MoodTag: "⚙️ Clockwork & Magic", RecommendedWeather: "☁️ Fog / Storm", Atmosphere: "Industrial brass clockwork gears and fantasy alchemy sanctum"}
	case 6:
		return SceneMeta{MoodTag: "🌅 Synthwave & Drive", RecommendedWeather: "🌆 Sunset", Atmosphere: "80s retro grid horizon with glowing neon sunsets and skyline drives"}
	default:
		return SceneMeta{MoodTag: "🌿 Focus Ambience", RecommendedWeather: "🌧️ Any Weather", Atmosphere: "Relaxing 8-bit atmospheric scene"}
	}
}
