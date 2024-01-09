package services

import (
	"boss-payback/internal/database/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleResponse(c *fiber.Ctx, role *models.Role) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": role.Description,
		},
		"message": fmt.Sprintf("%s role was created!", role.Name),
	})
}

func GetRolesResponse(c *fiber.Ctx, roles *[]models.Role) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    roles,
		"message": "Successfully fetched all roles",
	})
}

func UpdatedRoleNameResponse(c *fiber.Ctx, role *models.Role, updatedRoleName string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        updatedRoleName,
			"description": role.Description,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func UpdateRoleDescriptionResponse(c *fiber.Ctx, role *models.Role, updatedRoleDescription string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": updatedRoleDescription,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func DeleteRoleResponse(c *fiber.Ctx, role *models.Role) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", role.Name),
	})
}
