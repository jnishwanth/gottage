package features

import (
	"fmt"
	"time"
)

// TimerMode defines timer session phase.
type TimerMode int

const (
	ModeFocus TimerMode = iota
	ModeShortBreak
	ModeLongBreak
)

func (m TimerMode) String() string {
	switch m {
	case ModeFocus:
		return "Cozy Focus 📖"
	case ModeShortBreak:
		return "Tea Break ☕"
	case ModeLongBreak:
		return "Resting Break 🐈"
	default:
		return "Focus 📖"
	}
}

// FocusTimer manages pomodoro timer state.
type FocusTimer struct {
	Mode              TimerMode
	Duration          time.Duration
	Remaining         time.Duration
	Running           bool
	CompletedSessions int
	CurrentTask       string
	lastTick          time.Time
}

// NewFocusTimer creates a new timer with default 25-minute focus session.
func NewFocusTimer() *FocusTimer {
	return &FocusTimer{
		Mode:        ModeFocus,
		Duration:    25 * time.Minute,
		Remaining:   25 * time.Minute,
		Running:     false,
		CurrentTask: "Virtual Cottage Focus",
	}
}

// Toggle starts or pauses the timer.
func (ft *FocusTimer) Toggle() {
	ft.Running = !ft.Running
	if ft.Running {
		ft.lastTick = time.Now()
	}
}

// Reset resets timer back to full duration of current mode.
func (ft *FocusTimer) Reset() {
	ft.Running = false
	ft.Remaining = ft.Duration
}

// SetMode sets timer mode and updates duration.
func (ft *FocusTimer) SetMode(mode TimerMode) {
	ft.Mode = mode
	ft.Running = false
	switch mode {
	case ModeFocus:
		ft.Duration = 25 * time.Minute
	case ModeShortBreak:
		ft.Duration = 5 * time.Minute
	case ModeLongBreak:
		ft.Duration = 15 * time.Minute
	}
	ft.Remaining = ft.Duration
}

// Tick updates remaining time (called every tick ~100ms or 1s).
func (ft *FocusTimer) Tick() bool {
	if !ft.Running {
		return false
	}

	now := time.Now()
	elapsed := now.Sub(ft.lastTick)
	ft.lastTick = now

	ft.Remaining -= elapsed
	if ft.Remaining <= 0 {
		ft.Remaining = 0
		ft.Running = false

		if ft.Mode == ModeFocus {
			ft.CompletedSessions++
			if ft.CompletedSessions%4 == 0 {
				ft.SetMode(ModeLongBreak)
			} else {
				ft.SetMode(ModeShortBreak)
			}
		} else {
			ft.SetMode(ModeFocus)
		}
		return true // trigger alert chime/event
	}
	return false
}

// FormattedTime returns time as "MM:SS".
func (ft *FocusTimer) FormattedTime() string {
	mins := int(ft.Remaining.Minutes())
	secs := int(ft.Remaining.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", mins, secs)
}

// ProgressRatio returns completed fraction between 0.0 and 1.0.
func (ft *FocusTimer) ProgressRatio() float64 {
	if ft.Duration == 0 {
		return 0
	}
	ratio := 1.0 - (float64(ft.Remaining) / float64(ft.Duration))
	if ratio < 0 {
		return 0
	}
	if ratio > 1 {
		return 1
	}
	return ratio
}
