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
		id:        SceneCozyLibrary,
		name:      "Cozy Pixel Library 📚",
		gifPath:   "assets/animations/cozy_library.gif",
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
		id:        ScenePixelBakery,
		name:      "Pixel Cozy Bakery 🥐",
		gifPath:   "assets/animations/pixel_bakery.gif",
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
		id:        SceneLighthouseCliff,
		name:      "Lighthouse Cliff 🚨",
		gifPath:   "assets/animations/lighthouse_cliff.gif",
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
		id:        SceneCozyGreenhouse,
		name:      "Cozy Glasshouse Botanical Garden 🌿",
		gifPath:   "assets/animations/cozy_greenhouse.gif",
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

func (s *CityCafeScene) ID() SceneType { return SceneRainyCity }
func (s *CityCafeScene) Name() string { return "Rainy City Lofi Cafe 🏙️" }
func (s *CityCafeScene) WindowBounds() Rect { return Rect{X: 20, Y: 8, W: 120, H: 54} }

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

	for b := 0; b < 12; b++ {
		bx := wx + b*10 + 2
		bw := 9
		bh := 18 + ((b*17+7)%26)
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		// Warm illuminated apartment windows
		if tod == TimeNight || tod == TimeSunset || tod == TimeDawn {
			for wyPos := wy + wh - bh + 3; wyPos < wy+wh-3; wyPos += 4 {
				for wxPos := bx + 2; wxPos < bx+bw-2; wxPos += 2 {
					if (wxPos*3+wyPos*7+b)%5 != 0 {
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
		c.FillRect(wx+8, wy+8, 22, 5, hex("#101424"))
		c.DrawGlyph(wx+10, (wy+9)/2, 'C', neonCol)
		c.DrawGlyph(wx+12, (wy+9)/2, 'A', neonCol)
		c.DrawGlyph(wx+14, (wy+9)/2, 'F', neonCol)
		c.DrawGlyph(wx+16, (wy+9)/2, 'E', neonCol)
	}
}

func (s *CityCafeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Cafe Interior walls
	c.DrawVerticalGradient(0, 0, wx-1, 62, ColorWoodDark, ColorWallWarm)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 62, ColorWoodDark, ColorWallWarm)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Sleek steel window frame with condensation mullions
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWoodDark)

	// Window sill with potted succulent
	c.FillRect(wx, wy+wh-2, ww, 3, ColorWoodMedium)
	renderPottedSucculent(c, wx+8, wy+wh-6)

	// Polished Mahogany Cafe Bar Counter (Rows 60 to c.Height-1)
	c.FillRect(0, 60, c.Width, c.Height-60, ColorWoodDark)
	c.FillRect(0, 60, c.Width, 2, ColorWoodHighlight)
	c.FillRect(0, 62, c.Width, 1, ColorWoodPlankSep)

	// Espresso machine on left counter
	c.FillRect(4, 46, 14, 14, ColorWoodDark)
	c.FillRect(6, 48, 10, 10, hex("#343a40"))
	c.FillRect(8, 44, 4, 3, ColorLampBrass) // brass boiler dome
	c.SetPixel(7, 54, ColorWarmGlow)       // power lamp

	// Focus buddy studying at cafe bar
	if buddy != nil {
		buddy.Draw(c, 52, 34, frameCount)
	}

	// Glowing modern laptop on bar counter
	c.FillRect(82, 48, 24, 12, hex("#212529"))
	c.FillRect(84, 50, 20, 8, hex("#74c0fc"))
	// Animated code editor syntax lines
	for l := 0; l < 3; l++ {
		c.FillRect(86, 52+l*2, 6+l*4, 1, hex("#a6e3a1"))
	}
	c.FillRect(78, 59, 32, 2, hex("#ced4da")) // keyboard base

	// Steaming Latte / Matcha Mug
	c.FillRect(74, 53, 5, 6, ColorMugCeramic)
	c.SetPixel(75, 53, ColorMugSteam) // latte foam dot

	// Sleeping cafe cat curled on right bar stool
	renderCatHelper(c, 126, 64, catBreed, frameCount)
}

func (s *CityCafeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(86, 46, 28, 24, ColorWarmGlow, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 2. ForestCabinScene (16-Bit Forest Sanctuary Cabin)
// ---------------------------------------------------------
type ForestCabinScene struct{}

func (s *ForestCabinScene) ID() SceneType { return SceneForestCabin }
func (s *ForestCabinScene) Name() string { return "Forest Sanctuary Cabin 🌲" }
func (s *ForestCabinScene) WindowBounds() Rect { return Rect{X: 24, Y: 8, W: 112, H: 52} }

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
		for i := 0; i < 9; i++ {
			tx := wx + i*13 + (layer*6)
			th := 22 + ((i*11+layer*7)%18)
			// Trunk
			c.FillRect(tx+3, wy+wh-th, 3, th, ColorWoodDark)
			// Layered needle tiers
			for t := 0; t < 4; t++ {
				tierY := wy + wh - th + (t * 4)
				tierW := 4 + (t * 2)
				c.FillRect(tx+4-tierW/2, tierY, tierW, 3, treeCol)
			}
		}
	}
}

func (s *ForestCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Log cabin walls with timber log texture
	c.DrawVerticalGradient(0, 0, wx-1, 60, ColorWoodDark, ColorWallPlank)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 60, ColorWoodDark, ColorWallPlank)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)
	for y := 4; y < 60; y += 8 {
		c.FillRect(0, y, wx-1, 1, ColorWoodPlankSep)
		c.FillRect(wx+ww+1, y, c.Width-(wx+ww+1), 1, ColorWoodPlankSep)
	}

	// Timber window frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWoodDark)
	c.FillRect(wx+ww/3, wy, 2, wh, ColorWoodDark)
	c.FillRect(wx+(ww*2)/3, wy, 2, wh, ColorWoodDark)

	// Wood floor planks
	c.FillRect(0, 60, c.Width, c.Height-60, ColorWoodDark)
	for y := 60; y < c.Height; y += 7 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
	}

	// Large Monstera plant in terracotta pot on left
	renderPottedMonstera(c, 6, 38)

	// Focus buddy at rustic wood table
	if buddy != nil {
		buddy.Draw(c, 34, 34, frameCount)
	}

	// Rustic writing table
	c.FillRect(22, 52, 44, 4, ColorWoodMedium)
	c.FillRect(22, 52, 44, 1, ColorWoodHighlight)
	c.FillRect(24, 56, 4, 16, ColorWoodDark)
	c.FillRect(60, 56, 4, 16, ColorWoodDark)

	// Steaming mug & lantern on table
	c.FillRect(52, 46, 5, 6, ColorMugCeramic)
	if lampOn {
		c.FillRect(26, 42, 6, 9, ColorLampBrass)
		c.FillRect(27, 44, 4, 5, ColorWarmGlow)
	}

	// Crackling wood stove on right
	renderWoodStoveHelper(c, 114, 28, 34, 38, fireOn, frameCount)

	// Sleeping cat by the wood stove
	renderCatHelper(c, 84, 66, catBreed, frameCount)
}

