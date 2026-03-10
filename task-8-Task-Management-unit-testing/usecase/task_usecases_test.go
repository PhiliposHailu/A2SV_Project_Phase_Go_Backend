package usecase

import (
	"testing"

	"github.com/philipos/api/domain"
	"github.com/philipos/api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	task := &domain.Task{Title: "Test Task", Status: "Pending"}

	mockRepo.On("Create", task).Return(nil)

	err := uc.Create(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	task := &domain.Task{Title: "", Status: "Pending"}

	err := uc.Create(task)

	assert.Error(t, err)
	assert.Equal(t, "task title is required", err.Error())
	mockRepo.AssertNotCalled(t, "Create")
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)
	uc := NewTaskUsecase(mockRepo)

	expectedTask := &domain.Task{ID: "123", Title: "Test"}

	mockRepo.On("GetByID", "123").Return(expectedTask, nil)

	task, err := uc.GetByID("123")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)
	mockRepo.AssertExpectations(t)
}