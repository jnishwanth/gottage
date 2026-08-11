package art

// SceneInterface defines the modular contract for any Gottage visual scene.
// Implementing this interface allows adding new scenes seamlessly to Gottage!
type SceneInterface interface {
	// ID returns the SceneType enum identifier.
	ID() SceneType

	// Name returns the human-readable display name of the scene.
	Name() string

	// WindowBounds returns the window glass bounding box for weather clipping.
	WindowBounds() Rect

	// RenderSky renders Layer 1: Sky, stars, moon, distant mountains/buildings outside window.
	RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 })

	// RenderForeground renders Layer 3..5: Window frames, room walls, floor, desk, objects, and Focus Buddy companion.
	RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed)

	// ApplyLighting renders Layer 6: Dynamic light cones, fireplace warmth glows, and lightning flashes.
	ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool)
}

// SceneRegistry holds registered modular scene implementations.
var SceneRegistry = make(map[SceneType]SceneInterface)

// RegisterScene adds a modular scene implementation to the registry.
func RegisterScene(s SceneInterface) {
	SceneRegistry[s.ID()] = s
}

// GetScene returns the registered scene implementation for a given SceneType.
func GetScene(st SceneType) SceneInterface {
	if s, ok := SceneRegistry[st]; ok {
		return s
	}
	return SceneRegistry[SceneRainyCity]
}
