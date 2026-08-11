package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme defines colors for UI panels, borders, headers, and highlights.
type Theme struct {
	Name        string
	Border      lipgloss.Color
	Title       lipgloss.Color
	Accent      lipgloss.Color
	Highlight   lipgloss.Color
	Subdued     lipgloss.Color
	Background  lipgloss.Color
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
)

// Styles holds lipgloss style definitions.
type Styles struct {
	Theme        Theme
	Box          lipgloss.Style
	Title        lipgloss.Style
	Header       lipgloss.Style
	Footer       lipgloss.Style
	ActiveTab    lipgloss.Style
	InactiveTab  lipgloss.Style
	KeyHelp      lipgloss.Style
	ValueHelp    lipgloss.Style
	ProgressBar  lipgloss.Style
	StatusBadge  lipgloss.Style
}

// NewStyles creates lipgloss styles using given Theme.
func NewStyles(t Theme) Styles {
	return Styles{
		Theme: t,
		Box: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.Border).
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
	}
}
