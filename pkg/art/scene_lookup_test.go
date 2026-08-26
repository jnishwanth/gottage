package art

import (
	"strings"
	"testing"
)

func TestSceneNamesCoversTheRegistry(t *testing.T) {
	names := SceneNames()
	if len(names) != len(SceneOrder) {
		t.Errorf("SceneNames returned %d names for %d registered scenes", len(names), len(SceneOrder))
	}
	for i, n := range names {
		if strings.TrimSpace(n) == "" {
			t.Errorf("SceneOrder[%d] has a blank display name", i)
		}
	}
}

func TestSceneByName(t *testing.T) {
	cases := []struct {
		name  string
		query string
		want  SceneType
	}{
		{"exact", "Rainy City Lofi Cafe 🏙️", SceneRainyCity},
		{"exact is case insensitive", "rainy city lofi cafe 🏙️", SceneRainyCity},
		{"unique substring skips the emoji", "Rainy City Lofi Cafe", SceneRainyCity},
		{"substring is case insensitive", "kirokaze cozy room", SceneKirokazeRoom},
		{"surrounding space is trimmed", "  kirokaze cozy room  ", SceneKirokazeRoom},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SceneByName(tc.query)
			if err != nil {
				t.Fatalf("SceneByName(%q): %v", tc.query, err)
			}
			if got != tc.want {
				t.Errorf("SceneByName(%q) = %d (%q), want %d", tc.query, got, SceneRegistry[got].Name(), tc.want)
			}
		})
	}
}

// TestSceneByNameRejectsRatherThanGuessing is the point of having this function at
// all. GetScene answers an unregistered type with SceneRainyCity, which is how two
// built-in presets pointed at scenes Phase 1's curation had cut without anything
// noticing. A lookup by name must never resolve to a scene the caller did not ask
// for.
func TestSceneByNameRejectsRatherThanGuessing(t *testing.T) {
	for _, tc := range []struct {
		name  string
		query string
		want  string
	}{
		{"no match", "definitely not a scene", "no scene matches"},
		{"empty", "", "no scene name given"},
		{"whitespace only", "   ", "no scene name given"},
		{"ambiguous", "kirokaze", "matches 2 scenes"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SceneByName(tc.query)
			if err == nil {
				t.Fatalf("SceneByName(%q) returned %d (%q) instead of an error",
					tc.query, got, SceneRegistry[got].Name())
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error %q does not mention %q", err, tc.want)
			}
		})
	}
}

// TestSceneByNamePrefersExactOverSubstring guards the case where one scene's full
// name is a substring of another's: the exact match has to win regardless of which
// one SceneOrder reaches first.
func TestSceneByNamePrefersExactOverSubstring(t *testing.T) {
	for _, id := range SceneOrder {
		name := SceneRegistry[id].Name()
		got, err := SceneByName(name)
		if err != nil {
			t.Errorf("a scene's own name %q does not resolve: %v", name, err)
			continue
		}
		if got != id {
			t.Errorf("SceneByName(%q) = %d (%q), want %d", name, got, SceneRegistry[got].Name(), id)
		}
	}
}
