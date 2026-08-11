package features

// TaskItem represents a single to-do task.
type TaskItem struct {
	ID        int
	Title     string
	Completed bool
}

// TaskList manages to-do tasks.
type TaskList struct {
	Items     []TaskItem
	Cursor    int
	nextID    int
}

// NewTaskList creates task list with sample cozy tasks.
func NewTaskList() *TaskList {
	tl := &TaskList{
		nextID: 1,
	}
	tl.Add("Brew a warm cup of coffee / tea ☕")
	tl.Add("Write cozy Go code 💻")
	tl.Add("Listen to lofi beats 🎧")
	tl.Add("Pet the sleeping cat 🐈")
	return tl
}

// Add appends a new task.
func (tl *TaskList) Add(title string) {
	if title == "" {
		return
	}
	tl.Items = append(tl.Items, TaskItem{
		ID:        tl.nextID,
		Title:     title,
		Completed: false,
	})
	tl.nextID++
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
