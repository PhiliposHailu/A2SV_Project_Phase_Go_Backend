package usecase

import (
	"errors"
	"testing"

	"github.com/philipos/api/domain"
	"github.com/philipos/api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockPwd := new(mocks.PasswordService)
	mockJWT := new(mocks.JWTService)

	uc := NewUserUsecase(mockRepo, mockPwd, mockJWT)

	user := &domain.User{Username: "philip", Password: "password123"}

	mockRepo.On("GetByUsername", "philip").Return(nil, errors.New("not found"))
	mockPwd.On("HashPassword", "password123").Return("hashed_pw", nil)
	mockRepo.On("Create", user).Return(nil)

	err := uc.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, "hashed_pw", user.Password)
	assert.Equal(t, "user", user.Role)

	mockRepo.AssertExpectations(t)
	mockPwd.AssertExpectations(t)
}

func TestRegister_UsernameExists(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockPwd := new(mocks.PasswordService)
	mockJWT := new(mocks.JWTService)

	uc := NewUserUsecase(mockRepo, mockPwd, mockJWT)

	user := &domain.User{Username: "philip", Password: "123"}
	existingUser := &domain.User{Username: "philip"}

	mockRepo.On("GetByUsername", "philip").Return(existingUser, nil)

	err := uc.Register(user)

	assert.Error(t, err)
	assert.Equal(t, "username already exists", err.Error())
	
	mockPwd.AssertNotCalled(t, "HashPassword")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockPwd := new(mocks.PasswordService)
	mockJWT := new(mocks.JWTService)

	uc := NewUserUsecase(mockRepo, mockPwd, mockJWT)

	existingUser := &domain.User{ID: "1", Username: "philip", Password: "hashed_pw", Role: "admin"}

	mockRepo.On("GetByUsername", "philip").Return(existingUser, nil)
	mockPwd.On("ComparePassword", "hashed_pw", "password123").Return(nil)
	mockJWT.On("GenerateToken", "1", "admin").Return("mock_token", nil)

	token, err := uc.Login("philip", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "mock_token", token)

	mockRepo.AssertExpectations(t)
	mockPwd.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}