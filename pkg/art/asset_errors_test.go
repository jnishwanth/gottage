package art

import (
	"strings"
	"testing"
)

// These tests mutate package-global state (the error record and the loader
// cache), so none of them call t.Parallel.

func TestAssetLoadFailureIsRecorded(t *testing.T) {
	ClearAssetErrors()
	t.Cleanup(ClearAssetErrors)

	const missing = "assets/animations/no_such_scene.gif"
	if _, err := LoadAnimatedGif(missing, 8, 8); err == nil {
		t.Fatal("loading a missing asset returned no error")
	}

	recorded := AssetErrors()
	if _, ok := recorded[missing]; !ok {
		t.Errorf("%s failed to load but was not recorded; record holds %v", missing, keysOf(recorded))
	}
}

func TestSuccessfulLoadRecordsNothing(t *testing.T) {
	ClearAssetErrors()
	t.Cleanup(ClearAssetErrors)

	if _, err := LoadAnimatedGif("assets/animations/kirokaze_room.gif", 8, 8); err != nil {
		t.Fatalf("loading a real asset failed: %v", err)
	}
	if recorded := AssetErrors(); len(recorded) != 0 {
		t.Errorf("a successful load recorded %v", keysOf(recorded))
	}
}

// TestRenderSkyRecordsAssetFailure is the one that matters: it exercises the
// exact path that swallowed failures before -- a scene whose asset will not
// load, rendered through the interface the app calls.
func TestRenderSkyRecordsAssetFailure(t *testing.T) {
	ClearAssetErrors()
	t.Cleanup(ClearAssetErrors)

	broken := &AnimatedGifScene{
		id:        SceneRainyCity, // not registered; this literal is never handed to RegisterScene
		name:      "Broken Asset",
		gifPath:   "assets/animations/no_such_scene.gif",
		winBounds: Rect{X: 0, Y: 0, W: 4, H: 4},
	}

	c := NewCanvas(8, 8)
	broken.RenderSky(c, TimeNight, false, 0, nil)

	if _, ok := AssetErrors()[broken.gifPath]; !ok {
		t.Error("RenderSky fell back to the navy fill without recording why")
	}
	// The fallback itself is unchanged and still deliberate: replacing what the
	// user sees is Phase 16's call, not this phase's.
	if got := distinctColors(c); got != 1 {
		t.Errorf("fallback rendered %d distinct colors, want the flat fill's 1", got)
	}
}

func TestAssetErrorsReturnsACopy(t *testing.T) {
	ClearAssetErrors()
	t.Cleanup(ClearAssetErrors)

	const missing = "assets/animations/no_such_scene.gif"
	_, _ = LoadAnimatedGif(missing, 8, 8)

	got := AssetErrors()
	delete(got, missing)
	if _, ok := AssetErrors()[missing]; !ok {
		t.Error("mutating the returned map mutated the record")
	}
}

func keysOf(m map[string]error) string {
	if len(m) == 0 {
		return "(nothing)"
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
