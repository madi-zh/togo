package tasks

import (
	"db"
	"fmt"
)

type TasksRepository struct {
	session *db.DBSession
}

func InitRepo(session *db.DBSession) *TasksRepository {
	return &TasksRepository{session: session}
}

func (tr *TasksRepository) GetList() []Task {
	var tasks []Task
	rows, err := tr.session.Query("select id, title, description from tasks")
	for rows.Next() {
		var readTask Task
		if err := rows.Scan(&readTask.Id, &readTask.Title, &readTask.Description); err != nil {
			fmt.Println("Error while scanning")
		}
		tasks = append(tasks, readTask)
	}
	if err != nil {
		return tasks
	}
	return tasks
}

func (tr *TasksRepository) Add(t *Task) *Task {
	row := tr.session.QueryRow("insert into tasks (title, description) values ($1, $2) returning id, title, description", t.Title, t.Description)
	var newTask Task
	row.Scan(&newTask.Id, &newTask.Title, &newTask.Description)
	return &newTask
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
