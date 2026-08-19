package features

import (
	"fmt"
)

// DrinkType specifies the cozy hot beverage brewed in the cottage.
type DrinkType int

const (
	DrinkDripCoffee DrinkType = iota
	DrinkMatchaLatte
	DrinkEarlGrey
	DrinkHotCocoa
	DrinkSpicedChai
	DrinkChamomile
	DrinkBobaMilkTea
	DrinkLemonGinger
)

// DrinkInfo describes a handcrafted warm beverage.
type DrinkInfo struct {
	Type        DrinkType
	Name        string
	ShortName   string
	Icon        string
	Aroma       string
	ColorHex    string
	SteamSymbol string
	FlavorNotes string
}

// CuratedDrinks provides rich metadata for all cozy beverages.
var CuratedDrinks = []DrinkInfo{
	{
		Type:        DrinkDripCoffee,
		Name:        "Pour-Over Dark Roast Coffee ☕",
		ShortName:   "Dark Roast Coffee",
		Icon:        "☕",
		Aroma:       "Rich roasted hazelnut & dark cocoa aroma",
		ColorHex:    "#78350f",
		SteamSymbol: "♨️",
		FlavorNotes: "Slow-dripped Ethiopian beans with velvety crema and hazelnut clarity.",
	},
	{
		Type:        DrinkMatchaLatte,
		Name:        "Ceremonial Matcha Latte 🍵",
		ShortName:   "Matcha Latte",
		Icon:        "🍵",
		Aroma:       "Freshly whisked stone-ground green tea & oat milk froth",
		ColorHex:    "#4ade80",
		SteamSymbol: "🍃",
		FlavorNotes: "Uji ceremonial matcha whisked with warm oat foam for mindful focus.",
	},
	{
		Type:        DrinkEarlGrey,
		Name:        "Artisan Earl Grey Tea 🫖",
		ShortName:   "Earl Grey Tea",
		Icon:        "🫖",
		Aroma:       "Fragrant Italian bergamot citrus & blue cornflower petals",
		ColorHex:    "#d97706",
		SteamSymbol: "🌿",
		FlavorNotes: "Single-origin Ceylon black tea steeped with cold-pressed bergamot peel.",
	},
	{
		Type:        DrinkHotCocoa,
		Name:        "Marshmallow Hot Chocolate 🍫",
		ShortName:   "Hot Chocolate",
		Icon:        "🍫",
		Aroma:       "Thick Dutch dark chocolate with melting mini marshmallows",
		ColorHex:    "#b45309",
		SteamSymbol: "☁️",
		FlavorNotes: "Melted Valrhona cocoa simmered with sweet vanilla and fluffy marshmallows.",
	},
	{
		Type:        DrinkSpicedChai,
		Name:        "Honey Cardamom Spiced Chai 🍯",
		ShortName:   "Spiced Chai",
		Icon:        "🍯",
		Aroma:       "Crushed cardamom pods, cinnamon bark & wild blossom honey",
		ColorHex:    "#f59e0b",
		SteamSymbol: "✨",
		FlavorNotes: "Simmered Assam tea with cracked spices, fresh ginger, and clover honey.",
	},
	{
		Type:        DrinkChamomile,
		Name:        "Calming Chamomile Blossom 🌼",
		ShortName:   "Chamomile Blossom",
		Icon:        "🌼",
		Aroma:       "Sun-dried chamomile blossoms with a touch of lavender",
		ColorHex:    "#fde047",
		SteamSymbol: "🌸",
		FlavorNotes: "Caffeine-free whole flower chamomile infusion for deep tranquil calm.",
	},
	{
		Type:        DrinkBobaMilkTea,
		Name:        "Brown Sugar Boba Milk Tea 🧋",
		ShortName:   "Boba Milk Tea",
		Icon:        "🧋",
		Aroma:       "Caramelized brown sugar syrup with roasted oolong tea",
		ColorHex:    "#a16207",
		SteamSymbol: "🫧",
		FlavorNotes: "Warm slow-cooked brown sugar tapioca pearls in velvety milk tea.",
	},
	{
		Type:        DrinkLemonGinger,
		Name:        "Warm Lemon Ginger Honey Elixir 🍋",
		ShortName:   "Lemon Ginger Elixir",
		Icon:        "🍋",
		Aroma:       "Freshly squeezed Meyer lemon & grated spicy ginger root",
		ColorHex:    "#eab308",
		SteamSymbol: "💫",
		FlavorNotes: "Warm soothing lemon tonic with spicy raw ginger and golden wildflower honey.",
	},
}

// BeverageManager tracks the active cozy beverage, warmth level, and sips.
type BeverageManager struct {
	ActiveIdx   int
	Warmth      int // 0 to 100%
	SipCount    int
	TotalBrewed int
}

// NewBeverageManager creates a new beverage manager initialized with coffee.
func NewBeverageManager() *BeverageManager {
	return &BeverageManager{
		ActiveIdx:   0,
		Warmth:      100,
		SipCount:    0,
		TotalBrewed: 1,
	}
}

// CurrentDrink returns the current DrinkInfo.
func (bm *BeverageManager) CurrentDrink() DrinkInfo {
	if len(CuratedDrinks) == 0 {
		return DrinkInfo{Name: "Warm Tea 🍵", ShortName: "Tea", Icon: "🍵"}
	}
	return CuratedDrinks[bm.ActiveIdx%len(CuratedDrinks)]
}

// BrewDrink sets the active drink and resets warmth to 100%.
func (bm *BeverageManager) BrewDrink(idx int) DrinkInfo {
	if len(CuratedDrinks) == 0 {
		return DrinkInfo{}
	}
	bm.ActiveIdx = idx % len(CuratedDrinks)
	if bm.ActiveIdx < 0 {
		bm.ActiveIdx += len(CuratedDrinks)
	}
	bm.Warmth = 100
	bm.TotalBrewed++
	return bm.CurrentDrink()
}

// NextDrink cycles to the next beverage in the curated menu and brews it.
func (bm *BeverageManager) NextDrink() DrinkInfo {
	return bm.BrewDrink(bm.ActiveIdx + 1)
}

// PrevDrink cycles to the previous beverage in the curated menu and brews it.
func (bm *BeverageManager) PrevDrink() DrinkInfo {
	return bm.BrewDrink(bm.ActiveIdx - 1)
}

// TakeSip takes a sip from the current beverage, reducing warmth slightly and returning feedback.
func (bm *BeverageManager) TakeSip() (string, int) {
	bm.SipCount++
	// Warmth decreases by 3% per sip, but stays cozy at minimum 25%
	if bm.Warmth > 25 {
		bm.Warmth -= 3
	}
	drink := bm.CurrentDrink()
	msg := fmt.Sprintf("Took a warm sip of %s. %s", drink.ShortName, drink.FlavorNotes)
	return msg, bm.Warmth
}

// StatusBadge returns a compact status string for TUI header / status bar.
func (bm *BeverageManager) StatusBadge() string {
	drink := bm.CurrentDrink()
	return fmt.Sprintf("%s %s (%d%% ♨️)", drink.Icon, drink.ShortName, bm.Warmth)
}
