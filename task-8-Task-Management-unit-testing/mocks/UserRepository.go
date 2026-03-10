package mocks

import (
	"github.com/philipos/api/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}