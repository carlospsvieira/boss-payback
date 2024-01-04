package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	role := new(models.Role)

	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "",
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&role)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"id":   role.ID,
			"name": role.Name,
		},
		"message": fmt.Sprintf("%s role was created!", role.Name),
	})
}

// func GetUsersByRole(c *fiber.Ctx) error {

// }
