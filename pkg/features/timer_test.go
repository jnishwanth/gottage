package features

import (
	"testing"
	"time"
)

func TestTimerStateMachine(t *testing.T) {
	timer := NewFocusTimer()

	if timer.Running {
		t.Errorf("Timer should start paused")
	}
	if timer.Mode != ModeFocus {
		t.Errorf("Expected initial mode ModeFocus, got %v", timer.Mode)
	}

	// Test Toggle (Start)
	timer.Toggle()
	if !timer.Running {
		t.Errorf("Timer should be running after Toggle()")
	}

	// Test Tick
	completed := timer.Tick()
	if completed {
		t.Errorf("Timer should not be completed after 1 tick")
	}

	// Test Pause
	timer.Toggle()
	if timer.Running {
		t.Errorf("Timer should be paused after Toggle()")
	}

	// Test Reset
	timer.Reset()
	if timer.Remaining != 25*time.Minute {
		t.Errorf("Expected 25 minutes after Reset(), got %v", timer.Remaining)
	}
}
