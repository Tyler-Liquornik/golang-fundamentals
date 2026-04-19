package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// For production code, should be stored in an environment variable
const secretKey = "secret"

// GenerateJWTToken generates a JWT token from an email and userId that lasts for 2 hours
func GenerateJWTToken(email string, userId int64) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
