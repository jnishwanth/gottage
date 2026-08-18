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
