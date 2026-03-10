package usecase

import (
	"errors"
	"strings"

	"github.com/philipos/api/domain"
)

type userUsecase struct {
	userRepo        domain.UserRepository
	passwordService domain.PasswordService
	tokenService    domain.JWTService
}

func NewUserUsecase(repo domain.UserRepository, passService domain.PasswordService, givenTokenService domain.JWTService) domain.UserUsecase {
	return &userUsecase{
		userRepo:        repo,
		passwordService: passService,
		tokenService:    givenTokenService,
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

	hashedPassword, err := u.passwordService.HashPassword(user.Password)
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

	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := u.tokenService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate authentication token")
	}

	return token, nil
}
