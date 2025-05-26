package utils

import (
	"todo-backend/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// takes plain text password and returns hash value
// using one way hashing algorithm
// goal: store user password in database securely
func GenerateHashPassword(password string) (string, error) {
	// generate a hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashPassword(password, hash string) bool {
	// compare the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// takes a JWT token string and returns the claims
// Claims are key-value paris that represent info being transmitetd between parties
// func used to validate if token is legitimate and retrieve info contained in it
func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil

}
