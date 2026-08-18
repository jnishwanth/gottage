package features

import (
	"testing"
)

func TestTaskListOperations(t *testing.T) {
	tl := NewTaskList()

	initialCount := len(tl.Items)

	// Test Add
	tl.Add("Code Gottage TUI")
	tl.AddWithPriority("Listen to Lofi beats", PriorityHigh)

	if len(tl.Items) != initialCount+2 {
		t.Errorf("Expected %d tasks, got %d", initialCount+2, len(tl.Items))
	}

	// Test Priority cycle
	p := tl.CyclePriorityCurrent()
	if p != PriorityLow {
		t.Errorf("Expected PriorityLow after cycling PriorityHigh, got %v", p)
	}

	// Test Toggle Done
	tl.ToggleCurrent()
	if !tl.Items[0].Completed {
		t.Errorf("Expected first task to be completed")
	}

	// Test Filtering
	tl.CycleFilter() // FilterActive
	active := tl.FilteredItems()
	if len(active) == 0 {
		t.Errorf("Expected active tasks, got 0")
	}

	// Test Delete
	tl.DeleteCurrent()
	if len(tl.Items) != initialCount+1 {
		t.Errorf("Expected %d tasks after deletion, got %d", initialCount+1, len(tl.Items))
	}
}
