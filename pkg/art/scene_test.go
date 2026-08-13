package art

import (
	"testing"
)

func TestSceneRegistryAndCycling(t *testing.T) {
	cs := NewCottageScene(160, 90)

	if cs.CurrentScene != SceneRainyCity {
		t.Errorf("Expected initial scene SceneRainyCity, got %v", cs.CurrentScene)
	}

	initialScene := cs.CurrentScene
	count := len(SceneOrder)
	for i := 0; i < count; i++ {
		cs.CycleScene()
	}

	if cs.CurrentScene != initialScene {
		t.Errorf("Expected loop back to initial scene %v after %d cycles, got %v", initialScene, count, cs.CurrentScene)
	}
}

func TestWindowBounds(t *testing.T) {
	cs := NewCottageScene(160, 90)

	cs.CurrentScene = SceneRainyCity
	rCity := cs.WindowRect()
	if rCity.W != 120 || rCity.H != 58 {
		t.Errorf("Unexpected City window bounds: %+v", rCity)
	}

	cs.CurrentScene = SceneCottageClassic
	rCot := cs.WindowRect()
	if rCot.W != 80 || rCot.H != 46 {
		t.Errorf("Unexpected Cottage window bounds: %+v", rCot)
	}
}
