package art

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	RegisterScene(&CityCafeScene{})
	RegisterScene(&ForestCabinScene{})
	RegisterScene(&ClassicCottageScene{})
	RegisterScene(&ObservatoryScene{})
	RegisterScene(&ZenTeahouseScene{})
	RegisterScene(&CozyCampfireScene{})
	RegisterScene(&CozyAtticScene{})
	RegisterScene(&WinterCabinScene{})
	RegisterScene(&CozyGreenhouseScene{})
	RegisterScene(&MidnightLibraryScene{})
	RegisterScene(&LighthouseStudyScene{})
	RegisterScene(&PixelBakeryScene{})
	RegisterScene(&RetroLofiTrainScene{})
	RegisterScene(&ZenOnsenScene{})
	RegisterScene(&VinylLoungeScene{})
	RegisterScene(&SunkenHearthScene{})

	RegisterScene(&AnimatedGifScene{
		id:        SceneKirokazeRoom,
		name:      "Kirokaze Cozy Room 🎨",
		gifPath:   "assets/animations/kirokaze_room.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBalconyLofi,
		name:      "Lofi Balcony View 🏙️",
		gifPath:   "assets/animations/balcony_lofi.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRainSetup,
		name:      "Pixel Rain Setup 🌧️",
		gifPath:   "assets/animations/rain_setup.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneKirokazeCity,
		name:      "Kirokaze Retro City 🌃",
		gifPath:   "assets/animations/kirokaze_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneHomeSweetHome,
		name:      "Home Sweet Home 🏡",
		gifPath:   "assets/animations/home_sweet_home.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneWaneellaStreet,
		name:      "Waneella Pixel Street 🏮",
		gifPath:   "assets/animations/waneella_street.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSunsetCity,
		name:      "Sunset City Skyline 🌇",
		gifPath:   "assets/animations/sunset_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSnowCabin,
		name:      "Winter Cabin Snowfall ❄️",
		gifPath:   "assets/animations/snow_cabin.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberpunkCity,
		name:      "Cyberpunk Alley 🌆",
		gifPath:   "assets/animations/cyberpunk_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        ScenePixelTrain,
		name:      "Retro Pixel Train 🚂",
		gifPath:   "assets/animations/pixel_train.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneOceanWaves,
		name:      "Pixel Ocean Waves 🌊",
		gifPath:   "assets/animations/ocean_waves.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneStargazing,
		name:      "Stargazing Hill 🌌",
		gifPath:   "assets/animations/stargazing.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCoffeeShop,
		name:      "Pixel Coffee Shop ☕",
		gifPath:   "assets/animations/coffee_shop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAutumnPark,
		name:      "Autumn Amber Park 🍂",
		gifPath:   "assets/animations/autumn_park.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneMidnightOwl,
		name:      "Midnight Owl Vista 🦉",
		gifPath:   "assets/animations/midnight_owl.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRiverBridge,
		name:      "River Bridge Railway 🌉",
		gifPath:   "assets/animations/river_bridge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneNeonRunner,
		name:      "Neon Synthwave Runner 🏎️",
		gifPath:   "assets/animations/neon_runner.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneMagicSanctuary,
		name:      "Magic Alchemist Sanctuary 🔮",
		gifPath:   "assets/animations/magic_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberBot,
		name:      "Relaxing Cyber Bot 🤖",
		gifPath:   "assets/animations/cyber_bot.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSakuraShrine,
		name:      "Sakura Shrine Garden 🌸",
		gifPath:   "assets/animations/sakura_shrine.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCozyCampfire,
		name:      "Cozy Hearth Campfire 🔥",
		gifPath:   "assets/animations/cozy_campfire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSpaceStation,
		name:      "Orbital Space Station 🚀",
		gifPath:   "assets/animations/space_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRamenShop,
		name:      "Pixel Ramen Diner 🍜",
		gifPath:   "assets/animations/ramen_shop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneDeepSea,
		name:      "Deep Sea Aquarium 🪼",
		gifPath:   "assets/animations/deep_sea.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberArcade,
		name:      "Isometric Cyber Arcade 🕹️",
		gifPath:   "assets/animations/cyber_arcade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneNeonBistro,
		name:      "Neon Cyber Bistro 🏮",
		gifPath:   "assets/animations/neon_bistro.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSunsetStation,
		name:      "Sunset Train Platform 🚉",
		gifPath:   "assets/animations/sunset_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneMountainView,
		name:      "Cozy Mountain Peak 🏔️",
		gifPath:   "assets/animations/mountain_view.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneEnchantedForest,
		name:      "Enchanted Forest Waterfall 🌿",
		gifPath:   "assets/animations/enchanted_forest.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSkyIsland,
		name:      "Floating Sky Island ☁️",
		gifPath:   "assets/animations/sky_island.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBambooGrove,
		name:      "Zen Bamboo Grove 🎋",
		gifPath:   "assets/animations/bamboo_grove.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneVolcanoForge,
		name:      "Volcano Smithy Forge 🌋",
		gifPath:   "assets/animations/volcano_forge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCrystalCave,
		name:      "Glowing Crystal Cave 💎",
		gifPath:   "assets/animations/crystal_cave.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneDesertOasis,
		name:      "Moonlit Desert Oasis 🌙",
		gifPath:   "assets/animations/desert_oasis.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkLab,
		name:      "Steampunk Clockwork Lab ⚙️",
		gifPath:   "assets/animations/steampunk_lab.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAlienPlanet,
		name:      "Alien Planet Outpost 🛸",
		gifPath:   "assets/animations/alien_planet.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneHauntedManor,
		name:      "Haunted Pixel Manor 👻",
		gifPath:   "assets/animations/haunted_manor.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCloudTemple,
		name:      "Floating Cloud Temple ⛩️",
		gifPath:   "assets/animations/cloud_temple.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberBar,
		name:      "Neon Synthwave Cyber Bar 🍸",
		gifPath:   "assets/animations/cyber_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCoralReef,
		name:      "Submerged Coral Reef 🐠",
		gifPath:   "assets/animations/coral_reef.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneFrostPeak,
		name:      "Glacial Frost Peak 🏔️",
		gifPath:   "assets/animations/frost_peak.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberDojo,
		name:      "Cyberpunk Martial Dojo 🥋",
		gifPath:   "assets/animations/cyber_dojo.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneMushroomForest,
		name:      "Bioluminescent Mushroom Forest 🍄",
		gifPath:   "assets/animations/mushroom_forest.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroArcade,
		name:      "8-Bit Retro Pinball Arcade 🕹️",
		gifPath:   "assets/animations/retro_arcade_neon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        ScenePirateCove,
		name:      "Moonlight Pirate Ship Cove 🏴‍☠️",
		gifPath:   "assets/animations/pirate_cove.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCosmicObservatory,
		name:      "Deep Space Celestial Observatory 🔭",
		gifPath:   "assets/animations/cosmic_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberRooftop,
		name:      "Cyberpunk Neon Rooftop 🌆",
		gifPath:   "assets/animations/cyber_rooftop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCozyIgloo,
		name:      "Cozy Northern Igloo ❄️",
		gifPath:   "assets/animations/cozy_igloo.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAtlantisTemple,
		name:      "Atlantis Ocean Ruins 🔱",
		gifPath:   "assets/animations/atlantis_temple.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberMatrix,
		name:      "Cyber Matrix Terminal 💻",
		gifPath:   "assets/animations/cyber_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSunsetBeach,
		name:      "Sunset Beach Boardwalk 🏖️",
		gifPath:   "assets/animations/sunset_beach.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneFairytaleWindmill,
		name:      "Fairytale River Windmill 🌾",
		gifPath:   "assets/animations/fairytale_windmill.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberExpress,
		name:      "Neon Cyberpunk Express Train 🚄",
		gifPath:   "assets/animations/cyber_express.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumSwamp,
		name:      "Bioluminescent Willow Swamp 🐊",
		gifPath:   "assets/animations/biolum_swamp.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneClockworkTower,
		name:      "Clockwork Citadel Spire 🕰️",
		gifPath:   "assets/animations/clockwork_tower.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneStarlightCamp,
		name:      "Starlight Mountain Camp ⛺",
		gifPath:   "assets/animations/starlight_camp.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSpaceNebula,
		name:      "Deep Space Cosmic Nebula 🌌",
		gifPath:   "assets/animations/space_nebula.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberMarket,
		name:      "Pixel Cyber Market Alley 🏮",
		gifPath:   "assets/animations/cyber_market.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberDiner,
		name:      "Cyberpunk Neon Diner 🍔",
		gifPath:   "assets/animations/cyber_diner.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneVolcanoView,
		name:      "Pixel Volcanic Caldera 🌋",
		gifPath:   "assets/animations/volcano_view.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCrystalLagoon,
		name:      "Bioluminescent Crystal Lagoon 💎",
		gifPath:   "assets/animations/crystal_lagoon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkAirship,
		name:      "Steampunk Sky Airship Fleet ⚙️",
		gifPath:   "assets/animations/steampunk_airship.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCozyAttic,
		name:      "Cozy Rain-Slicked Attic Nook 🕯️",
		gifPath:   "assets/animations/cozy_attic.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroArcadeHall,
		name:      "Pixel Retro Neon Arcade Hall 🕹️",
		gifPath:   "assets/animations/retro_arcade_hall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberSubway,
		name:      "Cyberpunk Underground Subway 🚇",
		gifPath:   "assets/animations/cyber_subway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneFantasyCastle,
		name:      "Fantasy Castle Courtyard 🏰",
		gifPath:   "assets/animations/fantasy_castle.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyss,
		name:      "Bioluminescent Deep Abyss 🪼",
		gifPath:   "assets/animations/biolum_abyss.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveHighway,
		name:      "Retro Synthwave Highway 🏎️",
		gifPath:   "assets/animations/synthwave_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneWinterCafe,
		name:      "Cozy Winter Cafe ☕",
		gifPath:   "assets/animations/winter_cafe.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        ScenePixelLab,
		name:      "Isometric Pixel Laboratory 🧪",
		gifPath:   "assets/animations/pixel_lab.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberMatrixDeck,
		name:      "Cyberpunk Matrix Netrunner Deck 💾",
		gifPath:   "assets/animations/cyber_matrix_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneFairyMushroomHollow,
		name:      "Bioluminescent Fairy Mushroom Hollow 🍄",
		gifPath:   "assets/animations/fairy_mushroom_hollow.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroSpaceArcade,
		name:      "8-Bit Retro Space Arcade 👾",
		gifPath:   "assets/animations/retro_space_arcade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneMountainAuroraShack,
		name:      "Starlight Aurora Mountain Shack 🌌",
		gifPath:   "assets/animations/mountain_aurora_shack.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPort,
		name:      "Steampunk Celestial Sky Port ⚓",
		gifPath:   "assets/animations/steampunk_sky_port.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenRockGarden,
		name:      "Zen Moss Rock Garden 🪨",
		gifPath:   "assets/animations/zen_rock_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberBridge,
		name:      "Neon Cyberpunk Suspension Bridge 🌉",
		gifPath:   "assets/animations/cyber_bridge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCavern,
		name:      "Bioluminescent Underground Cavern 🪨",
		gifPath:   "assets/animations/biolum_cavern.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberRacer,
		name:      "8-Bit Retro Neon Grid Racer 🏎️",
		gifPath:   "assets/animations/cyber_racer.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneStarlightTreehouse,
		name:      "Starlight Ancient Oak Treehouse 🌳",
		gifPath:   "assets/animations/starlight_treehouse.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkClock,
		name:      "Steampunk Clockwork Cathedral 🕰️",
		gifPath:   "assets/animations/steampunk_clock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenTeahouse,
		name:      "Zen Bamboo Water Teahouse 🍵",
		gifPath:   "assets/animations/zen_teahouse.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberObservatory,
		name:      "Neon Cyberpunk Astronomical Observatory 🔭",
		gifPath:   "assets/animations/cyber_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumJungle,
		name:      "Bioluminescent Pixel Rainforest 🌴",
		gifPath:   "assets/animations/biolum_jungle.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveSunset,
		name:      "Retro Synthwave Sun & Grid Horizon 🌅",
		gifPath:   "assets/animations/synthwave_sunset.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneStarlightOasis,
		name:      "Starlight Desert Pyramid & Oasis 🐪",
		gifPath:   "assets/animations/starlight_oasis.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkWorkshop,
		name:      "Steampunk Automaton Workshop ⚙️",
		gifPath:   "assets/animations/steampunk_workshop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenCherryTemple,
		name:      "Zen Cherry Blossom Temple Courtyard 🌸",
		gifPath:   "assets/animations/zen_cherry_temple.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberDragon,
		name:      "Neon Cyber Dragon 🐉",
		gifPath:   "assets/animations/cyber_dragon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumRift,
		name:      "Bioluminescent Abyssal Rift 🌊",
		gifPath:   "assets/animations/biolum_rift.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkPlanetarium,
		name:      "Steampunk Celestial Planetarium 🪐",
		gifPath:   "assets/animations/steampunk_planetarium.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroLofiTrain,
		name:      "8-Bit Lofi Sunset Train 🚆",
		gifPath:   "assets/animations/retro_lofi_train.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBonsai,
		name:      "Zen Bonsai & Koi Pond 🪨",
		gifPath:   "assets/animations/zen_bonsai.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneWizardTower,
		name:      "Starlight Wizard Arcane Tower 🔮",
		gifPath:   "assets/animations/wizard_tower.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberDock,
		name:      "Neon Cyber Dock & Sky Traffic 🛸",
		gifPath:   "assets/animations/cyber_dock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumReefGate,
		name:      "Bioluminescent Reef Gate 🪸",
		gifPath:   "assets/animations/biolum_reef_gate.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkLocomotive,
		name:      "Steampunk Flying Locomotive 🚂",
		gifPath:   "assets/animations/steampunk_locomotive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        ScenePixelWinterCabin,
		name:      "Cozy Winter Forest Cabin 🌨️",
		gifPath:   "assets/animations/pixel_winter_cabin.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenWaterfallShrine,
		name:      "Zen Waterfall Torii Shrine ⛩️",
		gifPath:   "assets/animations/zen_waterfall_shrine.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralStargate,
		name:      "Astral Celestial Stargate 🌌",
		gifPath:   "assets/animations/astral_stargate.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberCloudCity,
		name:      "Neon Cyber Cloud City ☁️",
		gifPath:   "assets/animations/cyber_cloud_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCrystalDragon,
		name:      "Crystal Dragon Sanctuary 🐉",
		gifPath:   "assets/animations/crystal_dragon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSunset,
		name:      "Steampunk Airship Sunset 🌅",
		gifPath:   "assets/animations/steampunk_sunset.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveBeach,
		name:      "Synthwave Neon Beach 🏖️",
		gifPath:   "assets/animations/synthwave_beach.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenMoonShrine,
		name:      "Zen Moonlit Shrine 🌙",
		gifPath:   "assets/animations/zen_moon_shrine.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralClockwork,
		name:      "Astral Clockwork Universe 🌌",
		gifPath:   "assets/animations/astral_clockwork.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberPenthouse,
		name:      "Neon Cyber Penthouse 🌃",
		gifPath:   "assets/animations/cyber_penthouse.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumUnderwaterCity,
		name:      "Bioluminescent Sunken City 🧜",
		gifPath:   "assets/animations/biolum_underwater_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDock,
		name:      "Steampunk Sky Harbor Dock ⚓",
		gifPath:   "assets/animations/steampunk_sky_dock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroSynthwaveCity,
		name:      "Synthwave Cyber City Sunset 🌇",
		gifPath:   "assets/animations/retro_synthwave_city.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenAutumnTorii,
		name:      "Zen Autumn Torii Pathway 🍁",
		gifPath:   "assets/animations/zen_autumn_torii.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralNebulaGate,
		name:      "Astral Cosmic Nebula Portal 🔮",
		gifPath:   "assets/animations/astral_nebula_gate.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberGarage,
		name:      "Neon Cyber Hoverbike Garage 🏍️",
		gifPath:   "assets/animations/cyber_garage.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumEmeraldStream,
		name:      "Bioluminescent Emerald Crystal Stream 💎",
		gifPath:   "assets/animations/biolum_emerald_stream.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkAirshipDock,
		name:      "Steampunk Sky Airship Terminal ⚓",
		gifPath:   "assets/animations/steampunk_airship_dock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveRooftopPool,
		name:      "Retro Synthwave Rooftop Pool 🌴",
		gifPath:   "assets/animations/synthwave_rooftop_pool.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooWater,
		name:      "Zen Moonlit Bamboo Water Spout 🎋",
		gifPath:   "assets/animations/zen_bamboo_water.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralPrismPortal,
		name:      "Astral Celestial Prism Portal 🔮",
		gifPath:   "assets/animations/astral_prism_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberArcadeNeonBar,
		name:      "Neon Cyberpunk Synth Arcade Bar 🕹️",
		gifPath:   "assets/animations/cyber_arcade_neon_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalGrotto,
		name:      "Bioluminescent Crystal Grotto 💎",
		gifPath:   "assets/animations/biolum_crystal_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyStation,
		name:      "Steampunk Sky Airship Station ⚓",
		gifPath:   "assets/animations/steampunk_sky_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroSynthwaveHighwayNight,
		name:      "Retro Synthwave Highway Night Drive 🏎️",
		gifPath:   "assets/animations/retro_synthwave_highway_night.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooWaterfall,
		name:      "Zen Bamboo Waterfall Sanctum 🎋",
		gifPath:   "assets/animations/zen_bamboo_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacObservatory,
		name:      "Astral Celestial Zodiac Observatory 🔮",
		gifPath:   "assets/animations/astral_zodiac_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberRamenAlley,
		name:      "Neon Cyberpunk Ramen Alley 🍜",
		gifPath:   "assets/animations/cyber_ramen_alley.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalWaterfall,
		name:      "Bioluminescent Crystal Waterfall 💎",
		gifPath:   "assets/animations/biolum_crystal_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkClockworkCitadel,
		name:      "Steampunk Floating Citadel 🏰",
		gifPath:   "assets/animations/steampunk_clockwork_citadel.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroSynthwaveArcade,
		name:      "Retro Synthwave Arcade Night 🕹️",
		gifPath:   "assets/animations/retro_synthwave_arcade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooToriiGarden,
		name:      "Zen Moonlit Bamboo Torii ⛩️",
		gifPath:   "assets/animations/zen_bamboo_torii_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralPrismObservatory,
		name:      "Astral Celestial Prism Tower 🔮",
		gifPath:   "assets/animations/astral_prism_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberSpaceportLounge,
		name:      "Neon Cyberpunk Spaceport Lounge 🚀",
		gifPath:   "assets/animations/cyber_spaceport.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalCitadel,
		name:      "Bioluminescent Crystal Citadel 🏰",
		gifPath:   "assets/animations/biolum_citadel.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkFlyingGondola,
		name:      "Steampunk Airship Sky Gondola 🎈",
		gifPath:   "assets/animations/steampunk_gondola.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneRetroSynthwaveArcadeBar,
		name:      "8-Bit Retro Synthwave Arcade Bar 🍸",
		gifPath:   "assets/animations/synthwave_arcade_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonPagoda,
		name:      "Zen Moonlit Bamboo Pagoda ⛩️",
		gifPath:   "assets/animations/zen_bamboo_pagoda.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacPortalSpire,
		name:      "Astral Celestial Zodiac Portal Spire 🔮",
		gifPath:   "assets/animations/astral_zodiac_spire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberDockLounge,
		name:      "Neon Cyberpunk Dockside Lounge 🍹",
		gifPath:   "assets/animations/cyber_dock_lounge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalSanctuary,
		name:      "Bioluminescent Crystal Sanctuary 🔮",
		gifPath:   "assets/animations/biolum_crystal_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockBar,
		name:      "Steampunk Sky Dock Tavern 🍺",
		gifPath:   "assets/animations/steampunk_sky_dock_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridArcade,
		name:      "8-Bit Retro Synthwave Cyber Arcade 🎮",
		gifPath:   "assets/animations/synthwave_grid_arcade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonTemple,
		name:      "Zen Moonlit Bamboo Temple Shrine ⛩️",
		gifPath:   "assets/animations/zen_bamboo_moon_temple.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacPortalRing,
		name:      "Astral Celestial Zodiac Ring Portal 🌌",
		gifPath:   "assets/animations/astral_zodiac_portal_ring.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonSkyLounge,
		name:      "Neon Cyberpunk Sky Lounge 🏙️",
		gifPath:   "assets/animations/cyber_neon_sky_lounge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalLagoonSpire,
		name:      "Bioluminescent Crystal Lagoon Spire 💎",
		gifPath:   "assets/animations/biolum_crystal_lagoon_spire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkAirshipPortBar,
		name:      "Steampunk Airship Port Bar ⚙️",
		gifPath:   "assets/animations/steampunk_airship_port_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveNeonHighwayNight,
		name:      "8-Bit Retro Synthwave Neon Highway Night 🏎️",
		gifPath:   "assets/animations/synthwave_neon_highway_night.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonRiver,
		name:      "Zen Moonlit Bamboo River 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_river.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChamber,
		name:      "Astral Celestial Zodiac Chamber 🌌",
		gifPath:   "assets/animations/astral_zodiac_chamber.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberArcadeMatrixLounge,
		name:      "Neon Cyberpunk Matrix Arcade Lounge 👾",
		gifPath:   "assets/animations/cyber_arcade_matrix_lounge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumFairyGrotto,
		name:      "Bioluminescent Fairy Grotto 🧚",
		gifPath:   "assets/animations/biolum_fairy_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockHaven,
		name:      "Steampunk Sky Dock Haven ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridSunset,
		name:      "8-Bit Retro Synthwave Grid Sunset 🌅",
		gifPath:   "assets/animations/synthwave_grid_sunset.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrine,
		name:      "Zen Moonlit Bamboo Shrine 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacPortalChamber,
		name:      "Astral Celestial Zodiac Portal Chamber 🌌",
		gifPath:   "assets/animations/astral_zodiac_portal_chamber.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberRainRooftop,
		name:      "Neon Cyber Rain Rooftop 🌧️",
		gifPath:   "assets/animations/cyber_rain_rooftop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumUnderwaterTrench,
		name:      "Bioluminescent Trench Abyss 🪼",
		gifPath:   "assets/animations/biolum_underwater_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyAstronomy,
		name:      "Steampunk Sky Planetarium 🌌",
		gifPath:   "assets/animations/steampunk_sky_astronomy.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonHighway,
		name:      "Synthwave Grid Sunset Drive 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonBridge,
		name:      "Zen Moonlit Bamboo Bridge 🌉",
		gifPath:   "assets/animations/zen_bamboo_moon_bridge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacStarPortal,
		name:      "Astral Zodiac Star Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_star_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberRooftopBar,
		name:      "Neon Cyberpunk Skyline Rooftop Bar 🍸",
		gifPath:   "assets/animations/cyber_rooftop_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCoralGrotto,
		name:      "Bioluminescent Sunken Coral Grotto 🪸",
		gifPath:   "assets/animations/biolum_coral_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockLounge,
		name:      "Steampunk Sky Port Panorama ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_lounge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonBeach,
		name:      "8-Bit Retro Synthwave Neon Coastline 🌅",
		gifPath:   "assets/animations/synthwave_grid_neon_beach.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonFountain,
		name:      "Zen Moonlit Bamboo Water Cascade 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_fountain.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacStarChamber,
		name:      "Astral Celestial Zodiac Observatory Spire 🔮",
		gifPath:   "assets/animations/astral_zodiac_star_chamber.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonBistroAlley,
		name:      "Neon Cyberpunk Alley Bistro 🍸",
		gifPath:   "assets/animations/cyber_neon_bistro_alley.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalLagoonGrotto,
		name:      "Bioluminescent Crystal Lagoon Grotto 💎",
		gifPath:   "assets/animations/biolum_crystal_lagoon_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyObservatory,
		name:      "Steampunk Sky Harbor Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridSunsetHighway,
		name:      "8-Bit Retro Synthwave Sunset Highway 🏎️",
		gifPath:   "assets/animations/synthwave_grid_sunset_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonSpring,
		name:      "Zen Moonlit Bamboo Thermal Spring ♨️",
		gifPath:   "assets/animations/zen_bamboo_moon_spring.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacNexus,
		name:      "Astral Celestial Zodiac Nexus 🔮",
		gifPath:   "assets/animations/astral_zodiac_nexus.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonNoodleBar,
		name:      "Neon Cyberpunk Noodle Bar 🍜",
		gifPath:   "assets/animations/cyber_neon_noodle_bar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyssalCavernSpire,
		name:      "Bioluminescent Abyssal Cavern Spire 🪼",
		gifPath:   "assets/animations/biolum_abyssal_cavern_spire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyFortress,
		name:      "Steampunk Sky Fortress Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_fortress.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveCyberCityDrive,
		name:      "8-Bit Retro Synthwave Cyber City Drive 🏎️",
		gifPath:   "assets/animations/synthwave_cyber_city_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMountainSanctuary,
		name:      "Zen Moonlit Bamboo Mountain Sanctuary 🎋",
		gifPath:   "assets/animations/zen_bamboo_mountain_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacPrismMatrix,
		name:      "Astral Celestial Zodiac Prism Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_prism_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonArcadeAlley,
		name:      "Neon Cyberpunk Arcade Alley 🕹️",
		gifPath:   "assets/animations/cyber_neon_arcade_alley.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumDeepSeaTrench,
		name:      "Bioluminescent Deep Sea Abyss 🐋",
		gifPath:   "assets/animations/biolum_deep_sea_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkClockworkFortress,
		name:      "Steampunk Sky Clockwork Fortress ⚙️",
		gifPath:   "assets/animations/steampunk_clockwork_fortress.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridCosmoDrive,
		name:      "8-Bit Retro Synthwave Cosmo Drive 🏎️",
		gifPath:   "assets/animations/synthwave_grid_cosmo_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonWaterfall,
		name:      "Zen Moonlit Bamboo Water Spout 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChamberMatrix,
		name:      "Astral Celestial Zodiac Chamber Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_chamber_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixDiner,
		name:      "Neon Cyberpunk Matrix Diner 🍜",
		gifPath:   "assets/animations/cyber_neon_matrix_diner.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCoralAbyssGate,
		name:      "Bioluminescent Coral Abyss Gate 🪸",
		gifPath:   "assets/animations/biolum_coral_abyss_gate.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockCitadel,
		name:      "Steampunk Sky Dock Citadel ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_citadel.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonSkyline,
		name:      "8-Bit Retro Synthwave Neon Skyline 🌇",
		gifPath:   "assets/animations/synthwave_grid_neon_skyline.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooToriiWaterfall,
		name:      "Zen Moonlit Torii Waterfall ⛩️",
		gifPath:   "assets/animations/zen_bamboo_torii_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacPrismPortal,
		name:      "Astral Celestial Zodiac Prism Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_prism_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonCyberCasino,
		name:      "Neon Cyberpunk Cyber Casino 🎰",
		gifPath:   "assets/animations/cyber_neon_cyber_casino.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyssalTideTrench,
		name:      "Bioluminescent Abyssal Tide Trench 🌊",
		gifPath:   "assets/animations/biolum_abyssal_tide_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkClockworkAirship,
		name:      "Steampunk Clockwork Airship Dock ⚙️",
		gifPath:   "assets/animations/steampunk_clockwork_airship.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridSpaceStation,
		name:      "8-Bit Retro Synthwave Space Station 🌌",
		gifPath:   "assets/animations/synthwave_grid_space_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooPagodaSpring,
		name:      "Zen Moonlit Pagoda Hot Spring ⛩️",
		gifPath:   "assets/animations/zen_bamboo_pagoda_spring.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosVault,
		name:      "Astral Celestial Zodiac Chronos Vault 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonHoloArcade,
		name:      "Neon Cyberpunk Holo Arcade Lounge 🕹️",
		gifPath:   "assets/animations/cyber_neon_holo_arcade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalLeviathanReef,
		name:      "Bioluminescent Crystal Leviathan Reef 🐋",
		gifPath:   "assets/animations/biolum_crystal_leviathan_reef.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkObservatory,
		name:      "Steampunk Sky Clockwork Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCyberHighway,
		name:      "8-Bit Retro Synthwave Neon Cyber Highway 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_cyber_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooLotusPond,
		name:      "Zen Moonlit Bamboo Lotus Pond 🎋",
		gifPath:   "assets/animations/zen_bamboo_lotus_pond.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacAetherVault,
		name:      "Astral Celestial Zodiac Aether Vault 🔮",
		gifPath:   "assets/animations/astral_zodiac_aether_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixClub,
		name:      "Neon Cyberpunk Matrix Nightclub 🍸",
		gifPath:   "assets/animations/cyber_neon_matrix_club.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCoralCrystalTrench,
		name:      "Bioluminescent Coral Crystal Trench 🪸",
		gifPath:   "assets/animations/biolum_coral_crystal_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockWorkshop,
		name:      "Steampunk Sky Dock Workshop ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_workshop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonShallows,
		name:      "8-Bit Retro Synthwave Neon Shallows 🌅",
		gifPath:   "assets/animations/synthwave_grid_neon_shallows.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonLanternPond,
		name:      "Zen Moonlit Bamboo Lantern Pond 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_lantern_pond.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacSolarisVault,
		name:      "Astral Celestial Zodiac Solaris Vault 🔮",
		gifPath:   "assets/animations/astral_zodiac_solaris_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonHackerDen,
		name:      "Neon Cyberpunk Netrunner Den 💻",
		gifPath:   "assets/animations/cyber_neon_hacker_den.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalGlacierTrench,
		name:      "Bioluminescent Crystal Glacier Trench 🧊",
		gifPath:   "assets/animations/biolum_crystal_glacier_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortHangar,
		name:      "Steampunk Sky Port Hangar ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_hangar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCanyon,
		name:      "8-Bit Retro Synthwave Neon Canyon 🌄",
		gifPath:   "assets/animations/synthwave_grid_neon_canyon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrinePond,
		name:      "Zen Moonlit Bamboo Shrine Pond 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_pond.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacLunarisVault,
		name:      "Astral Celestial Zodiac Lunaris Vault 🔮",
		gifPath:   "assets/animations/astral_zodiac_lunaris_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonSkyDockPlaza,
		name:      "Cyberpunk Neon Sky Dock Plaza 🏙️",
		gifPath:   "assets/animations/cyber_neon_sky_dock_plaza.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyssalTrenchSpire,
		name:      "Bioluminescent Abyssal Trench Spire 🌊",
		gifPath:   "assets/animations/biolum_abyssal_trench_spire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortObservatory,
		name:      "Steampunk Sky Port Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonMetropolis,
		name:      "8-Bit Retro Synthwave Neon Metropolis 🌄",
		gifPath:   "assets/animations/synthwave_grid_neon_metropolis.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineWaterfall,
		name:      "Zen Moonlit Bamboo Shrine Waterfall 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosPortal,
		name:      "Astral Celestial Zodiac Chronos Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixTower,
		name:      "Cyberpunk Neon Matrix Tower 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_tower.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyssalTrenchHaven,
		name:      "Bioluminescent Abyssal Trench Haven 🌊",
		gifPath:   "assets/animations/biolum_abyssal_trench_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortFactory,
		name:      "Steampunk Sky Port Factory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_factory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonHighwayDrive,
		name:      "8-Bit Retro Synthwave Neon Highway Drive 🌄",
		gifPath:   "assets/animations/synthwave_grid_neon_highway_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineGarden,
		name:      "Zen Moonlit Bamboo Shrine Garden 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacLunarisPortal,
		name:      "Astral Celestial Zodiac Lunaris Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_lunaris_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonSkyTerminal,
		name:      "Neon Cyberpunk Sky Terminal 🏙️",
		gifPath:   "assets/animations/cyber_neon_sky_terminal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumAbyssalCrystalRift,
		name:      "Bioluminescent Abyssal Crystal Rift 💎",
		gifPath:   "assets/animations/biolum_abyssal_crystal_rift.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockFactory,
		name:      "Steampunk Sky Dock Factory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_factory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonExpressway,
		name:      "8-Bit Retro Synthwave Neon Expressway 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_expressway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonPavilion,
		name:      "Zen Moonlit Bamboo Pavilion 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_pavilion.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacAetherPortal,
		name:      "Astral Celestial Zodiac Aether Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_aether_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonSkyStation,
		name:      "Neon Cyberpunk Sky Station 🏙️",
		gifPath:   "assets/animations/cyber_neon_sky_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalSanctuaryGrotto,
		name:      "Bioluminescent Crystal Sanctuary Grotto 💎",
		gifPath:   "assets/animations/biolum_crystal_sanctuary_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyDockObservatory,
		name:      "Steampunk Sky Port Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_dock_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridCyberHighway,
		name:      "8-Bit Retro Synthwave Cyber Highway 🏎️",
		gifPath:   "assets/animations/synthwave_grid_cyber_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineFountain,
		name:      "Zen Moonlit Bamboo Shrine Fountain 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_fountain.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacSolarisPortal,
		name:      "Astral Celestial Zodiac Solaris Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_solaris_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonSkyBalcony,
		name:      "Neon Cyberpunk Sky Balcony 🌇",
		gifPath:   "assets/animations/cyber_neon_sky_balcony.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalCoralLagoon,
		name:      "Bioluminescent Coral Lagoon 🪸",
		gifPath:   "assets/animations/biolum_crystal_coral_lagoon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortTerminal,
		name:      "Steampunk Sky Port Terminal ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_terminal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridSkylineDrive,
		name:      "8-Bit Retro Synthwave Skyline Drive 🌅",
		gifPath:   "assets/animations/synthwave_grid_skyline_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValley,
		name:      "Zen Moonlit Bamboo Shrine Valley 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacLunarisMatrix,
		name:      "Astral Celestial Zodiac Lunaris Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_lunaris_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixSkyline,
		name:      "Neon Cyberpunk Matrix Skyline 🌇",
		gifPath:   "assets/animations/cyber_neon_matrix_skyline.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalReef,
		name:      "Bioluminescent Abyssal Reef 🪸",
		gifPath:   "assets/animations/biolum_crystal_abyssal_reef.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortObservatoryDeck,
		name:      "Steampunk Sky Port Observatory Deck ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_observatory_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalDrive,
		name:      "8-Bit Retro Synthwave Coastal Drive 🌅",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineWaterGarden,
		name:      "Zen Moonlit Bamboo Water Garden 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_water_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosMatrix,
		name:      "Astral Celestial Zodiac Chronos Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixSanctuary,
		name:      "Neon Cyberpunk Matrix Sanctuary 🌆",
		gifPath:   "assets/animations/cyber_neon_matrix_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTrenchGrotto,
		name:      "Bioluminescent Abyssal Trench Grotto 🌌",
		gifPath:   "assets/animations/biolum_crystal_abyssal_trench_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyPortClockworkDock,
		name:      "Steampunk Sky Port Clockwork Dock ⚙️",
		gifPath:   "assets/animations/steampunk_sky_port_clockwork_dock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonSunsetExpressway,
		name:      "8-Bit Retro Synthwave Sunset Expressway 🌅",
		gifPath:   "assets/animations/synthwave_grid_neon_sunset_expressway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineRockGarden,
		name:      "Zen Moonlit Bamboo Shrine Rock Garden 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_rock_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacAetherMatrix,
		name:      "Astral Celestial Zodiac Aether Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_aether_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixPortal,
		name:      "Neon Cyberpunk Matrix Portal 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalDeepTrench,
		name:      "Bioluminescent Crystal Deep Trench 🌊",
		gifPath:   "assets/animations/biolum_crystal_deep_trench.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkFortress,
		name:      "Steampunk Sky Port Clockwork Fortress ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_fortress.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonExpress,
		name:      "Synthwave Grid Neon Cyber Express 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_express.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonValleyStream,
		name:      "Zen Moonlit Bamboo Valley Stream 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_valley_stream.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosNexus,
		name:      "Astral Celestial Zodiac Chronos Nexus 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_nexus.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixTowerLounge,
		name:      "Neon Cyberpunk Matrix Tower Lounge 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_tower_lounge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideGrotto,
		name:      "Bioluminescent Crystal Abyssal Tide Grotto 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_grotto.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkHarbor,
		name:      "Steampunk Sky Port Clockwork Harbor ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_harbor.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalHighway,
		name:      "Synthwave Grid Neon Coastal Highway 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_highway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonValleyWaterfall,
		name:      "Zen Moonlit Bamboo Valley Waterfall 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_valley_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosVaultMatrix,
		name:      "Astral Celestial Zodiac Chronos Vault Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_vault_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixNexusDeck,
		name:      "Neon Cyberpunk Matrix Nexus Deck 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_nexus_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalReefVault,
		name:      "Bioluminescent Crystal Abyssal Reef Vault 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_reef_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkSkyport,
		name:      "Steampunk Sky Port Clockwork Skyport ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_skyport.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalExpressway,
		name:      "Synthwave Grid Neon Coastal Expressway 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_expressway.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineStream,
		name:      "Zen Moonlit Bamboo Shrine Stream 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_stream.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherMatrix,
		name:      "Astral Celestial Zodiac Chronos Aether Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixAetherDeck,
		name:      "Neon Cyberpunk Matrix Aether Deck 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_aether_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideVault,
		name:      "Bioluminescent Crystal Abyssal Tide Vault 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelDock,
		name:      "Steampunk Sky Port Citadel Dock ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_dock.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalExpress,
		name:      "Synthwave Grid Neon Coastal Express 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_express.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineCascade,
		name:      "Zen Moonlit Bamboo Shrine Cascade 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_cascade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherVault,
		name:      "Astral Celestial Zodiac Chronos Aether Vault 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_vault.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixSanctuaryDeck,
		name:      "Neon Cyberpunk Matrix Sanctuary Deck 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_sanctuary_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideNexus,
		name:      "Bioluminescent Crystal Abyssal Tide Nexus 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_nexus.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelSkyport,
		name:      "Steampunk Sky Port Citadel Skyport ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_skyport.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalHighwayExpress,
		name:      "Synthwave Grid Neon Coastal Highway Express 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_highway_express.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineCascadeFall,
		name:      "Zen Moonlit Bamboo Shrine Cascade Fall 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_cascade_fall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherVaultMatrix,
		name:      "Astral Celestial Zodiac Chronos Aether Vault Matrix 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_vault_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixOrbitDeck,
		name:      "Neon Cyberpunk Matrix Orbit Deck 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_orbit_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideSanctuary,
		name:      "Bioluminescent Crystal Abyssal Tide Sanctuary 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelHaven,
		name:      "Steampunk Sky Port Citadel Haven ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalBoulevard,
		name:      "Synthwave Grid Neon Coastal Boulevard 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_boulevard.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyWaterfall,
		name:      "Zen Moonlit Bamboo Shrine Valley Waterfall 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_waterfall.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortal,
		name:      "Astral Celestial Zodiac Chronos Aether Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixSkyDeck,
		name:      "Neon Cyberpunk Matrix Sky Deck 🏙️",
		gifPath:   "assets/animations/cyber_neon_matrix_sky_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTidePortal,
		name:      "Bioluminescent Crystal Abyssal Tide Portal 🌊",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_portal.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelPlaza,
		name:      "Steampunk Sky Port Citadel Plaza ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_plaza.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalAvenue,
		name:      "Synthwave Grid Neon Coastal Avenue 🌴",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_avenue.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyStream,
		name:      "Zen Moonlit Bamboo Shrine Valley Stream 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_stream.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherNexus,
		name:      "Astral Celestial Zodiac Chronos Aether Nexus 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_nexus.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixStarlightDeck,
		name:      "Cyberpunk Neon Matrix Starlight Deck 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_starlight_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideCitadel,
		name:      "Bioluminescent Abyssal Tide Citadel 🪼",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_citadel.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelObservatory,
		name:      "Steampunk Sky Port Citadel Observatory ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_observatory.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalExpressDrive,
		name:      "Synthwave Grid Neon Coastal Drive 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_express_drive.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyCascade,
		name:      "Zen Moonlit Bamboo Shrine Valley Cascade 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_cascade.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalMatrix,
		name:      "Astral Celestial Zodiac Chronos Aether Portal 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_matrix.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixStellarDeck,
		name:      "Cyberpunk Neon Matrix Stellar Deck 🚀",
		gifPath:   "assets/animations/cyber_neon_matrix_stellar_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideHaven,
		name:      "Bioluminescent Abyssal Tide Haven 🐚",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelWorkshop,
		name:      "Steampunk Sky Port Citadel Workshop ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_workshop.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalCruiser,
		name:      "Synthwave Grid Neon Coastal Cruiser 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_cruiser.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyPool,
		name:      "Zen Moonlit Bamboo Shrine Valley Pool 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_pool.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalNexus,
		name:      "Astral Celestial Zodiac Chronos Aether Nexus 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_nexus.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixGalaxyDeck,
		name:      "Neon Cyberpunk Matrix Galaxy Deck 🌌",
		gifPath:   "assets/animations/cyber_neon_matrix_galaxy_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideSanctum,
		name:      "Bioluminescent Abyssal Tide Sanctum 💎",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_sanctum.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelAirport,
		name:      "Steampunk Sky Port Citadel Airport ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_airport.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalRacer,
		name:      "Synthwave Grid Neon Coastal Racer 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_racer.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyFountain,
		name:      "Zen Moonlit Bamboo Shrine Valley Fountain 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_fountain.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalSanctum,
		name:      "Astral Celestial Zodiac Chronos Aether Sanctum 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_sanctum.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixCosmosDeck,
		name:      "Neon Cyberpunk Matrix Cosmos Deck 🌌",
		gifPath:   "assets/animations/cyber_neon_matrix_cosmos_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideTemple,
		name:      "Bioluminescent Abyssal Tide Temple 🛕",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_temple.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelStation,
		name:      "Steampunk Sky Port Citadel Station ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_station.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalFlier,
		name:      "Synthwave Grid Neon Coastal Flier 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_flier.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyGarden,
		name:      "Zen Moonlit Bamboo Shrine Valley Garden 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_garden.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalRealm,
		name:      "Astral Celestial Zodiac Chronos Aether Realm 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_realm.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixHorizonDeck,
		name:      "Neon Cyberpunk Matrix Horizon Deck 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_horizon_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideShrine,
		name:      "Bioluminescent Abyssal Tide Shrine ⛩️",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_shrine.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelDeck,
		name:      "Steampunk Sky Port Citadel Deck ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalGlide,
		name:      "Synthwave Grid Neon Coastal Glide 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_glide.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleySanctuary,
		name:      "Zen Moonlit Bamboo Shrine Valley Sanctuary 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_sanctuary.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalHaven,
		name:      "Astral Celestial Zodiac Chronos Aether Haven 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixZenithDeck,
		name:      "Neon Cyberpunk Matrix Zenith Deck 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_zenith_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideCanyon,
		name:      "Bioluminescent Abyssal Tide Canyon 🪸",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_canyon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelBastion,
		name:      "Steampunk Sky Clockwork Citadel Bastion ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_bastion.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalSprint,
		name:      "Synthwave Grid Neon Coastal Sprint 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_sprint.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyMeadow,
		name:      "Zen Moonlit Bamboo Shrine Valley Meadow 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_meadow.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalCelestial,
		name:      "Astral Celestial Zodiac Chronos Aether Portal Celestial 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_celestial.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixApexDeck,
		name:      "Neon Cyberpunk Matrix Apex Deck 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_apex_deck.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideAbyss,
		name:      "Bioluminescent Abyssal Tide Abyss 🪸",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_abyss.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelSpire,
		name:      "Steampunk Sky Clockwork Citadel Spire ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_spire.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalPulse,
		name:      "Synthwave Grid Neon Coastal Pulse 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_pulse.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyGrove,
		name:      "Zen Moonlit Bamboo Shrine Valley Grove 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_grove.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalApex,
		name:      "Astral Celestial Zodiac Chronos Aether Portal Apex 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_apex.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixZenithHorizon,
		name:      "Neon Cyberpunk Matrix Zenith Horizon 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_zenith_horizon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideDepths,
		name:      "Bioluminescent Abyssal Tide Depths 🪸",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_depths.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelTower,
		name:      "Steampunk Sky Clockwork Citadel Tower ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_tower.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalDrift,
		name:      "Synthwave Grid Neon Coastal Drift 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_drift.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleySanctum,
		name:      "Zen Moonlit Bamboo Shrine Valley Sanctum 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_sanctum.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalZenith,
		name:      "Astral Celestial Zodiac Chronos Aether Portal Zenith 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_zenith.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneCyberNeonMatrixStellarHorizon,
		name:      "Neon Cyberpunk Matrix Stellar Horizon 🌃",
		gifPath:   "assets/animations/cyber_neon_matrix_stellar_horizon.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneBiolumCrystalAbyssalTideChasm,
		name:      "Bioluminescent Abyssal Tide Chasm 🪸",
		gifPath:   "assets/animations/biolum_crystal_abyssal_tide_chasm.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSteampunkSkyClockworkCitadelBridge,
		name:      "Steampunk Sky Clockwork Citadel Bridge ⚙️",
		gifPath:   "assets/animations/steampunk_sky_clockwork_citadel_bridge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneSynthwaveGridNeonCoastalSurge,
		name:      "Synthwave Grid Neon Coastal Surge 🏎️",
		gifPath:   "assets/animations/synthwave_grid_neon_coastal_surge.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneZenBambooMoonShrineValleyHaven,
		name:      "Zen Moonlit Bamboo Shrine Valley Haven 🎋",
		gifPath:   "assets/animations/zen_bamboo_moon_shrine_valley_haven.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
	RegisterScene(&AnimatedGifScene{
		id:        SceneAstralZodiacChronosAetherPortalStellar,
		name:      "Astral Celestial Zodiac Chronos Aether Portal Stellar 🔮",
		gifPath:   "assets/animations/astral_zodiac_chronos_aether_portal_stellar.gif",
		winBounds: Rect{X: 10, Y: 10, W: 140, H: 60},
	})
}

// ---------------------------------------------------------
// AnimatedGifScene (Generic Decoded GIF Animation Scene)
// ---------------------------------------------------------
type AnimatedGifScene struct {
	id        SceneType
	name      string
	gifPath   string
	winBounds Rect
	showBuddy bool
	buddyX    int
	buddyY    int
}

func (s *AnimatedGifScene) ID() SceneType       { return s.id }
func (s *AnimatedGifScene) Name() string     { return s.name }
func (s *AnimatedGifScene) WindowBounds() Rect { return s.winBounds }

func (s *AnimatedGifScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	anim, err := LoadAnimatedGif(s.gifPath, c.Width, c.Height)
	if err == nil && anim != nil {
		anim.DrawFrame(c, frameCount)
	} else {
		c.Clear(PicoNavy)
	}
}

func (s *AnimatedGifScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	if s.showBuddy && buddy != nil {
		buddy.Draw(c, s.buddyX, s.buddyY, frameCount)
	}
}

func (s *AnimatedGifScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.50)
	}
}

