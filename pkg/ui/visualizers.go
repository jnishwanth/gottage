package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type SpectrumStyle int

const (
	SpectrumBars SpectrumStyle = iota
	SpectrumDots
	SpectrumBraille
	SpectrumEqualizer
	SpectrumWave
)

func (s SpectrumStyle) String() string {
	switch s {
	case SpectrumBars:
		return "Bars ▂▃▅█"
	case SpectrumDots:
		return "Dots .:*o"
	case SpectrumBraille:
		return "Braille ⡀⣀⣤⣿"
	case SpectrumEqualizer:
		return "Equalizer █▅▂"
	case SpectrumWave:
		return "Waveform ∿⌇∿"
	default:
		return "Bars"
	}
}

// RenderAudioSpectrum renders multi-mode spectrum visualizers.
func RenderAudioSpectrum(width int, isPlaying bool, frame uint64, mode SpectrumStyle, style lipgloss.Style) string {
	numBars := 16
	if width < 16 {
		numBars = 8
	}

	var charRungs []rune
	switch mode {
	case SpectrumDots:
		charRungs = []rune{' ', '.', ':', '*', 'o', 'O', '0', '@', '#'}
	case SpectrumBraille:
		charRungs = []rune{' ', '⡀', '⣀', '⡤', '⣤', '⡶', '⣶', '⣞', '⣿'}
	case SpectrumEqualizer:
		charRungs = []rune{' ', '▂', '▃', '▅', '▆', '▇', '█', '█', '█'}
	case SpectrumWave:
		charRungs = []rune{' ', '.', '~', '∿', '⌇', '〰', '≋', '≈', '🌊'}
	default: // SpectrumBars
		charRungs = []rune{' ', ' ', '▂', '▃', '▄', '▅', '▆', '▇', '█'}
	}

	var sb strings.Builder
	for i := 0; i < numBars; i++ {
		var level int
		if isPlaying {
			// Oscillate heights with harmonic sine waves
			phase := float64(frame)*0.2 + float64(i)*0.6
			val := (math.Sin(phase)*0.5 + 0.5) * (math.Cos(phase*0.7)*0.4 + 0.6)
			level = int(val * float64(len(charRungs)-1))
			if level < 1 {
				level = 1
			}
		} else {
			level = 1 // Resting quiet state
		}

		if level >= len(charRungs) {
			level = len(charRungs) - 1
		}
		sb.WriteRune(charRungs[level])
	}

	return style.Render(sb.String())
}

// ZenBreathState represents current phase of the 4-4-4-4 box breathing cycle.
type ZenBreathState struct {
	PhaseName   string
	Instruction string
	RingVisual  string
}

// GetZenBreathState returns box breathing state based on frame count.
func GetZenBreathState(frame uint64) ZenBreathState {
	cycle := (frame / 4) % 64
	sec := cycle / 16

	switch sec {
	case 0:
		ratio := float64(cycle%16) / 16.0
		dots := int(ratio*5) + 1
		return ZenBreathState{
			PhaseName:   "Breathe In 🌬️",
			Instruction: "Slow inhale through nose...",
			RingVisual:  fmt.Sprintf("(  %s  )", strings.Repeat("◯ ", dots)),
		}
	case 1:
		return ZenBreathState{
			PhaseName:   "Hold Breath 🧘",
			Instruction: "Hold gently and rest...",
			RingVisual:  "[ ◯ ◯ ◯ ◯ ◯ ◯ ]",
		}
	case 2:
		ratio := 1.0 - (float64(cycle%16) / 16.0)
		dots := int(ratio*5) + 1
		if dots < 1 {
			dots = 1
		}
		return ZenBreathState{
			PhaseName:   "Breathe Out 😮‍💨",
			Instruction: "Slow exhale through mouth...",
			RingVisual:  fmt.Sprintf("(  %s  )", strings.Repeat("◯ ", dots)),
		}
	default:
		return ZenBreathState{
			PhaseName:   "Rest & Calm 💤",
			Instruction: "Pause naturally...",
			RingVisual:  "(   •   )",
		}
	}
}
