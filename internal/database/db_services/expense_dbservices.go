package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExpenseInDB(c *fiber.Ctx, expense *models.Expense /*receiptImageURL string */) error {
	var user models.User
	if err := database.DB.Db.Model(&user).Where("id = ?", expense.UserID).First(&user).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	// expense.ReceiptImage = receiptImageURL

	if err := database.DB.Db.Create(&expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateExpenseAmountInDB(c *fiber.Ctx, id uint, updatedAmount float64) error {
	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", id).Update("amount", updatedAmount).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateExpenseDescriptionInDB(c *fiber.Ctx, id uint, updatedDesctiption string) error {
	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", id).Update("description", updatedDesctiption).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetExpensesInDB(c *fiber.Ctx, expenses *[]models.Expense) error {
	if err := database.DB.Db.Find(expenses).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetExpensesByUserInDB(c *fiber.Ctx, expenses *[]models.Expense, userId uint) error {
	if err := database.DB.Db.Where("user_id = ?", userId).Find(expenses).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteExpenseInDB(c *fiber.Ctx, expense *models.Expense) error {
	helpers.DeleteReceipt(c, expense.ID)

	if err := database.DB.Db.Model(expense).Where("id = ?", expense.ID).Unscoped().Delete(expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
