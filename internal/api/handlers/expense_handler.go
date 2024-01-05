package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateExpense(c *fiber.Ctx) error {
	var expenseRequest struct {
		UserID      uint
		Username    string  `json:"username"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	if err := helpers.ParseRequestBody(c, &expenseRequest); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.Db.Model(&user).Where("username = ? AND logged_in = ?", expenseRequest.Username, true).First(&user).Error; err != nil {
		return err
	}

	var expense models.Expense
	expense.Amount = expenseRequest.Amount
	expense.Description = expenseRequest.Description
	expense.UserID = user.ID

	database.DB.Db.Create(&expense)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"userId":      expense.UserID,
			"amount":      expense.Amount,
			"description": expense.Description,
		},
		"message": fmt.Sprintf("New expense created by %s!", expenseRequest.Username),
	})
}
