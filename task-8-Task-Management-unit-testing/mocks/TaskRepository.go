package mocks

import (
	"github.com/philipos/api/domain"
	"github.com/stretchr/testify/mock"
)

type TaskRepository struct {
	mock.Mock
}

func (m *TaskRepository) FetchAll() ([]domain.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskRepository) GetByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *TaskRepository) Create(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepository) Update(id string, task *domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *TaskRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}