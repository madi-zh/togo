package main

import (
	"database/sql"
	"slices"
	"tasks"
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
