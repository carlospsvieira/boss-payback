package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := utils.ParseRequestBody(c, &role); err != nil {
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

func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role

	if err := database.DB.Db.Find(&roles).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    roles,
		"message": "Successfully fetched all roles",
	})
}

func UpdateRoleName(c *fiber.Ctx) error {
	var roleRequest struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	if err := utils.ParseRequestBody(c, &roleRequest); err != nil {
		return err
	}

	if roleRequest.Name == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Role name cannot be empty")
	}

	role, err := helpers.FindRole(roleRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Db.Model(&role).Where("id = ?", role.ID).Update("name", roleRequest.Name).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        roleRequest.Name,
			"description": role.Description,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func UpdateRoleDescription(c *fiber.Ctx) error {
	var roleRequest struct {
		ID          uint   `json:"id"`
		Description string `json:"description"`
	}

	if err := utils.ParseRequestBody(c, &roleRequest); err != nil {
		return err
	}

	role, err := helpers.FindRole(roleRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Db.Model(&role).Where("id = ?", role.ID).Update("description", roleRequest.Description).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"name":        role.Name,
			"description": roleRequest.Description,
		},
		"message": fmt.Sprintf("Role with id %d was updated", role.ID),
	})
}

func DeleteRole(c *fiber.Ctx) error {
	var role models.Role
	if err := utils.ParseRequestBody(c, &role); err != nil {
		return err
	}

	role, err := helpers.FindRole(role.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Db.Unscoped().Delete(&role).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", role.Name),
	})
}
