package handlers

import (
	"boss-payback/internal/api/services"
	"boss-payback/internal/database/db_services"
	"boss-payback/internal/database/models"
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

	db_services.GetRolesInDB(c, &roles)

	return services.GetRolesResponse(c, &roles)
}

func UpdateRoleName(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateRoleNameRequest)

	if UpdateRoleNameRequest.Name == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Role name cannot be empty")
	}

	db_services.UpdateRoleNameInDB(c, UpdateRoleNameRequest.ID, UpdateRoleNameRequest.Name)

	return services.UpdatedRoleNameResponse(c, UpdateRoleNameRequest.ID, UpdateRoleNameRequest.Name)
}

func UpdateRoleDescription(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateRoleDescriptionRequest)

	db_services.UpdateRoleDescriptionInDB(c, UpdateRoleDescriptionRequest.ID, UpdateRoleDescriptionRequest.Description)

	return services.UpdateRoleDescriptionResponse(c, UpdateRoleDescriptionRequest.ID, UpdateRoleDescriptionRequest.Description)
}

func DeleteRole(c *fiber.Ctx) error {
	var role models.Role
	utils.ParseRequestBody(c, &role)

	db_services.DeleteRoleInDB(c, &role)

	return services.DeleteRoleResponse(c, &role)
}
