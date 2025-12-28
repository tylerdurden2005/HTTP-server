package app

import (
	"log"
	"net/http"
	"os"
	"webServerEx/internal/db/inmemory"
	"webServerEx/internal/handlers"
	"webServerEx/internal/middleware"
	"webServerEx/internal/service"
)

type App struct {
	handler *handlers.Handler
}

func NewApp() *App {
	storage := inmemory.NewStorage()
	repository := service.NewRepository(storage)
	serviceTasks := service.NewTasksService(repository)
	handler := handlers.NewHandler(serviceTasks)
	return &App{handler: handler}
}

func (a *App) Start() {
	log.SetOutput(os.Stdout)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos", a.handler.CreateTask)
	mux.HandleFunc("GET /todos", a.handler.GetAllTasks)
	mux.HandleFunc("GET /todos/{id}", a.handler.GetTask)
	mux.HandleFunc("PUT /todos/{id}", a.handler.UpdateTask)
	mux.HandleFunc("DELETE /todos/{id}", a.handler.DeleteTask)
	mux.HandleFunc("DELETE /todos", a.handler.DeleteTasks)
	loggedMux := middleware.LoggingMiddleware(mux)
	log.Println("HTTP-Server starting on :8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
