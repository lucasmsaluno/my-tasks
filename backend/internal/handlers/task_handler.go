package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/lucasmsaluno/my-tasks/internal/models"
	usecases "github.com/lucasmsaluno/my-tasks/internal/usecases"
)

type TaskHandler struct {
	UseCase usecases.TaskUsecase
}

func NewTaskHandler(uc usecases.TaskUsecase) *TaskHandler {
	return &TaskHandler{UseCase: uc}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := h.UseCase.CreateTask(&task); err != nil {
		http.Error(w, "error on create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.UseCase.GetAllTasks()
	if err != nil {
		http.Error(w, "error on get tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := h.UseCase.GetTaskByID(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	task.ID = id

	if err := h.UseCase.UpdateTask(&task); err != nil {
		http.Error(w, "error on update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.UseCase.DeleteTask(id); err != nil {
		http.Error(w, "error on delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
