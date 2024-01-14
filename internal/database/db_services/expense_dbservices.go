package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExpenseInDB(c *fiber.Ctx, expense *models.Expense) error {
	if err := database.Instance.Db.Create(&expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateExpenseAmountInDB(c *fiber.Ctx, id uint, updatedAmount float64) error {
	var expense models.Expense
	if err := database.Instance.Db.Model(&expense).Where("id = ?", id).Update("amount", updatedAmount).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateExpenseDescriptionInDB(c *fiber.Ctx, id uint, updatedDesctiption string) error {
	var expense models.Expense
	if err := database.Instance.Db.Model(&expense).Where("id = ?", id).Update("description", updatedDesctiption).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetExpensesInDB(c *fiber.Ctx, expenses *[]models.Expense) error {
	if err := database.Instance.Db.Find(expenses).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetExpensesByUserInDB(c *fiber.Ctx, expenses *[]models.Expense, userId uint) error {
	if err := database.Instance.Db.Where("user_id = ?", userId).Find(expenses).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteExpenseInDB(c *fiber.Ctx, id uint) error {
	helpers.DeleteReceipt(c, id)

	var expense models.Expense

	if err := database.Instance.Db.Model(&expense).Where("id = ?", id).Unscoped().Delete(&expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
