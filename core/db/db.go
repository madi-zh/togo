package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	id          uint32
	title       string
	description string
}

func (t Task) String() string {
	return fmt.Sprintf("task: %d %s d: %s", t.id, t.title, t.description)
}

func dbOps(db *sql.DB) {
	row := db.QueryRow("insert into tasks (title, description) values ('example1', 'test2') returning id, title, description")
	var newTask Task
	err := row.Scan(&newTask.id, &newTask.title, &newTask.description)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newTask)
	rows, err := db.Query("select id, title, description from tasks")
	if err != nil {
		log.Fatal("err")
	}
	for rows.Next() {
		var tempTask Task
		if err := rows.Scan(&tempTask.id, &tempTask.title, &tempTask.description); err != nil {
			log.Fatal("SCANEr")
		}
		fmt.Println(tempTask)
	}

	defer rows.Close()
}