// ---------------------------------------------------------
// 1. CityCafeScene (16-Bit Rainy City Cafe)
// ---------------------------------------------------------
type CityCafeScene struct{}

func (s *CityCafeScene) ID() SceneType       { return SceneRainyCity }
func (s *CityCafeScene) Name() string     { return "Rainy City Lofi Cafe 🏙️" }
func (s *CityCafeScene) WindowBounds() Rect { return Rect{X: 8, Y: 4, W: 58, H: 30} }

func (s *CityCafeScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = ColorSkyDayTop, ColorSkyDayBot
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Distant skyline buildings & atmospheric haze
	bldgColors := []colorful.Color{hex("#141824"), hex("#1a2238"), hex("#0d111a")}
	windowColors := []colorful.Color{ColorWarmGlow, ColorLampBrass, hex("#74c0fc"), hex("#ffd166")}

	for b := 0; b < 8; b++ {
		bx := wx + b*7 + 1
		bw := 6
		bh := 12 + ((b*13+5)%15)
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		// Warm illuminated apartment windows
		if tod == TimeNight || tod == TimeSunset || tod == TimeDawn {
			for wyPos := wy + wh - bh + 2; wyPos < wy+wh-2; wyPos += 3 {
				for wxPos := bx + 1; wxPos < bx+bw-1; wxPos += 2 {
					if (wxPos*3+wyPos*7+b)%4 != 0 {
						wCol := windowColors[(wxPos+wyPos+int(frameCount/30))%len(windowColors)]
						c.SetPixel(wxPos, wyPos, wCol)
					}
				}
			}
		}
	}

	// Glowing distant neon rooftop cafe sign
	if tod == TimeNight || tod == TimeSunset {
		neonCol := hex("#ff7597")
		if (frameCount/20)%2 == 0 {
			neonCol = hex("#ffd166")
		}
		c.FillRect(wx+4, wy+4, 14, 4, hex("#101424"))
		c.DrawGlyph(wx+5, (wy+5)/2, 'C', neonCol)
		c.DrawGlyph(wx+7, (wy+5)/2, 'A', neonCol)
		c.DrawGlyph(wx+9, (wy+5)/2, 'F', neonCol)
		c.DrawGlyph(wx+11, (wy+5)/2, 'E', neonCol)
	}
}

