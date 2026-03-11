package usecase

import (
	"errors"
	"testing"

	"github.com/philipos/api/domain"
	"github.com/philipos/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskUsecase_Create(t *testing.T) {
	tests := []struct {
		name          string
		inputTask     *domain.Task
		mockBehavior  func(m *mocks.TaskRepository)
		expectedError string
	}{
		{
			name:      "Success",
			inputTask: &domain.Task{Title: "Learn Testing", Status: "Pending"},
			mockBehavior: func(m *mocks.TaskRepository) {
				m.On("Create", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name:      "Failure - Empty Title",
			inputTask: &domain.Task{Title: "", Status: "Pending"},
			mockBehavior: func(m *mocks.TaskRepository) {
			},
			expectedError: "task title is required",
		},
		{
			name:      "Failure - DB Error",
			inputTask: &domain.Task{Title: "Valid Title", Status: "Pending"},
			mockBehavior: func(m *mocks.TaskRepository) {
				m.On("Create", mock.Anything).Return(errors.New("database down"))
			},
			expectedError: "database down",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TaskRepository)
			tc.mockBehavior(mockRepo)

			uc := NewTaskUsecase(mockRepo)

			err := uc.Create(tc.inputTask)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
