package art

import (
	"testing"
)

func TestSceneRegistryAndCycling(t *testing.T) {
	cs := NewCottageScene(160, 90)

	if cs.CurrentScene != SceneRainyCity {
		t.Errorf("Expected initial scene SceneRainyCity, got %v", cs.CurrentScene)
	}

	s1 := cs.CycleScene()
	if s1 != SceneForestCabin {
		t.Errorf("Expected next scene SceneForestCabin, got %v", s1)
	}

	s2 := cs.CycleScene()
	if s2 != SceneCottageClassic {
		t.Errorf("Expected next scene SceneCottageClassic, got %v", s2)
	}

	s3 := cs.CycleScene()
	if s3 != SceneObservatory {
		t.Errorf("Expected next scene SceneObservatory, got %v", s3)
	}

	s4 := cs.CycleScene()
	if s4 != SceneRainyCity {
		t.Errorf("Expected loop back to SceneRainyCity, got %v", s4)
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
