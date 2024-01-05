package handlers

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := helpers.ParseRequestBody(c, &user); err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
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

	if err := helpers.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("logged_in", true).Error; err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"roleId":   user.RoleID,
			"loggedIn": true,
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

	if err := helpers.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
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

	if err := helpers.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	hashedPassword, err := helpers.HashPassword(userRequest.NewPassword)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	if err := database.DB.Db.Model(&user).Where("username = ?", user.Username).Update("password", hashedPassword).Error; err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
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

	if err := helpers.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	var users []models.User
	if err := database.DB.Db.Where("role_id = ?", userRequest.RoleId).Find(&users).Error; err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	userCh := make(chan userResponse)
	defer close(userCh)

	for _, user := range users {
		go func(u models.User) {
			userCh <- userResponse{
				Username: u.Username,
				Email:    u.Email,
				RoleId:   u.RoleID,
			}
		}(user)
	}

	var response []userResponse
	for range users {
		user := <-userCh
		response = append(response, user)
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

	if err := helpers.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := database.DB.Db.Unscoped().Delete(&user).Error; err != nil {
		return helpers.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("%s was deleted!", user.Username),
	})
}