func (s *ForestCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(29, 44, 26, 22, ColorWarmGlow, 0.35)
	}
	if fireOn {
		c.ApplyRadialLight(130, 46, 28.0, ColorFireOrange, 0.38)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 3. ClassicCottageScene (Handcrafted Virtual Cottage)
// ---------------------------------------------------------
type ClassicCottageScene struct{}

func (s *ClassicCottageScene) ID() SceneType { return SceneCottageClassic }
func (s *ClassicCottageScene) Name() string { return "Classic Cozy Cottage 🏡" }
func (s *ClassicCottageScene) WindowBounds() Rect { return Rect{X: 42, Y: 8, W: 76, H: 48} }

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
	c.FillRect(wx, wy+wh-10, ww, 10, hillCol)
	for i := 0; i < 7; i++ {
		px := wx + i*11 + 3
		ph := 6 + (i*5)%8
		c.FillRect(px, wy+wh-10-ph, 3, ph, hillCol)
	}

	// Celestial objects outside window
	if tod == TimeNight && !lightning {
		// Crescent Moon
		c.FillRect(wx+ww-14, wy+6, 8, 6, ColorStars)
		c.FillRect(wx+ww-17, wy+6, 5, 6, skyTop)
		// Stars
		for _, st := range stars {
			px := wx + (st.X % ww)
			py := wy + (st.Y % (wh - 12))
			twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
			if twinkle > 0.3 {
				c.SetPixel(px, py, ColorStars)
			}
		}
	} else if tod == TimeDawn {
		// Glowing morning star
		c.DrawGlyph(wx+12, (wy+8)/2, '✦', ColorStars)
	} else if tod == TimeDay {
		// Soft golden sun
		c.FillRect(wx+ww-14, wy+6, 8, 6, ColorFireYellow)
	}
}

