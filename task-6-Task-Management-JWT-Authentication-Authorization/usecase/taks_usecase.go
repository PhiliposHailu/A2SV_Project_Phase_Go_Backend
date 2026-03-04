package usecase

import (
	"errors"
	"strings"

	"github.com/philipos/clean-task-api/domain" // Your module path
)


// class creation 
type taskUsecase struct {
	taskRepo domain.TaskRepository
}

// class constructor 
func NewTaskUsecase(repo domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepo: repo,
	}
}

func (u *taskUsecase) FetchAll() ([]domain.Task, error) {
	return u.taskRepo.FetchAll()
}

func (u *taskUsecase) GetByID(id string) (*domain.Task, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("task ID cannot be empty")
	}
	return u.taskRepo.GetByID(id)
}

func (u *taskUsecase) Create(task *domain.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("task title is required")
	}

	if task.Status == "" {
		task.Status = "Pending"
	}

	return u.taskRepo.Create(task)
}

func (u *taskUsecase) Update(id string, task *domain.Task) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("task ID cannot be empty")
	}
	return u.taskRepo.Update(id, task)
}

func (u *taskUsecase) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("task ID cannot be empty")
	}
	return u.taskRepo.Delete(id)
}