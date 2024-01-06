package helpers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
)

func FindRole(id uint) (models.Role, error) {
	var role models.Role
	if err := database.DB.Db.Where("id = ?", id).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}
