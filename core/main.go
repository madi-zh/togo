package main

import (
	"db"
	"encoding/json"
	"flag"
	"fmt"
	"tasks"
	"time"
)

func createTask(taskPtr *string, desPtr *string) *tasks.Task {
	return &tasks.Task{Title: *taskPtr, Description: *desPtr, Status: tasks.TaskIncomplete, CreatedAt: time.Now()}
}

func toggleTaskStatus(item *tasks.Task) {
	if item.Status == tasks.TaskCompleted {
		item.Status = tasks.TaskIncomplete
	} else {
		item.Status = tasks.TaskCompleted
	}
}

func processOp(opsPtr *string, id int64, item *tasks.Task, tasksRepo *tasks.TasksRepository) {
	switch *opsPtr {
	case "getList":
		result := tasksRepo.GetList()
		for _, res := range result {
			outputToJson(&res)
		}
	case "create":
		result := tasksRepo.Add(item)
		outputToJson(result)
	case "delete":
		tasksRepo.Delete(id)
	case "update":
		fmt.Println("update")
	case "getOne":
		result := tasksRepo.GetOne(id)
		outputToJson(result)
	default:
		fmt.Println("Available ops are: add, list, delete")
	}

}

func outputToJson(item *tasks.Task) {
	itemJson, _ := json.Marshal(item)
	fmt.Println(string(itemJson))
}

func main() {
	opsPtr := flag.String("op", "default", "operation for the todo list")
	taskPtr := flag.String("task", "<your task name>", "task name")
	deleteId := flag.Int64("id", -1, "Delete Id")
	descriptionPtr := flag.String("des", "<your description>", "please provide description")
	flag.Parse()
	// fetch postgresurl conn string
	// connStr := os.Getenv("POSTGRES_URL")
	psqlInfo := fmt.Sprintf("host=localhost port=5430 user=test " +
		"password=test dbname=database sslmode=disable")
	fmt.Println(psqlInfo)
	dbSession := db.CreateSession(psqlInfo)
	defer db.CloseSession(dbSession)
	tasksRepo := tasks.InitRepo(dbSession)
	todoItem1 := createTask(taskPtr, descriptionPtr)
	toggleTaskStatus(todoItem1)
	processOp(opsPtr, *deleteId, todoItem1, tasksRepo)
}
