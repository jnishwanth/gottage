package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, gottageBin, args...)
	cmd.Dir = t.TempDir()
	cmd.Env = append(os.Environ(), "GOTTAGE_ASSETS=")
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

// TestSceneFlagSelectsTheNamedScene checks the flag still does its job: an absent
// --scene must stay a legal no-op, and a named one must actually change what is
// rendered rather than being accepted and ignored.
func TestSceneFlagSelectsTheNamedScene(t *testing.T) {
	defaultOut, stderr, code := runGottage(t, "--demo")
	if code != 0 {
		t.Fatalf("--demo alone: exit %d (stderr: %q)", code, stderr)
	}
	if len(defaultOut) == 0 {
		t.Fatal("--demo alone rendered nothing")
	}

	namedOut, stderr, code := runGottage(t, "--demo", "--scene", "Kirokaze Cozy Room")
	if code != 0 {
		t.Fatalf("--demo --scene: exit %d (stderr: %q)", code, stderr)
	}
	if namedOut == defaultOut {
		t.Error("--scene rendered the default scene, so the name was accepted and ignored")
	}
}
