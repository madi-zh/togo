package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"tasks"
	"time"
)

func createTask(taskPtr *string) *tasks.TodoItem {
	return &tasks.TodoItem{Task: *taskPtr, Description: "Another description", State: tasks.TaskIncomplete, CreatedAt: time.Now()}
}

func toggleTaskStatus(item *tasks.TodoItem) {
	if item.State == tasks.TaskCompleted {
		item.State = tasks.TaskIncomplete
	} else {
		item.State = tasks.TaskCompleted
	}
}

func processOp(opsPtr *string, itemToProcess []byte) {
	switch *opsPtr {
	case "list":
		fmt.Println(string(itemToProcess))
	case "add":
		fmt.Println("adding")
	default:
		fmt.Println("Available ops are: add, list, delete")
	}
}

func main() {
	opsPtr := flag.String("op", "default", "operation for the todo list")
	taskPtr := flag.String("task", "<your task name>", "task name")
	flag.Parse()
	todoItem1 := createTask(taskPtr)
	toggleTaskStatus(todoItem1)
	itemJson, _ := json.Marshal(todoItem1)
	processOp(opsPtr, itemJson)
}
