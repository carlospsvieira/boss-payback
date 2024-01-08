package handlers

import (
	"boss-payback/internal/api/services"
	"boss-payback/internal/database/db_services"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	utils.ParseRequestBody(c, &role)

	db_services.CreateRoleInDB(c, &role)

	return services.CreateRoleResponse(c, &role)
}

func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role

	db_services.GetRolesInDB(c, roles)

	return services.GetRolesResponse(c, roles)
}

func UpdateRoleName(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateRoleNameRequest)

	if UpdateRoleNameRequest.Name == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Role name cannot be empty")
	}

	role, err := helpers.FindRole(UpdateRoleNameRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	db_services.UpdateRoleNameInDB(c, &role, UpdateRoleNameRequest.Name)

	return services.UpdatedRoleNameResponse(c, &role, UpdateRoleNameRequest.Name)
}

func UpdateRoleDescription(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateRoleDescriptionRequest)

	role, err := helpers.FindRole(UpdateRoleDescriptionRequest.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	db_services.UpdateRoleDescriptionInDB(c, &role, UpdateRoleDescriptionRequest.Description)

	return services.UpdateRoleDescriptionResponse(c, &role, UpdateRoleDescriptionRequest.Description)
}

func DeleteRole(c *fiber.Ctx) error {
	var role models.Role
	utils.ParseRequestBody(c, &role)

	role, err := helpers.FindRole(role.ID)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	db_services.DeleteRoleInDB(c, &role)

	return services.DeleteRoleResponse(c, &role)
}
