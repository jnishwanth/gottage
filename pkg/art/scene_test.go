package art

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSceneRegistryAndCycling(t *testing.T) {
	cs := NewCottageScene(160, 90)

	if cs.CurrentScene != SceneRainyCity {
		t.Errorf("Expected initial scene SceneRainyCity, got %v", cs.CurrentScene)
	}

	initialScene := cs.CurrentScene
	count := len(SceneOrder)
	for i := 0; i < count; i++ {
		cs.CycleScene()
	}

	if cs.CurrentScene != initialScene {
		t.Errorf("Expected loop back to initial scene %v after %d cycles, got %v", initialScene, count, cs.CurrentScene)
	}
}

func TestWindowBounds(t *testing.T) {
	cs := NewCottageScene(160, 90)

	cs.CurrentScene = SceneRainyCity
	rCity := cs.WindowRect()
	if rCity.W <= 0 || rCity.H <= 0 {
		t.Errorf("Unexpected City window bounds: %+v", rCity)
	}

	cs.CurrentScene = SceneCottageClassic
	rCot := cs.WindowRect()
	if rCot.W <= 0 || rCot.H <= 0 {
		t.Errorf("Unexpected Cottage window bounds: %+v", rCot)
	}
}

func TestCanvasRenderModes(t *testing.T) {
	cs := NewCottageScene(80, 40)

	if cs.RenderMode != RenderHalfBlock {
		t.Errorf("Expected default RenderMode RenderHalfBlock, got %v", cs.RenderMode)
	}

	modes := []CanvasRenderMode{RenderASCII, RenderBraille, RenderPhosphor, RenderHalfBlock}
	for _, expected := range modes {
		next := cs.CycleRenderMode()
		if next != expected {
			t.Errorf("Expected next RenderMode %v, got %v", expected, next)
		}
		rendered := cs.Render()
		if len(rendered) == 0 {
			t.Errorf("Expected non-empty rendered string for mode %v", expected)
		}
	}

	// Test Canvas clone preserves RenderMode
	c := NewCanvas(20, 20)
	c.RenderMode = RenderBraille
	cloned := c.Clone()
	if cloned.RenderMode != RenderBraille {
		t.Errorf("Expected cloned RenderMode RenderBraille, got %v", cloned.RenderMode)
	}
}

func TestProceduralShadersAndRipples(t *testing.T) {
	cs := NewCottageScene(80, 40)

	// Test 9 shader modes cycle clean
	for i := 0; i < 9; i++ {
		s := cs.CycleShaderEffect()
		if len(s.String()) == 0 {
			t.Errorf("Expected non-empty shader string description for index %d", i)
		}
		cs.RenderFrame()
	}

	// Test Canvas Energy Ripple trigger & rendering
	cs.TriggerCanvasRippleHex(40, 20, "#ff5588")
	if len(cs.Ripples) != 1 {
		t.Errorf("Expected 1 active canvas ripple, got %d", len(cs.Ripples))
	}
	cs.RenderFrame()
	cs.Update()
	if cs.Ripples[0].Life != 1 {
		t.Errorf("Expected ripple life 1 after update, got %v", cs.Ripples[0].Life)
	}
}

func TestPetCatCompanion(t *testing.T) {
	cs := NewCottageScene(80, 40)
	initialBliss := cs.CatBliss

	msg := cs.PetCat()
	if cs.CatBliss <= initialBliss {
		t.Errorf("Expected CatBliss to increase after petting, got %d (initial %d)", cs.CatBliss, initialBliss)
	}
	if len(msg) == 0 {
		t.Errorf("Expected non-empty purr message from PetCat")
	}
	if len(cs.OverlayParticles) == 0 {
		t.Errorf("Expected heart overlay particles spawned after PetCat")
	}
	if len(cs.Ripples) == 0 {
		t.Errorf("Expected rose ripple created on canvas after PetCat")
	}
}

