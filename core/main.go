package main

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tasks"
)

func getTasks(w http.ResponseWriter, req *http.Request) {
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)
	tasksRepo := tasks.InitRepo(dbSession)
	items := tasksRepo.GetList()
	response := map[string]any{
		"message": "Done",
		"items":   items,
	}
	jsonResponse(w, response, 200)
}

func createTask(w http.ResponseWriter, req *http.Request) {
	var task tasks.Task
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		jsonError(w, "Something wrong", http.StatusBadRequest)
		return
	}
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)
	tasksRepo := tasks.InitRepo(dbSession)

	createdTask := tasksRepo.Add(&task)
	jsonResponse(w, &createdTask, 200)
}

func getTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)

	tasksRepo := tasks.InitRepo(dbSession)

	task := tasksRepo.GetOne(taskId)
	jsonResponse(w, task, 200)
}

func deleteTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	if taskId < 1 {
		jsonError(w, "wrong taskid", http.StatusBadRequest)
		return
	}
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)

	tasksRepo := tasks.InitRepo(dbSession)

	tasksRepo.Delete(taskId)
	jsonResponse(w, nil, 204)
}

func updateTask(w http.ResponseWriter, req *http.Request) {
	var task tasks.Task
	taskId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if taskId < 1 || err != nil {
		jsonError(w, "Wrong id", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		jsonError(w, "Failed to parse body", http.StatusBadRequest)
		return
	}

	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)
	tasksRepo := tasks.InitRepo(dbSession)

	updatedTask := tasksRepo.Update(taskId, &task)

	jsonResponse(w, updatedTask, 200)

}

func jsonError(w http.ResponseWriter, message string, status int) {
	response := map[string]string{
		"error": message,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func jsonResponse(w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /tasks/", getTasks)
	router.HandleFunc("POST /tasks/", createTask)
	router.HandleFunc("GET /tasks/{id}", getTask)
	router.HandleFunc("DELETE /tasks/{id}", deleteTask)
	router.HandleFunc("PUT /tasks/{id}", updateTask)
	fmt.Println("Starting server at http://localhost:8090")
	http.ListenAndServe(":8090", router)
}
