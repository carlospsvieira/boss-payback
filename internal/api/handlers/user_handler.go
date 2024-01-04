package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var body struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewUsername string `json:"newUsername"`
	NewPassword string `json:"newPassword"`
}

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(body.Password)
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
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    "",
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
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "",
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("username", body.NewUsername).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}

func UpdatePassword(c *fiber.Ctx) error {
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	hashedPassword, err := utils.HashPassword(body.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("password", hashedPassword).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON("")
}

func DeleteUser(c *fiber.Ctx) error {
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "",
			"message": err.Error(),
		})
	}

	user, err := utils.FindUser(body.Username, body.Password)
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
