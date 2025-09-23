package model

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}
