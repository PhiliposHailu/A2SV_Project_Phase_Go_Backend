package mocks

import "github.com/stretchr/testify/mock"

type PasswordService struct {
	mock.Mock
}

func (m *PasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordService) ComparePassword(hash string, password string) error {
	args := m.Called(hash, password)
	return args.Error(0)
}