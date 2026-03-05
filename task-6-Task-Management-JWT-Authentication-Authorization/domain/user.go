package domain

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
