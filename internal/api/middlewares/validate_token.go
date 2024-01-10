package middlewares

import (
	"boss-payback/internal/api/auth"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func ValidateGeneralToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Token not provided")
	}

	_, err := auth.TokenCheck(token)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
	}

	return c.Next()
}

func ValidateAdminAndApproverToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Token not provided")
	}

	roleId, err := auth.TokenCheck(token)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
	}

	if roleId != 1 && roleId != 2 {
		return utils.HandleErrorResponse(c, fiber.StatusForbidden, "Forbidden: Role not allowed")
	}

	return c.Next()
}

func ValidateAdminToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Token not provided")
	}

	roleId, err := auth.TokenCheck(token)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
	}

	if roleId != 1 {
		return utils.HandleErrorResponse(c, fiber.StatusForbidden, "Forbidden: Role not allowed")
	}

	return c.Next()
}
