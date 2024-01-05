package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var role models.Role

func CreateRole(c *fiber.Ctx) error {
	if err := helpers.ParseRequestBody(c, &role); err != nil {
		return err
	}

	database.DB.Db.Create(&role)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": role.Description,
		},
		"message": fmt.Sprintf("%s role was created!", role.Name),
	})
}

func DeleteRole(c *fiber.Ctx) error {
	if err := helpers.ParseRequestBody(c, &role); err != nil {
		return err
	}

	role, err := helpers.FindRole(role.Name)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Db.Unscoped().Delete(&role).Error; err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", role.Name),
	})
}