func (s *CityCafeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Cafe Interior walls
	c.DrawVerticalGradient(0, 0, wx-1, 38, ColorWoodDark, ColorWallWarm)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 38, ColorWoodDark, ColorWallWarm)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Sleek steel window frame with condensation mullions
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)

	// Window sill with potted succulent
	c.FillRect(wx, wy+wh-1, ww, 2, ColorWoodMedium)
	renderPottedSucculent(c, wx+4, wy+wh-4)

	// Polished Mahogany Cafe Bar Counter (Rows 36 to c.Height-1)
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	c.FillRect(0, 36, c.Width, 1, ColorWoodHighlight)
	c.FillRect(0, 38, c.Width, 1, ColorWoodPlankSep)

	// Espresso machine on left counter
	c.FillRect(2, 26, 8, 10, ColorWoodDark)
	c.FillRect(3, 28, 6, 8, hex("#343a40"))
	c.FillRect(4, 24, 4, 2, ColorLampBrass) // brass boiler dome
	c.SetPixel(3, 33, ColorWarmGlow)        // power lamp

	// Focus buddy studying at cafe bar
	if buddy != nil {
		buddy.Draw(c, 18, 20, frameCount)
	}

	// Glowing modern laptop on bar counter
	c.FillRect(34, 28, 14, 8, hex("#212529"))
	c.FillRect(35, 29, 12, 5, hex("#74c0fc"))
	// Animated code editor syntax lines
	for l := 0; l < 3; l++ {
		c.FillRect(36, 30+l, 3+l*2, 1, hex("#a6e3a1"))
	}
	c.FillRect(32, 35, 18, 2, hex("#ced4da")) // keyboard base

	// Steaming Latte / Matcha Mug
	c.FillRect(30, 31, 3, 4, ColorMugCeramic)
	c.SetPixel(30, 31, ColorMugSteam) // latte foam dot

	// Sleeping cafe cat curled on right bar stool
	renderCatHelper(c, 56, 39, catBreed, frameCount)
}

func (s *CityCafeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(38, 26, 18, 16, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 2. ForestCabinScene (16-Bit Forest Sanctuary Cabin)
// ---------------------------------------------------------
type ForestCabinScene struct{}

func (s *ForestCabinScene) ID() SceneType       { return SceneForestCabin }
func (s *ForestCabinScene) Name() string     { return "Forest Sanctuary Cabin 🌲" }
func (s *ForestCabinScene) WindowBounds() Rect { return Rect{X: 12, Y: 4, W: 50, H: 28} }

func (s *ForestCabinScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#2a5c43"), hex("#81c3d7")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Layered Pine Forest Silhouettes with Depth
	for layer := 0; layer < 2; layer++ {
		treeCol := ColorLeafDark
		if layer == 0 {
			treeCol = hex("#10281b") // distant dark green
		}
		for i := 0; i < 7; i++ {
			tx := wx + i*8 + (layer*4)
			th := 14 + ((i*7+layer*5)%11)
			// Trunk
			c.FillRect(tx+2, wy+wh-th, 2, th, ColorWoodDark)
			// Layered needle tiers
			for t := 0; t < 3; t++ {
				tierY := wy + wh - th + (t * 3)
				tierW := 3 + (t * 2)
				c.FillRect(tx+3-tierW/2, tierY, tierW, 2, treeCol)
			}
		}
	}
}

func (s *ForestCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Log cabin walls with timber log texture
	c.DrawVerticalGradient(0, 0, wx-1, 36, ColorWoodDark, ColorWallPlank)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWoodDark, ColorWallPlank)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)
	for y := 4; y < 36; y += 6 {
		c.FillRect(0, y, wx-1, 1, ColorWoodPlankSep)
		c.FillRect(wx+ww+1, y, c.Width-(wx+ww+1), 1, ColorWoodPlankSep)
	}

	// Timber window frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/3, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx+(ww*2)/3, wy, 2, wh, ColorWoodDark)

	// Wood floor planks
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	for y := 36; y < c.Height; y += 5 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Large Monstera plant in terracotta pot on left
	renderPottedMonstera(c, 2, 22)

	// Focus buddy at rustic wood table
	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	// Rustic writing table
	c.FillRect(10, 32, 26, 3, ColorWoodMedium)
	c.FillRect(10, 32, 26, 1, ColorWoodHighlight)
	c.FillRect(11, 35, 2, 11, ColorWoodDark)
	c.FillRect(34, 35, 2, 11, ColorWoodDark)

	// Steaming mug & lantern on table
	c.FillRect(28, 29, 3, 3, ColorMugCeramic)
	if lampOn {
		c.FillRect(14, 24, 4, 7, ColorLampBrass)
		c.FillRect(15, 26, 2, 4, ColorWarmGlow)
	}

	// Crackling wood stove on right
	renderWoodStoveHelper(c, 52, 16, 20, 24, fireOn, frameCount)

	// Sleeping cat by the wood stove
	renderCatHelper(c, 38, 40, catBreed, frameCount)
}

