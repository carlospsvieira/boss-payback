package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/internal/database/services"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExpense(c *fiber.Ctx) error {
	var expense models.Expense
	if err := utils.ParseRequestBody(c, &expense); err != nil {
		return err
	}

	return services.CreateExpenseResponse(c, &expense)
}

func UpdateExpenseAmount(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateExpenseAmountRequest); err != nil {
		return err
	}

	return services.UpdateExpenseAmountResponse(c, UpdateExpenseAmountRequest.ID, UpdateExpenseAmountRequest.Amount)
}

func UpdateExpenseDescription(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateExpenseDescriptionRequest); err != nil {
		return err
	}

	return services.UpdateExpenseDescriptionResponse(c, UpdateExpenseDescriptionRequest.ID, UpdateExpenseDescriptionRequest.Description)
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
	if err := utils.ParseRequestBody(c, &DeleteExpenseRequest); err != nil {
		return err
	}

	expense, err := helpers.FindExpense(DeleteExpenseRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return services.DeleteExpenseResponse(c, &expense)
}
