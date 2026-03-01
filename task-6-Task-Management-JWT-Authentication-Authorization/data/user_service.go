package data

import (
	"context"
	"errors"

	"github.com/philipos/api/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterService(newUser *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)
	// Default Role 
	if newUser.Role == "" {
		newUser.Role = "user"
	}

	_, err = UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return errors.New("could not create user")
	}
	return nil
}
