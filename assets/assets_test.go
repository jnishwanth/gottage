package assets

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"
	"testing"
)

// wantAssetCount is the size of the curated set Phase 1 landed on.
//
// It is asserted exactly, not as a lower bound, so an asset that disappears
// together with its registration still fails: pkg/art's TestEverySceneHasAsset
// checks the registry and the asset set agree with each other, which stays true
// if both lose the same entry. Update this when Phases 8-10 add assets.
const wantAssetCount = 10

// minPlausibleAssetBytes is set above the ~130 bytes of a Git LFS pointer file
// and far below the smallest real asset (rain_setup.gif, 73KB).
//
// AGENTS.md section 5 forbids LFS precisely because //go:embed reads whatever is
// on disk at build time: with the smudge filter not run, the build succeeds,
// Open succeeds, and every scene renders flat navy. This is the tripwire.
const minPlausibleAssetBytes = 1024

func TestEmbeddedAssetsAreRealGifs(t *testing.T) {
	names, err := List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(names) != wantAssetCount {
		t.Errorf("embedded %d assets, want %d: %v", len(names), wantAssetCount, names)
	}
	if len(names) == 0 {
		t.Fatal("no embedded assets: every assertion below would pass vacuously")
	}

	for _, name := range names {
		t.Run(filepath.Base(name), func(t *testing.T) {
			b, err := ReadFile(name)
			if err != nil {
				t.Fatalf("ReadFile: %v", err)
			}
			if len(b) < minPlausibleAssetBytes {
				t.Errorf("%d bytes, below the %d-byte floor -- an LFS pointer or a truncated file looks like this",
					len(b), minPlausibleAssetBytes)
			}
			if !bytes.HasPrefix(b, []byte("GIF87a")) && !bytes.HasPrefix(b, []byte("GIF89a")) {
				head := b
				if len(head) > 16 {
					head = head[:16]
				}
				t.Errorf("does not start with GIF87a/GIF89a; first bytes are %q", head)
			}
		})
	}
}

func TestListPathsAreRegistrationKeys(t *testing.T) {
	names, err := List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	// Scenes register "assets/animations/x.gif". If List spoke the embedded FS's
	// own dialect instead, callers would have to translate and the point of a
	// single canonical key would be lost.
	for _, name := range names {
		if got := filepath.ToSlash(filepath.Dir(name)); got != "assets/animations" {
			t.Errorf("List returned %q, want a path under assets/animations", name)
		}
		if _, err := Open(name); err != nil {
			t.Errorf("List returned %q but Open rejects it: %v", name, err)
		}
	}
}

func TestOpenAcceptsBothPathForms(t *testing.T) {
	viaKey, err := ReadFile("assets/animations/kirokaze_room.gif")
	if err != nil {
		t.Fatalf("registration path: %v", err)
	}
	viaEmbed, err := ReadFile("animations/kirokaze_room.gif")
	if err != nil {
		t.Fatalf("embedded path: %v", err)
	}
	if !bytes.Equal(viaKey, viaEmbed) {
		t.Error("the two path forms resolved to different files")
	}
}

func TestOpenRejectsPathsOutsideTheAssetSet(t *testing.T) {
	for _, name := range []string{
		"../go.mod",
		"assets/../../go.mod",
		"/etc/passwd",
		"",
		".",
		"assets/",
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Open(name); err == nil {
				t.Errorf("Open(%q) succeeded; it must not resolve outside the asset set", name)
			}
		})
	}
}

// TestAssetOverrideEnv covers all three behaviours of GOTTAGE_ASSETS, because the
// interesting one is the fall-through: overriding a single asset must not require
// copying the other nine, and a directory that does not exist must be loud rather
// than silently ignored.
func TestAssetOverrideEnv(t *testing.T) {
	const overridden = "assets/animations/kirokaze_room.gif"
	const untouched = "assets/animations/snow_cabin.gif"

	embeddedOverridden, err := ReadFile(overridden)
	if err != nil {
		t.Fatalf("read embedded %s: %v", overridden, err)
	}
	embeddedUntouched, err := ReadFile(untouched)
	if err != nil {
		t.Fatalf("read embedded %s: %v", untouched, err)
	}

	t.Run("present in override wins", func(t *testing.T) {
		dir := t.TempDir()
		replacement := tinyGIF(t)
		writeAsset(t, dir, "animations/kirokaze_room.gif", replacement)
		t.Setenv(OverrideEnv, dir)

		got, err := ReadFile(overridden)
		if err != nil {
			t.Fatalf("ReadFile: %v", err)
		}
		if !bytes.Equal(got, replacement) {
			t.Errorf("got %d bytes, want the %d-byte override", len(got), len(replacement))
		}
		if bytes.Equal(got, embeddedOverridden) {
			t.Error("read the embedded copy while an override was present")
		}
	})

	t.Run("absent from override falls through", func(t *testing.T) {
		dir := t.TempDir()
		writeAsset(t, dir, "animations/kirokaze_room.gif", tinyGIF(t))
		t.Setenv(OverrideEnv, dir)

		got, err := ReadFile(untouched)
		if err != nil {
			t.Fatalf("ReadFile: %v", err)
		}
		if !bytes.Equal(got, embeddedUntouched) {
			t.Error("an asset absent from the override did not fall through to the embedded copy")
		}
	})

	t.Run("empty override directory falls through for everything", func(t *testing.T) {
		t.Setenv(OverrideEnv, t.TempDir())
		got, err := ReadFile(overridden)
		if err != nil {
			t.Fatalf("ReadFile: %v", err)
		}
		if !bytes.Equal(got, embeddedOverridden) {
			t.Error("empty override did not fall through to the embedded copy")
		}
	})

	t.Run("nonexistent directory is an error", func(t *testing.T) {
		t.Setenv(OverrideEnv, filepath.Join(t.TempDir(), "no-such-dir"))
		if _, err := Open(overridden); err == nil {
			t.Error("a GOTTAGE_ASSETS pointing nowhere resolved silently to the embedded copy")
		}
	})

	t.Run("file instead of directory is an error", func(t *testing.T) {
		dir := t.TempDir()
		notADir := filepath.Join(dir, "regular-file")
		if err := os.WriteFile(notADir, []byte("not a directory"), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
		t.Setenv(OverrideEnv, notADir)
		if _, err := Open(overridden); err == nil {
			t.Error("a GOTTAGE_ASSETS pointing at a file resolved silently to the embedded copy")
		}
	})
}

func writeAsset(t *testing.T, dir, rel string, data []byte) {
	t.Helper()
	full := filepath.Join(dir, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(full, data, 0o644); err != nil {
		t.Fatalf("write %s: %v", full, err)
	}
}

// tinyGIF encodes a real 2x2 GIF, so the override tests distinguish files by
// their bytes rather than by a sentinel that would not survive a decode.
func tinyGIF(t *testing.T) []byte {
	t.Helper()
	m := image.NewPaletted(image.Rect(0, 0, 2, 2), color.Palette{color.Black, color.White})
	m.SetColorIndex(1, 1, 1)
	var buf bytes.Buffer
	if err := gif.EncodeAll(&buf, &gif.GIF{Image: []*image.Paletted{m}, Delay: []int{10}}); err != nil {
		t.Fatalf("encode: %v", err)
	}
	return buf.Bytes()
}
