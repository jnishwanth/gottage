package features

// Priority defines task urgency levels.
type Priority int

const (
	PriorityNormal Priority = iota
	PriorityHigh
	PriorityLow
)

func (p Priority) String() string {
	switch p {
	case PriorityHigh:
		return "HIGH 🔴"
	case PriorityLow:
		return "LOW 🟢"
	default:
		return "MED 🟡"
	}
}

// TaskFilter defines task display filtering modes.
type TaskFilter int

const (
	FilterAll TaskFilter = iota
	FilterActive
	FilterCompleted
)

func (f TaskFilter) String() string {
	switch f {
	case FilterActive:
		return "Active 📖"
	case FilterCompleted:
		return "Completed ✓"
	default:
		return "All Tasks 📋"
	}
}

// TaskItem represents a single to-do task.
type TaskItem struct {
	ID        int
	Title     string
	Completed bool
	Priority  Priority
}

// TaskList manages to-do tasks.
type TaskList struct {
	Items     []TaskItem
	Cursor    int
	Filter    TaskFilter
	nextID    int
}

// NewTaskList creates task list with sample cozy tasks.
func NewTaskList() *TaskList {
	tl := &TaskList{
		Filter: FilterAll,
		nextID: 1,
	}
	tl.AddWithPriority("Brew a warm cup of coffee / tea ☕", PriorityHigh)
	tl.AddWithPriority("Write cozy Go code 💻", PriorityNormal)
	tl.AddWithPriority("Listen to lofi beats 🎧", PriorityLow)
	tl.AddWithPriority("Pet the sleeping cat 🐈", PriorityNormal)
	return tl
}

// Add appends a new task with normal priority.
func (tl *TaskList) Add(title string) {
	tl.AddWithPriority(title, PriorityNormal)
}

// AddWithPriority appends a new task with specified priority.
func (tl *TaskList) AddWithPriority(title string, priority Priority) {
	if title == "" {
		return
	}
	tl.Items = append(tl.Items, TaskItem{
		ID:        tl.nextID,
		Title:     title,
		Completed: false,
		Priority:  priority,
	})
	tl.nextID++
}

// CyclePriorityCurrent cycles priority of task at cursor.
func (tl *TaskList) CyclePriorityCurrent() Priority {
	if len(tl.Items) == 0 || tl.Cursor < 0 || tl.Cursor >= len(tl.Items) {
		return PriorityNormal
	}
	tl.Items[tl.Cursor].Priority = (tl.Items[tl.Cursor].Priority + 1) % 3
	return tl.Items[tl.Cursor].Priority
}

// CycleFilter cycles task filtering mode.
func (tl *TaskList) CycleFilter() TaskFilter {
	tl.Filter = (tl.Filter + 1) % 3
	return tl.Filter
}

// FilteredItems returns tasks matching current Filter setting.
func (tl *TaskList) FilteredItems() []TaskItem {
	if tl.Filter == FilterAll {
		return tl.Items
	}
	var res []TaskItem
	for _, item := range tl.Items {
		if tl.Filter == FilterActive && !item.Completed {
			res = append(res, item)
		} else if tl.Filter == FilterCompleted && item.Completed {
			res = append(res, item)
		}
	}
	return res
}

// ToggleCurrent toggles completed status of task at cursor.
func (tl *TaskList) ToggleCurrent() {
	if len(tl.Items) == 0 || tl.Cursor < 0 || tl.Cursor >= len(tl.Items) {
		return
	}
	tl.Items[tl.Cursor].Completed = !tl.Items[tl.Cursor].Completed
}

// DeleteCurrent removes task at cursor.
func (tl *TaskList) DeleteCurrent() {
	if len(tl.Items) == 0 || tl.Cursor < 0 || tl.Cursor >= len(tl.Items) {
		return
	}
	tl.Items = append(tl.Items[:tl.Cursor], tl.Items[tl.Cursor+1:]...)
	if tl.Cursor >= len(tl.Items) && len(tl.Items) > 0 {
		tl.Cursor = len(tl.Items) - 1
	}
}

// MoveUp moves cursor up.
func (tl *TaskList) MoveUp() {
	if tl.Cursor > 0 {
		tl.Cursor--
	}
}

// MoveDown moves cursor down.
func (tl *TaskList) MoveDown() {
	if tl.Cursor < len(tl.Items)-1 {
		tl.Cursor++
	}
}