func (s *ClassicCottageScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// Cozy wallpaper on left and right walls
	c.DrawVerticalGradient(0, 0, wx-1, 58, ColorWallPaper, ColorWallWarm)
	c.DrawVerticalGradient(wx+ww+1, 0, c.Width-(wx+ww+1), 58, ColorWallPaper, ColorWallWarm)
	c.FillRect(0, 0, c.Width, wy-1, ColorWoodDark)

	// Ornate Arched Timber Window Frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorWindowFrame)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorWindowFrame)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorWindowFrame)
	// Window sill with potted trailing plant
	c.FillRect(wx-2, wy+wh-2, ww+4, 3, ColorWindowSill)
	renderPottedSucculent(c, wx+ww-12, wy+wh-6)

	// Rich Wooden Floor Planks (Rows 58 to c.Height-1)
	c.FillRect(0, 58, c.Width, c.Height-58, ColorWoodDark)
	for y := 58; y < c.Height; y += 7 {
		c.FillRect(0, y, c.Width, 1, ColorWoodPlankSep)
		c.FillRect(0, y+1, c.Width, 1, ColorWoodMedium)
	}

	// Persian Ornate Floor Rug under desk & fireplace
	rugX, rugY, rugW, rugH := 36, 64, 88, 22
	c.FillRect(rugX, rugY, rugW, rugH, ColorFloorRugBase)
	c.DrawRectOutline(rugX, rugY, rugW, rugH, ColorFloorRugBorder)
	for rx := rugX + 6; rx < rugX+rugW-6; rx += 10 {
		c.FillRect(rx, rugY+rugH/2-2, 4, 4, ColorFloorRugGold)
	}

	// Left Wall: Bookshelf with colorful stacked books & trailing ivy
	renderBookshelfHelper(c, 2, 8, 20, 50)

	// Study Area: Oak Desk with Lamp & Focus Companion
	renderDeskHelper(c, 22, 50, 40, lampOn)

	if buddy != nil {
		buddy.Draw(c, 24, 32, frameCount)
	}

	// Right Wall: Multi-course Brick Fireplace Hearth
	renderFireplaceHelper(c, 114, 24, 42, 42, fireOn, frameCount)

	// Sleeping Cat Companion on Velvet Cushion
	renderCatHelper(c, 74, 66, catBreed, frameCount)
}

