package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBSession struct {
	*sql.DB
}

func CreateSession() *DBSession {
	psqlInfo := fmt.Sprintf("host=localhost port=5430 user=test " +
		"password=test dbname=database sslmode=disable")
	if conn, err := sql.Open("postgres", psqlInfo); err != nil {
		log.Fatal("Issue while opening conn", err)
	} else {
		return &DBSession{conn}
	}
	return nil
}

func CloseSession(session *DBSession) {
	session.Close()
}

// func dbOps(db *sql.DB) {

// 	fmt.Println(newTask)
// 	rows, err := db.Query("select id, title, description from tasks")
// 	if err != nil {
// 		log.Fatal("err")
// 	}
// 	for rows.Next() {
// 		var tempTask tasks.Task
// 		if err := rows.Scan(&tempTask.Id, &tempTask.Title, &tempTask.Description); err != nil {
// 			fmt.Println("Error while scanning")
// 		}
// 		fmt.Println(tempTask)
// 	}

// 	defer rows.Close()
// }
