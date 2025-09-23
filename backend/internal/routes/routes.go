package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmsaluno/my-tasks/internal/handlers"
	"github.com/lucasmsaluno/my-tasks/internal/repositories"
	"github.com/lucasmsaluno/my-tasks/internal/usecases"
)

func SetupRoutes(repo repositories.TaskRepository) http.Handler {

	usecase := usecases.NewTaskUsecase(repo)
	taskHandler := handlers.NewTaskHandler(usecase)
	r := mux.NewRouter()

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	return r
}