func (s *ClassicCottageScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(26, 44, 28, 24, ColorWarmGlow, 0.38)
	}
	if fireOn {
		c.ApplyRadialLight(135, 46, 32.0, ColorFireOrange, 0.40)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 4. ObservatoryScene (Starlight Celestial Observatory)
// ---------------------------------------------------------
type ObservatoryScene struct{}

func (s *ObservatoryScene) ID() SceneType { return SceneObservatory }
func (s *ObservatoryScene) Name() string { return "Celestial Observatory 🌌" }
func (s *ObservatoryScene) WindowBounds() Rect { return Rect{X: 20, Y: 6, W: 120, H: 54} }

func (s *ObservatoryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, ColorSkyNightTop, ColorSkyNightBot)

	// Swirling cosmic nebula dust
	nebulaCols := []colorful.Color{hex("#3d1a45"), hex("#1a2b4c"), hex("#153b3d")}
	for n := 0; n < 3; n++ {
		nx := wx + 18 + n*30
		ny := wy + 8 + n*10
		c.ApplyRadialLight(nx, ny, 16.0, nebulaCols[n], 0.25)
	}

	// Bright Crescent Moon
	c.FillRect(wx+84, wy+8, 12, 10, ColorStars)
	c.FillRect(wx+80, wy+8, 8, 10, ColorSkyNightTop)

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
	c.FillRect(0, 0, wx-1, 60, hex("#10141f"))
	c.FillRect(wx+ww+1, 0, c.Width-(wx+ww+1), 60, hex("#10141f"))
	c.FillRect(0, 0, c.Width, wy-1, hex("#10141f"))

	// Domed brass skylight frame
	c.DrawRectOutline(wx-1, wy-1, ww+2, wh+2, ColorLampBrass)
	c.FillRect(wx+ww/2-1, wy, 2, wh, ColorLampBrass)
	c.FillRect(wx, wy+wh/2-1, ww, 2, ColorLampBrass)

	// Observatory dark stone floor
	c.FillRect(0, 60, c.Width, c.Height-60, hex("#161b26"))
	c.FillRect(0, 60, c.Width, 2, ColorLampBrass)

	// Large Brass Telescope on heavy wooden tripod facing skylight
	telX, telY := 96, 28
	c.FillRect(telX, telY, 24, 6, ColorLampBrass) // brass tube
	c.FillRect(telX+20, telY-2, 6, 10, ColorLampBrass) // objective lens
	c.FillRect(telX-2, telY+2, 4, 4, hex("#212529")) // eyepiece
	// Tripod legs
	c.FillRect(telX+10, telY+6, 3, 30, ColorWoodDark)
	c.FillRect(telX+4, telY+6, 3, 30, ColorWoodDark)
	c.FillRect(telX+16, telY+6, 3, 30, ColorWoodDark)

	// Focus Buddy at star chart desk on left
	if buddy != nil {
		buddy.Draw(c, 30, 34, frameCount)
	}

	// Star Chart Table
	c.FillRect(20, 52, 40, 4, ColorWoodMedium)
	c.FillRect(20, 52, 40, 1, ColorWoodHighlight)
	c.FillRect(24, 56, 4, 16, ColorWoodDark)
	c.FillRect(54, 56, 4, 16, ColorWoodDark)

	// Celestial map & brass lantern
	c.FillRect(36, 50, 14, 3, ColorNotebookPaper)
	if lampOn {
		c.FillRect(24, 42, 6, 9, ColorLampBrass)
		c.FillRect(25, 44, 4, 5, ColorWarmGlow)
	}

	// Sleeping cat by observation desk
	renderCatHelper(c, 70, 66, catBreed, frameCount)
}

func (s *ObservatoryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(27, 44, 26, 22, ColorWarmGlow, 0.35)
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
		sy := by + 14 + s*13
		c.FillRect(bx+1, sy, bw-2, 2, ColorWoodMedium)

		// Stacked and standing books on each shelf
		for b := 0; b < 6; b++ {
			bX := bx + 2 + b*3
			bH := 6 + ((b*3+s*5)%5)
			bCol := bookColors[(b+s)%len(bookColors)]
			c.FillRect(bX, sy-bH, 2, bH, bCol)
		}
	}

	// Top shelf: Terracotta pot with trailing ivy vines
	c.FillRect(bx+4, by-2, 8, 4, ColorPotClay)
	c.FillRect(bx+5, by-3, 6, 1, ColorLeafLight)
	// Trailing green ivy vines dripping down the side
	c.SetPixel(bx+2, by+2, ColorLeafMedium)
	c.SetPixel(bx+2, by+4, ColorLeafLight)
	c.SetPixel(bx+1, by+6, ColorLeafMedium)
	c.SetPixel(bx+1, by+8, ColorLeafDark)
	c.SetPixel(bx, by+10, ColorLeafMedium)
}

