package tasks

import "fmt"

type NotFoundError struct {
	Id int64
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Task %d was not found", e.Id)
}
