package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func TokenCheck(tokenString string) (uint, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	roleID, ok := claims["roleID"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid roleID in token")
	}

	return uint(roleID), nil
}