// renderPottedSucculent draws a tiny terracotta pot with green leaves.
func renderPottedSucculent(c *Canvas, px, py int) {
	c.FillRect(px+1, py+2, 4, 3, ColorPotClay)
	c.FillRect(px, py+2, 6, 1, ColorPotClay)
	c.SetPixel(px+2, py+1, ColorLeafLight)
	c.SetPixel(px+3, py, ColorLeafMedium)
	c.SetPixel(px+1, py, ColorLeafDark)
	c.SetPixel(px+4, py+1, ColorLeafMedium)
}

// renderPottedMonstera draws a large potted monstera plant with split leaves.
func renderPottedMonstera(c *Canvas, px, py int) {
	c.FillRect(px+3, py+14, 10, 8, ColorPotClay)
	c.FillRect(px+2, py+14, 12, 2, ColorPotClay)

	// Broad green monstera leaves
	c.FillRect(px, py+4, 6, 6, ColorLeafDark)
	c.FillRect(px+8, py+2, 7, 7, ColorLeafMedium)
	c.FillRect(px+4, py, 6, 5, ColorLeafLight)
	// Leaf cutout details
	c.SetPixel(px+2, py+6, ColorWallPlank)
	c.SetPixel(px+10, py+5, ColorWallPlank)
}

// renderDeskHelper draws an oak writing desk with green banker's lamp, journal, and mug.
func renderDeskHelper(c *Canvas, dx, dy, dw int, lampOn bool) {
	// Oak desk tabletop
	c.FillRect(dx, dy, dw, 4, ColorWoodMedium)
	c.FillRect(dx, dy, dw, 1, ColorWoodHighlight)
	// Desk drawer & legs
	c.FillRect(dx+2, dy+4, dw-4, 4, ColorWoodDark)
	c.FillRect(dx+dw/2-2, dy+5, 4, 1, ColorLampBrass) // brass drawer pull
	c.FillRect(dx+2, dy+8, 3, 16, ColorWoodDark)
	c.FillRect(dx+dw-5, dy+8, 3, 16, ColorWoodDark)

	// Classic Banker's Lamp (Brass neck + Emerald Green glass shade)
	if lampOn {
		c.FillRect(dx+4, dy-2, 4, 2, ColorLampBrass)
		c.FillRect(dx+5, dy-8, 2, 6, ColorLampBrass)
		c.FillRect(dx+2, dy-12, 8, 4, ColorLampShade)
		c.FillRect(dx+3, dy-9, 6, 1, ColorWarmGlow)
	} else {
		c.FillRect(dx+4, dy-2, 4, 2, hex("#495057"))
		c.FillRect(dx+5, dy-8, 2, 6, hex("#495057"))
		c.FillRect(dx+2, dy-12, 8, 4, hex("#212529"))
	}

	// Open Notebook Journal with scribbled lines
	c.FillRect(dx+16, dy-1, 10, 3, ColorNotebookPaper)
	c.SetPixel(dx+18, dy, ColorWoodDark)
	c.SetPixel(dx+20, dy, ColorWoodDark)
	c.SetPixel(dx+22, dy, ColorWoodDark)

	// Steaming Ceramic Terracotta Mug
	c.FillRect(dx+28, dy-4, 5, 5, ColorMugCeramic)
	c.SetPixel(dx+33, dy-3, ColorMugCeramic) // handle
}

