package art

import (
	"strings"
)

// SceneCategories lists display names for categorized scene browser.
var SceneCategories = []string{
	"🏡 Cozy & Cottage",
	"⚡ Cyberpunk & Neon",
	"🌿 Zen & Nature",
	"🌌 Sci-Fi & Cosmic",
	"⚙️ Steampunk & Fantasy",
	"🌅 Synthwave & Retro",
}

// CategorizeScene determines category index (0..5) for a given SceneType.
func CategorizeScene(st SceneType) int {
	name := strings.ToLower(st.String())
	switch {
	case strings.Contains(name, "cyber") || strings.Contains(name, "matrix") || strings.Contains(name, "neon") || strings.Contains(name, "kirokaze"):
		return 1
	case strings.Contains(name, "zen") || strings.Contains(name, "bamboo") || strings.Contains(name, "shrine") || strings.Contains(name, "torii") || strings.Contains(name, "waterfall") || strings.Contains(name, "cherry") || strings.Contains(name, "bonsai") || strings.Contains(name, "grove") || strings.Contains(name, "lotus") || strings.Contains(name, "garden"):
		return 2
	case strings.Contains(name, "observatory") || strings.Contains(name, "space") || strings.Contains(name, "nebula") || strings.Contains(name, "astral") || strings.Contains(name, "zodiac") || strings.Contains(name, "alien") || strings.Contains(name, "stargate") || strings.Contains(name, "cosmos") || strings.Contains(name, "stargazing") || strings.Contains(name, "abyssal") || strings.Contains(name, "deep sea"):
		return 3
	case strings.Contains(name, "steampunk") || strings.Contains(name, "clockwork") || strings.Contains(name, "citadel") || strings.Contains(name, "airship") || strings.Contains(name, "castle") || strings.Contains(name, "magic") || strings.Contains(name, "wizard") || strings.Contains(name, "haunted") || strings.Contains(name, "volcano") || strings.Contains(name, "forge") || strings.Contains(name, "pirate"):
		return 4
	case strings.Contains(name, "synthwave") || strings.Contains(name, "highway") || strings.Contains(name, "arcade") || strings.Contains(name, "retro") || strings.Contains(name, "grid") || strings.Contains(name, "racer") || strings.Contains(name, "coastal"):
		return 5
	default:
		return 0 // Cozy & Cottage
	}
}

// GetScenesByCategory returns slice of SceneType in given category index.
func GetScenesByCategory(catIdx int) []SceneType {
	var list []SceneType
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
