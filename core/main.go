package main

import (
	"db"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tasks"
)

type Server struct {
	repo tasks.Repository
}

func (s *Server) getTasks(w http.ResponseWriter, req *http.Request) {
	items := s.repo.GetList()
	response := map[string]any{
		"message": "Done",
		"items":   items,
	}
	jsonResponse(w, response, 200)
}

func (s *Server) createTask(w http.ResponseWriter, req *http.Request) {
	var task tasks.Task
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		jsonError(w, "Something wrong", http.StatusBadRequest)
		return
	}

	createdTask := s.repo.Add(&task)
	jsonResponse(w, &createdTask, 200)
}

func (s *Server) getTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	task := s.repo.GetOne(taskId)
	jsonResponse(w, task, 200)
}

func (s *Server) deleteTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	if taskId < 1 {
		jsonError(w, "wrong taskid", http.StatusBadRequest)
		return
	}

	s.repo.Delete(taskId)
	jsonResponse(w, nil, 204)
}

func (s *Server) updateTask(w http.ResponseWriter, req *http.Request) {
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

	updatedTask := s.repo.Update(taskId, &task)

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
	dbSession := db.CreateSession()
	defer db.CloseSession(dbSession)
	s := &Server{repo: tasks.InitRepo(dbSession)}

	router.HandleFunc("GET /tasks/", s.getTasks)
	router.HandleFunc("POST /tasks/", s.createTask)
	router.HandleFunc("GET /tasks/{id}", s.getTask)
	router.HandleFunc("DELETE /tasks/{id}", s.deleteTask)
	router.HandleFunc("PUT /tasks/{id}", s.updateTask)
	fmt.Println("Starting server at http://localhost:8090")
	http.ListenAndServe(":8090", router)
}
