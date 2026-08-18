package features

// InspirationQuote represents a tranquil mindfulness quote for terminal focus sessions.
type InspirationQuote struct {
	Quote  string
	Author string
}

// CuratedQuotes holds 24 focus-enhancing mindfulness prompts.
var CuratedQuotes = []InspirationQuote{
	{Quote: "In tranquility, clarity blossoms.", Author: "Zen Wisdom"},
	{Quote: "Step by step, line by line.", Author: "Focus Mantra"},
	{Quote: "Quiet mind, sharp focus.", Author: "Terminal Focus"},
	{Quote: "Breathe in peace, breathe out clutter.", Author: "Mindfulness"},
	{Quote: "Slow down to go fast.", Author: "Craftsmanship"},
	{Quote: "Deep work is a form of art.", Author: "Cal Newport"},
	{Quote: "Find joy in the simple process.", Author: "Cozy Life"},
	{Quote: "Small steps every day lead to great paths.", Author: "Proverb"},
	{Quote: "Focus is the art of knowing what to ignore.", Author: "R.D. Sherwin"},
	{Quote: "Serenity is not freedom from storm, but peace within.", Author: "Anonymous"},
	{Quote: "Craft each line with care and intention.", Author: "Developer Zen"},
	{Quote: "Silence is where your best ideas grow.", Author: "Creative Mind"},
	{Quote: "Create a space where your mind can soar.", Author: "Gottage Sanctuary"},
	{Quote: "Flow like water, steady like a mountain.", Author: "Taoist Maxim"},
	{Quote: "Rest is the field that gathers strength.", Author: "Ovid"},
	{Quote: "Simplicity is the ultimate sophistication.", Author: "Leonardo da Vinci"},
	{Quote: "Patience and persistence build wonders.", Author: "Focus Philosophy"},
	{Quote: "Keep your eyes on the horizon, feet on the ground.", Author: "Mindful Living"},
	{Quote: "The quiet mind hears the whisper of genius.", Author: "Focus Proverb"},
	{Quote: "Every small effort builds a sanctuary of focus.", Author: "Atmosphere"},
	{Quote: "Warm tea, rain on glass, deep code.", Author: "Cozy Focus"},
	{Quote: "Stars shine brightest in the darkest skies.", Author: "Night Work"},
	{Quote: "Embrace the rhythm of steady creation.", Author: "Focus Rhythm"},
	{Quote: "Presence is the greatest gift to your craft.", Author: "Mindful Code"},
}

// InspirationManager tracks active mindfulness inspiration quotes.
type InspirationManager struct {
	ActiveIdx int
}

// NewInspirationManager initializes the inspiration quote manager.
func NewInspirationManager() *InspirationManager {
	return &InspirationManager{
		ActiveIdx: 0,
	}
}

// CurrentQuote returns the active inspiration quote.
func (im *InspirationManager) CurrentQuote() InspirationQuote {
	if len(CuratedQuotes) == 0 {
		return InspirationQuote{Quote: "Focus on the present.", Author: "Mindfulness"}
	}
	return CuratedQuotes[im.ActiveIdx%len(CuratedQuotes)]
}

// NextQuote cycles to the next inspiration quote and returns it.
func (im *InspirationManager) NextQuote() InspirationQuote {
	if len(CuratedQuotes) == 0 {
		return InspirationQuote{Quote: "Focus on the present.", Author: "Mindfulness"}
	}
	im.ActiveIdx = (im.ActiveIdx + 1) % len(CuratedQuotes)
	return im.CurrentQuote()
}
