package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"gottage/assets"
	"gottage/pkg/art"
)

// gottageBin is the real binary, built once. The --scene contract is about exit
// codes and stderr, neither of which is reachable from inside the test process:
// main calls os.Exit.
var gottageBin string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "gottage-cli")
	if err != nil {
		fmt.Fprintf(os.Stderr, "temp dir: %v\n", err)
		os.Exit(1)
	}
	gottageBin = filepath.Join(dir, "gottage")
	build := exec.Command("go", "build", "-o", gottageBin, ".")
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "build gottage: %v\n", err)
		os.RemoveAll(dir)
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

// runGottage runs the binary from a directory that is not the repo, which is the
// condition Phase 2 exists to make safe, and with GOTTAGE_ASSETS neutralised so a
// developer's override cannot change what these cases mean.
func runGottage(t *testing.T, args ...string) (stdout, stderr string, code int) {
	t.Helper()
	return runGottageWithAssets(t, "", args...)
}

// runGottageWithAssets sets GOTTAGE_ASSETS to exactly override, replacing any
// value already in the environment rather than appending a second one: which of
// two same-named entries getenv answers with is not something to rely on.
//
// It also drops every variable art.TerminalProtocolEnv names. CottageScene.Render
// answers a Kitty-class terminal with a base64 graphics payload carrying no SGR
// colour sequences at all, so inheriting the developer's TERM_PROGRAM would make
// the colour counts below read 0 against a perfectly good binary. Taking the list
// from the detector rather than restating it here is what stops a variable added
// to that detector from silently reaching the child.
func runGottageWithAssets(t *testing.T, override string, args ...string) (stdout, stderr string, code int) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	stripped := append(art.TerminalProtocolEnv(), assets.OverrideEnv)
	env := slices.DeleteFunc(os.Environ(), func(kv string) bool {
		name, _, _ := strings.Cut(kv, "=")
		return slices.Contains(stripped, name)
	})

	cmd := exec.CommandContext(ctx, gottageBin, args...)
	cmd.Dir = t.TempDir()
	cmd.Env = append(env, assets.OverrideEnv+"="+override)
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if ctx.Err() != nil {
		t.Fatalf("gottage %q never exited; it is waiting on a terminal it was not asked for", args)
	}
	var exitErr *exec.ExitError
	switch {
	case err == nil:
		code = 0
	case errors.As(err, &exitErr):
		code = exitErr.ExitCode()
	default:
		t.Fatalf("run gottage %q: %v", args, err)
	}
	return out.String(), errOut.String(), code
}

// TestSceneFlagRejectsAnEmptyValue is the regression: both guards used to test
// *sceneFlag != "", so `gottage --demo --scene "$SCENE"` with SCENE unset skipped
// scene resolution entirely and rendered the default -- SceneRainyCity, the
// handcrafted scene that reads no asset -- with no diagnostic, and `gottage
// --scene ""` opened the interactive TUI instead of exiting 2.
func TestSceneFlagRejectsAnEmptyValue(t *testing.T) {
	t.Run("without --demo", func(t *testing.T) {
		_, stderr, code := runGottage(t, "--scene", "")
		if code != 2 {
			t.Errorf("exit %d, want 2 (stderr: %q)", code, stderr)
		}
		if !strings.Contains(stderr, "--scene requires --demo") {
			t.Errorf("stderr does not say why: %q", stderr)
		}
	})

	t.Run("with --demo", func(t *testing.T) {
		_, stderr, code := runGottage(t, "--demo", "--scene", "")
		if code != 2 {
			t.Errorf("exit %d, want 2 (stderr: %q)", code, stderr)
		}
		if !strings.Contains(stderr, "Available scenes:") {
			t.Errorf("stderr does not list the valid names: %q", stderr)
		}
	})

	t.Run("with --demo and --scene=", func(t *testing.T) {
		_, stderr, code := runGottage(t, "--demo", "--scene=")
		if code != 2 {
			t.Errorf("exit %d, want 2 (stderr: %q)", code, stderr)
		}
	})
}

