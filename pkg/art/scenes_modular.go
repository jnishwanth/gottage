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
func (s *CityCafeScene) Name() string { return "16-Bit Rainy City Cafe 🏙️" }
func (s *CityCafeScene) WindowBounds() Rect { return Rect{X: 20, Y: 8, W: 120, H: 58} }

func (s *CityCafeScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	c.DrawVerticalGradient(0, 0, c.Width, 66, PicoNavy, PicoBlack)

	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoDarkPurple, PicoOrange
	case TimeDay:
		skyTop, skyBot = PicoNavy, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoRed
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// 16-Bit Skyscrapers Silhouettes
	bldgColors := []colorful.Color{PicoBlack, PicoNavy, PicoDarkPurple}
	neonColors := []colorful.Color{PicoPink, PicoCyan, PicoYellow, PicoLime}

	for b := 0; b < 10; b++ {
		bx := wx + b*12 + 2
		bw := 10
		bh := 24 + (b*13)%28
		c.FillRect(bx, wy+wh-bh, bw, bh, bldgColors[b%3])

		if tod == TimeNight || tod == TimeSunset {
			neonCol := neonColors[(b+int(frameCount/15))%4]
			for wyPos := wy + wh - bh + 4; wyPos < wy+wh-4; wyPos += 6 {
				for wxPos := bx + 2; wxPos < bx+bw-2; wxPos += 3 {
					if (wxPos+wyPos)%2 == 0 {
						c.SetPixel(wxPos, wyPos, neonCol)
					}
				}
			}
		}
	}

	// 16-Bit Neon Sign
	if tod == TimeNight || tod == TimeSunset {
		neonSignCol := neonColors[int(frameCount/10)%4]
		c.DrawGlyph(wx+8, wy+10, 'L', neonSignCol)
		c.DrawGlyph(wx+10, wy+10, 'O', neonSignCol)
		c.DrawGlyph(wx+12, wy+10, '-', neonSignCol)
		c.DrawGlyph(wx+14, wy+10, 'F', neonSignCol)
		c.DrawGlyph(wx+16, wy+10, 'I', neonSignCol)
	}
}

func (s *CityCafeScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	// 16-Bit Window Frame (Renders ON TOP of rain drops!)
	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoDarkGrey)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoDarkGrey)

	// Sleek 16-bit Bar Counter Table (Rows 66 to Height-1)
	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoPeach)

	// 16-Bit Focus Companion at Desk (Renders ON TOP of rain drops!)
	if buddy != nil {
		buddy.Draw(c, 60, 40, frameCount)
	}

	// Glowing 16-Bit Laptop on Bar Counter
	c.FillRect(88, 52, 28, 12, PicoDarkGrey)
	c.FillRect(92, 54, 20, 8, PicoCyan)
	for l := 0; l < 4; l++ {
		c.FillRect(94, 56+l*2, 6+l*3, 1, PicoLime)
	}
	c.FillRect(84, 64, 36, 4, PicoLightGrey)

	// Coffee Mug on Bar Counter
	c.FillRect(80, 58, 6, 6, PicoRed)
}

func (s *CityCafeScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(92, 52, 32, 24, PicoYellow, 0.35)
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
func (s *ForestCabinScene) Name() string { return "16-Bit Forest Cabin 🌲" }
func (s *ForestCabinScene) WindowBounds() Rect { return Rect{X: 24, Y: 10, W: 112, H: 54} }

func (s *ForestCabinScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoDarkPurple, PicoOrange
	case TimeDay:
		skyTop, skyBot = PicoGreen, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoPeach
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	// Pine Trees
	for i := 0; i < 8; i++ {
		tx := wx + i*14 + 4
		th := 20 + (i*9)%18
		c.FillRect(tx, wy+wh-th, 10, th, PicoBlack)
		c.FillRect(tx+2, wy+wh-th-8, 6, 8, PicoGreen)
	}
}

func (s *ForestCabinScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(0, 0, wx-2, 66, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoBrown, PicoDarkGrey)
	c.DrawVerticalGradient(0, 0, c.Width, wy-2, PicoBrown, PicoDarkGrey)
	for y := 0; y < 66; y += 10 {
		c.FillRect(0, y, c.Width, 2, PicoBlack)
	}

	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoDarkGrey)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoBrown)
	c.FillRect(wx+ww/3, wy, 4, wh, PicoBrown)
	c.FillRect(wx+(ww*2)/3, wy, 4, wh, PicoBrown)

	if buddy != nil {
		buddy.Draw(c, 36, 40, frameCount)
	}

	renderFireplaceHelper(c, 108, 32, 36, 34, fireOn, frameCount)
	renderCatHelper(c, 76, 70, catBreed, frameCount)
}

