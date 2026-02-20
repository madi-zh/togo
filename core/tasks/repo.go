package tasks

import (
	"db"
	"fmt"
)

type Repository interface {
	GetList() ([]Task, error)
	GetOne(id int64) (*Task, error)
	Add(t *Task) (*Task, error)
	Delete(id int64) bool
	Update(id int64, t *Task) (*Task, error)
}

type TasksRepository struct {
	session *db.DBSession
}

func InitRepo(session *db.DBSession) *TasksRepository {
	return &TasksRepository{session: session}
}

func (tr *TasksRepository) GetList() ([]Task, error) {
	var tasks []Task
	rows, err := tr.session.Query("select id, title, description from tasks")
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

func (tr *TasksRepository) GetOne(id int64) (*Task, error) {
	row := tr.session.QueryRow("select id, title, description from tasks where id = $1", id)
	var newTask Task
	err := row.Scan(&newTask.Id, &newTask.Title, &newTask.Description)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (tr *TasksRepository) Add(t *Task) (*Task, error) {
	row := tr.session.QueryRow("insert into tasks (title, description) values ($1, $2) returning id, title, description", t.Title, t.Description)
	var newTask Task
	err := row.Scan(&newTask.Id, &newTask.Title, &newTask.Description)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (tr *TasksRepository) Delete(id int64) bool {
	if id <= 0 {
		return false
	}
	result, err := tr.session.Exec("delete from tasks where id = $1", id)
	if err != nil {
		fmt.Println("Issue while deletion:", err)
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		fmt.Println("Issue while rowsAffected:", err)
	} else {
		return rowsAffected == 1
	}
	return false
}

func (tr *TasksRepository) Update(id int64, t *Task) (*Task, error) {
	if id <= 0 {
		return nil, nil
	}
	row := tr.session.QueryRow("update tasks set title = $1, description = $2 where id = $3 returning id, title, description", t.Title, t.Description, id)
	var updatedTask Task
	err := row.Scan(&updatedTask.Id, &updatedTask.Title, &updatedTask.Description)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}
