package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
	}

	user.Password = string(hashedPassword)

	database.DB.Db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"roleId":   user.RoleID,
		},
		"message": fmt.Sprintf("%s was created!", user.Username),
	})
}

func Login(c *fiber.Ctx) error {
	var userRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"roleId":   user.RoleID,
		},
		"message": fmt.Sprintf("%s logged in successfully!", user.Username),
	})
}

func UpdateUsername(c *fiber.Ctx) error {
	var userRequest struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		NewUsername string `json:"newUsername"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("username", userRequest.NewUsername).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}

func UpdatePassword(c *fiber.Ctx) error {
	var userRequest struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		NewPassword string `json:"newPassword"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	hashedPassword, err := utils.HashPassword(userRequest.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("password", hashedPassword).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}

func GetUsersByRole(c *fiber.Ctx) error {
	var userRequest struct {
		RoleId uint `json:"roleId"`
	}

	type userResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		RoleId   uint   `json:"roleId"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var users []models.User
	if err := database.DB.Db.Where("role_id = ?", userRequest.RoleId).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	var response []userResponse
	for _, user := range users {
		response = append(response, userResponse{
			Username: user.Username,
			Email:    user.Email,
			RoleId:   user.RoleID,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"users": response,
		},
		"message": fmt.Sprintf("Successfully fetched all users with role id %d", userRequest.RoleId),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var userRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := database.DB.Db.Unscoped().Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", user.Username),
	})
}
