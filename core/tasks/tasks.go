package tasks

import "time"

type TaskState int

const (
	TaskIncomplete TaskState = iota
	TaskCompleted
)

var stateName = map[TaskState]string{
	TaskIncomplete: "Incomplete",
	TaskCompleted:  "Done",
}

func (ts TaskState) String() string {
	return stateName[ts]
}

type TodoItem struct {
	Task        string    `json:"name"`
	Description string    `json:"description"`
	State       TaskState `json:"status"`
	CreatedAt   time.Time
}