func (s *ForestCabinScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if fireOn {
		c.ApplyRadialLight(126, 48, 28.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 3. ClassicCottageScene (16-Bit Classic Cottage)
// ---------------------------------------------------------
type ClassicCottageScene struct{}

func (s *ClassicCottageScene) ID() SceneType { return SceneCottageClassic }
func (s *ClassicCottageScene) Name() string { return "16-Bit Classic Cottage 🏡" }
func (s *ClassicCottageScene) WindowBounds() Rect { return Rect{X: 40, Y: 10, W: 80, H: 46} }

func (s *ClassicCottageScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	var skyTop, skyBot colorful.Color
	switch tod {
	case TimeDawn:
		skyTop, skyBot = PicoLavender, PicoPeach
	case TimeDay:
		skyTop, skyBot = PicoNavy, PicoCyan
	case TimeSunset:
		skyTop, skyBot = PicoDarkPurple, PicoRed
	case TimeNight:
		skyTop, skyBot = PicoBlack, PicoNavy
	}

	if lightning {
		skyTop, skyBot = PicoWhite, PicoWhite
	}
	c.DrawVerticalGradient(wx, wy, ww, wh, skyTop, skyBot)

	if tod == TimeNight && !lightning {
		c.FillRect(wx+48, wy+6, 8, 6, PicoWhite)
		c.FillRect(wx+46, wy+6, 4, 6, skyTop)
	} else if tod == TimeDay {
		c.FillRect(wx+48, wy+6, 8, 6, PicoYellow)
	}
}

func (s *ClassicCottageScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(0, 0, wx-2, 66, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoDarkPurple, PicoNavy)
	c.DrawVerticalGradient(0, 0, c.Width, wy-2, PicoDarkPurple, PicoNavy)

	c.FillRect(0, 66, c.Width, c.Height-66, PicoBrown)
	c.FillRect(0, 66, c.Width, 2, PicoBlack)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoBrown)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoBrown)

	renderDeskHelper(c, 12, 58, 48, lampOn)

	if buddy != nil {
		buddy.Draw(c, 20, 38, frameCount)
	}

	renderFireplaceHelper(c, 108, 34, 36, 32, fireOn, frameCount)
	renderCatHelper(c, 74, 70, catBreed, frameCount)
}

func (s *ClassicCottageScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lampOn {
		c.ApplyLightCone(20, 52, 32, 24, PicoYellow, 0.35)
	}
	if fireOn {
		c.ApplyRadialLight(126, 50, 28.0, PicoOrange, 0.35)
	}
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// ---------------------------------------------------------
// 4. ObservatoryScene (16-Bit Starlight Observatory)
// ---------------------------------------------------------
type ObservatoryScene struct{}

func (s *ObservatoryScene) ID() SceneType { return SceneObservatory }
func (s *ObservatoryScene) Name() string { return "16-Bit Starlight Observatory 🌌" }
func (s *ObservatoryScene) WindowBounds() Rect { return Rect{X: 20, Y: 8, W: 120, H: 54} }

func (s *ObservatoryScene) RenderSky(c *Canvas, tod TimeOfDay, lightning bool, frameCount uint64, stars []struct{ X, Y int; Phase float64 }) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.DrawVerticalGradient(wx, wy, ww, wh, PicoBlack, PicoNavy)

	c.FillRect(wx+72, wy+8, 12, 10, PicoWhite)
	c.FillRect(wx+68, wy+8, 8, 10, PicoBlack)

	for i, st := range stars {
		px := wx + (st.X % ww)
		py := wy + (st.Y % wh)
		twinkle := math.Sin(float64(frameCount)*0.1 + st.Phase)
		if twinkle > 0.2 {
			c.SetPixel(px, py, PicoWhite)
			if i%3 == 0 {
				c.DrawGlyph(px, py/2, '✦', PicoCyan)
			}
		}
	}
}

func (s *ObservatoryScene) RenderForeground(c *Canvas, frameCount uint64, buddy *FocusBuddy, lampOn, fireOn bool, catBreed CatBreed) {
	r := s.WindowBounds()
	wx, wy, ww, wh := r.X, r.Y, r.W, r.H

	c.FillRect(0, 0, wx-2, 66, PicoBlack)
	c.FillRect(wx+ww+2, 0, c.Width-(wx+ww+2), 66, PicoBlack)
	c.FillRect(0, 0, c.Width, wy-2, PicoBlack)
	c.FillRect(0, wy+wh+2, c.Width, c.Height-(wy+wh+2), PicoBlack)

	c.DrawRectOutline(wx-2, wy-2, ww+4, wh+4, PicoDarkGrey)
	c.FillRect(wx+ww/2-2, wy, 4, wh, PicoDarkGrey)
	c.FillRect(wx, wy+wh/2-2, ww, 4, PicoDarkGrey)

	if buddy != nil {
		buddy.Draw(c, 40, 40, frameCount)
	}

	c.FillRect(64, 68, 24, 4, PicoPeach)
	c.FillRect(72, 72, 8, 18, PicoDarkGrey)
}

func (s *ObservatoryScene) ApplyLighting(c *Canvas, lampOn, fireOn, lightning bool) {
	if lightning {
		c.BlendRect(0, 0, c.Width, c.Height, PicoWhite, 0.70)
	}
}

// Helper drawing functions
func renderDeskHelper(c *Canvas, dx, dy, dw int, lampOn bool) {
	c.FillRect(dx, dy, dw, 6, PicoBrown)
	c.FillRect(dx, dy, dw, 2, PicoPeach)
	c.FillRect(dx+4, dy+6, 6, 24, PicoBrown)

	if lampOn {
		c.FillRect(dx+6, dy-2, 6, 2, PicoYellow)
		c.FillRect(dx+8, dy-10, 2, 8, PicoYellow)
		c.FillRect(dx+4, dy-16, 10, 6, PicoGreen)
	} else {
		c.FillRect(dx+6, dy-2, 6, 2, PicoDarkGrey)
		c.FillRect(dx+8, dy-10, 2, 8, PicoDarkGrey)
		c.FillRect(dx+4, dy-16, 10, 6, PicoDarkGrey)
	}

	c.FillRect(dx+24, dy-6, 6, 6, PicoRed)
}

func renderFireplaceHelper(c *Canvas, fx, fy, fw, fh int, fireOn bool, frameCount uint64) {
	c.FillRect(fx, fy, fw, fh, PicoBrown)
	c.FillRect(fx+4, fy+8, fw-8, fh-8, PicoBlack)

	if !fireOn {
		return
	}

	fireCenter := fx + fw/2
	fireTick := (frameCount / 3) % 4

	for i := -6; i <= 6; i++ {
		px := fireCenter + i
		flameH := int(8.0 + 4.0*math.Sin(float64(i)+float64(fireTick)*1.5))

		for fyPos := 0; fyPos < flameH; fyPos++ {
			py := fy + fh - 4 - fyPos
			if py < fy {
				continue
			}
			if fyPos < 3 {
				c.SetPixel(px, py, PicoYellow)
			} else if fyPos < 6 {
				c.SetPixel(px, py, PicoOrange)
			} else {
				c.SetPixel(px, py, PicoRed)
			}
		}
	}
}

func renderCatHelper(c *Canvas, cx, cy int, breed CatBreed, frameCount uint64) {
	c.FillRect(cx-4, cy+8, 24, 6, PicoDarkPurple)

	var furMain colorful.Color
	switch breed {
	case CatOrange:
		furMain = PicoOrange
	case CatTuxedo:
		furMain = PicoNavy
	case CatCalico:
		furMain = PicoPeach
	case CatBlack:
		furMain = PicoBlack
	}

	breathOffset := 0
	if (frameCount/12)%2 == 0 {
		breathOffset = -2
	}

	c.FillRect(cx, cy+breathOffset, 16, 8, furMain)
	c.FillRect(cx-2, cy+2+breathOffset, 6, 6, furMain)

	if (frameCount/10)%3 == 1 {
		c.DrawGlyph(cx-4, (cy-4)/2, 'z', PicoCyan)
	}
}
