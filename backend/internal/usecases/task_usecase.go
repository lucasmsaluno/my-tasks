package usecases

import (
	model "github.com/lucasmsaluno/my-tasks/internal/models"
	"github.com/lucasmsaluno/my-tasks/internal/repositories"
)

type TaskUsecase interface {
	CreateTask(task *model.Task) error
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	repo repositories.TaskRepository
}

func NewTaskUsecase(repo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) CreateTask(task *model.Task) error {
	return u.repo.Create(task)
}

func (u *taskUsecase) GetAllTasks() ([]model.Task, error) {
	return u.repo.GetAllTasks()
}

func (u *taskUsecase) GetTaskByID(id int) (*model.Task, error) {
	return u.repo.GetByID(id)
}

func (u *taskUsecase) UpdateTask(task *model.Task) error {
	return u.repo.Update(task)
}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.Delete(id)
}
