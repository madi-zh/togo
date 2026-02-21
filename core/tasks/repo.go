package tasks

import (
	"context"
	"db"
	"fmt"
)

type Repository interface {
	GetList(ctx context.Context) ([]Task, error)
	GetOne(ctx context.Context, id int64) (*Task, error)
	Add(ctx context.Context, t *Task) (*Task, error)
	Delete(ctx context.Context, id int64) (bool, error)
	Update(ctx context.Context, id int64, t *Task) (*Task, error)
}

type TasksRepository struct {
	session *db.DBSession
}

func InitRepo(session *db.DBSession) *TasksRepository {
	return &TasksRepository{session: session}
}

func (tr *TasksRepository) GetList(ctx context.Context) ([]Task, error) {
	var tasks []Task
	rows, err := tr.session.QueryContext(ctx, "select id, title, description from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var readTask Task
		if err := rows.Scan(&readTask.Id, &readTask.Title, &readTask.Description); err != nil {
			fmt.Println("Error while scanning")
		}
		tasks = append(tasks, readTask)
	}
	return tasks, nil
}

func (tr *TasksRepository) GetOne(ctx context.Context, id int64) (*Task, error) {
	row := tr.session.QueryRowContext(ctx, "select id, title, description from tasks where id = $1", id)
	var newTask Task
	err := row.Scan(&newTask.Id, &newTask.Title, &newTask.Description)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (tr *TasksRepository) Add(ctx context.Context, t *Task) (*Task, error) {
	row := tr.session.QueryRowContext(ctx, "insert into tasks (title, description) values ($1, $2) returning id, title, description", t.Title, t.Description)
	var newTask Task
	err := row.Scan(&newTask.Id, &newTask.Title, &newTask.Description)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (tr *TasksRepository) Delete(ctx context.Context, id int64) (bool, error) {
	if id <= 0 {
		return false, nil
	}
	result, err := tr.session.ExecContext(ctx, "delete from tasks where id = $1", id)
	if err != nil {
		return false, err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return false, err
	} else {
		return rowsAffected == 1, nil
	}
}

func (tr *TasksRepository) Update(ctx context.Context, id int64, t *Task) (*Task, error) {
	if id <= 0 {
		return nil, nil
	}
	row := tr.session.QueryRowContext(ctx, "update tasks set title = $1, description = $2 where id = $3 returning id, title, description", t.Title, t.Description, id)
	var updatedTask Task
	err := row.Scan(&updatedTask.Id, &updatedTask.Title, &updatedTask.Description)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}
