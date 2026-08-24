package ui

import (
	"strconv"
	"strings"
	"testing"

	"gottage/pkg/art"

	tea "github.com/charmbracelet/bubbletea"
)

// TestSceneCountStringsTrackRegistry pins every user-facing surface that
// advertises a scene count to len(art.SceneOrder).
//
// Eight literals hardcoded "370+ scenes" while the registry held 380
// registrations. Phase 1's curation shrank the registry by roughly 8x, turning
// a mild understatement into a large overstatement that nothing detected. The
// assertions below render the real surfaces through their real constructors
// rather than grepping the source, so they describe what a user reads: any
// count that stops tracking the registry fails here, whether it was typed in
// as a literal or cached before init() finished populating SceneOrder.
func TestSceneCountStringsTrackRegistry(t *testing.T) {
	if len(art.SceneOrder) == 0 {
		t.Fatal("art.SceneOrder is empty; scene registration did not run")
	}
	count := strconv.Itoa(len(art.SceneOrder))

	// Rendered surfaces, each keyed by the substring a user would see.
	surfaces := []struct {
		name   string
		render func() string
		want   []string
	}{
		{
			name: "help modal",
			render: func() string {
				m := newSizedModel()
				m.ShowHelp = true
				return m.View()
			},
			want: []string{
				"Open Categorized Scene Catalog Modal (" + count + " Scenes)",
				"Toggle Auto-Cycle Scope (Favorites Only / All " + count + " Scenes)",
			},
		},
		{
			name: "ambiance view auto-cycle scope",
			render: func() string {
				m := newSizedModel()
				return m.renderAmbianceView()
			},
			want: []string{"All " + count + " Scenes"},
		},
		{
			name: "scene catalog modal",
			render: func() string {
				m := newSizedModel()
				m.ShowScenePicker = true
				return m.View()
			},
			want: []string{"(" + count + " Total Scenes)"},
		},
		{
			name: "scene catalog search placeholder",
			render: func() string {
				m := newSizedModel()
				m.ShowScenePicker = true
				m.PickerSearchMode = true
				return m.View()
			},
			want: []string{"filter " + count + " scenes"},
		},
		{
			name: "command palette entries",
			render: func() string {
				m := newSizedModel()
				var sb strings.Builder
				for _, cmd := range GetCommandList(&m) {
					sb.WriteString(cmd.Title + "\n" + cmd.Description + "\n")
				}
				return sb.String()
			},
			want: []string{
				"or all " + count + " scenes",
				"catalog with " + count + " pixel-art presets",
			},
		},
		{
			name: "auto-cycle scope status message via [U]",
			render: func() string {
				m := newSizedModel()
				// First press selects favorites-only; the second returns to
				// the all-scenes scope whose message carries the count.
				updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'U'}})
				updated, _ = updated.(Model).Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'U'}})
				return updated.(Model).StatusMsg
			},
			want: []string{"Auto-Cycle Scope: All " + count + " Scenes"},
		},
		{
			name: "auto-cycle scope status message via command palette",
			render: func() string {
				m := newSizedModel()
				cmds := FilterCommands(GetCommandList(&m), "Auto-Cycle: Toggle Favorites Scope")
				if len(cmds) == 0 {
					return "no matching command"
				}
				cmds[0].Action(&m)
				cmds[0].Action(&m)
				return m.StatusMsg
			},
			want: []string{"Auto-Cycle Scope: All " + count + " Scenes"},
		},
	}

	for _, s := range surfaces {
		out := s.render()
		for _, want := range s.want {
			if !strings.Contains(out, want) {
				t.Errorf("%s: expected rendered output to contain %q (registry holds %s scenes)", s.name, want, count)
			}
		}
		// Only meaningful while the live count is not itself a 370-substring.
		if !strings.Contains(count, "370") && strings.Contains(out, "370") {
			t.Errorf("%s: rendered output still advertises a hardcoded 370 scene count; derive it from len(art.SceneOrder)", s.name)
		}
	}
}

func newSizedModel() Model {
	m := NewModel()
	m.Width = 120
	m.Height = 40
	return m
}
