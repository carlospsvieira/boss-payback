package services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleResponse(c *fiber.Ctx, role *models.Role) error {
	database.DB.Db.Create(role)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": role.Description,
		},
		"message": fmt.Sprintf("%s role was created!", role.Name),
	})
}

func UpdateRoleNameResponse(c *fiber.Ctx, role *models.Role, updatedRoleName string) error {
	if err := database.DB.Db.Model(role).Where("id = ?", role.ID).Update("name", updatedRoleName).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        updatedRoleName,
			"description": role.Description,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func UpdateRoleDescriptionResponse(c *fiber.Ctx, role *models.Role, updatedRoleDescription string) error {
	if err := database.DB.Db.Model(role).Where("id = ?", role.ID).Update("description", updatedRoleDescription).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": updatedRoleDescription,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func DeleteRoleResponse(c *fiber.Ctx, role *models.Role) error {
	if err := database.DB.Db.Unscoped().Delete(role).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", role.Name),
	})
}
