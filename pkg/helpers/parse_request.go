package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func ParseRequestBody(c *fiber.Ctx, target interface{}) error {
	if err := c.BodyParser(target); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return nil
}