func TestSceneFlagRejectsAnUnknownName(t *testing.T) {
	_, stderr, code := runGottage(t, "--demo", "--scene", "no such scene")
	if code != 2 {
		t.Errorf("exit %d, want 2 (stderr: %q)", code, stderr)
	}
	if !strings.Contains(stderr, "Available scenes:") {
		t.Errorf("stderr does not list the valid names: %q", stderr)
	}
}

// fgColor matches the 24-bit foreground sequence Canvas.renderHalfBlock writes.
// The rendered frame is the demo's whole output contract, so counting the
// distinct colours in it is reading published output, not implementation text.
var fgColor = regexp.MustCompile(`\x1b\[38;2;(\d+;\d+;\d+)m`)

func distinctForegroundColors(frame string) int {
	seen := make(map[string]struct{}, 2048)
	for _, m := range fgColor.FindAllStringSubmatch(frame, -1) {
		seen[m[1]] = struct{}{}
	}
	return len(seen)
}

// Both floors are measured on the built binary from a temp directory, pinned to
// --time night --weather rain so neither the wall clock nor the weather default
// moves them. Only the star and particle RNG still does, by about 5 either way.
//
// minDemoGifColors: the ten GIF-backed scenes score 1017 (Pixel Rain Setup) to
// 1581 (Home Sweet Home); the default handcrafted CityCafeScene scores 579-584;
// the navy fallback a failed asset produces scores 107. 850 sits between the two
// live bands, 1.45x over the default and 1.2x under the weakest GIF.
//
// minTimeOfDayColorGap: --time day scores 682-686 against night's 579-584. If
// --time does not reach the frame both runs resolve TimeAuto against the same
// clock and the gap collapses to RNG noise.
const (
	minDemoGifColors     = 850
	minTimeOfDayColorGap = 50
)

// TestSceneFlagSelectsTheNamedScene checks the flag does its job: an absent
// --scene stays a legal no-op, and a named one actually changes what is rendered
// rather than being accepted and ignored.
//
// The two frames cannot be compared for inequality. NewCottageScene seeds its rng
// from the clock, so two runs of the same scene already differ; that comparison
// would hold with the assignment of the resolved SceneType deleted. A colour
// floor survives the noise, and asserting the default stays under it is what
// stops the floor from passing for a reason other than the one claimed.
func TestSceneFlagSelectsTheNamedScene(t *testing.T) {
	pinned := []string{"--demo", "--time", "night", "--weather", "rain"}

	defaultOut, stderr, code := runGottage(t, pinned...)
	if code != 0 {
		t.Fatalf("--demo alone: exit %d (stderr: %q)", code, stderr)
	}
	namedOut, stderr, code := runGottage(t, append(slices.Clone(pinned), "--scene", "Kirokaze Cozy Room")...)
	if code != 0 {
		t.Fatalf("--demo --scene: exit %d (stderr: %q)", code, stderr)
	}

	named := distinctForegroundColors(namedOut)
	fallback := distinctForegroundColors(defaultOut)
	t.Logf("Kirokaze Cozy Room: %d distinct colours; default scene: %d", named, fallback)

	if named < minDemoGifColors {
		t.Errorf("--scene rendered %d distinct colours, want at least %d -- the GIF never reached the frame",
			named, minDemoGifColors)
	}
	if fallback >= minDemoGifColors {
		t.Errorf("the default scene renders %d distinct colours, at or above the %d floor, so the floor above no longer distinguishes the two",
			fallback, minDemoGifColors)
	}
}

