package infrastructure

import (
	"github.com/philipos/api/domain"
	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() domain.PasswordService {
	return &passwordService{}
}

func (s *passwordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func (s *passwordService) ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
