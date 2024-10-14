package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//! TODO Generate JWT Key

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)

	//? JWT Token Claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
