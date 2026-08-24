package art

import (
	"testing"
)

func TestCategorizeScene(t *testing.T) {
	if len(SceneOrder) == 0 {
		t.Fatal("Expected non-empty SceneOrder")
	}

	for i := 0; i < len(SceneCategories); i++ {
		scenes := GetScenesByCategory(i)
		if len(scenes) == 0 {
			t.Errorf("Category %d (%s) returned 0 scenes", i, SceneCategories[i])
		}
	}
}

// TestEveryCategoryMatchesARegisteredScene asserts each non-favorites category
// actually holds at least one registered scene, and that CategorizeScene never
// returns an index outside SceneCategories.
//
// It counts CategorizeScene results directly rather than calling
// GetScenesByCategory, because GetScenesByCategoryAndFavorites returns the whole
// of SceneOrder when a category matches nothing. That fallback keeps
// TestCategorizeScene green for a category that matches zero scenes, while the
// picker silently lists every scene in the app under that category's header --
// the state curation left "Steampunk & Fantasy" in once its scenes were cut.
//
// The range check guards the other direction: view.go indexes SceneCategories
// with a CategorizeScene result, so a switch case returning an index the slice
// does not have panics the picker.
func TestEveryCategoryMatchesARegisteredScene(t *testing.T) {
	counts := make([]int, len(SceneCategories))
	for _, st := range SceneOrder {
		idx := CategorizeScene(st)
		if idx < 1 || idx >= len(SceneCategories) {
			t.Errorf("CategorizeScene(%q) = %d, outside the valid range 1..%d; indexing SceneCategories with it would panic",
				st.String(), idx, len(SceneCategories)-1)
			continue
		}
		counts[idx]++
	}

	for i := 1; i < len(SceneCategories); i++ {
		if counts[i] == 0 {
			t.Errorf("category %d (%s) matches no registered scene, so the picker falls back to listing all %d scenes under that header; add matching keywords or drop the category",
				i, SceneCategories[i], len(SceneOrder))
		}
	}
}