// renderFireplaceHelper draws a multi-course brick hearth with mantle, candle, and animated flame engine.
func renderFireplaceHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	// Brick chimney masonry
	c.FillRect(fx, fy, fw, fh, ColorHearthBrick)
	// Brick mortar lines
	for by := fy + 3; by < fy+fh; by += 4 {
		c.FillRect(fx, by, fw, 1, ColorHearthMortar)
		offset := ((by - fy) / 4) % 2 * 4
		for bx := fx + offset; bx < fx+fw; bx += 8 {
			c.FillRect(bx, by-3, 1, 3, ColorHearthMortar)
		}
	}

	// Timber Mantlepiece at top of fireplace
	c.FillRect(fx-2, fy, fw+4, 3, ColorWoodDark)
	c.FillRect(fx-2, fy, fw+4, 1, ColorWoodHighlight)

	// Candle on mantlepiece
	c.FillRect(fx+6, fy-5, 3, 5, ColorCandleWax)
	if fireOn {
		c.SetPixel(fx+7, fy-6, ColorFireYellow)
	}

	// Arched firebox interior
	boxX, boxY, boxW, boxH := fx+6, fy+10, fw-12, fh-12
	c.FillRect(boxX, boxY, boxW, boxH, ColorHearthInterior)
	c.DrawRectOutline(boxX, boxY, boxW, boxH, ColorHearthBrickDark)

	if !fireOn {
		// Cold ash bed
		c.FillRect(boxX+4, boxY+boxH-3, boxW-8, 2, hex("#343a40"))
		return
	}

	// Hot glowing charcoal / ember bed
	c.FillRect(boxX+4, boxY+boxH-3, boxW-8, 2, ColorFireEmber)

	// Multi-layered animated dancing flames
	fireCenter := boxX + boxW/2
	fireTick := (frameCount / 3) % 6

	for i := -6; i <= 6; i++ {
		px := fireCenter + i
		flameH := int(7.0 + 4.0*math.Sin(float64(i)*0.8+float64(fireTick)*1.2))

		for fyPos := 0; fyPos < flameH; fyPos++ {
			py := boxY + boxH - 3 - fyPos
			if py < boxY {
				continue
			}
			if fyPos < 2 {
				c.SetPixel(px, py, ColorFireCore)
			} else if fyPos < 5 {
				c.SetPixel(px, py, ColorFireYellow)
			} else if fyPos < 8 {
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
	c.FillRect(fx+4, fy+10, fw-8, fh-14, hex("#212529"))
	c.DrawRectOutline(fx+4, fy+10, fw-8, fh-14, hex("#343a40"))
	// Stove legs
	c.FillRect(fx+6, fy+fh-4, 3, 4, hex("#10141f"))
	c.FillRect(fx+fw-9, fy+fh-4, 3, 4, hex("#10141f"))
	// Stove pipe chimney
	c.FillRect(fx+fw/2-2, fy, 4, 10, hex("#212529"))

	// Glowing stove window
	if fireOn {
		c.FillRect(fx+10, fy+16, fw-20, 8, ColorFireOrange)
		c.FillRect(fx+12, fy+18, fw-24, 4, ColorFireYellow)
	} else {
		c.FillRect(fx+10, fy+16, fw-20, 8, hex("#10141f"))
	}
}

// renderCatHelper draws an organic, cute sleeping cat companion with smooth micro-animations.
func renderCatHelper(c *Canvas, cx, cy int, breed CatBreed, frameCount uint64) {
	// Plush Indigo Velvet Cushion
	c.FillRect(cx-4, cy+6, 24, 6, ColorCatCushion)
	c.DrawRectOutline(cx-4, cy+6, 24, 6, hex("#2a2356"))

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
	c.FillRect(cx, cy+breathOffset, 14, 7, furMain)
	c.FillRect(cx+2, cy+1+breathOffset, 10, 5, furDetail)

	// Head tucked in
	c.FillRect(cx-2, cy+1+breathOffset, 5, 5, furMain)
	// Cute inner ear pink details
	c.SetPixel(cx-2, cy+breathOffset, hex("#ff7597"))
	c.SetPixel(cx+1, cy+breathOffset, hex("#ff7597"))

	// Dynamic Tail Wag
	tailWag := int((frameCount / 18) % 3)
	c.FillRect(cx+14, cy+3+breathOffset-tailWag, 3, 2, furMain)
	c.FillRect(cx+16, cy+1+breathOffset-tailWag*2, 2, 3, furMain)

	// Purr & Sleep glyphs
	if (frameCount/24)%2 == 1 {
		c.DrawGlyph(cx+5, (cy-2)/2, '♥', hex("#ff7597"))
	} else if (frameCount/14)%3 == 1 {
		c.DrawGlyph(cx-3, (cy-3)/2, 'z', hex("#74c0fc"))
	}
}