// TestDemoRendersAnsiWhateverTerminalRanTheSuite reproduces a false failure the
// colour floors used to take from their surroundings: with TERM_PROGRAM=ghostty
// in the environment, CottageScene.Render returns EncodeKittyGPU's base64 payload,
// distinctForegroundColors reads 0, and every assertion above reports a GIF that
// never reached the frame on a binary where it did.
func TestDemoRendersAnsiWhateverTerminalRanTheSuite(t *testing.T) {
	for _, name := range art.TerminalProtocolEnv() {
		t.Setenv(name, "ghostty")
	}

	out, stderr, code := runGottage(t, "--demo", "--time", "night", "--weather", "rain", "--scene", "Kirokaze Cozy Room")
	if code != 0 {
		t.Fatalf("exit %d (stderr: %q)", code, stderr)
	}
	if got := distinctForegroundColors(out); got < minDemoGifColors {
		t.Errorf("%d distinct colours: the host terminal's protocol reached the child and the frame is not ANSI", got)
	}
}

// TestTimeFlagReachesTheDemoFrame covers --time being accepted and dropped: it
// was read only on the interactive path, after the demo block had already exited,
// so `gottage --demo --time night` rendered at whatever hour it happened to be.
func TestTimeFlagReachesTheDemoFrame(t *testing.T) {
	nightOut, stderr, code := runGottage(t, "--demo", "--time", "night", "--weather", "rain")
	if code != 0 {
		t.Fatalf("--time night: exit %d (stderr: %q)", code, stderr)
	}
	dayOut, stderr, code := runGottage(t, "--demo", "--time", "day", "--weather", "rain")
	if code != 0 {
		t.Fatalf("--time day: exit %d (stderr: %q)", code, stderr)
	}

	night := distinctForegroundColors(nightOut)
	day := distinctForegroundColors(dayOut)
	t.Logf("night: %d distinct colours; day: %d", night, day)

	if day-night < minTimeOfDayColorGap {
		t.Errorf("night rendered %d distinct colours and day %d; a gap under %d means both frames resolved the same time of day and --time never reached the render",
			night, day, minTimeOfDayColorGap)
	}
}

// TestUnknownFlagValuesAreRejected covers the other half of the same class: both
// switches used to have no default arm, so a value no case claimed was accepted
// and silently dropped.
func TestUnknownFlagValuesAreRejected(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want string
	}{
		{"weather typo", []string{"--demo", "--weather", "hail"}, "unknown --weather"},
		{"time typo", []string{"--demo", "--time", "midnight"}, "unknown --time"},
		{"empty weather", []string{"--demo", "--weather", ""}, "unknown --weather"},
		{"empty time", []string{"--demo", "--time", ""}, "unknown --time"},
		{"weather typo without --demo", []string{"--weather", "hail"}, "unknown --weather"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, stderr, code := runGottage(t, tc.args...)
			if code != 2 {
				t.Errorf("exit %d, want 2 (stderr: %q)", code, stderr)
			}
			if !strings.Contains(stderr, tc.want) {
				t.Errorf("stderr does not name the bad value: %q", stderr)
			}
		})
	}
}

// TestDemoReportsAssetFailure is the reason --demo is worth running at all: the
// phase doc uses it as the proof that the binary resolves its assets off the repo
// root. It exited 0 with an empty stderr on a GOTTAGE_ASSETS that pointed
// nowhere, rendering the flat navy fallback, so the command could not tell an
// operator apart from the failure it exists to detect.
func TestDemoReportsAssetFailure(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-override")

	stdout, stderr, code := runGottageWithAssets(t, missing, "--demo", "--scene", "Kirokaze Cozy Room")
	if code != 1 {
		t.Errorf("exit %d, want 1 (stderr: %q)", code, stderr)
	}
	if !strings.Contains(stderr, "failed to load") {
		t.Errorf("stderr does not report the failure: %q", stderr)
	}
	if !strings.Contains(stderr, "assets/animations/kirokaze_room.gif") {
		t.Errorf("stderr does not name the asset that failed: %q", stderr)
	}
	if strings.Contains(stdout, "Demo Frame Rendered") {
		t.Error("the demo announced a rendered frame while its scene was the navy fallback")
	}
	if got := distinctForegroundColors(stdout); got >= minDemoGifColors {
		t.Errorf("%d distinct colours: the asset resolved after all, so this case proves nothing", got)
	}
}
