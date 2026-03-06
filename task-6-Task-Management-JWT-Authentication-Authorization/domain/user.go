package domain

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
}

type UserUsecase interface {
	Register(user *User) error
	Login(username string, password string) (string, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash string, password string) error
}

type JWTService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(token string) (jwt.MapClaims, error)
}
