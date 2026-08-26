// Package assets is the single resolver for every file gottage reads out of
// assets/. The curated GIFs are compiled into the binary with //go:embed, so a
// build runs correctly from any working directory.
//
// Before this package existed, scenes registered CWD-relative string literals
// ("assets/animations/kirokaze_room.gif") that pkg/art opened with a bare
// os.Open. Run the binary from anywhere but the repo root and every open failed,
// AnimatedGifScene.RenderSky swallowed the error, and the scene rendered as a
// flat navy rectangle -- indistinguishable from a scene that is merely dark.
// pkg/art's own tests run from pkg/art/, so they never once exercised a
// successful GIF load.
//
// The package deliberately has no CWD fallback. Silent CWD resolution is the bug
// being fixed; GOTTAGE_ASSETS is the explicit, supported escape hatch that
// replaces it.
package assets

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"slices"
	"strings"
)

// OverrideEnv names an optional directory that takes precedence over the
// embedded copies, laid out the same way as assets/ itself (so an override for
// "assets/animations/x.gif" lives at "$GOTTAGE_ASSETS/animations/x.gif").
//
// It exists so Phases 8-10 can drop in a candidate frame and look at it without
// a rebuild. A file absent from the override falls through to the embedded copy,
// which is what makes overriding one asset out of ten practical.
const OverrideEnv = "GOTTAGE_ASSETS"

// keyPrefix is the leading segment of a registration path. Scenes register
// "assets/animations/x.gif"; the embedded FS is rooted at assets/, so its own
// path for that file is "animations/x.gif".
const keyPrefix = "assets/"

// The build fails if this pattern matches nothing, which makes an emptied or
// renamed assets/animations/ impossible to ship as a green build. That is a free
// non-vacuity guard and the reason the pattern is not written defensively.
//
//go:embed animations/*.gif
var embedded embed.FS

// Open resolves an asset by its registration path, e.g.
// "assets/animations/kirokaze_room.gif". The embedded-relative form
// ("animations/kirokaze_room.gif") is accepted too.
//
// Resolution order is $GOTTAGE_ASSETS, then the embedded FS.
func Open(name string) (fs.File, error) {
	rel, err := embedPath(name)
	if err != nil {
		return nil, err
	}

	override, err := overrideFS()
	if err != nil {
		return nil, err
	}
	if override != nil {
		f, err := override.Open(rel)
		switch {
		case err == nil:
			return f, nil
		case errors.Is(err, fs.ErrNotExist):
			// Fall through to the embedded copy: overriding one asset must not
			// require copying the other nine into the override directory.
		default:
			// A permission error or a directory where a file belongs is a
			// misconfigured override, not an absent one. Surface it.
			return nil, fmt.Errorf("%s override: %w", OverrideEnv, err)
		}
	}

	f, err := embedded.Open(rel)
	if err != nil {
		return nil, fmt.Errorf("asset %s: %w", name, err)
	}
	return f, nil
}

// ReadFile returns an asset's full contents, resolved exactly as Open resolves it.
func ReadFile(name string) ([]byte, error) {
	f, err := Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// List returns every embedded asset's registration path, sorted.
//
// It reports the *embedded* set on purpose: that is what ships in the binary and
// what the registry is checked against. An override directory can shadow those
// files but cannot add to the shipped set.
func List() ([]string, error) {
	var names []string
	err := fs.WalkDir(embedded, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		names = append(names, keyPrefix+p)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk embedded assets: %w", err)
	}
	slices.Sort(names)
	return names, nil
}

// embedPath converts a registration path into the embedded FS's own path.
func embedPath(name string) (string, error) {
	rel := strings.TrimPrefix(path.Clean(name), keyPrefix)
	if !fs.ValidPath(rel) || rel == "." {
		return "", fmt.Errorf("asset %q is not a valid asset path", name)
	}
	return rel, nil
}

// overrideFS returns an FS rooted at $GOTTAGE_ASSETS, or nil when unset.
//
// os.DirFS enforces fs.ValidPath, so "..", absolute paths and empty segments are
// rejected by construction rather than by a hand-rolled check here.
//
// The environment is read on every call rather than cached behind a sync.Once:
// t.Setenv changes it mid-process, and a cached lookup would silently ignore
// that -- making the override untestable and its own failure mode invisible.
// Ten assets behind a decode cache means this stat is not on any hot path.
func overrideFS() (fs.FS, error) {
	dir := os.Getenv(OverrideEnv)
	if dir == "" {
		return nil, nil
	}
	info, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("%s=%q: %w", OverrideEnv, dir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s=%q is not a directory", OverrideEnv, dir)
	}
	return os.DirFS(dir), nil
}
