package usecase

import (
	"errors"
	"testing"

	"github.com/philipos/api/domain"
	"github.com/philipos/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Register(t *testing.T) {
	tests := []struct {
		name          string
		inputUser     *domain.User
		mockBehavior  func(repo *mocks.UserRepository, pwd *mocks.PasswordService)
		expectedError string
	}{
		{
			name:      "Success",
			inputUser: &domain.User{Username: "philip", Password: "123"},
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService) {
				repo.On("GetByUsername", "philip").Return(nil, errors.New("not found"))
				pwd.On("HashPassword", "123").Return("hashed_123", nil)
				repo.On("Create", mock.Anything).Return(nil)
			},
			expectedError: "",
		},
		{
			name:      "Failure - Username Exists",
			inputUser: &domain.User{Username: "philip", Password: "123"},
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService) {
				repo.On("GetByUsername", "philip").Return(&domain.User{Username: "philip"}, nil)
			},
			expectedError: "username already exists",
		},
		{
			name:      "Failure - Empty Fields",
			inputUser: &domain.User{Username: "", Password: ""},
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService) {},
			expectedError: "username and password cannot be empty",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.UserRepository)
			mockPwd := new(mocks.PasswordService)
			mockJWT := new(mocks.JWTService)

			tc.mockBehavior(mockRepo, mockPwd)

			uc := NewUserUsecase(mockRepo, mockPwd, mockJWT)
			err := uc.Register(tc.inputUser)

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			}

			mockRepo.AssertExpectations(t)
			mockPwd.AssertExpectations(t)
		})
	}
}

func TestUserUsecase_Login(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		password      string
		mockBehavior  func(repo *mocks.UserRepository, pwd *mocks.PasswordService, jwt *mocks.JWTService)
		expectedToken string
		expectedError string
	}{
		{
			name:     "Success",
			username: "philip",
			password: "123",
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService, jwt *mocks.JWTService) {
				existingUser := &domain.User{ID: "1", Username: "philip", Password: "hashed_123", Role: "admin"}
				repo.On("GetByUsername", "philip").Return(existingUser, nil)
				pwd.On("ComparePassword", "hashed_123", "123").Return(nil)
				jwt.On("GenerateToken", "1", "admin").Return("valid_token", nil)
			},
			expectedToken: "valid_token",
			expectedError: "",
		},
		{
			name:     "Failure - User Not Found",
			username: "unknown",
			password: "123",
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService, jwt *mocks.JWTService) {
				repo.On("GetByUsername", "unknown").Return(nil, errors.New("not found"))
			},
			expectedToken: "",
			expectedError: "invalid username or password",
		},
		{
			name:     "Failure - Wrong Password",
			username: "philip",
			password: "wrong",
			mockBehavior: func(repo *mocks.UserRepository, pwd *mocks.PasswordService, jwt *mocks.JWTService) {
				existingUser := &domain.User{ID: "1", Username: "philip", Password: "hashed_123"}
				repo.On("GetByUsername", "philip").Return(existingUser, nil)
				pwd.On("ComparePassword", "hashed_123", "wrong").Return(errors.New("mismatch"))
			},
			expectedToken: "",
			expectedError: "invalid username or password",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.UserRepository)
			mockPwd := new(mocks.PasswordService)
			mockJWT := new(mocks.JWTService)

			tc.mockBehavior(mockRepo, mockPwd, mockJWT)

			uc := NewUserUsecase(mockRepo, mockPwd, mockJWT)
			token, err := uc.Login(tc.username, tc.password)

			if tc.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedToken, token)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
				assert.Empty(t, token)
			}

			mockRepo.AssertExpectations(t)
			mockPwd.AssertExpectations(t)
			mockJWT.AssertExpectations(t)
		})
	}
}