func (s *ForestCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(16, 26, 16, 14, ColorWarmGlow, 0.35)
	}
	if fireOn {
		c.ApplyRadialLight(62, 26, 18.0, ColorFireOrange, 0.38)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 3. ClassicCottageScene (Handcrafted Virtual Cottage)
// ---------------------------------------------------------
type ClassicCottageScene struct{}

func (s *ClassicCottageScene) ID() SceneType       { return SceneCottageClassic }
func (s *ClassicCottageScene) Name() string     { return "Classic Cozy Cottage 🏡" }
func (s *ClassicCottageScene) WindowBounds() Rect { return Rect{X: 18, Y: 4, W: 34, H: 26} }

func (s *ClassicCottageScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = ColorSkyDayTop, ColorSkyDayBot
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Rolling hill silhouettes & pine trees outside window
	hillCol := hex("#121b2d")
	if tod == TimeDay {
		hillCol = hex("#2d6a4f")
	} else if tod == TimeSunset {
		hillCol = hex("#3a1c28")
	}
	c.FillRect(wx, wy+wh-6, ww, 6, hillCol)
	for i := 0; i < 5; i++ {
		px := wx + i*7 + 2
		ph := 4 + (i*3)%5
		c.FillRect(px, wy+wh-6-ph, 2, ph, hillCol)
	}

	// Celestial objects outside window
	if tod == TimeNight && !lightning {
		// Crescent Moon
		c.FillRect(wx+ww-7, wy+3, 4, 3, ColorStars)
		c.FillRect(wx+ww-9, wy+3, 3, 3, skyTop)
		// Stars
		for _, st := range stars {
			px := wx + (st.X % ww)
			py := wy + (st.Y % (wh - 8))
			twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
			if twinkle > 0.3 {
				c.SetPixel(px, py, ColorStars)
			}
		}
	} else if tod == TimeDawn {
		// Glowing morning star
		c.DrawGlyph(wx+6, (wy+4)/2, '✦', ColorStars)
	} else if tod == TimeDay {
		// Soft golden sun
		c.FillRect(wx+ww-7, wy+3, 5, 4, ColorFireYellow)
	}
}

func (s *ClassicCottageScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Cozy wallpaper on left and right walls
	c.DrawVerticalGradient(0, 0, wx-1, 36, ColorWallPaper, ColorWallWarm)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWallPaper, ColorWallWarm)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Ornate Arched Timber Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWindowFrame)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWindowFrame)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWindowFrame)
	// Window sill with potted trailing plant
	c.FillRect(wx-1, wy+wh-1, ww+2, 2, ColorWindowSill)
	renderPottedSucculent(c, wx+ww-6, wy+wh-4)

	// Rich Wooden Floor Planks (Rows 36 to c.Height-1)
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	for y := 36; y < c.Height; y += 5 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
		c.FillRect(0, y+1, c.Width, 1, ColorWoodMedium)
	}

	// Persian Ornate Floor Rug under desk & fireplace
	rugX, rugY, rugW, rugH := 14, 39, 44, 11
	c.FillRect(rugX, rugY, rugW, rugH, ColorFloorRugBase)
	c.DrawRectOutline(rugX, rugY, rugW, rugH, ColorFloorRugBorder)
	for rx := rugX + 4; rx < rugX+rugW-4; rx += 6 {
		c.FillRect(rx, rugY+rugH/2-1, 2, 2, ColorFloorRugGold)
	}

	// Left Wall: Bookshelf with colorful stacked books & trailing ivy
	renderBookshelfHelper(c, 2, 4, 14, 32)

	// Study Area: Oak Desk with Lamp & Focus Companion
	renderDeskHelper(c, 16, 32, 24, lampOn)

	if buddy != nil {
		buddy.Draw(c, 16, 20, frameCount)
	}

	// Right Wall: Multi-course Brick Fireplace Hearth
	renderFireplaceHelper(c, 52, 12, 20, 26, fireOn, frameCount)

	// Sleeping Cat Companion on Velvet Cushion
	renderCatHelper(c, 38, 40, catBreed, frameCount)
}

func (s *ClassicCottageScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(18, 26, 16, 14, ColorWarmGlow, 0.38)
	}
	if fireOn {
		c.ApplyRadialLight(62, 24, 20.0, ColorFireOrange, 0.40)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 4. ObservatoryScene (Starlight Celestial Observatory)
// ---------------------------------------------------------
type ObservatoryScene struct{}

func (s *ObservatoryScene) ID() SceneType       { return SceneObservatory }
func (s *ObservatoryScene) Name() string     { return "Celestial Observatory 🌌" }
func (s *ObservatoryScene) WindowBounds() Rect { return Rect{X: 10, Y: 3, W: 54, H: 29} }

func (s *ObservatoryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, ColorSkyNightTop, ColorSkyNightBot)

	// Swirling cosmic nebula dust
	nebulaCols := []colorful.Color{hex("#3d1a45"), hex("#1a2b4c"), hex("#153b3d")}
	for n := 0; n < 3; n++ {
		nx := wx + 10 + n*16
		ny := wy + 4 + n*6
		c.ApplyRadialLight(nx, ny, 10.0, nebulaCols[n], 0.25)
	}

	// Bright Crescent Moon
	c.FillRect(wx+ww-10, wy+4, 6, 5, ColorStars)
	c.FillRect(wx+ww-12, wy+4, 4, 5, ColorSkyNightTop)

	// Constellation clusters
	for i, st := range stars {
		px := wx + (st.X % ww)
		py := wy + (st.Y % wh)
		twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
		if twinkle > 0.15 {
			c.SetPixel(px, py, ColorStars)
			if i%4 == 0 {
				c.DrawGlyph(px, py/2, '✦', hex("#74c0fc"))
			}
		}
	}
}

func (s *ObservatoryScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Observatory dome architecture
	c.FillRect(0, 0, wx-1, 36, hex("#10141f"))
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, hex("#10141f"))
	c.FillRect(0, 0, c.Width, wy-1, hex("#10141f"))

	// Domed brass skylight frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorLampBrass)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorLampBrass)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorLampBrass)

	// Observatory dark stone floor
	c.FillRect(0, 36, c.Width, c.Height-36, hex("#161b26"))
	c.FillRect(0, 36, c.Width, 1, ColorLampBrass)

	// Large Brass Telescope on heavy wooden tripod facing skylight
	telX, telY := 46, 16
	c.FillRect(telX, telY, 16, 4, ColorLampBrass)        // brass tube
	c.FillRect(telX+14, telY-1, 4, 6, ColorLampBrass)    // objective lens
	c.FillRect(telX-2, telY+1, 3, 2, hex("#212529"))     // eyepiece
	// Tripod legs
	c.FillRect(telX+6, telY+4, 2, 18, ColorWoodDark)
	c.FillRect(telX+2, telY+4, 2, 18, ColorWoodDark)
	c.FillRect(telX+10, telY+4, 2, 18, ColorWoodDark)

	// Focus Buddy at star chart desk on left
	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	// Star Chart Table
	c.FillRect(10, 32, 26, 3, ColorWoodMedium)
	c.FillRect(10, 32, 26, 1, ColorWoodHighlight)
	c.FillRect(11, 35, 2, 11, ColorWoodDark)
	c.FillRect(34, 35, 2, 11, ColorWoodDark)

	// Celestial map & brass lantern
	c.FillRect(20, 30, 8, 2, ColorNotebookPaper)
	if lampOn {
		c.FillRect(14, 24, 4, 7, ColorLampBrass)
		c.FillRect(15, 26, 2, 4, ColorWarmGlow)
	}

	// Sleeping cat by observation desk
	renderCatHelper(c, 34, 40, catBreed, frameCount)
}

func (s *ObservatoryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(16, 26, 16, 14, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// Helper drawing functions for realistic, cozy interiors
// ---------------------------------------------------------

// renderBookshelfHelper draws a detailed wooden bookshelf with colorful book spines and trailing ivy.
func renderBookshelfHelper(c *Canvas, bx, by, bw, bh int) {
	// Bookshelf dark wood frame
	c.FillRect(bx, by, bw, bh, ColorWoodDark)
	c.DrawRectOutline(bx, by, bw, bh, ColorWoodPlankSep)

	// Shelves (3 levels)
	bookColors := []colorful.Color{ColorBookSpine1, ColorBookSpine2, ColorBookSpine3, ColorBookSpine4}
	for s := 0; s < 3; s++ {
		sy := by + 9 + s*9
		c.FillRect(bx+1, sy, bw-2, 1, ColorWoodMedium)

		// Stacked and standing books on each shelf
		for b := 0; b < 4; b++ {
			bX := bx + 2 + b*3
			bH := 4 + ((b*3+s*5)%4)
			bCol := bookColors[(b+s)%len(bookColors)]
			c.FillRect(bX, sy-bH, 2, bH, bCol)
		}
	}

	// Top shelf: Terracotta pot with trailing ivy vines
	c.FillRect(bx+3, by-2, 6, 3, ColorPotClay)
	c.FillRect(bx+4, by-3, 4, 1, ColorLeafLight)
	// Trailing green ivy vines dripping down the side
	c.SetPixel(bx+1, by+1, ColorLeafMedium)
	c.SetPixel(bx+1, by+3, ColorLeafLight)
	c.SetPixel(bx, by+5, ColorLeafMedium)
	c.SetPixel(bx, by+7, ColorLeafDark)
}

// renderPottedSucculent draws a tiny terracotta pot with green leaves.
func renderPottedSucculent(c *Canvas, px, py int) {
	c.FillRect(px+1, py+1, 3, 2, ColorPotClay)
	c.FillRect(px, py+1, 5, 1, ColorPotClay)
	c.SetPixel(px+1, py, ColorLeafLight)
	c.SetPixel(px+2, py, ColorLeafMedium)
	c.SetPixel(px+3, py, ColorLeafDark)
}

// renderPottedMonstera draws a large potted monstera plant with split leaves.
func renderPottedMonstera(c *Canvas, px, py int) {
	c.FillRect(px+2, py+10, 6, 6, ColorPotClay)
	c.FillRect(px+1, py+10, 8, 2, ColorPotClay)

	// Broad green monstera leaves
	c.FillRect(px, py+3, 4, 4, ColorLeafDark)
	c.FillRect(px+5, py+1, 5, 5, ColorLeafMedium)
	c.FillRect(px+2, py, 4, 4, ColorLeafLight)
	// Leaf cutout details
	c.SetPixel(px+1, py+4, ColorWallPlank)
	c.SetPixel(px+7, py+3, ColorWallPlank)
}

// renderDeskHelper draws an oak writing desk with green banker's lamp, journal, and mug.
func renderDeskHelper(c *Canvas, dx, dy, dw int, lampOn bool) {
	// Oak desk tabletop
	c.FillRect(dx, dy, dw, 3, ColorWoodMedium)
	c.FillRect(dx, dy, dw, 1, ColorWoodHighlight)
	// Desk drawer & legs
	c.FillRect(dx+2, dy+3, dw-4, 3, ColorWoodDark)
	c.FillRect(dx+dw/2-2, dy+4, 4, 1, ColorLampBrass) // brass drawer pull
	c.FillRect(dx+2, dy+6, 2, 10, ColorWoodDark)
	c.FillRect(dx+dw-4, dy+6, 2, 10, ColorWoodDark)

	// Classic Banker's Lamp (Brass neck + Emerald Green glass shade)
	if lampOn {
		c.FillRect(dx+2, dy-2, 3, 2, ColorLampBrass)
		c.FillRect(dx+3, dy-6, 1, 4, ColorLampBrass)
		c.FillRect(dx+1, dy-9, 5, 3, ColorLampShade)
		c.FillRect(dx+2, dy-7, 3, 1, ColorWarmGlow)
	} else {
		c.FillRect(dx+2, dy-2, 3, 2, hex("#495057"))
		c.FillRect(dx+3, dy-6, 1, 4, hex("#495057"))
		c.FillRect(dx+1, dy-9, 5, 3, hex("#212529"))
	}

	// Open Notebook Journal with scribbled lines
	c.FillRect(dx+10, dy-1, 7, 2, ColorNotebookPaper)
	c.SetPixel(dx+12, dy, ColorWoodDark)
	c.SetPixel(dx+14, dy, ColorWoodDark)

	// Steaming Ceramic Terracotta Mug
	c.FillRect(dx+19, dy-3, 3, 3, ColorMugCeramic)
	c.SetPixel(dx+22, dy-2, ColorMugCeramic) // handle
}

// renderFireplaceHelper draws a multi-course brick hearth with mantle, candle, and animated flame engine.
func renderFireplaceHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	// Brick chimney masonry
	c.FillRect(fx, fy, fw, fh, ColorHearthBrick)
	// Brick mortar lines
	for by := fy + 3; by < fy+fh; by += 4 {
		c.FillRect(fx, by, fw, 1, ColorHearthMortar)
		offset := ((by - fy) / 4) % 2 * 3
		for bx := fx + offset; bx < fx+fw; bx += 6 {
			c.FillRect(bx, by-3, 1, 3, ColorHearthMortar)
		}
	}

	// Timber Mantlepiece at top of fireplace
	c.FillRect(fx-2, fy, fw+4, 2, ColorWoodDark)
	c.FillRect(fx-2, fy, fw+4, 1, ColorWoodHighlight)

	// Candle on mantlepiece
	c.FillRect(fx+3, fy-4, 2, 4, ColorCandleWax)
	if fireOn {
		c.SetPixel(fx+3, fy-5, ColorFireYellow)
	}

	// Arched firebox interior
	boxX, boxY, boxW, boxH := fx+3, fy+7, fw-6, fh-8
	c.FillRect(boxX, boxY, boxW, boxH, ColorHearthInterior)
	c.DrawRectOutline(boxX, boxY, boxW, boxH, ColorHearthBrickDark)

	if !fireOn {
		// Cold ash bed
		c.FillRect(boxX+2, boxY+boxH-2, boxW-4, 1, hex("#343a40"))
		return
	}

	// Hot glowing charcoal / ember bed
	c.FillRect(boxX+2, boxY+boxH-2, boxW-4, 1, ColorFireEmber)

	// Multi-layered animated dancing flames
	fireCenter := boxX + boxW/2
	fireTick := (frameCount / 3) % 4

	for i := -3; i <= 3; i++ {
		px := fireCenter + i
		flameH := int(5.0 + 3.0*math.Sin(float64(i)*0.8+float64(fireTick)*1.4))

		for fyPos := 0; fyPos < flameH; fyPos++ {
			py := boxY + boxH - 2 - fyPos
			if py < boxY {
				continue
			}
			if fyPos < 2 {
				c.SetPixel(px, py, ColorFireCore)
			} else if fyPos < 4 {
				c.SetPixel(px, py, ColorFireYellow)
			} else if fyPos < 6 {
				c.SetPixel(px, py, ColorFireOrange)
			} else {
				c.SetPixel(px, py, ColorFireRed)
			}
		}
	}
}

// renderWoodStoveHelper draws a rustic cast-iron wood stove.
func renderWoodStoveHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	// Cast iron stove body
	c.FillRect(fx+2, fy+6, fw-4, fh-8, hex("#212529"))
	c.DrawRectOutline(fx+2, fy+6, fw-4, fh-8, hex("#343a40"))
	// Stove legs
	c.FillRect(fx+3, fy+fh-2, 2, 2, hex("#10141f"))
	c.FillRect(fx+fw-5, fy+fh-2, 2, 2, hex("#10141f"))
	// Stove pipe chimney
	c.FillRect(fx+fw/2-1, fy, 3, 6, hex("#212529"))

	// Glowing stove window
	if fireOn {
		c.FillRect(fx+5, fy+10, fw-10, 6, ColorFireOrange)
		c.FillRect(fx+6, fy+11, fw-12, 3, ColorFireYellow)
	} else {
		c.FillRect(fx+5, fy+10, fw-10, 6, hex("#10141f"))
	}
}

