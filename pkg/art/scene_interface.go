package art

import (
	"fmt"
	"strconv"
	"strings"
)

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
	RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct {
		X, Y  int
		Phase float64
	})

	// RenderForeground renders Layer 3..5: Window frames, room walls, floor, desk, objects, and Focus Buddy companion.
	RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed)

	// ApplyLighting renders Layer 6: Dynamic light cones, fireplace warmth glows, and lightning flashes.
	ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool)
}

// SceneRegistry holds registered modular scene implementations.
var SceneRegistry = make(map[SceneType]SceneInterface)
var SceneOrder = []SceneType{}

// RegisterScene adds a modular scene implementation to the registry.
//
// A duplicate ID is a panic, not a warning. Registration happens in init(), so
// this fails the process and every test binary immediately and unmissably. The
// previous version assigned unconditionally, which let 20 AnimatedGifScene
// literals silently shadow 20 handcrafted scenes across 8 commits; log-and-skip
// would only have converted that silent-overwrite bug into a silent-skip bug.
func RegisterScene(s SceneInterface) {
	if existing, ok := SceneRegistry[s.ID()]; ok {
		panic(fmt.Sprintf("duplicate scene ID %d (%q): %T would shadow %T",
			s.ID(), s.Name(), s, existing))
	}
	SceneOrder = append(SceneOrder, s.ID())
	SceneRegistry[s.ID()] = s
}

// GetScene returns the registered scene implementation for a given SceneType.
func GetScene(st SceneType) SceneInterface {
	if s, ok := SceneRegistry[st]; ok {
		return s
	}
	return SceneRegistry[SceneRainyCity]
}

// SceneNames returns every registered scene's display name, in SceneOrder.
func SceneNames() []string {
	names := make([]string, 0, len(SceneOrder))
	for _, id := range SceneOrder {
		names = append(names, SceneRegistry[id].Name())
	}
	return names
}

// SceneByName resolves a display name to a registered SceneType.
//
// Deliberately not built on GetScene, which answers an unregistered type with
// SceneRainyCity. A caller that asks for one scene and quietly receives another
// is the same silent failure Phase 1 spent itself removing -- and it masked two
// built-in presets pointing at scenes curation had cut.
//
// An exact case-insensitive match wins. Failing that, a unique case-insensitive
// substring match wins, so a name carrying an emoji ("Kirokaze Cozy Room \U0001F3A8")
// stays reachable from a shell. Ambiguity is an error naming the candidates,
// never a pick.
func SceneByName(name string) (SceneType, error) {
	want := strings.ToLower(strings.TrimSpace(name))
	if want == "" {
		return 0, fmt.Errorf("no scene name given")
	}

	var partial []SceneType
	for _, id := range SceneOrder {
		have := strings.ToLower(SceneRegistry[id].Name())
		if have == want {
			return id, nil
		}
		if strings.Contains(have, want) {
			partial = append(partial, id)
		}
	}

	switch len(partial) {
	case 0:
		return 0, fmt.Errorf("no scene matches %q", name)
	case 1:
		return partial[0], nil
	default:
		matched := make([]string, 0, len(partial))
		for _, id := range partial {
			matched = append(matched, strconv.Quote(SceneRegistry[id].Name()))
		}
		return 0, fmt.Errorf("%q matches %d scenes: %s", name, len(partial), strings.Join(matched, ", "))
	}
}
