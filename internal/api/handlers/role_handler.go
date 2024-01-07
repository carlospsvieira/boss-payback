package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/internal/database/services"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := utils.ParseRequestBody(c, &role); err != nil {
		return err
	}

	return services.CreateRoleResponse(c, &role)
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
	if err := utils.ParseRequestBody(c, &UpdateRoleNameRequest); err != nil {
		return err
	}

	if UpdateRoleNameRequest.NewName == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Role name cannot be empty")
	}

	role, err := helpers.FindRole(UpdateRoleNameRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return services.UpdateRoleNameResponse(c, &role, UpdateRoleNameRequest.NewName)
}

func UpdateRoleDescription(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateRoleDescriptionRequest); err != nil {
		return err
	}

	role, err := helpers.FindRole(UpdateRoleDescriptionRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return services.UpdateRoleDescriptionResponse(c, &role, UpdateRoleDescriptionRequest.NewDescription)
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

	return services.DeleteRoleResponse(c, &role)
}
