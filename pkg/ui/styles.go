package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme defines colors for UI panels, borders, headers, and highlights.
type Theme struct {
	Name       string
	Border     lipgloss.Color
	Title      lipgloss.Color
	Accent     lipgloss.Color
	Highlight  lipgloss.Color
	Subdued    lipgloss.Color
	Background lipgloss.Color
}

var (
	ThemeCozyAutumn = Theme{
		Name:       "Cozy Autumn 🍁",
		Border:     lipgloss.Color("#d97736"),
		Title:      lipgloss.Color("#ffd166"),
		Accent:     lipgloss.Color("#ef476f"),
		Highlight:  lipgloss.Color("#06d6a0"),
		Subdued:    lipgloss.Color("#8d99ae"),
		Background: lipgloss.Color("#1e171b"),
	}

	ThemeRainyIndigo = Theme{
		Name:       "Rainy Indigo 🌧️",
		Border:     lipgloss.Color("#5c7cfa"),
		Title:      lipgloss.Color("#a5d8ff"),
		Accent:     lipgloss.Color("#74c0fc"),
		Highlight:  lipgloss.Color("#63e6be"),
		Subdued:    lipgloss.Color("#868e96"),
		Background: lipgloss.Color("#101424"),
	}

	ThemeSunsetPastel = Theme{
		Name:       "Sunset Pastel 🌇",
		Border:     lipgloss.Color("#f783ac"),
		Title:      lipgloss.Color("#ffdeeb"),
		Accent:     lipgloss.Color("#da77f2"),
		Highlight:  lipgloss.Color("#ffec99"),
		Subdued:    lipgloss.Color("#b197fc"),
		Background: lipgloss.Color("#211326"),
	}

	ThemeCyberpunk = Theme{
		Name:       "Cyberpunk 8-Bit ⚡",
		Border:     lipgloss.Color("#ff0055"),
		Title:      lipgloss.Color("#00f0ff"),
		Accent:     lipgloss.Color("#ffe600"),
		Highlight:  lipgloss.Color("#00ff66"),
		Subdued:    lipgloss.Color("#7000ff"),
		Background: lipgloss.Color("#050510"),
	}

	ThemeNordicFrost = Theme{
		Name:       "Nordic Frost ❄️",
		Border:     lipgloss.Color("#88c0d0"),
		Title:      lipgloss.Color("#eceff4"),
		Accent:     lipgloss.Color("#81a1c1"),
		Highlight:  lipgloss.Color("#a3be8c"),
		Subdued:    lipgloss.Color("#4c566a"),
		Background: lipgloss.Color("#2e3440"),
	}

	ThemeEmeraldForest = Theme{
		Name:       "Emerald Forest 🌲",
		Border:     lipgloss.Color("#2b8a3e"),
		Title:      lipgloss.Color("#69db7c"),
		Accent:     lipgloss.Color("#38d9a9"),
		Highlight:  lipgloss.Color("#ffd43b"),
		Subdued:    lipgloss.Color("#495057"),
		Background: lipgloss.Color("#121b16"),
	}

	AllThemes = []Theme{
		ThemeCozyAutumn,
		ThemeRainyIndigo,
		ThemeSunsetPastel,
		ThemeCyberpunk,
		ThemeNordicFrost,
		ThemeEmeraldForest,
	}
)

// CycleTheme returns the next theme in the slice.
func CycleTheme(current Theme) Theme {
	for i, t := range AllThemes {
		if t.Name == current.Name {
			return AllThemes[(i+1)%len(AllThemes)]
		}
	}
	return AllThemes[0]
}

// Styles holds lipgloss style definitions.
type Styles struct {
	Theme       Theme
	Box         lipgloss.Style
	Card        lipgloss.Style
	Title       lipgloss.Style
	Header      lipgloss.Style
	Footer      lipgloss.Style
	ActiveTab   lipgloss.Style
	InactiveTab lipgloss.Style
	KeyCap      lipgloss.Style
	KeyHelp     lipgloss.Style
	ValueHelp   lipgloss.Style
	ProgressBar lipgloss.Style
	StatusBadge lipgloss.Style
	PillBadge   lipgloss.Style
	MutedText   lipgloss.Style
	BoldAccent  lipgloss.Style
}

// NewStyles creates lipgloss styles using given Theme.
func NewStyles(t Theme) Styles {
	return Styles{
		Theme: t,
		Box: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.Border).
			Padding(0, 1),
		Card: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(t.Subdued).
			Padding(0, 1),
		Title: lipgloss.NewStyle().
			Foreground(t.Title).
			Bold(true),
		Header: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),
		Footer: lipgloss.NewStyle().
			Foreground(t.Subdued).
			Italic(true),
		ActiveTab: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(t.Border).
			Bold(true).
			Padding(0, 1),
		InactiveTab: lipgloss.NewStyle().
			Foreground(t.Subdued).
			Padding(0, 1),
		KeyCap: lipgloss.NewStyle().
			Foreground(t.Title).
			Background(lipgloss.Color("#282c34")).
			Bold(true).
			Padding(0, 1),
		KeyHelp: lipgloss.NewStyle().
			Foreground(t.Title).
			Bold(true),
		ValueHelp: lipgloss.NewStyle().
			Foreground(t.Subdued),
		ProgressBar: lipgloss.NewStyle().
			Foreground(t.Highlight),
		StatusBadge: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#101424")).
			Background(t.Highlight).
			Bold(true).
			Padding(0, 1),
		PillBadge: lipgloss.NewStyle().
			Foreground(t.Title).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.Border).
			Padding(0, 1),
		MutedText: lipgloss.NewStyle().
			Foreground(t.Subdued),
		BoldAccent: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true),
	}
}
