package tasks

import (
	"fmt"
	"time"
)

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

type Task struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      TaskState `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TaskList struct {
	Items []Task `json:"items"`
}

func (t Task) String() string {
	return fmt.Sprintf("task: %d %s d: %s", t.Id, t.Title, t.Description)
}
