package services

import (
	"boss-payback/internal/api/auth"
	"boss-payback/internal/database/models"
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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"email":    user.Email,
			"roleId":   user.RoleID,
		},
		"message": fmt.Sprintf("%s was created!", user.Username),
	})
}

func LoginUserResponse(c *fiber.Ctx, user *models.User) error {
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

func UpdateUsernameResponse(c *fiber.Ctx, updatedUsername string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"updatedUsername": updatedUsername,
		},
		"message": fmt.Sprintf("Username updated to %s", updatedUsername),
	})
}

func UpdateUserRoleResponse(c *fiber.Ctx, updatedRole uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"updatedRole": updatedRole,
		},
		"message": fmt.Sprintf("Username updated to %d", updatedRole),
	})
}

func UsersByRoleResponse(c *fiber.Ctx, roleId uint, users []models.User) error {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", user.Username),
	})
}
