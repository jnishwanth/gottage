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
	case strings.Contains(name, "zen") || strings.Contains(name, "bamboo") || strings.Contains(name, "shrine") || strings.Contains(name, "torii") || strings.Contains(name, "waterfall") || strings.Contains(name, "cherry") || strings.Contains(name, "bonsai") || strings.Contains(name, "grove") || strings.Contains(name, "lotus") || strings.Contains(name, "garden"):
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