// handcraftedScenes maps every SceneType with a hand-drawn implementation to the
// concrete type that must answer for it.
//
// The `shadowed` column marks the 20 IDs that an AnimatedGifScene literal silently
// claimed before RegisterScene was made to panic. For those 20, GetScene returned a
// looping GIF while presets and the command palette advertised the hand-drawn scene
// by name. That went unnoticed across 8 commits.
var handcraftedScenes = []struct {
	id       SceneType
	wantType string
	shadowed bool
}{
	{SceneRainyCity, "*art.CityCafeScene", false},
	{SceneForestCabin, "*art.ForestCabinScene", false},
	{SceneCottageClassic, "*art.ClassicCottageScene", false},
	{SceneObservatory, "*art.ObservatoryScene", false},
	{SceneZenTeahouse, "*art.ZenTeahouseScene", true},
	{SceneCozyCampfire, "*art.CozyCampfireScene", true},
	{SceneCozyAttic, "*art.CozyAtticScene", true},
	{ScenePixelWinterCabin, "*art.WinterCabinScene", true},
	{SceneCozyGreenhouse, "*art.CozyGreenhouseScene", false},
	{SceneCozyLibrary, "*art.MidnightLibraryScene", false},
	{SceneLighthouseCliff, "*art.LighthouseStudyScene", false},
	{ScenePixelBakery, "*art.PixelBakeryScene", false},
	{SceneRetroLofiTrain, "*art.RetroLofiTrainScene", true},
	{SceneZenWaterfallShrine, "*art.ZenOnsenScene", true},
	{SceneNeonBistro, "*art.VinylLoungeScene", true},
	{SceneStarlightCamp, "*art.SunkenHearthScene", true},
	{SceneRamenShop, "*art.MidnightRamenBarScene", false},
	{SceneMagicSanctuary, "*art.BotanicalApothecaryScene", false},
	{SceneMountainAuroraShack, "*art.NordicAuroraPierScene", false},
	{SceneStarlightTreehouse, "*art.StarlitRooftopStudyScene", false},
	{SceneCoffeeShop, "*art.AntiqueBookshopScene", true},
	{SceneEnchantedForest, "*art.GlasshouseBotanicalScene", true},
	{SceneRiverBridge, "*art.RiversideCampScene", true},
	{SceneOceanWaves, "*art.CoastalBalconyScene", true},
	{SceneSunsetStation, "*art.LofiStudyDeskScene", true},
	{SceneBambooGrove, "*art.ZenRainGardenScene", true},
	{SceneWinterCafe, "*art.CozyCatCafeScene", true},
	{SceneAutumnPark, "*art.AutumnLakeDockScene", true},
	{SceneFairytaleWindmill, "*art.HobbitBurrowScene", false},
	{SceneCyberBar, "*art.JapaneseJazzKissaScene", false},
	{SceneCloudTemple, "*art.StarlitLoftBedScene", false},
	{SceneWizardTower, "*art.AlchemistTowerScene", false},
	{SceneZenBonsai, "*art.MachiyaGardenScene", true},
	{SceneCyberDiner, "*art.CyberLofiDinerScene", true},
	{SceneCozyIgloo, "*art.NordicBayWindowScene", true},
	{SceneSkyIsland, "*art.ForestTreehouseScene", true},
}

// distinctColors counts unique pixel colors on a rendered canvas, quantized to 8
// bits per channel to match what the renderer actually emits.
//
// Read Pixels directly rather than parsing Render(): the half-block renderer
// run-length-suppresses repeated SGR codes and re-quantizes, so the output string
// undercounts.
func distinctColors(c *Canvas) int {
	seen := make(map[[3]uint8]struct{}, 1024)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			p := c.Pixels[y][x]
			seen[[3]uint8{uint8(p.R * 255), uint8(p.G * 255), uint8(p.B * 255)}] = struct{}{}
		}
	}
	return len(seen)
}

// renderSceneLayers draws only the scene's own three layers onto a bare canvas.
//
// It deliberately bypasses CottageScene.RenderFrame's post-processing (weather
// particles, overlay particles, vignette, CRT). Those layers paint over 100
// distinct colors on their own, so a scene stubbed out entirely still clears any
// plausible threshold when measured on the composited frame -- which would make
// this assertion exactly as hollow as the len(rendered) != 0 it replaces.
// Measured: with all three ZenTeahouseScene methods replaced by c.Clear(PicoNavy),
// the composited frame still yields 111 colors; the isolated layers yield 1.
func renderSceneLayers(sc SceneInterface, tod TimeOfDay, lit bool) *Canvas {
	c := NewCanvas(74, 52)
	cs := NewCottageScene(74, 52)
	sc.RenderSky(c, tod, false, cs.FrameCount, cs.Stars)
	sc.RenderForeground(c, cs.FrameCount, cs.Buddy, lit, lit, cs.CatBreed)
	sc.ApplyLighting(c, lit, lit, false)
	return c
}

// minDistinctColors is the floor a real scene's own layers must clear.
//
// Measured over all 36 handcrafted scenes x 4 times of day x lamp/fire on and off
// (288 renders): min 45, 5th percentile 50, median 807, max 1910. The floor sits
// well under the real minimum so ordinary art changes never trip it, while the
// failure this exists to catch cannot slip through: a scene that does not draw
// yields exactly 1, which is how 20 broken scenes shipped under a green suite.
//
// The 45 is not an outlier to design around -- unlit scenes legitimately cluster
// in the 45-61 range while lit ones run into the hundreds. If a change drops a
// scene below 16, that is a broken render, not a threshold that needs lowering.
const minDistinctColors = 16

