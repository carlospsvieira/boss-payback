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
	form, err := c.MultipartForm()
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	amount, userId, filePath, description, err := helpers.ExpenseForm(c, form)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	expense := models.Expense{
		UserID:       userId,
		Description:  description,
		Amount:       amount,
		ReceiptImage: filePath,
	}

	db_services.CreateExpenseInDB(c, &expense)

	return services.CreateExpenseResponse(c, &expense)
}

func UpdateExpenseAmount(c *fiber.Ctx) error {
	var request UpdateExpenseAmountRequest
	utils.ParseRequestBody(c, &request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if request.Amount == 0 {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Amount cannot be missing or zero")
	}

	db_services.UpdateExpenseAmountInDB(c, id, request.Amount)

	return services.UpdateExpenseAmountResponse(c, request.Amount, id)
}

func UpdateExpenseDescription(c *fiber.Ctx) error {
	var request UpdateExpenseDescriptionRequest
	utils.ParseRequestBody(c, &request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.UpdateExpenseDescriptionInDB(c, id, request.Description)

	return services.UpdateExpenseDescriptionResponse(c, request.Description, id)
}

func GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense

	db_services.GetExpensesInDB(c, &expenses)

	return services.GetExpensesResponse(c, &expenses)
}

func GetExpensesByUser(c *fiber.Ctx) error {
	userId, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var expenses []models.Expense

	db_services.GetExpensesByUserInDB(c, &expenses, userId)

	return services.GetExpensesByUserResponse(c, &expenses, userId)
}

func DeleteExpense(c *fiber.Ctx) error {
	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.DeleteExpenseInDB(c, id)

	return services.DeleteExpenseResponse(c, id)
}
