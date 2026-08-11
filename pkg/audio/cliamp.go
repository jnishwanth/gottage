package audio

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// CLIAMPStatus holds parsed details from `cliamp status`.
type CLIAMPStatus struct {
	Installed     bool
	DaemonRunning bool
	Playing       bool
	TrackTitle    string
	Artist        string
	Album         string
	Position      string
	Duration      string
	Volume        string
	Provider      string
	Raw           string
}

// PresetTrack represents a curated lofi stream / track preset.
type PresetTrack struct {
	Name string
	URL  string
	Tag  string
}

// CuratedLofiPresets provides ready-to-play lofi radio streams & tracks.
var CuratedLofiPresets = []PresetTrack{
	{Name: "Lofi Girl 24/7 Chill Radio", URL: "https://stream.zeno.fm/f3wvbbqmdg8uv", Tag: "Lofi Girl"},
	{Name: "Cozy Bedroom Lofi Beats", URL: "https://stream.zeno.fm/0r0xa792kwzuv", Tag: "Cozy Beats"},
	{Name: "Japan Lofi & Synthwave", URL: "https://stream.zeno.fm/79836371380ef", Tag: "Chillhop"},
	{Name: "Rainy Day Piano & Lofi", URL: "https://stream.zeno.fm/a823p32y9d0uv", Tag: "Rainy Piano"},
	{Name: "Chiptune & 8-Bit Chill", URL: "https://stream.zeno.fm/u841vbg2kwzuv", Tag: "Chiptune"},
}

// CLIAMPController handles interaction with cliamp.
type CLIAMPController struct {
	mu           sync.Mutex
	cliampPath   string
	status       CLIAMPStatus
	lastChecked  time.Time
	ActivePreset int
}

// NewCLIAMPController initializes the CLIAMP integration controller.
func NewCLIAMPController() *CLIAMPController {
	c := &CLIAMPController{
		ActivePreset: 0,
	}
	path, err := exec.LookPath("cliamp")
	if err == nil {
		c.cliampPath = path
		c.status.Installed = true
	} else {
		// Fallback common homebrew path
		c.cliampPath = "/opt/homebrew/bin/cliamp"
		if _, err := exec.Command(c.cliampPath, "--version").Output(); err == nil {
			c.status.Installed = true
		}
	}
	c.PollStatus()
	return c
}

// PollStatus queries cliamp for current state.
func (c *CLIAMPController) PollStatus() CLIAMPStatus {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.status.Installed {
		return c.status
	}

	cmd := exec.Command(c.cliampPath, "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	rawOutput := out.String()
	c.status.Raw = strings.TrimSpace(rawOutput)

	if err != nil || strings.Contains(rawOutput, "cliamp is not running") || strings.Contains(rawOutput, "no socket") {
		c.status.DaemonRunning = false
		c.status.Playing = false
		return c.status
	}

	c.status.DaemonRunning = true
	c.parseStatusOutput(rawOutput)
	c.lastChecked = time.Now()
	return c.status
}

func (c *CLIAMPController) parseStatusOutput(raw string) {
	c.status.Playing = strings.Contains(raw, "playing") || strings.Contains(raw, "State: playing") || strings.Contains(raw, "▶")

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Track:") || strings.HasPrefix(line, "Title:") {
			c.status.TrackTitle = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "Track:"), "Title:"))
		} else if strings.HasPrefix(line, "Artist:") {
			c.status.Artist = strings.TrimSpace(strings.TrimPrefix(line, "Artist:"))
		} else if strings.HasPrefix(line, "Album:") {
			c.status.Album = strings.TrimSpace(strings.TrimPrefix(line, "Album:"))
		} else if strings.HasPrefix(line, "Volume:") {
			c.status.Volume = strings.TrimSpace(strings.TrimPrefix(line, "Volume:"))
		} else if strings.HasPrefix(line, "Provider:") {
			c.status.Provider = strings.TrimSpace(strings.TrimPrefix(line, "Provider:"))
		}
	}

	if c.status.TrackTitle == "" && c.status.DaemonRunning {
		if c.status.Playing {
			c.status.TrackTitle = CuratedLofiPresets[c.ActivePreset].Name
		} else {
			c.status.TrackTitle = "Paused / Idle"
		}
	}
}

// EnsureDaemon ensures the cliamp daemon is running.
func (c *CLIAMPController) EnsureDaemon() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.status.Installed {
		return fmt.Errorf("cliamp binary not found on system")
	}

	if c.status.DaemonRunning {
		return nil
	}

	// Launch daemon asynchronously
	cmd := exec.Command(c.cliampPath, "--daemon")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start cliamp daemon: %w", err)
	}

	// Wait briefly for socket creation
	time.Sleep(300 * time.Millisecond)
	c.status.DaemonRunning = true
	return nil
}

// TogglePlayPause toggles playback state.
func (c *CLIAMPController) TogglePlayPause() error {
	if !c.status.Installed {
		return fmt.Errorf("cliamp not installed")
	}
	if err := c.EnsureDaemon(); err != nil {
		return err
	}

	cmd := exec.Command(c.cliampPath, "toggle")
	_ = cmd.Run()

	c.PollStatus()
	return nil
}

// NextTrack switches to next track or preset.
func (c *CLIAMPController) NextTrack() error {
	c.ActivePreset = (c.ActivePreset + 1) % len(CuratedLofiPresets)
	return c.PlayPreset(c.ActivePreset)
}

// PrevTrack switches to previous track or preset.
func (c *CLIAMPController) PrevTrack() error {
	c.ActivePreset = (c.ActivePreset - 1 + len(CuratedLofiPresets)) % len(CuratedLofiPresets)
	return c.PlayPreset(c.ActivePreset)
}

// PlayPreset loads and plays a curated preset stream.
func (c *CLIAMPController) PlayPreset(idx int) error {
	if idx < 0 || idx >= len(CuratedLofiPresets) {
		idx = 0
	}
	c.ActivePreset = idx
	preset := CuratedLofiPresets[idx]

	if !c.status.Installed {
		return fmt.Errorf("cliamp not installed")
	}
	if err := c.EnsureDaemon(); err != nil {
		return err
	}

	// Queue or load track via cliamp
	cmd := exec.Command(c.cliampPath, "load", preset.URL)
	_ = cmd.Run()

	c.PollStatus()
	return nil
}

// SetVolume adjusts volume in dB or percentage.
func (c *CLIAMPController) SetVolume(db string) error {
	if !c.status.Installed {
		return fmt.Errorf("cliamp not installed")
	}
	cmd := exec.Command(c.cliampPath, "volume", db)
	return cmd.Run()
}

// GetStatus returns the cached status.
func (c *CLIAMPController) GetStatus() CLIAMPStatus {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.status
}
