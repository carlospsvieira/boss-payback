package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleInDB(c *fiber.Ctx, role *models.Role) error {
	if err := database.Instance.Db.Create(role).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetRolesInDB(c *fiber.Ctx, roles *[]models.Role) error {
	if err := database.Instance.Db.Find(roles).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateRoleNameInDB(c *fiber.Ctx, id uint, updatedRoleName string) error {
	var role models.Role
	if err := database.Instance.Db.Model(&role).Where("id = ?", id).Update("name", updatedRoleName).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateRoleDescriptionInDB(c *fiber.Ctx, id uint, updatedRoleDescription string) error {
	var role models.Role
	if err := database.Instance.Db.Model(&role).Where("id = ?", id).Update("description", updatedRoleDescription).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteRoleInDB(c *fiber.Ctx, id uint) error {
	var role models.Role
	if err := database.Instance.Db.Model(&role).Where("id = ?", id).Unscoped().Delete(&role).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
