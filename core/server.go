package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tasks"
	"time"
)

type Server struct {
	repo tasks.Repository
}

func (s *Server) getTasks(w http.ResponseWriter, req *http.Request) {
	items, err := s.repo.GetList(req.Context())
	if err != nil {
		jsonError(w, "Issue while fetching", http.StatusInternalServerError)
		return
	}
	response := map[string]any{
		"message": "Done",
		"items":   items,
	}
	jsonResponse(w, response, http.StatusOK)
}

func (s *Server) createTask(w http.ResponseWriter, req *http.Request) {
	var task tasks.Task
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&task)

	if err != nil {
		jsonError(w, "Something wrong", http.StatusBadRequest)
		return
	}

	createdTask, err := s.repo.Add(req.Context(), &task)
	if err != nil {
		jsonError(w, "Couldn't find", http.StatusBadRequest)
		return
	}
	jsonResponse(w, &createdTask, http.StatusOK)
}

func (s *Server) getTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := parseId(req)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	task, err := s.repo.GetOne(req.Context(), taskId)
	if err != nil {
		jsonError(w, "Couldn't find", http.StatusBadRequest)
		return
	}
	jsonResponse(w, task, http.StatusOK)
}

func (s *Server) deleteTask(w http.ResponseWriter, req *http.Request) {
	taskId, err := parseId(req)
	if err != nil {
		jsonError(w, "Id is not int", http.StatusBadRequest)
		return
	}
	if taskId < 1 {
		jsonError(w, "wrong taskid", http.StatusBadRequest)
		return
	}

	s.repo.Delete(req.Context(), taskId)
	jsonResponse(w, nil, http.StatusNoContent)
}

func (s *Server) updateTask(w http.ResponseWriter, req *http.Request) {
	var task tasks.Task
	taskId, err := parseId(req)
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

	updatedTask, err := s.repo.Update(req.Context(), taskId, &task)
	if err != nil {
		jsonError(w, "Couldn't find", http.StatusBadRequest)
		return
	}
	jsonResponse(w, updatedTask, http.StatusOK)

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

func parseId(req *http.Request) (int64, error) {
	return strconv.ParseInt(req.PathValue("id"), 10, 64)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fmt.Printf("Method: %s\n", req.Method)
		fmt.Printf("Path: %s\n", req.URL.Path)
		next.ServeHTTP(w, req)
		duration := time.Since(start)
		fmt.Printf("Duration %v\n", duration)
	})
}
