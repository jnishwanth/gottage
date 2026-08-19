package art

import (
	"github.com/lucasb-eyer/go-colorful"
)

// PICO-8 16-Color Retro Palette Definitions
var (
	PicoBlack      = hex("#000000") // 0: Black
	PicoNavy       = hex("#1D2B53") // 1: Dark Blue
	PicoDarkPurple = hex("#7E2553") // 2: Dark Purple
	PicoGreen      = hex("#008751") // 3: Dark Green
	PicoBrown      = hex("#AB5236") // 4: Brown
	PicoDarkGrey   = hex("#5F574F") // 5: Dark Grey
	PicoLightGrey  = hex("#C2C3C7") // 6: Light Grey
	PicoWhite      = hex("#FFF1E8") // 7: White
	PicoRed        = hex("#FF004D") // 8: Red
	PicoOrange     = hex("#FFA300") // 9: Orange
	PicoYellow     = hex("#FFEC27") // 10: Yellow
	PicoLime       = hex("#00E436") // 11: Green
	PicoCyan       = hex("#29ADFF") // 12: Blue
	PicoLavender   = hex("#83769C") // 13: Indigo / Lavender
	PicoPink       = hex("#FF77A8") // 14: Pink
	PicoPeach      = hex("#FFCCAA") // 15: Peach

	// Cozy Cottage Natural Palette
	ColorWoodDark       = hex("#2c1810") // Dark Walnut Wood
	ColorWoodMedium     = hex("#5c3826") // Warm Oak Wood
	ColorWoodLight      = hex("#8c5835") // Honey Pine Wood
	ColorWoodHighlight  = hex("#b87d4b") // Birch Highlight
	ColorWoodPlankSep   = hex("#1a0f0a") // Wood Floor Seam
	ColorFloorRugBase   = hex("#702632") // Persian Rug Maroon
	ColorFloorRugBorder = hex("#d4a373") // Rug Cream Border
	ColorFloorRugGold   = hex("#e9c46a") // Rug Gold Pattern
	ColorWallWarm       = hex("#231f20") // Warm Interior Shadow
	ColorWallPlank      = hex("#3d281d") // Wall Wainscoting
	ColorWallPaper      = hex("#342b38") // Cozy Night Wallpaper
	ColorWallPaperDay   = hex("#6b5b6e") // Cozy Day Wallpaper

	// Fireplace & Warm Light
	ColorHearthBrick    = hex("#6e2d22") // Terracotta Brick
	ColorHearthBrickDark= hex("#441b15") // Brick Shadow
	ColorHearthMortar   = hex("#8f6a5b") // Brick Mortar
	ColorHearthInterior = hex("#140a08") // Fireplace Backwall
	ColorFireCore       = hex("#fff3b0") // White-Hot Flame Core
	ColorFireYellow     = hex("#ffcc00") // Bright Golden Flame
	ColorFireOrange     = hex("#ff7b00") // Rich Orange Flame
	ColorFireRed        = hex("#d9381e") // Deep Red Flame
	ColorFireEmber      = hex("#801100") // Smoldering Embers
	ColorWarmGlow       = hex("#ffd166") // Ambient Lamp/Fire Glow
	ColorLampShade      = hex("#2a9d8f") // Classic Banker's Green / Brass Shade
	ColorLampBrass      = hex("#e7c169") // Polished Brass Lamp Base
	ColorCandleWax      = hex("#faedcd") // Cream Candle Wax

	// Nature & Plants
	ColorLeafDark       = hex("#1b4332") // Dark Monstera Leaf
	ColorLeafMedium     = hex("#2d6a4f") // Pothos Green
	ColorLeafLight      = hex("#52b788") // Fresh Vine Sprout
	ColorLeafHighlight  = hex("#95d5b2") // Leaf Sun Bleed
	ColorPotClay        = hex("#b05a3c") // Terracotta Plant Pot

	// Sky & Atmosphere
	ColorSkyNightTop    = hex("#0a0e17") // Deep Cosmic Obsidian
	ColorSkyNightMid    = hex("#121b2d") // Midnight Indigo
	ColorSkyNightBot    = hex("#1e2942") // Distant Horizon Blue
	ColorSkySunsetTop   = hex("#2b1b3d") // Violet Sunset Dusk
	ColorSkySunsetMid   = hex("#a23e48") // Rose Sunset Horizon
	ColorSkySunsetBot   = hex("#f38d68") // Golden Amber Horizon
	ColorSkyDawnTop     = hex("#1f1a3a") // Twilight Purple
	ColorSkyDawnMid     = hex("#6c4f70") // Morning Lilac
	ColorSkyDawnBot     = hex("#e8998d") // Peach Dawn Sky
	ColorSkyDayTop      = hex("#3a7ca5") // Crisp Blue Sky
	ColorSkyDayMid      = hex("#81c3d7") // Soft Atmosphere
	ColorSkyDayBot      = hex("#d9dcd6") // Misty Distant Horizon

	// Window & Rain
	ColorWindowFrame    = hex("#3e2723") // Dark Timber Window Frame
	ColorWindowSill     = hex("#5d4037") // Window Sill Wood
	ColorGlassTint      = hex("#1b263b") // Rainwashed Glass
	ColorRainDrop       = hex("#a0c4e2") // Rain Streak
	ColorRainSplash     = hex("#e0fbfc") // Window Sill Rain Mist
	ColorSnowFlake      = hex("#ffffff") // Crisp Snowflake
	ColorStars          = hex("#fff8dc") // Warm Twinkling Star

	// Books, Mug & Details
	ColorBookSpine1     = hex("#9b2226") // Burgundy Book
	ColorBookSpine2     = hex("#005f73") // Deep Teal Book
	ColorBookSpine3     = hex("#bb3e03") // Rust Orange Book
	ColorBookSpine4     = hex("#588157") // Sage Green Book
	ColorMugCeramic     = hex("#e76f51") // Warm Ceramic Terracotta Mug
	ColorMugSteam       = hex("#e2eafc") // Soft Drifting Steam
	ColorNotebookPaper  = hex("#fdf0d5") // Cream Journal Pages
	ColorCatCushion     = hex("#3d348b") // Velvet Indigo Cushion

	// Zen Japanese Porch & Tea Garden
	ColorTatamiBase     = hex("#c2b280") // Warm Golden Straw Tatami
	ColorTatamiBorder   = hex("#1c1917") // Dark Brocade Tatami Border
	ColorShojiPaper     = hex("#f7f4ea") // Translucent Rice Paper
	ColorShojiLattice   = hex("#44281d") // Dark Cedar Shoji Grid
	ColorPondDeep       = hex("#14213d") // Deep Tranquil Koi Pond
	ColorPondShallow    = hex("#2a4365") // Pond Water Horizon
	ColorKoiGold        = hex("#ff9f1c") // Golden Koi Fish
	ColorKoiWhite       = hex("#fdf0ed") // Pearl White Koi
	ColorSakuraPink     = hex("#ffb7b2") // Falling Cherry Blossom Petal
	ColorSakuraSoft     = hex("#ffe5ec") // Soft Petal Highlight
	ColorStoneLantern   = hex("#5c677d") // Garden Stone Lantern
	ColorMatchaGreen    = hex("#588157") // Rich Whisked Matcha Tea

	// Alpine Campfire & Lake
	ColorLakeDeep       = hex("#0b132b") // Alpine Night Lake
	ColorLakeReflect    = hex("#1c2541") // Mountain Moonlit Water
	ColorPineSilhouette = hex("#0f172a") // Starlit Pine Forest
	ColorCampTent       = hex("#b08968") // Canvas A-Frame Tent
	ColorPlaidRed       = hex("#9e2a2b") // Cozy Plaid Wool Blanket
	ColorPlaidNavy      = hex("#1b263b") // Plaid Weave Contrast
	ColorCampSmoke      = hex("#94a3b8") // Drifting Campfire Smoke

	// Artist Atelier & Winter Hearth
	ColorCanvasEasel    = hex("#7f5539") // Studio Pine Easel
	ColorCanvasCloth    = hex("#ede0d4") // Stretched Canvas Fabric
	ColorLavenderDried  = hex("#7b6d8d") // Hanging Dried Lavender
	ColorEucalyptus     = hex("#84a98c") // Dried Eucalyptus Leaves
	ColorCocoaHot       = hex("#3b1f14") // Rich Hot Chocolate
	ColorMarshmallow    = hex("#fffdfa") // Fluffy Marshmallow
	ColorFrostIce       = hex("#e0f2fe") // Window Frost Crystals

	// Greenhouse & Glass Solarium
	ColorGreenhouseFrame= hex("#1f2d16") // Dark Forest Green Iron Ribs
	ColorMonsteraDark   = hex("#0f2411") // Deep Emerald Monstera Leaf
	ColorFernGreen      = hex("#38b000") // Lush Trailing Fern
	ColorHangingPotClay = hex("#cb6843") // Terracotta Hanging Planter
	ColorCopperWaterCan = hex("#b87333") // Polished Copper Watering Can

	// Midnight Stained Glass Library
	ColorStainedGlassBlue = hex("#1d3557") // Sapphire Stained Glass
	ColorStainedGlassGold = hex("#e9c46a") // Warm Amber Stained Glass
	ColorStainedGlassRuby = hex("#9b2226") // Deep Ruby Stained Glass
	ColorMahoganyBookcase = hex("#311811") // Deep Rich Mahogany
	ColorGoldFoilSpine    = hex("#e5b869") // Gilded Book Lettering
	ColorGrandfatherClock = hex("#442211") // Antique Grandfather Clock
	ColorBrassPendulum    = hex("#e0a96d") // Swinging Brass Pendulum
	ColorBankerGreen      = hex("#2d6a4f") // Emerald Banker's Lamp Glass

	// Coastal Ocean Lighthouse Study
	ColorOceanDeepBlue  = hex("#09111e") // Midnight Abyssal Ocean
	ColorOceanWaveMid   = hex("#1b2838") // Rolling Swell Water
	ColorOceanWaveCrest = hex("#e2eafc") // Frothing Whitecap Foam
	ColorLighthouseWhite= hex("#edf2f4") // Whitewashed Lighthouse Tower
	ColorLighthouseRed  = hex("#d90429") // Classic Red Spiral Band
	ColorLighthouseBeam = hex("#fff3b0") // Sweeping Amber Searchlight Beam
	ColorNauticalChart  = hex("#f4ebd9") // Antique Maritime Map
	ColorStoveCastIron  = hex("#1e2124") // Potbelly Cast Iron Stove

	// Hearthside Autumn Bakery
	ColorBreadCrustGold = hex("#c67d3b") // Artisan Sourdough Crust
	ColorBreadFlourDust = hex("#fefae0") // Dusted Baker's Flour
	ColorBrickOvenStone = hex("#584b42") // Rustic Oven Masonry
	ColorGarlicBraid    = hex("#e9ecef") // Braided Garlic Bulbs
	ColorCopperKettle   = hex("#ca6702") // Whistling Copper Kettle
)

func hex(s string) colorful.Color {
	c, _ := colorful.Hex(s)
	return c
}