// renderCatHelper draws an organic, cute sleeping cat companion with smooth micro-animations.
func renderCatHelper(c *Canvas, cx, cy int, breed CatBreed, frameCount uint64) {
	// Plush Indigo Velvet Cushion
	c.FillRect(cx-2, cy+4, 16, 5, ColorCatCushion)
	c.DrawRectOutline(cx-2, cy+4, 16, 5, hex("#2a2356"))

	var furMain, furDetail colorful.Color
	switch breed {
	case CatOrange:
		furMain = hex("#e76f51")
		furDetail = hex("#f4a261")
	case CatTuxedo:
		furMain = hex("#1a1a24")
		furDetail = hex("#ffffff")
	case CatCalico:
		furMain = hex("#f4a261")
		furDetail = hex("#264653")
	case CatBlack:
		furMain = hex("#121218")
		furDetail = hex("#2b2b36")
	}

	// Soft breathing offset (expands and contracts every 16 frames)
	breathOffset := 0
	if (frameCount/16)%2 == 0 {
		breathOffset = -1
	}

	// Curled Body
	c.FillRect(cx, cy+breathOffset, 10, 5, furMain)
	c.FillRect(cx+2, cy+1+breathOffset, 6, 3, furDetail)

	// Head tucked in
	c.FillRect(cx-1, cy+1+breathOffset, 4, 4, furMain)
	// Cute inner ear pink details
	c.SetPixel(cx-1, cy+breathOffset, hex("#ff7597"))
	c.SetPixel(cx+1, cy+breathOffset, hex("#ff7597"))

	// Dynamic Tail Wag
	tailWag := int((frameCount / 18) % 3)
	c.FillRect(cx+10, cy+2+breathOffset-tailWag, 2, 2, furMain)
	c.FillRect(cx+11, cy+1+breathOffset-tailWag, 2, 2, furMain)

	// Purr & Sleep glyphs
	if (frameCount/24)%2 == 1 {
		c.DrawGlyph(cx+4, (cy-2)/2, '♥', hex("#ff7597"))
	} else if (frameCount/14)%3 == 1 {
		c.DrawGlyph(cx-2, (cy-2)/2, 'z', hex("#74c0fc"))
	}
}

// ---------------------------------------------------------
// 5. ZenTeahouseScene (Traditional Japanese Engawa & Tea Garden)
// ---------------------------------------------------------
type ZenTeahouseScene struct{}

func (s *ZenTeahouseScene) ID() SceneType       { return SceneZenTeahouse }
func (s *ZenTeahouseScene) Name() string     { return "Zen Engawa Tea Sanctuary 🍵" }
func (s *ZenTeahouseScene) WindowBounds() Rect { return Rect{X: 12, Y: 4, W: 50, H: 28} }

func (s *ZenTeahouseScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#508991"), hex("#b8e0d2")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Distant Mount Fuji Silhouette with snow peak
	fujiBase := wy + wh - 8
	fujiPeakX := wx + ww/2 - 2
	fujiPeakY := wy + 5
	fujiSnowH := 4
	for y := fujiPeakY; y < fujiBase; y++ {
		span := (y - fujiPeakY) * 3 / 2
		mCol := hex("#1e293b")
		if tod == TimeDay {
			mCol = hex("#334155")
		} else if tod == TimeSunset {
			mCol = hex("#4c1d4f")
		}
		if y < fujiPeakY+fujiSnowH {
			mCol = hex("#f8fafc")
		}
		startX := fujiPeakX - span
		endX := fujiPeakX + span
		if startX < wx {
			startX = wx
		}
		if endX > wx+ww {
			endX = wx + ww
		}
		if endX > startX {
			c.FillRect(startX, y, endX-startX, 1, mCol)
		}
	}

	// Garden Greenery & Moss Hills in midground
	c.FillRect(wx, wy+wh-8, ww, 3, hex("#1b4332"))

	// Gentle Koi Pond Water with subtle reflections
	for y := wy + wh - 5; y < wy+wh; y++ {
		waterCol := ColorPondDeep
		if y%2 == 0 {
			waterCol = ColorPondShallow
		}
		c.FillRect(wx, y, ww, 1, waterCol)
	}

	// Animated Golden Koi Fish swimming in pond
	koiTick := int((frameCount / 6) % uint64(ww-10))
	koiX := wx + 4 + koiTick
	koiY := wy + wh - 3 + int(math.Sin(float64(frameCount)*0.1)*1.2)
	if koiX < wx+ww-5 && koiY >= wy+wh-5 && koiY < wy+wh {
		c.FillRect(koiX, koiY, 3, 1, ColorKoiGold)
		c.SetPixel(koiX+3, koiY, ColorKoiWhite)
	}

	// Traditional Stone Garden Lantern (Tōrō) on right edge
	toroX, toroY := wx+ww-8, wy+wh-11
	c.FillRect(toroX+1, toroY, 4, 1, ColorStoneLantern) // cap
	c.FillRect(toroX, toroY+1, 6, 1, ColorStoneLantern)
	c.FillRect(toroX+1, toroY+2, 4, 3, hex("#1e293b")) // chamber
	if tod == TimeNight || tod == TimeSunset {
		c.FillRect(toroX+2, toroY+3, 2, 1, ColorWarmGlow) // glowing paper light
	}
	c.FillRect(toroX+1, toroY+5, 4, 4, ColorStoneLantern) // pedestal

	// Bamboo Grove stalks on left side of garden
	for b := 0; b < 3; b++ {
		bx := wx + 2 + b*3
		sway := int(math.Sin(float64(frameCount)*0.05 + float64(b))*0.8)
		c.FillRect(bx+sway, wy+3, 2, wh-6, ColorLeafDark)
		for ny := wy + 6; ny < wy+wh-5; ny += 5 {
			c.SetPixel(bx+sway, ny, ColorLeafLight)
			c.SetPixel(bx+1+sway, ny, ColorLeafLight)
		}
	}

	// Drifting Sakura Cherry Blossom Petals
	for i := 0; i < 6; i++ {
		px := wx + ((int(frameCount)*1 + i*13) % ww)
		py := wy + ((int(frameCount/2) + i*7) % (wh - 2))
		c.SetPixel(px, py, ColorSakuraPink)
		if i%2 == 0 {
			c.SetPixel(px+1, py, ColorSakuraSoft)
		}
	}

	// Celestial: Moon or Sun
	if tod == TimeNight && !lightning {
		c.FillRect(wx+8, wy+3, 4, 3, ColorStars)
		c.FillRect(wx+6, wy+3, 3, 3, skyTop)
	} else if tod == TimeSunset {
		c.FillRect(wx+ww-10, wy+6, 6, 4, ColorFireOrange)
	}
}

func (s *ZenTeahouseScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww := r.X, r.Y, r.W

	// Shoji sliding door lattice walls on left & right
	c.FillRect(0, 0, wx-1, 34, ColorShojiPaper)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 34, ColorShojiPaper)
	for y := 2; y < 34; y += 4 {
		c.FillRect(0, y, wx-1, 1, ColorShojiLattice)
		c.FillRect(wx+ww+1, y, c.Width-(wx+ww+1), 1, ColorShojiLattice)
	}
	for x := 3; x < wx-1; x += 4 {
		c.FillRect(x, 0, 1, 34, ColorShojiLattice)
	}
	for x := wx + ww + 4; x < c.Width; x += 4 {
		c.FillRect(x, 0, 1, 34, ColorShojiLattice)
	}

	// Cedar Eaves / Overhang at top
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)
	c.FillRect(0, wy-1, c.Width, 1, ColorWoodHighlight)

	// Hanging bronze wind chime (Furin) with swaying paper slip
	furinX := wx + 10
	furinSway := int(math.Sin(float64(frameCount)*0.08)*1.2)
	c.FillRect(furinX, wy, 1, 2, ColorLampBrass)
	c.FillRect(furinX-1, wy+2, 3, 2, ColorLampBrass)
	c.FillRect(furinX+furinSway, wy+4, 1, 4, ColorShojiPaper)

	// Veranda Engawa Wooden Deck (Rows 32 to 36)
	c.FillRect(0, 32, c.Width, 4, ColorWoodMedium)
	c.FillRect(0, 32, c.Width, 1, ColorWoodHighlight)
	c.FillRect(0, 34, c.Width, 1, ColorWoodPlankSep)

	// Tatami Straw Mat Room Floor (Rows 36 to c.Height-1)
	c.FillRect(0, 36, c.Width, c.Height-36, ColorTatamiBase)
	c.FillRect(0, 36, c.Width, 1, ColorTatamiBorder)
	c.FillRect(0, 44, c.Width, 1, ColorTatamiBorder)
	c.FillRect(36, 36, 2, c.Height-36, ColorTatamiBorder)

	// Low dark wood Chabudai tea table (center-left)
	tblX, tblY, tblW := 12, 34, 20
	c.FillRect(tblX, tblY, tblW, 3, ColorWoodDark)
	c.FillRect(tblX, tblY, tblW, 1, ColorWoodMedium)
	c.FillRect(tblX+2, tblY+3, 2, 5, ColorWoodDark)
	c.FillRect(tblX+tblW-4, tblY+3, 2, 5, ColorWoodDark)

	// Cast-Iron Tetsubin Teapot on table
	c.FillRect(tblX+2, tblY-4, 5, 4, hex("#1a1a1a"))
	c.FillRect(tblX+3, tblY-5, 3, 1, ColorLampBrass)
	c.SetPixel(tblX+1, tblY-3, hex("#1a1a1a")) // spout
	c.SetPixel(tblX+7, tblY-3, hex("#1a1a1a")) // handle

	// Vibrant Green Matcha Bowl (Chawan)
	c.FillRect(tblX+11, tblY-3, 4, 3, ColorPotClay)
	c.FillRect(tblX+12, tblY-3, 2, 1, ColorMatchaGreen)

	// Focus buddy enjoying tea
	if buddy != nil {
		buddy.Draw(c, 22, 22, frameCount)
	}

	// Sleeping Calico Cat on soft cushion (center-right)
	renderCatHelper(c, 48, 38, catBreed, frameCount)
}

func (s *ZenTeahouseScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(58, 22, 16, 14, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 6. CozyCampfireScene (Starlit Alpine Campfire)
// ---------------------------------------------------------
type CozyCampfireScene struct{}

func (s *CozyCampfireScene) ID() SceneType       { return SceneCozyCampfire }
func (s *CozyCampfireScene) Name() string     { return "Starlit Alpine Campfire 🏕️" }
func (s *CozyCampfireScene) WindowBounds() Rect { return Rect{X: 4, Y: 2, W: 66, H: 34} }

func (s *CozyCampfireScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#1a365d"), hex("#79addc")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Rugged Alpine Mountain Silhouette with Snow Caps
	mountainCols := []colorful.Color{hex("#0f172a"), hex("#1e293b")}
	for m := 0; m < 3; m++ {
		mx := wx + m*22 + 4
		mw := 24
		mh := 14 + (m*5)%7
		for y := 0; y < mh; y++ {
			span := y * mw / (mh * 2)
			py := wy + wh - 12 - mh + y
			mCol := mountainCols[m%2]
			if y < 4 {
				mCol = hex("#e2e8f0")
			}
			c.FillRect(mx+mw/2-span, py, span*2, 1, mCol)
		}
	}

	// Alpine Mirror Lake Water
	c.FillRect(wx, wy+wh-10, ww, 10, ColorLakeDeep)
	for y := wy + wh - 10; y < wy+wh; y += 2 {
		c.FillRect(wx, y, ww, 1, ColorLakeReflect)
	}

	// Pine tree shoreline silhouettes
	for i := 0; i < 9; i++ {
		px := wx + i*7 + 1
		ph := 8 + ((i*11)%6)
		c.FillRect(px+1, wy+wh-10-ph, 2, ph, ColorPineSilhouette)
		for t := 0; t < 3; t++ {
			c.FillRect(px+2-t, wy+wh-10-ph+t*2, 1+t*2, 2, ColorPineSilhouette)
		}
	}

	// Twinkling stars & occasional shooting star
	if tod == TimeNight && !lightning {
		for _, st := range stars {
			px := wx + (st.X % ww)
			py := wy + (st.Y % (wh - 12))
			twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
			if twinkle > 0.25 {
				c.SetPixel(px, py, ColorStars)
			}
		}
		shootTick := int((frameCount / 2) % 120)
		if shootTick < 20 {
			sx := wx + 10 + shootTick*2
			sy := wy + 4 + shootTick
			c.SetPixel(sx, sy, PicoWhite)
			c.SetPixel(sx-1, sy-1, PicoCyan)
		}
		c.FillRect(wx+ww-8, wy+3, 4, 4, ColorStars)
		c.FillRect(wx+ww-10, wy+3, 3, 4, skyTop)
	}
}

func (s *CozyCampfireScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	// Campsite earth & grassy clearing
	c.FillRect(0, 34, c.Width, c.Height-34, hex("#1a120c"))
	c.FillRect(0, 34, c.Width, 2, hex("#283618"))

	// Canvas A-Frame Camp Tent on left
	tentX, tentY, tentW, tentH := 2, 24, 18, 16
	for y := 0; y < tentH; y++ {
		span := y * tentW / (tentH * 2)
		c.FillRect(tentX+tentW/2-span, tentY+y, span*2, 1, ColorCampTent)
	}
	c.FillRect(tentX+tentW/2-2, tentY+tentH-8, 4, 8, hex("#140d07"))
	c.DrawLine(tentX+tentW/2, tentY, tentX-2, tentY+tentH, hex("#606c38"))
	c.DrawLine(tentX+tentW/2, tentY, tentX+tentW+2, tentY+tentH, hex("#606c38"))

	// Rustic Split-Log Bench (center-left)
	c.FillRect(18, 38, 14, 3, ColorWoodMedium)
	c.FillRect(18, 38, 14, 1, ColorWoodHighlight)
	c.FillRect(20, 41, 2, 4, ColorWoodDark)
	c.FillRect(29, 41, 2, 4, ColorWoodDark)

	// Tree Stump Table with enamel mug & lantern
	c.FillRect(34, 36, 8, 8, ColorWoodDark)
	c.FillRect(34, 36, 8, 2, ColorWoodMedium)
	c.FillRect(36, 33, 3, 3, hex("#457b9d"))
	c.SetPixel(39, 34, hex("#457b9d"))

	// Focus buddy sitting at camp bench
	if buddy != nil {
		buddy.Draw(c, 18, 26, frameCount)
	}

	// Stone Ring Campfire Hearth (center-right)
	firePitX, firePitY := 48, 38
	stoneCols := []colorful.Color{hex("#495057"), hex("#6c757d"), hex("#343a40")}
	for s := 0; s < 7; s++ {
		sx := firePitX - 6 + s*2
		c.FillRect(sx, firePitY+4, 3, 2, stoneCols[s%3])
	}
	c.FillRect(firePitX-5, firePitY+2, 12, 3, hex("#1a0f0a"))

	if fireOn {
		c.FillRect(firePitX-4, firePitY+3, 10, 2, ColorFireEmber)
		fireTick := (frameCount / 3) % 4
		for i := -3; i <= 3; i++ {
			px := firePitX + 1 + i
			flameH := int(6.0 + 3.5*math.Sin(float64(i)*0.9+float64(fireTick)*1.5))
			for fy := 0; fy < flameH; fy++ {
				py := firePitY + 2 - fy
				if py < 10 {
					continue
				}
				if fy < 2 {
					c.SetPixel(px, py, ColorFireCore)
				} else if fy < 4 {
					c.SetPixel(px, py, ColorFireYellow)
				} else if fy < 7 {
					c.SetPixel(px, py, ColorFireOrange)
				} else {
					c.SetPixel(px, py, ColorFireRed)
				}
			}
		}

		for sm := 0; sm < 3; sm++ {
			smX := firePitX + int(math.Sin(float64(frameCount)*0.08+float64(sm))*3.0)
			smY := firePitY - 8 - sm*3 - int((frameCount/4)%5)
			if smY > 0 {
				c.SetPixel(smX, smY, ColorCampSmoke)
			}
		}
	}

	// Cozy Plaid Wool Camp Blanket with Sleeping Pet
	c.FillRect(58, 42, 14, 8, ColorPlaidRed)
	c.FillRect(58, 45, 14, 2, ColorPlaidNavy)
	c.FillRect(64, 42, 2, 8, ColorPlaidNavy)
	renderCatHelper(c, 58, 38, catBreed, frameCount)
}

func (s *CozyCampfireScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(49, 36, 24.0, ColorFireOrange, 0.45)
	}
	if lampOn {
		c.ApplyRadialLight(38, 34, 14.0, ColorWarmGlow, 0.30)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.75)
	}
}

// ---------------------------------------------------------
// 7. CozyAtticScene (Sunlit Attic Artist Atelier)
// ---------------------------------------------------------
type CozyAtticScene struct{}

func (s *CozyAtticScene) ID() SceneType       { return SceneCozyAttic }
func (s *CozyAtticScene) Name() string     { return "Sunlit Attic Artist Atelier 🎨" }
func (s *CozyAtticScene) WindowBounds() Rect { return Rect{X: 18, Y: 4, W: 36, H: 24} }

func (s *CozyAtticScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = ColorSkyDayTop, ColorSkyDayBot
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	for t := 0; t < 5; t++ {
		tx := wx + t*8 + 1
		th := 6 + (t*7)%5
		c.FillRect(tx, wy+wh-th, 8, th, ColorLeafDark)
		c.FillRect(tx+1, wy+wh-th-1, 6, 2, ColorLeafMedium)
	}

	if tod == TimeNight && !lightning {
		for _, st := range stars {
			px := wx + (st.X % ww)
			py := wy + (st.Y % (wh - 6))
			twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
			if twinkle > 0.3 {
				c.SetPixel(px, py, ColorStars)
			}
		}
		c.FillRect(wx+ww-8, wy+3, 4, 4, ColorStars)
		c.FillRect(wx+ww-10, wy+3, 3, 4, skyTop)
	} else if tod == TimeDay {
		cloudX := wx + int((frameCount/3)%uint64(ww+10)) - 8
		c.FillRect(cloudX, wy+4, 10, 3, hex("#ffffff"))
		c.FillRect(cloudX+2, wy+3, 6, 1, hex("#ffffff"))
	}
}

