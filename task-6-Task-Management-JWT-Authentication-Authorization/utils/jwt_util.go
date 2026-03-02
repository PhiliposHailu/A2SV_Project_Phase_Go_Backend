package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my_super_secret_key_123")

func GenerateToken(userID string, role string) (string, error) {
	// Payload
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// add the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	return token.SignedString(secretKey)
}
