package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"tasks"
	"testing"
)

type MockRepo struct {
	tasks []tasks.Task
}

func (r *MockRepo) GetList(ctx context.Context) ([]tasks.Task, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return r.tasks, nil
}

func (r *MockRepo) GetOne(ctx context.Context, id int64) (*tasks.Task, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			return &r.tasks[i], nil
		}
	}
	return nil, &tasks.NotFoundError{Id: id}
}

func (r *MockRepo) Add(ctx context.Context, newTask *tasks.Task) (*tasks.Task, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	r.tasks = append(r.tasks, *newTask)
	return &r.tasks[len(r.tasks)-1], nil
}

func (r *MockRepo) Update(ctx context.Context, id int64, updatedTask *tasks.Task) (*tasks.Task, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			r.tasks[i].Title = updatedTask.Title
			r.tasks[i].Description = updatedTask.Description
			return &r.tasks[i], nil
		}
	}
	return nil, &tasks.NotFoundError{Id: id}
}

func (r *MockRepo) Delete(ctx context.Context, id int64) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	for i := 0; i < len(r.tasks); i++ {
		if r.tasks[i].Id == id {
			r.tasks = slices.Delete(r.tasks, i, i+1)
			return true, nil
		}
	}
	return false, nil
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

func TestGetTask(t *testing.T) {
	mockRepo := &MockRepo{
		tasks: []tasks.Task{
			{Id: 1, Title: "test", Description: "Description"},
		},
	}
	s := &Server{repo: mockRepo}

	tests := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{"valid task", "1", 200},
		{"not found", "999", 404},
		{"invalid id", "abc", 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/tasks/"+tt.id, nil)
			req.SetPathValue("id", tt.id)
			w := httptest.NewRecorder()

			s.getTask(w, req)
			if w.Result().StatusCode != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, w.Result().StatusCode)
			}
			if w.Result().StatusCode == http.StatusOK {
				var responseTask tasks.Task
				err := json.NewDecoder(w.Result().Body).Decode(&responseTask)
				if err != nil {
					t.Errorf("Error fetching task")
				}
			}
		})
	}

}