func (s *CozyAtticScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.FillRect(0, 0, wx-1, 36, ColorWoodDark)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWoodDark)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	for x := 0; x < c.Width; x += 16 {
		c.DrawLine(x, 0, x+8, 36, ColorWoodMedium)
	}

	c.FillRect(6, 4, 3, 6, ColorLavenderDried)
	c.SetPixel(7, 3, ColorWoodDark)
	c.FillRect(12, 5, 3, 5, ColorEucalyptus)
	c.SetPixel(13, 4, ColorWoodDark)

	fairyCols := []colorful.Color{ColorWarmGlow, hex("#ff7597"), hex("#74c0fc"), hex("#63e6be")}
	for f := 0; f < 8; f++ {
		fx := 4 + f*8
		fy := 6 + int(math.Sin(float64(f)*0.8)*2.0)
		fCol := fairyCols[(f+int(frameCount/20))%len(fairyCols)]
		c.SetPixel(fx, fy, fCol)
	}

	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodHighlight)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodMedium)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodMedium)
	c.FillRect(wx-1, wy+wh-1, ww+2, 2, ColorWoodHighlight)

	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodMedium)
	for y := 36; y < c.Height; y += 5 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	easelX, easelY := 48, 18
	c.FillRect(easelX+6, easelY, 2, 24, ColorCanvasEasel)
	c.DrawLine(easelX+6, easelY+4, easelX, easelY+24, ColorCanvasEasel)
	c.DrawLine(easelX+6, easelY+4, easelX+12, easelY+24, ColorCanvasEasel)
	c.FillRect(easelX+1, easelY+6, 12, 10, ColorCanvasCloth)
	c.DrawRectOutline(easelX+1, easelY+6, 12, 10, ColorWoodDark)
	c.FillRect(easelX+2, easelY+7, 10, 4, hex("#74c0fc"))
	c.FillRect(easelX+2, easelY+11, 10, 4, hex("#2d6a4f"))
	c.FillRect(easelX+7, easelY+8, 2, 2, ColorFireYellow)

	renderDeskHelper(c, 14, 32, 22, lampOn)
	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	bookCols := []colorful.Color{ColorBookSpine1, ColorBookSpine2, ColorBookSpine3, ColorBookSpine4}
	for b := 0; b < 3; b++ {
		c.FillRect(37, 43-b*2, 12, 2, bookCols[b%len(bookCols)])
	}
	renderCatHelper(c, 36, 36, catBreed, frameCount)
}

func (s *CozyAtticScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(16, 26, 16, 14, ColorWarmGlow, 0.38)
	}
	c.ApplyRadialLight(24, 6, 18.0, ColorWarmGlow, 0.20)
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 8. WinterCabinScene (Winter Snowfall Hearth Cabin)
// ---------------------------------------------------------
type WinterCabinScene struct{}

func (s *WinterCabinScene) ID() SceneType       { return ScenePixelWinterCabin }
func (s *WinterCabinScene) Name() string     { return "Winter Snowfall Hearth Cabin ❄️" }
func (s *WinterCabinScene) WindowBounds() Rect { return Rect{X: 16, Y: 4, W: 34, H: 26} }

func (s *WinterCabinScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#8da9c4"), hex("#eef4f8")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	c.FillRect(wx, wy+wh-8, ww, 8, hex("#e2e8f0"))
	c.FillRect(wx, wy+wh-9, ww, 1, hex("#ffffff"))

	for p := 0; p < 4; p++ {
		px := wx + p*8 + 2
		ph := 12 + (p*7)%6
		c.FillRect(px+2, wy+wh-8-ph, 2, ph, ColorWoodDark)
		for t := 0; t < 3; t++ {
			c.FillRect(px+3-t*2, wy+wh-8-ph+t*3, 1+t*4, 2, hex("#1b4332"))
			c.FillRect(px+3-t*2, wy+wh-8-ph+t*3, 1+t*4, 1, hex("#ffffff"))
		}
	}

	for i := 0; i < 14; i++ {
		sx := wx + ((int(frameCount)*1 + i*9) % ww)
		sy := wy + ((int(frameCount/2) + i*5) % wh)
		c.SetPixel(sx, sy, hex("#ffffff"))
	}
}

func (s *WinterCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.FillRect(0, 0, wx-1, 36, ColorWallPlank)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWallPlank)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)
	for y := 4; y < 36; y += 5 {
		c.FillRect(0, y, wx-1, 1, ColorWoodPlankSep)
		c.FillRect(wx+ww+1, y, c.Width-(wx+ww+1), 1, ColorWoodPlankSep)
	}

	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)
	c.FillRect(wx-1, wy+wh-1, ww+2, 2, ColorWoodMedium)
	c.SetPixel(wx, wy, ColorFrostIce)
	c.SetPixel(wx+1, wy, ColorFrostIce)
	c.SetPixel(wx, wy+1, ColorFrostIce)
	c.SetPixel(wx+ww-1, wy, ColorFrostIce)
	c.SetPixel(wx+ww-1, wy+1, ColorFrostIce)

	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	for y := 36; y < c.Height; y += 5 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	c.FillRect(14, 40, 44, 9, hex("#7f4f24"))
	c.DrawRectOutline(14, 40, 44, 9, hex("#936639"))
	c.FillRect(18, 42, 36, 5, hex("#a68a64"))

	renderDeskHelper(c, 14, 32, 22, lampOn)
	c.FillRect(14+18, 29, 3, 3, ColorCocoaHot)
	c.SetPixel(14+19, 29, ColorMarshmallow)

	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	renderFireplaceHelper(c, 50, 10, 22, 28, fireOn, frameCount)

	c.FillRect(44, 32, 5, 4, hex("#d6ccc2"))
	c.FillRect(44, 32, 1, 4, hex("#212529"))
	c.FillRect(47, 33, 1, 2, hex("#212529"))

	renderCatHelper(c, 36, 40, catBreed, frameCount)
}

func (s *WinterCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(61, 24, 22.0, ColorFireOrange, 0.45)
	}
	if lampOn {
		c.ApplyLightCone(16, 26, 16, 14, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.75)
	}
}

// ---------------------------------------------------------
// 9. CozyGreenhouseScene (Glasshouse Solarium Sanctuary)
// ---------------------------------------------------------
type CozyGreenhouseScene struct{}

func (s *CozyGreenhouseScene) ID() SceneType       { return SceneCozyGreenhouse }
func (s *CozyGreenhouseScene) Name() string     { return "Glasshouse Solarium Sanctuary 🌿" }
func (s *CozyGreenhouseScene) WindowBounds() Rect { return Rect{X: 6, Y: 2, W: 62, H: 32} }

func (s *CozyGreenhouseScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#40916c"), hex("#b7e4c7")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Distant misty garden tree canopies in layers
	for t := 0; t < 6; t++ {
		tx := wx + t*10 + 2
		th := 14 + (t*7)%8
		canopyCol := hex("#1b4332")
		if tod == TimeSunset {
			canopyCol = hex("#3f2d3d")
		} else if tod == TimeNight {
			canopyCol = hex("#0f1c15")
		}
		c.FillCircle(tx+4, wy+wh-th, 6, canopyCol)
	}

	// Garden blooming flowerbed (Lavender, Hydrangea, Camellia)
	c.FillRect(wx, wy+wh-6, ww, 6, hex("#2d6a4f"))
	flowerColors := []colorful.Color{ColorLavenderDried, hex("#74c0fc"), hex("#ff7597"), hex("#ffd166")}
	for fx := wx + 1; fx < wx+ww-1; fx += 2 {
		fCol := flowerColors[(fx*3+int(frameCount/20))%len(flowerColors)]
		c.SetPixel(fx, wy+wh-4+int((fx%3)*1), fCol)
	}

	// Rain droplets on glass panes
	for i := 0; i < 8; i++ {
		dropX := wx + int((frameCount*2+uint64(i)*11)%uint64(ww-2))
		dropY := wy + int((frameCount*3+uint64(i)*7)%uint64(wh-4))
		c.SetPixel(dropX, dropY, ColorRainDrop)
		c.SetPixel(dropX, dropY+1, ColorRainSplash)
	}

	// Arched iron greenhouse ceiling ribs
	archApexX := wx + ww/2
	archApexY := wy + 2
	for i := 0; i < 5; i++ {
		baseX := wx + i*(ww/4)
		c.DrawLine(archApexX, archApexY, baseX, wy+wh, ColorGreenhouseFrame)
	}
}

func (s *CozyGreenhouseScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Greenhouse Iron Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorGreenhouseFrame)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorGreenhouseFrame)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorGreenhouseFrame)

	// Terracotta Tiled Solarium Floor (Rows 36 to Height)
	c.FillRect(0, 36, c.Width, c.Height-36, hex("#7f4f24"))
	for y := 36; y < c.Height; y += 4 {
		c.FillRect(0, y, c.Width, 1, hex("#582f0e"))
	}
	for x := 0; x < c.Width; x += 8 {
		c.FillRect(x, 36, 1, c.Height-36, hex("#582f0e"))
	}

	// Hanging Macrame Planters from ceiling truss
	// Planter 1 (Left)
	p1Sway := int(math.Sin(float64(frameCount)*0.06) * 1.0)
	c.DrawLine(14, wy, 14+p1Sway, wy+8, hex("#d6ccc2"))
	c.FillCircle(14+p1Sway, wy+9, 2, ColorHangingPotClay)
	for l := 0; l < 4; l++ {
		c.SetPixel(13+p1Sway, wy+11+l*2, ColorFernGreen)
		c.SetPixel(15+p1Sway, wy+12+l*2, ColorLeafMedium)
	}

	// Planter 2 (Right)
	p2Sway := int(math.Cos(float64(frameCount)*0.05) * 1.0)
	c.DrawLine(58, wy, 58+p2Sway, wy+7, hex("#d6ccc2"))
	c.FillCircle(58+p2Sway, wy+8, 2, ColorHangingPotClay)
	for l := 0; l < 5; l++ {
		c.SetPixel(57+p2Sway, wy+10+l*2, ColorLeafLight)
		c.SetPixel(59+p2Sway, wy+11+l*2, ColorFernGreen)
	}

	// Large Potted Monstera Deliciosa on far left
	c.FillRect(2, 34, 8, 8, ColorPotClay)
	c.FillRect(1, 33, 10, 2, ColorHangingPotClay)
	c.FillCircle(5, 26, 4, ColorMonsteraDark)
	c.FillCircle(9, 23, 3, ColorLeafMedium)
	c.SetPixel(6, 25, hex("#52b788"))
	c.SetPixel(8, 23, hex("#52b788"))

	// Large Fiddle Leaf Fig on far right
	c.FillRect(62, 34, 8, 8, ColorPotClay)
	c.FillRect(61, 33, 10, 2, ColorHangingPotClay)
	c.FillCircle(65, 25, 4, ColorLeafDark)
	c.FillCircle(68, 22, 3, ColorLeafLight)

	// Rustic Cedar Potting Table / Desk (Center-left)
	renderDeskHelper(c, 16, 32, 24, lampOn)

	// Polished Copper Watering Can on Potting Table
	c.FillRect(34, 27, 4, 4, ColorCopperWaterCan)
	c.FillRect(32, 28, 2, 1, ColorLampBrass) // spout
	c.FillRect(38, 26, 1, 4, ColorLampBrass) // handle

	// Steaming Herbal Tea Mug
	c.FillRect(20, 29, 3, 3, ColorMugCeramic)
	c.SetPixel(20, 29, ColorMatchaGreen) // green tea foam

	// FocusBuddy companion gardening & studying
	if buddy != nil {
		buddy.Draw(c, 18, 20, frameCount)
	}

	// Fairy Light Strand strung along top truss
	if lampOn {
		for fx := wx + 2; fx < wx+ww-2; fx += 5 {
			fy := wy + 3 + int(math.Sin(float64(fx)*0.2)*1.5)
			lightCol := ColorWarmGlow
			if (uint64(fx)+frameCount/10)%2 == 0 {
				lightCol = ColorFireYellow
			}
			c.SetPixel(fx, fy, lightCol)
		}
	}

	// Sleeping Calico Cat on Woven Cushion
	renderCatHelper(c, 46, 38, catBreed, frameCount)
}

func (s *CozyGreenhouseScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyRadialLight(37, 20, 22.0, ColorWarmGlow, 0.40)
	}
	if fireOn {
		c.ApplyRadialLight(48, 36, 18.0, ColorFireOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 10. MidnightLibraryScene (Midnight Stained Glass Book Nook)
// ---------------------------------------------------------
type MidnightLibraryScene struct{}

func (s *MidnightLibraryScene) ID() SceneType       { return SceneCozyLibrary }
func (s *MidnightLibraryScene) Name() string     { return "Midnight Stained Glass Library 📚" }
func (s *MidnightLibraryScene) WindowBounds() Rect { return Rect{X: 20, Y: 3, W: 34, H: 26} }

func (s *MidnightLibraryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = ColorSkyDayTop, ColorSkyDayBot
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Arched Stained Glass Window border
	glassColors := []colorful.Color{ColorStainedGlassRuby, ColorStainedGlassGold, ColorStainedGlassBlue, hex("#2a9d8f")}
	for bx := wx; bx < wx+ww; bx++ {
		col := glassColors[(bx*3)%len(glassColors)]
		c.SetPixel(bx, wy, col)
		c.SetPixel(bx, wy+wh-1, col)
	}
	for by := wy; by < wy+wh; by++ {
		col := glassColors[(by*3)%len(glassColors)]
		c.SetPixel(wx, by, col)
		c.SetPixel(wx+ww-1, by, col)
	}

	// Crescent Moon through central glass
	if tod == TimeNight && !lightning {
		c.FillCircle(wx+ww/2+4, wy+6, 3, ColorStars)
		c.FillCircle(wx+ww/2+2, wy+5, 3, skyTop)
	}

	// Twinkling stars & floating library dust motes
	for i := 0; i < 8; i++ {
		sx := wx + 2 + ((int(frameCount/3) + i*7) % (ww - 4))
		sy := wy + 2 + ((int(frameCount/4) + i*11) % (wh - 4))
		if (uint64(i)+frameCount/15)%2 == 0 {
			c.SetPixel(sx, sy, ColorStars)
		}
	}
}

func (s *MidnightLibraryScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Deep Mahogany Wall Paneling
	c.FillRect(0, 0, wx-1, 36, ColorMahoganyBookcase)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorMahoganyBookcase)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Arched Window Frame with Leaded Mullions
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)

	// Full-Height Mahogany Bookshelves on Left (Columns 0 to 18)
	c.FillRect(0, 4, 18, 32, ColorMahoganyBookcase)
	c.DrawRectOutline(0, 4, 18, 32, ColorWoodDark)
	// 3 Tiers of Books with Gilded Spines
	bookCols := []colorful.Color{ColorBookSpine1, ColorBookSpine2, ColorBookSpine3, ColorBookSpine4, ColorStainedGlassRuby}
	for tier := 0; tier < 3; tier++ {
		shelfY := 8 + tier*10
		c.FillRect(1, shelfY+8, 16, 2, ColorWoodMedium)
		for b := 0; b < 6; b++ {
			bx := 2 + b*2
			bh := 6 + (b%3)*1
			col := bookCols[(b+tier*3)%len(bookCols)]
			c.FillRect(bx, shelfY+8-bh, 2, bh, col)
			c.SetPixel(bx, shelfY+8-bh+2, ColorGoldFoilSpine)
		}
	}

	// Full-Height Bookshelves on Right (Columns 56 to Width)
	c.FillRect(56, 4, c.Width-56, 32, ColorMahoganyBookcase)
	c.DrawRectOutline(56, 4, c.Width-56, 32, ColorWoodDark)
	for tier := 0; tier < 3; tier++ {
		shelfY := 8 + tier*10
		c.FillRect(57, shelfY+8, 15, 2, ColorWoodMedium)
		for b := 0; b < 5; b++ {
			bx := 58 + b*3
			bh := 5 + ((b*2)%4)
			col := bookCols[(b+tier*2+2)%len(bookCols)]
			c.FillRect(bx, shelfY+8-bh, 2, bh, col)
			c.SetPixel(bx+1, shelfY+8-bh+2, ColorGoldFoilSpine)
		}
	}

	// Antique Grandfather Clock on left of desk (Columns 18 to 24)
	c.FillRect(18, 12, 6, 24, ColorGrandfatherClock)
	c.DrawRectOutline(18, 12, 6, 24, ColorWoodDark)
	c.FillCircle(21, 15, 2, ColorNotebookPaper)
	c.SetPixel(21, 15, PicoBlack)
	// Animated swinging brass pendulum
	pendulumOffset := int(math.Sin(float64(frameCount)*0.08) * 1.5)
	c.DrawLine(21, 18, 21+pendulumOffset, 28, ColorWoodDark)
	c.FillCircle(21+pendulumOffset, 29, 1, ColorBrassPendulum)

	// Polished Oak Study Desk & Oriental Rug
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	c.FillRect(16, 38, 42, 10, ColorFloorRugBase)
	c.DrawRectOutline(16, 38, 42, 10, ColorFloorRugBorder)
	c.DrawRectOutline(18, 40, 38, 6, ColorFloorRugGold)

	// Solid Oak Desk
	renderDeskHelper(c, 24, 32, 26, lampOn)

	// Emerald Banker's Lamp
	c.FillRect(26, 25, 6, 3, ColorBankerGreen)
	c.FillRect(28, 28, 2, 4, ColorLampBrass)
	c.FillRect(27, 32, 4, 1, ColorLampBrass)
	if lampOn {
		c.SetPixel(29, 27, ColorWarmGlow)
	}

	// Open Parchment Journal & Quill Inkwell
	c.FillRect(36, 29, 6, 3, ColorNotebookPaper)
	c.SetPixel(37, 30, PicoBlack)
	c.SetPixel(39, 30, PicoBlack)
	c.FillRect(44, 29, 2, 3, hex("#1d3557"))
	c.SetPixel(45, 27, PicoWhite)

	// Steaming Earl Grey Mug
	c.FillRect(33, 29, 3, 3, ColorMugCeramic)
	c.SetPixel(33, 29, ColorMugSteam)

	// FocusBuddy companion reading at desk
	if buddy != nil {
		buddy.Draw(c, 26, 20, frameCount)
	}

	// Sleeping Black Cat on velvet rug
	renderCatHelper(c, 44, 38, catBreed, frameCount)
}

