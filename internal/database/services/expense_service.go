package services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateExpenseResponse(c *fiber.Ctx, expense *models.Expense) error {
	var user models.User
	if err := database.DB.Db.Model(&user).Where("id = ? AND logged_in = ?", expense.UserID, true).First(&user).Error; err != nil {
		return err
	}

	database.DB.Db.Create(&expense)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"userId":      expense.UserID,
			"amount":      expense.Amount,
			"description": expense.Description,
		},
		"message": "New expense created!",
	})
}

func UpdateExpenseAmountResponse(c *fiber.Ctx, id uint, updatedAmount float64) error {
	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", id).Update("amount", updatedAmount).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"amount": updatedAmount,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", id),
	})
}

func UpdateExpenseDescriptionResponse(c *fiber.Ctx, id uint, updatedDesctiption string) error {
	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", id).Update("description", updatedDesctiption).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"description": updatedDesctiption,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", id),
	})
}

func DeleteExpenseResponse(c *fiber.Ctx, expense *models.Expense) error {
	if err := database.DB.Db.Unscoped().Delete(expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Expense with id %d was deleted!", expense.ID),
	})
}
