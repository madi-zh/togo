package main

import (
	"db"
	"fmt"
	"net/http"
	"tasks"
)

func main() {
	router := http.NewServeMux()
	dbSession := db.CreateSession()
	defer dbSession.Close()
	s := &Server{repo: tasks.InitRepo(dbSession)}

	router.HandleFunc("GET /tasks/", s.getTasks)
	router.HandleFunc("POST /tasks/", s.createTask)
	router.HandleFunc("GET /tasks/{id}", s.getTask)
	router.HandleFunc("DELETE /tasks/{id}", s.deleteTask)
	router.HandleFunc("PUT /tasks/{id}", s.updateTask)
	fmt.Println("Starting server at http://localhost:8090")
	http.ListenAndServe(":8090", router)
}
