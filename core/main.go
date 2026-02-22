package main

import (
	"context"
	"db"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tasks"
	"time"
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
	server := http.Server{
		Addr:    ":8090",
		Handler: logging(router),
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		server.ListenAndServe()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)
}
