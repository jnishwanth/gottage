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

// RenderAudioSpectrum2D renders a multi-line 2D audio spectrum visualizer matrix.
func RenderAudioSpectrum2D(width int, height int, isPlaying bool, frame uint64, mode SpectrumStyle, styles Styles) string {
	if width < 16 {
		width = 16
	}
	if height < 4 {
		height = 4
	}

	numBands := width / 2
	if numBands > 32 {
		numBands = 32
	}

	levels := make([]int, numBands)
	peaks := make([]int, numBands)

	for i := 0; i < numBands; i++ {
		if isPlaying {
			phase := float64(frame)*0.25 + float64(i)*0.45
			hVal := (math.Sin(phase)*0.5 + 0.5) * (math.Cos(phase*0.85)*0.35 + 0.65)
			hVal += 0.15 * math.Abs(math.Sin(float64(frame)*0.4 + float64(i*2)))
			if hVal < 0.05 {
				hVal = 0.05
			}
			if hVal > 1.0 {
				hVal = 1.0
			}
			levels[i] = int(hVal * float64(height))
			peakOffset := int(math.Abs(math.Sin(float64(frame)*0.2+float64(i)))*2.0) + 1
			peaks[i] = levels[i] + peakOffset
			if peaks[i] >= height {
				peaks[i] = height - 1
			}
		} else {
			phase := float64(frame)*0.1 + float64(i)*0.3
			hVal := (math.Sin(phase)*0.5 + 0.5) * 0.2
			levels[i] = int(hVal * float64(height))
			if levels[i] < 1 {
				levels[i] = 1
			}
			peaks[i] = levels[i]
		}
	}

	var rows []string
	topColor := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0055")).Bold(true)
	midColor := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffd166")).Bold(true)
	lowColor := lipgloss.NewStyle().Foreground(lipgloss.Color("#06d6a0")).Bold(true)
	peakColor := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true)
	dimColor := styles.MutedText

	for r := height - 1; r >= 0; r-- {
		var rowBuf strings.Builder
		for i := 0; i < numBands; i++ {
			if mode == SpectrumWave {
				// Oscilloscope scope curve
				sineY := int((math.Sin(float64(frame)*0.2 + float64(i)*0.4)*0.5 + 0.5) * float64(height-1))
				if r == sineY {
					rowBuf.WriteString(styles.BoldAccent.Render("∿ "))
				} else if r == sineY-1 || r == sineY+1 {
					rowBuf.WriteString(styles.ProgressBar.Render("· "))
				} else {
					rowBuf.WriteString(dimColor.Render("  "))
				}
			} else if mode == SpectrumBraille {
				// Braille matrix density
				if r < levels[i] {
					rowBuf.WriteString(styles.ProgressBar.Render("⣿ "))
				} else if r == peaks[i] {
					rowBuf.WriteString(peakColor.Render("⠁ "))
				} else {
					rowBuf.WriteString(dimColor.Render("⠀ "))
				}
			} else if mode == SpectrumDots {
				// Stereo Mirror Spectrum
				half := numBands / 2
				distFromCenter := i - half
				if distFromCenter < 0 {
					distFromCenter = -distFromCenter
				}
				effLevel := levels[distFromCenter]
				if r < effLevel {
					rowBuf.WriteString(midColor.Render("● "))
				} else if r == peaks[distFromCenter] {
					rowBuf.WriteString(peakColor.Render("○ "))
				} else {
					rowBuf.WriteString(dimColor.Render("· "))
				}
			} else {
				// 2D Graphic Equalizer Towers (Default & Bars)
				if r == peaks[i] && peaks[i] > levels[i] {
					rowBuf.WriteString(peakColor.Render("▔ "))
				} else if r < levels[i] {
					var cellStyle lipgloss.Style
					ratio := float64(r) / float64(height)
					if ratio > 0.7 {
						cellStyle = topColor
					} else if ratio > 0.35 {
						cellStyle = midColor
					} else {
						cellStyle = lowColor
					}
					rowBuf.WriteString(cellStyle.Render("█ "))
				} else {
					rowBuf.WriteString(dimColor.Render("· "))
				}
			}
		}
		rows = append(rows, rowBuf.String())
	}

	return strings.Join(rows, "\n")
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
