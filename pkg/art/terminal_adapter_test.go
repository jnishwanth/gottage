package art

import (
	"os"
	"strings"
	"testing"
)

func TestDetectTerminalProtocol(t *testing.T) {
	origTermProg := os.Getenv("TERM_PROGRAM")
	origTerm := os.Getenv("TERM")
	origKitty := os.Getenv("KITTY_WINDOW_ID")

	defer func() {
		os.Setenv("TERM_PROGRAM", origTermProg)
		os.Setenv("TERM", origTerm)
		os.Setenv("KITTY_WINDOW_ID", origKitty)
	}()

	os.Setenv("TERM_PROGRAM", "ghostty")
	os.Setenv("TERM", "xterm-ghostty")
	if proto := DetectTerminalProtocol(); proto != ProtocolKitty {
		t.Errorf("Expected ProtocolKitty for Ghostty, got %v", proto)
	}

	os.Setenv("TERM_PROGRAM", "iTerm.app")
	os.Setenv("TERM", "xterm-256color")
	if proto := DetectTerminalProtocol(); proto != ProtocolITerm2 {
		t.Errorf("Expected ProtocolITerm2 for iTerm.app, got %v", proto)
	}

	os.Setenv("TERM_PROGRAM", "alacritty")
	os.Setenv("TERM", "alacritty")
	os.Setenv("KITTY_WINDOW_ID", "")
	if proto := DetectTerminalProtocol(); proto != ProtocolHalfBlock {
		t.Errorf("Expected ProtocolHalfBlock for Alacritty, got %v", proto)
	}

	os.Setenv("KITTY_WINDOW_ID", "1")
	if proto := DetectTerminalProtocol(); proto != ProtocolKitty {
		t.Errorf("Expected ProtocolKitty for a bare KITTY_WINDOW_ID, got %v", proto)
	}
}

// TestTerminalProtocolEnvNamesEveryVariableConsulted pins the contract callers
// rely on to pin the protocol: neutralising every name TerminalProtocolEnv
// returns has to leave the ANSI half-block renderer selected, whatever terminal
// the test process itself was launched from.
func TestTerminalProtocolEnvNamesEveryVariableConsulted(t *testing.T) {
	names := TerminalProtocolEnv()
	if len(names) == 0 {
		t.Fatal("TerminalProtocolEnv named nothing: the assertions below would pass vacuously")
	}

	for _, name := range names {
		t.Setenv(name, "ghostty")
	}
	if got := DetectTerminalProtocol(); got == ProtocolHalfBlock {
		t.Fatalf("a terminal set through every named variable still detected as %v, so clearing them proves nothing", got)
	}

	for _, name := range names {
		t.Setenv(name, "")
	}
	if got := DetectTerminalProtocol(); got != ProtocolHalfBlock {
		t.Errorf("clearing every name in TerminalProtocolEnv left %v; the detector reads a variable the list does not name", got)
	}
}

func TestEncodeKittyGPU(t *testing.T) {
	c := NewCanvas(160, 90)
	payload := EncodeKittyGPU(c, 74, 34)

	if !strings.HasPrefix(payload, "\x1b_Ga=T,f=32,i=1,q=2,C=1,s=320,v=180,c=74,r=34") {
		t.Errorf("Kitty GPU transmission escape sequence header invalid: %s", payload[:40])
	}
}
