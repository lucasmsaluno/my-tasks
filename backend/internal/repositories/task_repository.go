package repositories

import (
	"database/sql"

	model "github.com/lucasmsaluno/my-tasks/internal/models"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetAllTasks() ([]model.Task, error)
	GetByID(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2) 
	          RETURNING id, title, description, updated_at`
	return r.db.QueryRow(query, task.Title, task.Description).
		Scan(&task.ID, &task.Title, &task.Description, &task.UpdatedAt)
}

func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	query := `SELECT id, title, description, updated_at FROM tasks WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *model.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2 
	          WHERE id = $3 RETURNING id, title, description, updated_at`
	return r.db.QueryRow(query, task.Title, task.Description, task.ID).
		Scan(&task.ID, &task.Title, &task.Description, &task.UpdatedAt)
}

func (r *taskRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
