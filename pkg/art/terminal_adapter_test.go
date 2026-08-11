package art

import (
	"os"
	"strings"
	"testing"
)

func TestDetectTerminalProtocol(t *testing.T) {
	// Backup original env
	origTermProg := os.Getenv("TERM_PROGRAM")
	origTerm := os.Getenv("TERM")
	origKitty := os.Getenv("KITTY_WINDOW_ID")

	defer func() {
		os.Setenv("TERM_PROGRAM", origTermProg)
		os.Setenv("TERM", origTerm)
		os.Setenv("KITTY_WINDOW_ID", origKitty)
	}()

	// Test Ghostty
	os.Setenv("TERM_PROGRAM", "ghostty")
	os.Setenv("TERM", "xterm-ghostty")
	if proto := DetectTerminalProtocol(); proto != ProtocolKitty {
		t.Errorf("Expected ProtocolKitty for Ghostty, got %v", proto)
	}

	// Test iTerm2
	os.Setenv("TERM_PROGRAM", "iTerm.app")
	os.Setenv("TERM", "xterm-256color")
	if proto := DetectTerminalProtocol(); proto != ProtocolITerm2 {
		t.Errorf("Expected ProtocolITerm2 for iTerm.app, got %v", proto)
	}

	// Test Alacritty / Fallback
	os.Setenv("TERM_PROGRAM", "alacritty")
	os.Setenv("TERM", "alacritty")
	os.Setenv("KITTY_WINDOW_ID", "")
	if proto := DetectTerminalProtocol(); proto != ProtocolHalfBlock {
		t.Errorf("Expected ProtocolHalfBlock for Alacritty, got %v", proto)
	}
}

func TestEncodeKittyGPU(t *testing.T) {
	c := NewCanvas(10, 10)
	payload := EncodeKittyGPU(c, 74, 34)

	if !strings.HasPrefix(payload, "\x1b_Ga=T,f=32,s=10,v=10,c=74,r=34") {
		t.Errorf("Kitty GPU transmission escape sequence header invalid: %s", payload[:40])
	}
	if !strings.HasSuffix(payload, "\x1b\\") {
		t.Errorf("Kitty GPU transmission escape sequence footer invalid")
	}
}
