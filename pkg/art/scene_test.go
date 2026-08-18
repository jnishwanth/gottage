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

func TestCanvasRenderModes(t *testing.T) {
	cs := NewCottageScene(80, 40)

	if cs.RenderMode != RenderHalfBlock {
		t.Errorf("Expected default RenderMode RenderHalfBlock, got %v", cs.RenderMode)
	}

	modes := []CanvasRenderMode{RenderASCII, RenderBraille, RenderPhosphor, RenderHalfBlock}
	for _, expected := range modes {
		next := cs.CycleRenderMode()
		if next != expected {
			t.Errorf("Expected next RenderMode %v, got %v", expected, next)
		}
		rendered := cs.Render()
		if len(rendered) == 0 {
			t.Errorf("Expected non-empty rendered string for mode %v", expected)
		}
	}

	// Test Canvas clone preserves RenderMode
	c := NewCanvas(20, 20)
	c.RenderMode = RenderBraille
	cloned := c.Clone()
	if cloned.RenderMode != RenderBraille {
		t.Errorf("Expected cloned RenderMode RenderBraille, got %v", cloned.RenderMode)
	}
}

func TestProceduralShadersAndRipples(t *testing.T) {
	cs := NewCottageScene(80, 40)

	// Test 9 shader modes cycle clean
	for i := 0; i < 9; i++ {
		s := cs.CycleShaderEffect()
		if len(s.String()) == 0 {
			t.Errorf("Expected non-empty shader string description for index %d", i)
		}
		cs.RenderFrame()
	}

	// Test Canvas Energy Ripple trigger & rendering
	cs.TriggerCanvasRippleHex(40, 20, "#ff5588")
	if len(cs.Ripples) != 1 {
		t.Errorf("Expected 1 active canvas ripple, got %d", len(cs.Ripples))
	}
	cs.RenderFrame()
	cs.Update()
	if cs.Ripples[0].Life != 1 {
		t.Errorf("Expected ripple life 1 after update, got %v", cs.Ripples[0].Life)
	}
}

func TestPetCatCompanion(t *testing.T) {
	cs := NewCottageScene(80, 40)
	initialBliss := cs.CatBliss

	msg := cs.PetCat()
	if cs.CatBliss <= initialBliss {
		t.Errorf("Expected CatBliss to increase after petting, got %d (initial %d)", cs.CatBliss, initialBliss)
	}
	if len(msg) == 0 {
		t.Errorf("Expected non-empty purr message from PetCat")
	}
	if len(cs.OverlayParticles) == 0 {
		t.Errorf("Expected heart overlay particles spawned after PetCat")
	}
	if len(cs.Ripples) == 0 {
		t.Errorf("Expected rose ripple created on canvas after PetCat")
	}
}
