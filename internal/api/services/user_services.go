package services

import (
	"boss-payback/internal/api/auth"
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserByRole struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint   `json:"roleId"`
}

func CreateUserResponse(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Create(user).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"email":    user.Email,
			"roleId":   user.RoleID,
		},
		"message": fmt.Sprintf("%s was created!", user.Username),
	})
}

func LoginResponse(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("logged_in", true).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"roleId":   user.RoleID,
			"token":    token,
			"loggedIn": true,
		},
		"message": fmt.Sprintf("%s logged in successfully!", user.Username),
	})
}

func UpdateUsernameResponse(c *fiber.Ctx, user *models.User, updatedUsername string) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("username", updatedUsername).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"updatedUsername": updatedUsername,
		},
		"message": fmt.Sprintf("Username updated to %s", updatedUsername),
	})
}

func UpdatePasswordResponse(c *fiber.Ctx, user *models.User, updatedPassword string) error {
	hashedPassword, err := helpers.HashPassword(updatedPassword)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	if err := database.DB.Db.Model(&user).Where("id = ?", user.ID).Update("password", hashedPassword).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusNoContent).SendString("Password updated!")
}

func UpdateUserRoleResponse(c *fiber.Ctx, user *models.User, updatedRole uint) error {
	if err := database.DB.Db.Model(user).Where("id = ?", user.ID).Update("role_id", updatedRole).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"updatedRole": updatedRole,
		},
		"message": fmt.Sprintf("Username updated to %d", updatedRole),
	})
}

func UsersByRoleResponse(c *fiber.Ctx, roleId uint) error {
	var users []models.User
	if err := database.DB.Db.Where("role_id = ?", roleId).Find(&users).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	userCh := make(chan UserByRole)
	defer close(userCh)

	for _, user := range users {
		go func(u models.User) {
			userCh <- UserByRole{
				Username: u.Username,
				Email:    u.Email,
				RoleId:   u.RoleID,
			}
		}(user)
	}

	var response []UserByRole
	for range users {
		user := <-userCh
		response = append(response, user)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"users": response,
		},
		"message": fmt.Sprintf("Successfully fetched all users with role id %d", roleId),
	})
}

func DeleteUserResponse(c *fiber.Ctx, user *models.User) error {
	if err := database.DB.Db.Unscoped().Delete(user).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", user.Username),
	})
}
