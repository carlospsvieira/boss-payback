package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateExpense(c *fiber.Ctx) error {
	var expenseRequest struct {
		UserID      uint
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	if err := utils.ParseRequestBody(c, &expenseRequest); err != nil {
		return err
	}

	var user models.User
	if err := database.DB.Db.Model(&user).Where("id = ? AND logged_in = ?", expenseRequest.UserID, true).First(&user).Error; err != nil {
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
		"message": "New expense created!",
	})
}

func UpdateExpenseAmount(c *fiber.Ctx) error {
	var expenseRequest struct {
		ID     uint    `json:"id"`
		Amount float64 `json:"amount"`
	}

	if err := utils.ParseRequestBody(c, &expenseRequest); err != nil {
		return err
	}

	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", expenseRequest.ID).Update("amount", expenseRequest.Amount).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"description": expenseRequest.Amount,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", expenseRequest.ID),
	})
}

func UpdateExpenseDescription(c *fiber.Ctx) error {
	var expenseRequest struct {
		ID          uint    `json:"id"`
		Description float64 `json:"description"`
	}

	if err := utils.ParseRequestBody(c, &expenseRequest); err != nil {
		return err
	}

	var expense models.Expense
	if err := database.DB.Db.Model(&expense).Where("id = ?", expenseRequest.ID).Update("description", expenseRequest.Description).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"description": expenseRequest.Description,
		},
		"message": fmt.Sprintf("Expense with id %d was updated.", expenseRequest.ID),
	})
}

func GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	if err := database.DB.Db.Find(&expenses).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    expenses,
		"message": "Successfully fetched all expenses",
	})
}

func DeleteExpense(c *fiber.Ctx) error {
	var expenseRequest struct {
		ID uint `json:"id"`
	}

	if err := utils.ParseRequestBody(c, &expenseRequest); err != nil {
		return err
	}

	expense, err := helpers.FindExpense(expenseRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Db.Unscoped().Delete(&expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Expense with id %d was deleted!", expense.ID),
	})
}
