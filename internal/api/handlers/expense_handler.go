package handlers

import (
	"boss-payback/internal/api/services"
	"boss-payback/internal/database/db_services"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateExpense(c *fiber.Ctx) error {
	// file, err := c.FormFile("receiptImage")
	// if err != nil {
	// 	return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	// }

	// receiptImageURL, err := helpers.SaveUploadedFile(file)
	// if err != nil {
	// 	return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	// }

	var expense models.Expense
	if err := utils.ParseRequestBody(c, &expense); err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := db_services.CreateExpenseInDB(c, &expense /*receiptImageURL */); err != nil {
		return err
	}

	return services.CreateExpenseResponse(c, &expense)
}

func UpdateExpenseAmount(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateExpenseAmountRequest)

	db_services.UpdateExpenseAmountInDB(c, UpdateExpenseAmountRequest.ID, UpdateExpenseAmountRequest.Amount)

	return services.UpdateExpenseAmountResponse(c, UpdateExpenseAmountRequest.Amount, UpdateExpenseAmountRequest.ID)
}

func UpdateExpenseDescription(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateExpenseDescriptionRequest)

	db_services.UpdateExpenseDescriptionInDB(c, UpdateExpenseDescriptionRequest.ID, UpdateExpenseDescriptionRequest.Description)

	return services.UpdateExpenseDescriptionResponse(c, UpdateExpenseDescriptionRequest.Description, UpdateExpenseDescriptionRequest.ID)
}

func GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	db_services.GetExpensesInDB(c, expenses)

	return services.GetExpensesResponse(c, expenses)
}

func GetExpensesByUser(c *fiber.Ctx) error {
	var expenses []models.Expense
	utils.ParseRequestBody(c, &GetExpensesByUserRequest)

	db_services.GetExpensesByUserInDB(c, expenses, GetExpensesByUserRequest.UserID)

	return services.GetExpensesByUserResponse(c, expenses, GetExpensesByUserRequest.UserID)
}

func DeleteExpense(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &DeleteExpenseRequest)

	expense, err := helpers.FindExpense(DeleteExpenseRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	db_services.DeleteExpenseInDB(c, &expense)

	return services.DeleteExpenseResponse(c, &expense)
}
