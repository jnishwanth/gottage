package art

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// TerminalProtocol defines supported terminal graphics protocols.
type TerminalProtocol int

const (
	ProtocolHalfBlock TerminalProtocol = iota // Standard ANSI half-block '▀' (Alacritty / Apple Terminal)
	ProtocolKitty                             // Kitty GPU Graphics Protocol (Ghostty / Kitty / WezTerm)
	ProtocolITerm2                            // iTerm2 Inline Graphics Protocol (iTerm2 / WezTerm)
)

func (p TerminalProtocol) String() string {
	switch p {
	case ProtocolKitty:
		return "Ghostty / Kitty GPU Protocol ⚡"
	case ProtocolITerm2:
		return "iTerm2 Inline Protocol 🖼️"
	case ProtocolHalfBlock:
		return "8-Bit Half-Block Canvas ▀"
	default:
		return "8-Bit Half-Block Canvas ▀"
	}
}

// DetectTerminalProtocol inspects environment variables to determine terminal graphics capabilities.
func DetectTerminalProtocol() TerminalProtocol {
	termProg := strings.ToLower(os.Getenv("TERM_PROGRAM"))
	term := strings.ToLower(os.Getenv("TERM"))
	kittyWindow := os.Getenv("KITTY_WINDOW_ID")

	// Ghostty, Kitty, WezTerm natively support Kitty GPU Graphics Protocol
	if termProg == "ghostty" || term == "xterm-ghostty" || termProg == "kitty" || kittyWindow != "" || termProg == "wezterm" {
		return ProtocolKitty
	}

	// iTerm2 protocol
	if termProg == "iterm.app" {
		return ProtocolITerm2
	}

	// Fallback to standard ANSI half-blocks (Alacritty, Apple Terminal, XTerm)
	return ProtocolHalfBlock
}

// EncodeKittyGPU converts a Canvas pixel grid into Kitty GPU Graphics Protocol Base64 payload.
// Format: \x1b_Ga=T,f=32,s=<width>,v=<height>,c=<cols>,r=<rows>;<base64_payload>\x1b\
func EncodeKittyGPU(c *Canvas, cols, rows int) string {
	w, h := c.Width, c.Height
	rgbaBuf := make([]byte, w*h*4)

	idx := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			col := c.Pixels[y][x]
			rgbaBuf[idx] = byte(col.R * 255)
			rgbaBuf[idx+1] = byte(col.G * 255)
			rgbaBuf[idx+2] = byte(col.B * 255)
			rgbaBuf[idx+3] = 255 // Opaque alpha
			idx += 4
		}
	}

	// Base64 encode raw RGBA payload
	b64Payload := base64.StdEncoding.EncodeToString(rgbaBuf)

	var sb strings.Builder

	// Kitty Graphics Transmission Escape Sequence
	// a=T (Transmit & Display), f=32 (32-bit RGBA), s=width, v=height, c=columns, r=rows
	hdr := fmt.Sprintf("\x1b_Ga=T,f=32,s=%d,v=%d,c=%d,r=%d", w, h, cols, rows)

	// Send in 4096-byte chunk chunks per Kitty protocol specification
	payloadBytes := []byte(b64Payload)
	chunkSize := 4096

	for i := 0; i < len(payloadBytes); i += chunkSize {
		end := i + chunkSize
		m := 1
		if end >= len(payloadBytes) {
			end = len(payloadBytes)
			m = 0
		}
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%s,m=%d;%s\x1b\\", hdr, m, string(payloadBytes[i:end])))
		} else {
			sb.WriteString(fmt.Sprintf("\x1b_Gm=%d;%s\x1b\\", m, string(payloadBytes[i:end])))
		}
	}

	return sb.String()
}
