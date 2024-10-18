package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//! TODO Generate JWT Key

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

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

func GenerateRefreshToken(username string) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func GenerateNewAccessToken(refreshToken *jwt.Token) (string, error) {
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		return "", jwt.ErrInvalidKey
	}
	expirationTime := time.Now().Add(45 * time.Second)
	newClaims := jwt.MapClaims{
		"username": claims["username"],
		"exp":      expirationTime.Unix(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newTokenString, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return newTokenString, nil
}
