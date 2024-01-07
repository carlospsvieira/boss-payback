package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func TokenCheck(tokenString string) error {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
