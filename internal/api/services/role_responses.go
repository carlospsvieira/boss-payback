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

func UpdatedRoleNameResponse(c *fiber.Ctx, id uint, updatedRoleName string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name": updatedRoleName,
		},
		"message": fmt.Sprintf("Role with id %d was updated", id),
	})
}

func UpdateRoleDescriptionResponse(c *fiber.Ctx, id uint, updatedRoleDescription string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"description": updatedRoleDescription,
		},
		"message": fmt.Sprintf("Role with id %d was updated", id),
	})
}

func DeleteRoleResponse(c *fiber.Ctx, role *models.Role) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Role with id %d was deleted!", role.ID),
	})
}
