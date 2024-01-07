package middlewares

import (
	"boss-payback/internal/api/auth"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Token not provided")
	}

	if err := auth.TokenCheck(token); err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
	}

	return c.Next()
}
