package main

import (
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"slices"
	"tasks"
	"testing"
)

type MockRepo struct {
	tasks []tasks.Task
}

func (r *MockRepo) GetList() []tasks.Task {
	return r.tasks
}

func (r *MockRepo) GetOne(id int64) (*tasks.Task, error) {
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			return &r.tasks[i], nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *MockRepo) Add(newTask *tasks.Task) (*tasks.Task, error) {
	r.tasks = append(r.tasks, *newTask)
	return &r.tasks[len(r.tasks)-1], nil
}

func (r *MockRepo) Update(id int64, updatedTask *tasks.Task) (*tasks.Task, error) {
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			r.tasks[i].Title = updatedTask.Title
			r.tasks[i].Description = updatedTask.Description
			return &r.tasks[i], nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *MockRepo) Delete(id int64) bool {
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			r.tasks = slices.Delete(r.tasks, i, i+1)
			return true
		}
	}
	return false
}

func TestGetTasks(t *testing.T) {
	mockRepo := &MockRepo{
		tasks: []tasks.Task{
			{Id: 1, Title: "test", Description: "Description"},
		},
	}
	s := &Server{repo: mockRepo}

	req := httptest.NewRequest("GET", "/tasks/", nil)
	w := httptest.NewRecorder()
	s.getTasks(w, req)
	result := w.Result()
	if result.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", result.StatusCode)
	}

	defer result.Body.Close()
	var response struct {
		Message string       `json:"message"`
		Items   []tasks.Task `json:"items"`
	}
	err := json.NewDecoder(result.Body).Decode(&response)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	for i, task := range response.Items {
		if mockRepo.tasks[i].Id != task.Id {
			t.Errorf("Got error: %d", i)
		}
	}
}
