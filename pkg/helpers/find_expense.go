package helpers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
)

func FindExpense(id uint) (models.Expense, error) {
	var expense models.Expense
	if err := database.DB.Db.Where("id = ?", id).First(&expense).Error; err != nil {
		return expense, err
	}

	return expense, nil
}
