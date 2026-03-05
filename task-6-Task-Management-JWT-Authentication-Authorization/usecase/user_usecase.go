package usecase

import (
	"errors"
	"strings"

	"github.com/philipos/api/domain"
	"github.com/philipos/api/utils" 
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}


func (u *userUsecase) Register(user *domain.User) error {
	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("username and password cannot be empty")
	}

	existingUser, _ := u.userRepo.GetByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to secure password")
	}

	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	return u.userRepo.Create(user)
}

func (u *userUsecase) Login(username string, password string) (string, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password") 
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate authentication token")
	}

	return token, nil
}