package helpers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
)

func FindRole(name string) (models.Role, error) {
	var role models.Role
	if err := database.DB.Db.Where("name = ?", name).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}