func (s *MidnightLibraryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(29, 28, 16, 16, ColorWarmGlow, 0.45)
	}
	if fireOn {
		c.ApplyRadialLight(21, 25, 20.0, ColorFireYellow, 0.30)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 11. LighthouseStudyScene (Coastal Ocean Lighthouse Study)
// ---------------------------------------------------------
type LighthouseStudyScene struct{}

func (s *LighthouseStudyScene) ID() SceneType       { return SceneLighthouseCliff }
func (s *LighthouseStudyScene) Name() string     { return "Coastal Ocean Lighthouse Study 🌊" }
func (s *LighthouseStudyScene) WindowBounds() Rect { return Rect{X: 8, Y: 3, W: 58, H: 28} }

func (s *LighthouseStudyScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#1d3557"), hex("#457b9d")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorOceanDeepBlue, hex("#1e293b")
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh/2+2, skyTop, skyBot)

	// Stars
	if tod == TimeNight && !lightning {
		for i := 0; i < 10; i++ {
			sx := wx + ((int(frameCount/4) + i*13) % ww)
			sy := wy + ((int(frameCount/6) + i*7) % (wh/2 - 2))
			c.SetPixel(sx, sy, ColorStars)
		}
	}

	// Deep Ocean Swell Water (Lower half of window)
	oceanY := wy + wh/2
	c.DrawVerticalGradient(wx, oceanY, ww, wh/2, ColorOceanWaveMid, ColorOceanDeepBlue)

	// Animated Undulating Waves with Whitecap Crests
	for w := 0; w < 4; w++ {
		waveY := oceanY + 2 + w*3
		for x := wx; x < wx+ww; x++ {
			waveOffset := int(math.Sin(float64(x)*0.2 + float64(frameCount)*0.08 + float64(w)*1.5)*1.2)
			py := waveY + waveOffset
			if py >= oceanY && py < wy+wh {
				if (x+int(frameCount/4))%6 == 0 {
					c.SetPixel(x, py, ColorOceanWaveCrest)
				}
			}
		}
	}

	// Distant Rocky Cliff & Lighthouse on right side
	cliffX := wx + ww - 14
	c.FillRect(cliffX, oceanY-2, 14, wh/2+2, hex("#0f172a"))

	// Whitewashed Lighthouse Tower
	c.FillRect(cliffX+4, oceanY-12, 5, 10, ColorLighthouseWhite)
	c.FillRect(cliffX+4, oceanY-9, 5, 3, ColorLighthouseRed)
	c.FillRect(cliffX+3, oceanY-14, 7, 2, hex("#1e293b"))
	c.FillCircle(cliffX+6, oceanY-15, 2, ColorWarmGlow)

	// Rotating Lighthouse Searchlight Beam across the sky & ocean!
	beamTick := (frameCount / 3) % 40
	if beamTick < 20 {
		beamAngle := (float64(beamTick) / 20.0) * math.Pi
		beamEndX := cliffX + 6 - int(math.Cos(beamAngle)*35)
		beamEndY := (oceanY - 14) + int(math.Sin(beamAngle)*18)
		c.DrawLine(cliffX+6, oceanY-14, beamEndX, beamEndY, ColorLighthouseBeam)
	}
}

func (s *LighthouseStudyScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Weathered Maritime Timber Walls
	c.FillRect(0, 0, wx-1, 36, ColorWoodDark)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWoodDark)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Arched Window Frame with Brass Rivet Accents
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)
	for x := wx + 4; x < wx+ww; x += 10 {
		c.SetPixel(x, wy-1, ColorLampBrass)
		c.SetPixel(x, wy+wh, ColorLampBrass)
	}

	// Teak Wood Plank Floor (Rows 36 to Height)
	c.FillRect(0, 36, c.Width, c.Height-36, hex("#442211"))
	for y := 36; y < c.Height; y += 4 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Heavy Cast-Iron Potbelly Stove on Left (Columns 4 to 12)
	c.FillRect(4, 18, 8, 18, ColorStoveCastIron)
	c.DrawRectOutline(4, 18, 8, 18, hex("#10141f"))
	c.FillRect(7, 2, 2, 16, ColorStoveCastIron)
	if fireOn {
		c.FillRect(6, 26, 4, 4, ColorFireOrange)
		c.SetPixel(7, 27, ColorFireYellow)
	} else {
		c.FillRect(6, 26, 4, 4, hex("#10141f"))
	}

	// Maritime Teak Drafting Table (Columns 18 to 42)
	renderDeskHelper(c, 18, 32, 24, lampOn)

	// Nautical Charts, Sextant & Brass Telescope on Table
	c.FillRect(28, 29, 8, 3, ColorNauticalChart)
	c.SetPixel(30, 30, hex("#1d3557"))
	c.SetPixel(34, 30, hex("#e63946"))
	c.FillRect(38, 26, 2, 5, ColorLampBrass)
	c.DrawLine(38, 26, 42, 22, ColorLampBrass)

	// Steaming Spiced Cider Mug
	c.FillRect(22, 29, 3, 3, ColorMugCeramic)
	c.SetPixel(22, 29, ColorFireOrange)

	// FocusBuddy companion studying navigation
	if buddy != nil {
		buddy.Draw(c, 18, 20, frameCount)
	}

	// Sleeping Tuxedo Cat on knit wool blanket by stove
	c.FillRect(48, 39, 18, 8, ColorPlaidRed)
	c.DrawRectOutline(48, 39, 18, 8, ColorPlaidNavy)
	renderCatHelper(c, 52, 38, catBreed, frameCount)
}

func (s *LighthouseStudyScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(8, 26, 22.0, ColorFireOrange, 0.45)
	}
	if lampOn {
		c.ApplyLightCone(22, 26, 16, 14, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.75)
	}
}

// ---------------------------------------------------------
// 12. PixelBakeryScene (Hearthside Autumn Village Bakery)
// ---------------------------------------------------------
type PixelBakeryScene struct{}

func (s *PixelBakeryScene) ID() SceneType       { return ScenePixelBakery }
func (s *PixelBakeryScene) Name() string     { return "Hearthside Autumn Village Bakery 🥐" }
func (s *PixelBakeryScene) WindowBounds() Rect { return Rect{X: 16, Y: 4, W: 36, H: 24} }

func (s *PixelBakeryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#e76f51"), hex("#f4a261")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Distant Autumn Village Rooftops & Cobblestone
	c.FillRect(wx, wy+wh-8, ww, 8, hex("#4a3525"))
	for h := 0; h < 4; h++ {
		hx := wx + h*9 + 1
		c.FillRect(hx, wy+wh-12, 7, 5, hex("#6f4e37"))
		c.FillRect(hx+2, wy+wh-14, 2, 3, hex("#2b2b2b"))
		smokeY := wy + wh - 15 - int((frameCount/4)%5)
		c.SetPixel(hx+2, smokeY, hex("#e2e8f0"))
	}

	// Golden & Crimson Autumn Maple Foliage
	c.FillCircle(wx+4, wy+wh-14, 5, hex("#d90429"))
	c.FillCircle(wx+ww-5, wy+wh-14, 6, hex("#f77f00"))

	// Drifting falling autumn leaves
	for i := 0; i < 8; i++ {
		lx := wx + ((int(frameCount*1) + i*11) % ww)
		ly := wy + ((int(frameCount/2) + i*7) % (wh - 2))
		leafCol := ColorFireOrange
		if i%2 == 0 {
			leafCol = ColorFireYellow
		}
		c.SetPixel(lx, ly, leafCol)
	}
}

func (s *PixelBakeryScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Rustic Stone Masonry Walls
	c.FillRect(0, 0, wx-1, 36, ColorBrickOvenStone)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorBrickOvenStone)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Arched Rustic Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)

	// Hanging Braids from Ceiling Beams (Garlic, Lavender, Dried Chilies)
	for i := 0; i < 3; i++ {
		gx := 4 + i*4
		c.DrawLine(gx, 0, gx, 8, ColorWoodDark)
		c.SetPixel(gx, 6, ColorGarlicBraid)
		c.SetPixel(gx, 8, ColorGarlicBraid)
	}
	for i := 0; i < 2; i++ {
		lx := 54 + i*6
		c.DrawLine(lx, 0, lx, 8, ColorWoodDark)
		c.SetPixel(lx, 6, ColorLavenderDried)
		c.SetPixel(lx, 8, ColorLavenderDried)
	}

	// Rustic Wood-Fired Brick Baking Oven on Right (Columns 50 to Width)
	renderFireplaceHelper(c, 50, 10, 22, 28, fireOn, frameCount)
	// Whistling Copper Tea Kettle on top of oven
	c.FillRect(54, 8, 4, 3, ColorCopperKettle)
	c.SetPixel(53, 9, ColorLampBrass)
	if fireOn {
		c.SetPixel(52, 7, ColorMugSteam)
	}

	// Baker's Wooden Workbench in Center (Columns 12 to 38)
	renderDeskHelper(c, 12, 32, 26, lampOn)

	// Freshly Baked Golden Sourdough Boules & Braided Bread on Baker's Bench
	c.FillCircle(26, 29, 3, ColorBreadCrustGold)
	c.SetPixel(26, 28, ColorBreadFlourDust)
	c.FillRect(31, 28, 6, 3, ColorBreadCrustGold)
	c.SetPixel(33, 28, ColorBreadFlourDust)
	c.FillRect(20, 29, 4, 2, ColorWoodHighlight)

	// Steaming Hot Cocoa Mug
	c.FillRect(14, 29, 3, 3, ColorMugCeramic)
	c.SetPixel(14, 29, ColorMarshmallow)

	// FocusBuddy companion baking
	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	// Warm Hardwood Bakery Floor
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	for y := 36; y < c.Height; y += 4 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Sleeping Tabby Cat curled on Burlap Grain Sack
	c.FillRect(38, 39, 14, 7, hex("#c2b280"))
	renderCatHelper(c, 40, 38, catBreed, frameCount)
}

func (s *PixelBakeryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(61, 24, 24.0, ColorFireOrange, 0.45)
	}
	if lampOn {
		c.ApplyLightCone(16, 26, 16, 14, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 13. RetroLofiTrainScene (Vintage Lofi Train Carriage)
// ---------------------------------------------------------
type RetroLofiTrainScene struct{}

func (s *RetroLofiTrainScene) ID() SceneType       { return SceneRetroLofiTrain }
func (s *RetroLofiTrainScene) Name() string     { return "Vintage Lofi Train Carriage 🚂" }
func (s *RetroLofiTrainScene) WindowBounds() Rect { return Rect{X: 12, Y: 4, W: 50, H: 26} }

func (s *RetroLofiTrainScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#3a7ca5"), hex("#81c3d7")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Stars
	if tod == TimeNight && !lightning {
		for i := 0; i < 8; i++ {
			sx := wx + ((int(frameCount/3) + i*9) % ww)
			sy := wy + ((int(frameCount/5) + i*7) % (wh/2 - 1))
			c.SetPixel(sx, sy, ColorStars)
		}
	}

	// Parallax Layer 1: Distant Rolling Mountain Ridges (scrolling slowly)
	scroll1 := int((frameCount / 6) % uint64(ww))
	for x := 0; x < ww; x++ {
		mountainY := wy + wh - 12 + int(math.Sin(float64(x+scroll1)*0.08)*3.0)
		for y := mountainY; y < wy+wh; y++ {
			c.SetPixel(wx+x, y, hex("#1e293b"))
		}
	}

	// Parallax Layer 2: Passing Utility / Telegraph Poles & Wire Spans
	poleSpacing := 24
	scroll2 := int((frameCount * 2) % uint64(poleSpacing*4))
	for p := -1; p < 4; p++ {
		px := wx + p*poleSpacing - (scroll2 % poleSpacing)
		if px >= wx && px < wx+ww {
			// Telegraph Pole Timber
			c.FillRect(px, wy+6, 2, wh-6, ColorTrainTelegraph)
			c.FillRect(px-2, wy+8, 6, 1, ColorTrainTelegraph) // crossarm
			c.SetPixel(px-2, wy+7, ColorTrainBrass)
			c.SetPixel(px+3, wy+7, ColorTrainBrass)
		}
	}
	// Bouncing Power Lines connecting poles
	for x := wx; x < wx+ww; x++ {
		sag := int(math.Abs(math.Sin(float64(x+scroll2)*0.13)) * 2.0)
		wireY := wy + 8 + sag
		if wireY >= wy && wireY < wy+wh {
			c.SetPixel(x, wireY, hex("#475569"))
		}
	}

	// Parallax Layer 3: Passing Railway Embankment Trees / Hedges (rapid scroll)
	scroll3 := int((frameCount * 3) % uint64(ww))
	for t := 0; t < 6; t++ {
		tx := wx + ((t*10 + scroll3) % ww)
		c.FillCircle(tx, wy+wh-3, 3, hex("#0f2411"))
	}

	// Rain Streaks on the moving carriage window
	for i := 0; i < 6; i++ {
		dropX := wx + int((frameCount*4+uint64(i)*13)%uint64(ww-2))
		dropY := wy + int((frameCount*2+uint64(i)*7)%uint64(wh-4))
		c.SetPixel(dropX, dropY, ColorRainDrop)
		c.SetPixel(dropX+1, dropY+1, ColorRainSplash)
	}
}

func (s *RetroLofiTrainScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Mahogany Carriage Interior Walls
	c.FillRect(0, 0, wx-1, 36, ColorWoodDark)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWoodDark)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Arched Carriage Window Frame with Brass Rivets
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodMedium)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)
	for x := wx + 4; x < wx+ww; x += 8 {
		c.SetPixel(x, wy-1, ColorTrainBrass)
		c.SetPixel(x, wy+wh, ColorTrainBrass)
	}

	// Overhead Polished Brass Luggage Rack (Columns 12 to 62, Rows 0 to 4)
	c.FillRect(12, 1, 50, 2, ColorTrainBrass)
	for x := 12; x < 62; x += 4 {
		c.SetPixel(x, 2, ColorWoodDark)
	}
	// Vintage Leather Traveling Valise & Strapped Umbrella on rack
	c.FillRect(20, 0, 10, 3, ColorTrainLuggage)
	c.SetPixel(22, 1, ColorTrainBrass)
	c.SetPixel(27, 1, ColorTrainBrass)
	c.FillRect(36, 1, 8, 1, hex("#1e293b")) // rolled umbrella

	// Deep Burgundy Velvet Tufted Carriage Seat (Right side, Columns 46 to Width)
	c.FillRect(46, 22, c.Width-46, 26, ColorTrainVelvet)
	c.DrawRectOutline(46, 22, c.Width-46, 26, ColorTrainVelvetDark)
	// Tufted button pattern on seat back
	for y := 24; y < 34; y += 4 {
		for x := 48; x < c.Width-4; x += 6 {
			c.SetPixel(x, y, ColorTrainVelvetDark)
			c.SetPixel(x+1, y, ColorTrainBrass)
		}
	}

	// Polished Teak Window Folding Table (Columns 14 to 44)
	renderDeskHelper(c, 14, 32, 28, lampOn)

	// Stainless Steel / Ceramic Thermos Mug with Rising Steam
	c.FillRect(36, 28, 4, 5, ColorMugCeramic)
	c.FillRect(35, 29, 1, 3, ColorLampBrass)
	c.SetPixel(37, 28, ColorMugSteam)

	// Open Traveler's Journal & Fountain Pen
	c.FillRect(22, 29, 8, 3, ColorNotebookPaper)
	c.SetPixel(24, 30, PicoBlack)
	c.SetPixel(27, 30, PicoBlack)
	c.SetPixel(31, 29, ColorLampBrass) // pen

	// Vintage Portable Cassette Player with Animated Spinning Tape Reels
	c.FillRect(14, 28, 7, 4, ColorCassetteTape)
	c.DrawRectOutline(14, 28, 7, 4, hex("#10141f"))
	// Dual spinning reels
	reelTick := (frameCount / 3) % 4
	reelColors := []colorful.Color{ColorCassetteReel, hex("#94a3b8")}
	c.SetPixel(16, 29, reelColors[reelTick%2])
	c.SetPixel(18, 29, reelColors[(reelTick+1)%2])
	// Blinking green battery indicator LED
	c.SetPixel(20, 29, PicoLime)

	// FocusBuddy companion looking out the window / studying
	if buddy != nil {
		buddy.Draw(c, 14, 20, frameCount)
	}

	// Carriage Carpeted Floor (Rows 36 to Height)
	c.FillRect(0, 36, c.Width, c.Height-36, hex("#311811"))
	for y := 36; y < c.Height; y += 4 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Sleeping Cat companion curled comfortably on velvet seat
	renderCatHelper(c, 50, 38, catBreed, frameCount)
}

