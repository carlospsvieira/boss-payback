package utils

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func FindUser(username string, password string) (models.User, error) {
	var user models.User
	if err := database.DB.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, fmt.Errorf("invalid credentials")
	}

	return user, nil
}
