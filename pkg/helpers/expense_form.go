package helpers

import (
	"boss-payback/pkg/utils"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ExpenseForm(c *fiber.Ctx, form *multipart.Form) (float64, uint, string, string, error) {
	description := form.Value["description"][0]
	amountStr := form.Value["amount"][0]
	userIDStr := form.Value["userId"][0]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, 0, "", "", err
	}

	userId, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, 0, "", "", err
	}

	file, err := c.FormFile("receiptImage")
	if err != nil {
		return 0, 0, "", "", err
	}

	filePath := os.Getenv("UPLOADS_DIR_PATH") + utils.GenerateUUID() + file.Filename
	if err := c.SaveFile(file, filePath); err != nil {
		return 0, 0, "", "", err
	}

	return amount, uint(userId), filePath, description, nil
}
