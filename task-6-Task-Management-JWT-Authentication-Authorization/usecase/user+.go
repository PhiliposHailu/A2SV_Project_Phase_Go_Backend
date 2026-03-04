package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/philipos/api/models"
	"github.com/philipos/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterService(newUser *models.User) error {
	newUser.Username = strings.ToLower(newUser.Username)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)
	if newUser.Role == "" {
		newUser.Role = "user"
	}

	_, err = UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}
	return nil
}

func LoginService(username string, password string) (string, error) {
	var user models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(user.ID.Hex(), user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
