package services

import (
	"boss-payback/internal/database/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateExpenseResponse(c *fiber.Ctx, expense *models.Expense) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"userId":      expense.UserID,
			"amount":      expense.Amount,
			"description": expense.Description,
		},
		"message": "New expense created!",
	})
}

func UpdateExpenseAmountResponse(c *fiber.Ctx, updatedAmount float64, id uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"amount": updatedAmount,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", id),
	})
}

func UpdateExpenseDescriptionResponse(c *fiber.Ctx, updatedDescription string, id uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"description": updatedDescription,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", id),
	})
}

func GetExpensesResponse(c *fiber.Ctx, expenses *[]models.Expense) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    expenses,
		"message": "Successfully fetched all expenses",
	})
}

func GetExpensesByUserResponse(c *fiber.Ctx, expenses *[]models.Expense, userId uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    expenses,
		"message": fmt.Sprintf("Successfully fetched all expenses from user with id %d", userId),
	})
}

func DeleteExpenseResponse(c *fiber.Ctx, id uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Expense with id %d was deleted!", id),
	})
}
