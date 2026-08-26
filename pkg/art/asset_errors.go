package art

import (
	"maps"
	"sync"
)

// Asset load failures are recorded here instead of being swallowed.
//
// AnimatedGifScene.RenderSky answered a failed load with c.Clear(PicoNavy) and
// nothing else. That is how 20 broken scenes shipped across 8 commits: a flat
// navy rectangle is indistinguishable from a scene that is merely dark, and the
// only test covering it asserted the rendered string was non-empty.
//
// Deliberately in memory, with nothing written to stderr. Bubbletea holds the
// terminal in alt-screen mode, so a log line at frame rate would corrupt the
// render it is reporting on. Phase 16 owns surfacing these to the status line
// and rate-limiting them; this is the record it reads.
var assetErrors struct {
	mu   sync.Mutex
	seen map[string]error
}

// recordAssetError notes the first failure seen for a path.
//
// First-error-wins rather than last: RenderSky calls the loader 16.7 times a
// second, so a per-frame overwrite would churn the map for no added information,
// and the first failure is the one with the original cause attached.
func recordAssetError(path string, err error) {
	if err == nil {
		return
	}
	assetErrors.mu.Lock()
	defer assetErrors.mu.Unlock()
	if assetErrors.seen == nil {
		assetErrors.seen = make(map[string]error)
	}
	if _, ok := assetErrors.seen[path]; ok {
		return
	}
	assetErrors.seen[path] = err
}

// AssetErrors returns a copy of every asset that has failed to load, keyed by its
// registration path.
//
// TODO(phase-16): surface these on the status line. plan-phase-16.md's
// TestNoSilentFallbacks is the test that closes the loop.
func AssetErrors() map[string]error {
	assetErrors.mu.Lock()
	defer assetErrors.mu.Unlock()
	return maps.Clone(assetErrors.seen)
}

// ClearAssetErrors discards the record. Tests use it to isolate; Phase 16 will
// use it to dismiss a surfaced warning.
func ClearAssetErrors() {
	assetErrors.mu.Lock()
	defer assetErrors.mu.Unlock()
	assetErrors.seen = nil
}
