package domain

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type TaskRepository interface {
	FetchAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task *Task) error
	Update(id string, task *Task) error
	Delete(id string) error
}

type TaskUsecase interface {
	FetchAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task *Task) error
	Update(id string, task *Task) error
	Delete(id string) error
}