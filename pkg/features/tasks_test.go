package features

import (
	"testing"
)

func TestTaskListOperations(t *testing.T) {
	tl := NewTaskList()

	initialCount := len(tl.Items)

	// Test Add
	tl.Add("Code Gottage TUI")
	tl.Add("Listen to Lofi beats")

	if len(tl.Items) != initialCount+2 {
		t.Errorf("Expected %d tasks, got %d", initialCount+2, len(tl.Items))
	}

	// Test Toggle Done
	tl.ToggleCurrent()
	if !tl.Items[0].Completed {
		t.Errorf("Expected first task to be completed")
	}

	// Test Delete
	tl.DeleteCurrent()
	if len(tl.Items) != initialCount+1 {
		t.Errorf("Expected %d tasks after deletion, got %d", initialCount+1, len(tl.Items))
	}
}
