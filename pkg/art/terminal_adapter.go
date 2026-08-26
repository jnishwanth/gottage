package art

import (
	"encoding/base64"
	"fmt"
	"os"
	"slices"
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

const (
	envTermProgram = "TERM_PROGRAM"
	envTerm        = "TERM"
	envKittyWindow = "KITTY_WINDOW_ID"
)

// terminalProtocolEnv is the only route detection has to the environment.
//
// Detection reads through this list, so a variable consulted by the matcher but
// missing from it resolves to the empty string and cannot influence the answer.
// That is what keeps TerminalProtocolEnv complete: a caller that neutralises
// every name it returns gets ProtocolHalfBlock, and stays getting it when the
// matcher grows a case.
var terminalProtocolEnv = []string{envTermProgram, envTerm, envKittyWindow}

// TerminalProtocolEnv returns every environment variable DetectTerminalProtocol
// consults.
//
// CottageScene.Render answers ProtocolKitty with a base64 graphics payload rather
// than ANSI, so anything asserting on rendered colour has to pin the protocol
// instead of inheriting whichever terminal the process was launched from.
func TerminalProtocolEnv() []string {
	return slices.Clone(terminalProtocolEnv)
}

// DetectTerminalProtocol inspects environment variables to determine terminal graphics capabilities.
func DetectTerminalProtocol() TerminalProtocol {
	env := make(map[string]string, len(terminalProtocolEnv))
	for _, name := range terminalProtocolEnv {
		env[name] = os.Getenv(name)
	}
	return detectTerminalProtocol(env)
}

func detectTerminalProtocol(env map[string]string) TerminalProtocol {
	termProg := strings.ToLower(env[envTermProgram])
	term := strings.ToLower(env[envTerm])
	kittyWindow := env[envKittyWindow]

	if termProg == "ghostty" || term == "xterm-ghostty" || termProg == "kitty" || kittyWindow != "" || termProg == "wezterm" {
		return ProtocolKitty
	}

	if termProg == "iterm.app" {
		return ProtocolITerm2
	}

	return ProtocolHalfBlock
}

// ClearKittyGPU returns the escape sequence to delete Kitty GPU Image ID 1 immediately.
func ClearKittyGPU() string {
	return "\x1b_Ga=d,d=i,i=1\x1b\\"
}

// EncodeKittyGPU converts a Canvas pixel grid into Kitty GPU Graphics Protocol Base64 payload.
// Applies 2x integer scaling (320x180 HD retro pixels) with 16:9 aspect locking to eliminate all stretch/blur!
func EncodeKittyGPU(c *Canvas, cols, rows int) string {
	scale := 2
	srcW, srcH := c.Width, c.Height
	dstW, dstH := srcW*scale, srcH*scale

	rgbaBuf := make([]byte, dstW*dstH*4)

	// 2x Nearest-Neighbor Integer Scaling for 320x180 HD Retro Pixels
	for y := 0; y < srcH; y++ {
		for x := 0; x < srcW; x++ {
			col := c.Pixels[y][x]
			r := byte(col.R * 255)
			g := byte(col.G * 255)
			b := byte(col.B * 255)

			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					px := x*scale + dx
					py := y*scale + dy
					idx := (py*dstW + px) * 4
					rgbaBuf[idx] = r
					rgbaBuf[idx+1] = g
					rgbaBuf[idx+2] = b
					rgbaBuf[idx+3] = 255
				}
			}
		}
	}

	b64Payload := base64.StdEncoding.EncodeToString(rgbaBuf)
	var sb strings.Builder

	// Kitty Escape Header:
	// a=T (Transmit & Display), f=32 (32-bit RGBA), i=1 (Image ID 1 for in-place GPU overwrite),
	// q=2 (Quiet mode - suppress terminal ACKs), C=1 (Do NOT move cursor)
	hdr := fmt.Sprintf("\x1b_Ga=T,f=32,i=1,q=2,C=1,s=%d,v=%d,c=%d,r=%d", dstW, dstH, cols, rows)

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

	for r := 0; r < rows; r++ {
		if r > 0 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