func (s *RetroLofiTrainScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(24, 24, 16, 14, ColorWarmGlow, 0.40)
	}
	if fireOn {
		c.ApplyRadialLight(40, 24, 22.0, ColorWarmGlow, 0.30)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 14. ZenOnsenScene (Moonlit Mountain Hot Spring Onsen)
// ---------------------------------------------------------
type ZenOnsenScene struct{}

func (s *ZenOnsenScene) ID() SceneType       { return SceneZenWaterfallShrine }
func (s *ZenOnsenScene) Name() string     { return "Moonlit Mountain Hot Spring Onsen ♨️" }
func (s *ZenOnsenScene) WindowBounds() Rect { return Rect{X: 6, Y: 2, W: 62, H: 30} }

func (s *ZenOnsenScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#1d3557"), hex("#457b9d")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh/2+4, skyTop, skyBot)

	// Distant snowcapped Mount Fuji peak in center
	fujiApexX := wx + ww/2
	fujiApexY := wy + 6
	c.DrawLine(fujiApexX, fujiApexY, fujiApexX-14, wy+wh/2+4, hex("#1b2838"))
	c.DrawLine(fujiApexX, fujiApexY, fujiApexX+14, wy+wh/2+4, hex("#1b2838"))
	// Snowcap on Fuji peak
	c.FillCircle(fujiApexX, fujiApexY+2, 3, PicoWhite)

	// Crescent Moon
	if tod == TimeNight && !lightning {
		c.FillCircle(wx+ww-10, wy+5, 3, ColorStars)
		c.FillCircle(wx+ww-12, wy+4, 3, skyTop)
	}

	// Twinkling stars
	if tod == TimeNight && !lightning {
		for i := 0; i < 8; i++ {
			sx := wx + ((int(frameCount/3) + i*11) % ww)
			sy := wy + ((int(frameCount/5) + i*7) % (wh/2 - 2))
			c.SetPixel(sx, sy, ColorStars)
		}
	}

	// Layered Pine & Bamboo silhouettes in mountain breeze
	for b := 0; b < 6; b++ {
		bx := wx + b*10 + 2
		bSway := int(math.Sin(float64(frameCount)*0.05+float64(b))*1.2)
		c.FillCircle(bx+bSway, wy+wh/2+2, 5, hex("#0f2411"))
	}
}

func (s *ZenOnsenScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	// Open-Air Cedar Veranda Posts & Crossbeams
	c.FillRect(0, 0, 8, 36, ColorHinokiDark)
	c.FillRect(c.Width-8, 0, 8, 36, ColorHinokiDark)
	c.FillRect(0, 0, c.Width, 3, ColorHinokiDark)

	// Steaming Hinoki Cedar Hot Spring Onsen Pool (Columns 8 to 44, Rows 26 to 48)
	// Outer Cedar Rim
	c.FillRect(8, 26, 36, 22, ColorHinokiWood)
	c.DrawRectOutline(8, 26, 36, 22, ColorHinokiDark)

	// Mineral Onsen Water with Undulating Surface Waves
	c.FillRect(10, 28, 32, 18, ColorOnsenWaterDeep)
	for w := 0; w < 3; w++ {
		waveY := 30 + w*5
		for x := 10; x < 42; x++ {
			wOffset := int(math.Sin(float64(x)*0.25+float64(frameCount)*0.08+float64(w))*1.0)
			py := waveY + wOffset
			if py >= 28 && py < 46 {
				c.SetPixel(x, py, ColorOnsenWaterMid)
				if (x+int(frameCount/3))%5 == 0 {
					c.SetPixel(x, py, hex("#a8dadc")) // mineral sparkle
				}
			}
		}
	}

	// Rising Geothermal Thermal Steam Vapors from Onsen
	for i := 0; i < 5; i++ {
		stX := 14 + i*6 + int(math.Sin(float64(frameCount)*0.08+float64(i))*1.5)
		stY := 26 - int((frameCount/2+uint64(i)*7)%12)
		if stY >= 14 && stY < 26 {
			c.SetPixel(stX, stY, ColorOnsenSteam)
			c.SetPixel(stX+1, stY+1, hex("#e2eafc"))
		}
	}

	// Traditional Stone Garden Lantern (Tōrō) on Left (Columns 2 to 7)
	c.FillRect(2, 18, 5, 14, ColorStoneLantern)
	c.DrawRectOutline(2, 18, 5, 14, hex("#2b2b36"))
	c.FillRect(1, 17, 7, 2, ColorStoneLantern) // lantern roof
	// Flickering candle flame inside lantern window
	candleCol := ColorFireYellow
	if (frameCount/10)%2 == 0 {
		candleCol = ColorFireOrange
	}
	c.SetPixel(4, 22, candleCol)

	// Bamboo Shishi-Odoshi Water Spout pouring into onsen
	c.FillRect(6, 24, 6, 2, ColorShojiLattice)
	c.SetPixel(11, 25, hex("#a8dadc"))
	c.SetPixel(11, 26, hex("#a8dadc")) // water drop

	// Mossy River Stones around pool edge
	c.FillCircle(8, 46, 3, ColorMossGreen)
	c.FillCircle(42, 46, 3, ColorMossGreen)

	// Cypress Bath Bucket with Folded White Towel on pool ledge
	c.FillRect(40, 24, 4, 3, ColorHinokiWood)
	c.FillRect(41, 23, 2, 1, PicoWhite) // folded towel

	// Cedar Veranda Deck (Columns 44 to Width, Rows 32 to Height)
	c.FillRect(44, 32, c.Width-44, c.Height-32, ColorTatamiBase)
	for y := 32; y < c.Height; y += 4 {
		c.FillRect(44, y, c.Width-44, 1, ColorTatamiBorder)
	}

	// Whisked Ceramic Matcha Tea Bowl on low cedar tray
	c.FillRect(46, 30, 4, 3, ColorMugCeramic)
	c.SetPixel(47, 30, ColorMatchaGreen)
	c.SetPixel(48, 30, ColorMatchaGreen)

	// FocusBuddy companion relaxing on veranda deck
	if buddy != nil {
		buddy.Draw(c, 44, 20, frameCount)
	}

	// Sleeping Cat curled on woven tatami floor cushion
	renderCatHelper(c, 52, 38, catBreed, frameCount)
}

func (s *ZenOnsenScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyRadialLight(5, 22, 20.0, ColorWarmGlow, 0.45)
	}
	if fireOn {
		c.ApplyRadialLight(26, 34, 22.0, ColorOnsenWaterMid, 0.30)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 15. VinylLoungeScene (Rainy Midnight Vinyl Record Lounge)
// ---------------------------------------------------------
type VinylLoungeScene struct{}

func (s *VinylLoungeScene) ID() SceneType       { return SceneNeonBistro }
func (s *VinylLoungeScene) Name() string     { return "Rainy Midnight Vinyl Lounge 📻" }
func (s *VinylLoungeScene) WindowBounds() Rect { return Rect{X: 18, Y: 4, W: 36, H: 24} }

func (s *VinylLoungeScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#1a2238"), hex("#2a3b5c")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Blurred City Rain Bokeh Lights outside window
	bokehColors := []colorful.Color{hex("#ff7597"), hex("#ffd166"), hex("#74c0fc"), hex("#06d6a0")}
	for b := 0; b < 6; b++ {
		bx := wx + 4 + b*5
		by := wy + 6 + (b*7)%12
		bCol := bokehColors[(b+int(frameCount/30))%len(bokehColors)]
		c.FillCircle(bx, by, 2, bCol)
	}

	// Rain Streaks on window glass
	for i := 0; i < 8; i++ {
		dropX := wx + int((frameCount*2+uint64(i)*11)%uint64(ww-2))
		dropY := wy + int((frameCount*3+uint64(i)*7)%uint64(wh-4))
		c.SetPixel(dropX, dropY, ColorRainDrop)
		c.SetPixel(dropX, dropY+1, ColorRainSplash)
	}
}

func (s *VinylLoungeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Mid-Century Acoustic Slat Walnut Wall Paneling
	c.FillRect(0, 0, wx-1, 36, ColorWoodDark)
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 36, ColorWoodDark)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)
	for x := 0; x < wx-1; x += 3 {
		c.FillRect(x, 0, 1, 36, ColorWoodPlankSep)
	}
	for x := wx + ww + 1; x < c.Width; x += 3 {
		c.FillRect(x, 0, 1, 36, ColorWoodPlankSep)
	}

	// Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodMedium)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)

	// Mid-Century Walnut Hi-Fi Credenza (Columns 10 to 52, Rows 30 to 42)
	renderDeskHelper(c, 10, 32, 38, lampOn)

	// 1. ANIMATED SPINNING VINYL TURNTABLE (Columns 12 to 24)
	c.FillRect(12, 28, 12, 4, ColorTurntableWood)
	c.DrawRectOutline(12, 28, 12, 4, hex("#1a1a24"))
	// Black Vinyl LP Disc
	c.FillCircle(17, 27, 3, ColorVinylBlack)
	// Grooved Vinyl Texture Rings
	c.DrawCircle(17, 27, 2, hex("#262626"))
	// Spinning Colorful Center Label
	labelTick := (frameCount / 2) % 4
	labelCols := []colorful.Color{ColorVinylLabelGold, hex("#ef476f"), hex("#06d6a0"), hex("#118ab2")}
	c.SetPixel(17, 27, labelCols[labelTick])
	// Tonearm in place resting on vinyl groove
	c.DrawLine(22, 24, 18, 27, ColorLampBrass)
	c.SetPixel(22, 24, ColorLampBrass) // pivot

	// Floating Musical Motes when spinning
	if (frameCount/20)%2 == 1 {
		c.DrawGlyph(19, 11, '♪', ColorWarmGlow)
	} else {
		c.DrawGlyph(21, 10, '♫', hex("#ff7597"))
	}

	// 2. GLOWING VACUUM TUBE STEREO AMPLIFIER (Columns 26 to 36)
	c.FillRect(26, 27, 10, 5, hex("#2b2d42"))
	c.DrawRectOutline(26, 27, 10, 5, hex("#10141f"))
	// 2 Glass Vacuum Tubes with Glowing Orange Filaments
	c.FillRect(28, 24, 2, 3, hex("#8d99ae"))
	c.FillRect(32, 24, 2, 3, hex("#8d99ae"))
	tubeGlowCol := ColorVacuumTubeGlow
	if (frameCount/8)%2 == 0 {
		tubeGlowCol = ColorFireYellow
	}
	c.SetPixel(28, 25, tubeGlowCol)
	c.SetPixel(32, 25, tubeGlowCol)
	// Volume Knob & VU Meter
	c.SetPixel(30, 29, ColorLampBrass)
	c.SetPixel(34, 29, PicoLime)

	// 3. Wooden Crate of Vintage Vinyl LPs with colorful spines
	c.FillRect(38, 26, 8, 6, ColorWoodLight)
	c.DrawRectOutline(38, 26, 8, 6, ColorWoodDark)
	bookCols := []colorful.Color{ColorBookSpine1, ColorBookSpine2, ColorBookSpine3, ColorBookSpine4}
	for i := 0; i < 4; i++ {
		c.FillRect(39+i*2, 27, 1, 4, bookCols[i%len(bookCols)])
	}

	// 4. Acoustic Guitar leaning against wall on far right (Columns 54 to 62)
	c.FillEllipse(58, 28, 3, 5, ColorGuitarWood)
	c.FillCircle(58, 28, 1, PicoBlack) // soundhole
	c.FillRect(57, 16, 2, 10, ColorWoodDark) // neck
	c.FillRect(56, 14, 4, 3, ColorWoodDark)  // headstock

	// Steaming Pour-Over Coffee Mug
	c.FillRect(10, 29, 3, 3, ColorMugCeramic)
	c.SetPixel(10, 29, ColorMugSteam)

	// FocusBuddy companion with studio headphones listening & typing
	if buddy != nil {
		buddy.Draw(c, 10, 20, frameCount)
	}

	// Hardwood Herringbone Floor (Rows 36 to Height)
	c.FillRect(0, 36, c.Width, c.Height-36, ColorWoodDark)
	for y := 36; y < c.Height; y += 4 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Sleeping Cat curled beside warm vacuum tube amp
	renderCatHelper(c, 46, 38, catBreed, frameCount)
}

func (s *VinylLoungeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(20, 22, 16, 14, ColorWarmGlow, 0.40)
	}
	if fireOn {
		c.ApplyRadialLight(30, 26, 18.0, ColorVacuumTubeGlow, 0.45)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 16. SunkenHearthScene (Sunken Hearth & Blanket Fort)
// ---------------------------------------------------------
type SunkenHearthScene struct{}

func (s *SunkenHearthScene) ID() SceneType       { return SceneStarlightCamp }
func (s *SunkenHearthScene) Name() string     { return "Sunken Hearth & Blanket Fort 🏕️" }
func (s *SunkenHearthScene) WindowBounds() Rect { return Rect{X: 14, Y: 3, W: 46, H: 23} }

func (s *SunkenHearthScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = ColorSkyDawnTop, ColorSkyDawnBot
	case TimeDay:
		skyTop, skyBot = hex("#1e3a8a"), hex("#60a5fa")
	case TimeSunset:
		skyTop, skyBot = ColorSkySunsetTop, ColorSkySunsetBot
	case TimeNight:
		skyTop, skyBot = ColorSkyNightTop, ColorSkyNightBot
	}
	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Twinkling stars
	if tod == TimeNight && !lightning {
		for i := 0; i < 8; i++ {
			sx := wx + ((int(frameCount/3) + i*11) % ww)
			sy := wy + ((int(frameCount/5) + i*7) % (wh - 2))
			c.SetPixel(sx, sy, ColorStars)
		}
	}

	// Distant Snow-dusted Evergreen Mountains
	for t := 0; t < 5; t++ {
		tx := wx + t*10 + 2
		c.FillCircle(tx, wy+wh-4, 5, hex("#0f172a"))
	}

	// Gentle Drifting Snowflakes
	for i := 0; i < 8; i++ {
		sx := wx + int((frameCount+uint64(i)*9)%uint64(ww-2))
		sy := wy + int((frameCount/2+uint64(i)*5)%uint64(wh-2))
		c.SetPixel(sx, sy, PicoWhite)
	}

	// Sheer Draped Curtains framing window sides
	c.FillRect(wx, wy, 4, wh, hex("#334155"))
	c.FillRect(wx+ww-4, wy, 4, wh, hex("#334155"))
}

func (s *SunkenHearthScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Heavy Timber Ceiling Rafters
	c.FillRect(0, 0, c.Width, 4, ColorWoodDark)
	for x := 0; x < c.Width; x += 12 {
		c.FillRect(x, 0, 2, 6, ColorWoodMedium)
	}

	// Strung Fairy Lights along ceiling rafters with warm twinkle
	if lampOn {
		for x := 4; x < c.Width-4; x += 5 {
			fy := 4 + int(math.Sin(float64(x)*0.3)*1.5)
			lightCol := ColorFairyLightAmber
			if (uint64(x)+frameCount/10)%2 == 0 {
				lightCol = ColorFireYellow
			}
			c.SetPixel(x, fy, lightCol)
		}
	}

	// Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)

	// Sunken Pit Stone Masonry Edge (Rows 30 to Height)
	c.FillRect(0, 30, c.Width, 2, ColorSunkenPitStone)

	// Sunken Floor with Plush Velvet Cushions & Patchwork Quilts
	c.FillRect(0, 32, c.Width, c.Height-32, ColorWoodDark)

	// Oversized Mustard & Teal Floor Cushions on Left
	c.FillEllipse(12, 36, 8, 4, ColorBlanketGold)
	c.FillEllipse(18, 38, 6, 3, ColorBlanketTeal)

	// Oversized Purple & Burgundy Cushions on Right
	c.FillEllipse(56, 36, 8, 4, ColorCushionPlush)
	c.FillEllipse(50, 38, 6, 3, ColorPlaidRed)

	// Central Sunken Stone Hearth / Fire Pit (Columns 26 to 48, Rows 26 to 40)
	c.FillRect(26, 28, 22, 12, ColorSunkenPitStone)
	c.DrawRectOutline(26, 28, 22, 12, hex("#212529"))
	c.FillRect(28, 30, 18, 9, hex("#100806"))

	// Dynamic Multi-Tier Hearth Flames & Embers
	if fireOn {
		flameTick := frameCount % 6
		coreCol := ColorFireCore
		if flameTick%2 == 0 {
			coreCol = ColorFireYellow
		}
		c.FillRect(34, 34, 6, 4, ColorFireOrange)
		c.FillRect(35, 33, 4, 3, coreCol)
		c.SetPixel(36, 31-int(flameTick%3), ColorFireYellow)
		c.SetPixel(38, 32-int((flameTick+1)%3), ColorFireOrange)

		// Suspended Cast-Iron Kettle over Fire Pit on Iron Tripod
		c.DrawLine(30, 26, 37, 30, hex("#1a1a24"))
		c.DrawLine(44, 26, 37, 30, hex("#1a1a24"))
		c.FillRect(35, 29, 4, 3, ColorStoveCastIron)
		c.SetPixel(34, 30, ColorLampBrass) // spout
		c.SetPixel(34, 28, ColorMugSteam)  // steam
	} else {
		c.FillRect(34, 35, 6, 2, hex("#343a40"))
	}

	// Low Chabudai Wooden Table with Teapot and Ceramic Cups
	c.FillRect(18, 32, 10, 3, ColorWoodMedium)
	c.FillRect(18, 32, 10, 1, ColorWoodHighlight)
	c.FillRect(20, 30, 3, 3, ColorMugCeramic)
	c.SetPixel(20, 30, ColorMugSteam)

	// FocusBuddy wrapped in cozy blanket reading by the fire
	if buddy != nil {
		buddy.Draw(c, 10, 20, frameCount)
	}

	// Sleeping Cat curled happily on plush mustard cushion pile
	renderCatHelper(c, 48, 36, catBreed, frameCount)
}

func (s *SunkenHearthScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(37, 34, 26.0, ColorFireOrange, 0.50)
	}
	if lampOn {
		c.ApplyRadialLight(20, 8, 16.0, ColorWarmGlow, 0.25)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}



