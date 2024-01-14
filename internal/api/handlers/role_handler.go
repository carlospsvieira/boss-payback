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
	var request UpdateRoleNameRequest
	utils.ParseRequestBody(c, &request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if request.Name == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Role name cannot be empty")
	}

	db_services.UpdateRoleNameInDB(c, id, request.Name)

	return services.UpdatedRoleNameResponse(c, id, request.Name)
}

func UpdateRoleDescription(c *fiber.Ctx) error {
	var request UpdateRoleDescriptionRequest
	utils.ParseRequestBody(c, &request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.UpdateRoleDescriptionInDB(c, id, request.Description)

	return services.UpdateRoleDescriptionResponse(c, id, request.Description)
}

func DeleteRole(c *fiber.Ctx) error {
	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.DeleteRoleInDB(c, id)

	return services.DeleteRoleResponse(c, id)
}
