package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserByRole struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint   `json:"roleId"`
}

func CreateUserInDB(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Create(user).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user")
	}

	return nil
}

func LoginUserInDB(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("logged_in", true).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func UpdateUsernameInDB(c *fiber.Ctx, user *models.User, updatedUsername string) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("username", updatedUsername).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdatePasswordInDB(c *fiber.Ctx, user *models.User, hashedPassword []byte) error {

	if err := database.DB.Db.Model(&user).Where("id = ?", user.ID).Update("password", hashedPassword).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateUserRoleInDB(c *fiber.Ctx, user *models.User, updatedRole uint) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("role_id", updatedRole).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UsersByRoleInDB(c *fiber.Ctx, roleId uint) ([]models.User, error) {
	var users []models.User

	if err := database.DB.Db.Where("role_id = ?", roleId).Find(&users).Error; err != nil {
		return nil, utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return users, nil
}

func DeleteUserInDB(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Unscoped().Delete(user).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
