package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type Task struct {
	id          uint32
	title       string
	description string
}

func (t Task) String() string {
	return fmt.Sprintf("task: %s d: %s", t.title, t.description)
}

func main() {
	dbConnStr := os.Getenv("POSTGRES_URL")
	fmt.Println(dbConnStr)
	if dbConnStr == "" {
		log.Fatal("POSTGRES_URL is not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal("{DBERR}")
	}
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("PWITHS", err)
	}
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://migrations",
	// 	"database", driver)
	if err != nil {
		log.Fatal(err)
	}
	// if err := m.Up(); err != nil {
	// 	log.Fatal("UP")
	// 	log.Fatal(err)
	// }

	row := db.QueryRow("insert into tasks (title, description) values ('example1', 'test2') returning id, title, description")
	var newTask Task
	err = row.Scan(&newTask.id, &newTask.title, &newTask.description)
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