// TestHandcraftedScenesResolveToHandcrafted is the regression test for Phase 1.
//
// GetScene falls back to SceneRegistry[SceneRainyCity] on a miss, so a `!= nil`
// check can never fail -- the type assertion is what has teeth here.
func TestHandcraftedScenesResolveToHandcrafted(t *testing.T) {
	if len(handcraftedScenes) != 36 {
		t.Fatalf("expected 36 handcrafted scenes in the table, got %d", len(handcraftedScenes))
	}
	for _, tc := range handcraftedScenes {
		t.Run(strings.TrimPrefix(tc.wantType, "*art."), func(t *testing.T) {
			sc, ok := SceneRegistry[tc.id]
			if !ok {
				t.Fatalf("SceneType %d is not registered at all", tc.id)
			}
			if got := fmt.Sprintf("%T", sc); got != tc.wantType {
				extra := ""
				if tc.shadowed {
					extra = " (this ID was shadowed by a GIF before Phase 1; it has regressed)"
				}
				t.Errorf("SceneType %d resolves to %s, want %s%s", tc.id, got, tc.wantType, extra)
			}
			if sc.Name() == "" {
				t.Error("scene has an empty display name")
			}
			if r := sc.WindowBounds(); r.W <= 0 || r.H <= 0 {
				t.Errorf("non-positive window bounds %+v", r)
			}
		})
	}
}

// TestNoDuplicateSceneRegistrations guards the invariant RegisterScene now panics
// on. SceneOrder only ever appended on first registration, so a duplicate used to
// leave the map one entry short of the slice while silently replacing a value.
func TestNoDuplicateSceneRegistrations(t *testing.T) {
	if len(SceneRegistry) != len(SceneOrder) {
		t.Errorf("len(SceneRegistry)=%d but len(SceneOrder)=%d", len(SceneRegistry), len(SceneOrder))
	}
	seen := make(map[SceneType]int, len(SceneOrder))
	for i, id := range SceneOrder {
		if prev, dup := seen[id]; dup {
			t.Errorf("SceneType %d appears at SceneOrder[%d] and [%d]", id, prev, i)
		}
		seen[id] = i
	}
}

// TestEverySceneHasAsset checks the registry and assets/ agree in both directions:
// no registration points at a missing GIF, and no GIF on disk is unregistered.
//
// A missing asset is invisible at runtime -- LoadAnimatedGif fails and the scene
// falls back to a flat navy fill -- so this has to be caught in a test.
func TestEverySceneHasAsset(t *testing.T) {
	const animDir = "assets/animations"
	registered := make(map[string]SceneType)
	for id, sc := range SceneRegistry {
		gs, ok := sc.(*AnimatedGifScene)
		if !ok {
			continue
		}
		if _, err := os.Stat(filepath.Join("..", "..", gs.gifPath)); err != nil {
			t.Errorf("SceneType %d (%q) registers %s, which does not exist", id, gs.name, gs.gifPath)
			continue
		}
		registered[filepath.Base(gs.gifPath)] = id
	}

	entries, err := os.ReadDir(filepath.Join("..", "..", animDir))
	if err != nil {
		t.Fatalf("read %s: %v", animDir, err)
	}
	for _, e := range entries {
		if filepath.Ext(e.Name()) != ".gif" {
			continue
		}
		if _, ok := registered[e.Name()]; !ok {
			t.Errorf("%s/%s is on disk but no scene registers it", animDir, e.Name())
		}
	}
}

// TestHandcraftedProceduralScenes renders every hand-drawn scene at each time of
// day and asserts the result actually contains art.
//
// The previous version asserted len(Canvas.Render()) != 0. That can only fail if
// the renderer emits nothing at all: a scene that fell back to c.Clear(PicoNavy)
// -- a solid navy rectangle -- passed it, which is how 20 broken scenes shipped.
func TestHandcraftedProceduralScenes(t *testing.T) {
	for _, tc := range handcraftedScenes {
		name := strings.TrimPrefix(tc.wantType, "*art.")
		sc := GetScene(tc.id)
		for _, tod := range []TimeOfDay{TimeDawn, TimeDay, TimeSunset, TimeNight} {
			for _, lit := range []bool{true, false} {
				t.Run(fmt.Sprintf("%s/%v/lit=%v", name, tod, lit), func(t *testing.T) {
					if n := distinctColors(renderSceneLayers(sc, tod, lit)); n < minDistinctColors {
						t.Errorf("scene layers produced %d distinct colors, want >= %d (1 means a flat Clear())",
							n, minDistinctColors)
					}
				})
			}
		}
	}
}

func TestCanvasGeometryPrimitives(t *testing.T) {
	c := NewCanvas(40, 40)
	c.DrawCircle(20, 20, 8, PicoRed)
	c.FillCircle(20, 20, 4, PicoYellow)
	c.DrawEllipse(20, 20, 10, 5, PicoCyan)
	c.FillEllipse(20, 20, 6, 3, PicoGreen)
	c.DrawArc(20, 20, 12, 0, 3.14, PicoPink)

	rendered := c.Render()
	if len(rendered) == 0 {
		t.Errorf("Expected non-empty rendered canvas for geometry primitives")
	}
}
