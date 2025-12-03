package main

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
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

func outputToJson(item *tasks.Task) []byte {
	itemJson, _ := json.Marshal(item)
	return itemJson
}

func outputToJsonList(items []tasks.Task) []byte {
	itemsJson, _ := json.Marshal(tasks.TaskList{Items: items})
	return itemsJson
}

func hello(w http.ResponseWriter, req *http.Request) {
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)
	tasksRepo := tasks.InitRepo(dbSession)
	items := tasksRepo.GetList()
	res := outputToJsonList(items)
	fmt.Fprintf(w, string(res))
}

func main() {
	// opsPtr := flag.String("op", "default", "operation for the todo list")
	// taskPtr := flag.String("task", "<your task name>", "task name")
	// deleteId := flag.Int64("id", -1, "Delete Id")
	// descriptionPtr := flag.String("des", "<your description>", "please provide description")
	// flag.Parse()

	http.HandleFunc("/", hello)
	fmt.Println("Starting server at http://localhost:8090")
	http.ListenAndServe(":8090", nil)
	// todoItem1 := createTask(taskPtr, descriptionPtr)
	// toggleTaskStatus(todoItem1)
	// processOp(opsPtr, *deleteId, todoItem1, tasksRepo)
}
