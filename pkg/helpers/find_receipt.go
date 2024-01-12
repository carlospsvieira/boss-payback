package helpers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func DeleteReceipt(c *fiber.Ctx, expenseId uint) error {
	var expense models.Expense
	if err := database.Instance.Db.Where("id = ?", expenseId).First(&expense).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := os.Remove(expense.ReceiptImage); err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
