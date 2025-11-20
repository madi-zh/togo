package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"tasks"
	"time"
)

func createTask(taskPtr *string, desPtr *string) *tasks.TodoItem {
	return &tasks.TodoItem{Task: *taskPtr, Description: *desPtr, State: tasks.TaskIncomplete, CreatedAt: time.Now()}
}

func toggleTaskStatus(item *tasks.TodoItem) {
	if item.State == tasks.TaskCompleted {
		item.State = tasks.TaskIncomplete
	} else {
		item.State = tasks.TaskCompleted
	}
}

func processOp(opsPtr *string, item *tasks.TodoItem) {
	switch *opsPtr {
	case "getList":
		fmt.Println(string(item.Task))
	case "create":
		fmt.Println("create")
	case "delete":
		fmt.Println("delete")
	case "update":
		fmt.Println("update")
	case "getOne":
		fmt.Println("getOne")
	default:
		fmt.Println("Available ops are: add, list, delete")
	}
}

func outputToJson(item *tasks.TodoItem) {
	itemJson, _ := json.Marshal(item)
	fmt.Println(itemJson)
}

func main() {
	opsPtr := flag.String("op", "default", "operation for the todo list")
	taskPtr := flag.String("task", "<your task name>", "task name")
	descriptionPtr := flag.String("des", "<your description>", "please provide description")
	flag.Parse()

	todoItem1 := createTask(taskPtr, descriptionPtr)
	toggleTaskStatus(todoItem1)
	processOp(opsPtr, todoItem1)
	outputToJson(todoItem1)
}
