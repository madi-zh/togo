package main

import (
	"db"
	"fmt"
	"net/http"
	"os"
	"tasks"
)

func main() {
	router := http.NewServeMux()
	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	dbSession := db.CreateSession(cfg)
	defer dbSession.Close()
	s := &Server{repo: tasks.InitRepo(dbSession)}

	router.HandleFunc("GET /tasks/", s.getTasks)
	router.HandleFunc("POST /tasks/", s.createTask)
	router.HandleFunc("GET /tasks/{id}", s.getTask)
	router.HandleFunc("DELETE /tasks/{id}", s.deleteTask)
	router.HandleFunc("PUT /tasks/{id}", s.updateTask)
	fmt.Println("Starting server at http://localhost:8090")
	http.ListenAndServe(":8090", logging(router))
}
