package handlers

import (
	"boss-payback/internal/database/models"
	"boss-payback/internal/database/services"
	"boss-payback/pkg/helpers"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := utils.ParseRequestBody(c, &user); err != nil {
		return err
	}

	if user.Username == "" && user.Email == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Fields empty or missing")
	}

	if !helpers.ValidatePassword(user.Password) {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "Invalid password")
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	user.Password = string(hashedPassword)

	return services.CreateUserResponse(c, &user)
}

func Login(c *fiber.Ctx) error {
	var request UserRequest

	if err := utils.ParseRequestBody(c, &request); err != nil {
		return err
	}

	user, err := helpers.FindUser(request.Username, request.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return services.LoginResponse(c, &user)
}

func UpdateUsername(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateUsernameRequest); err != nil {
		return err
	}

	if UpdateUsernameRequest.NewUsername == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "New username cannot be empty")
	}

	user, err := helpers.FindUser(UpdateUsernameRequest.Username, UpdateUsernameRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return services.UpdateUsernameResponse(c, &user, UpdateUsernameRequest.NewUsername)
}

func UpdatePassword(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdatePasswordRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(UpdatePasswordRequest.Username, UpdatePasswordRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !helpers.ValidatePassword(UpdatePasswordRequest.Password) {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "New password does not meet criteria")
	}

	return services.UpdatePasswordResponse(c, &user, UpdatePasswordRequest.Password)
}

func UpdateUserRole(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateUserRoleRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(UpdateUserRoleRequest.Username, UpdateUserRoleRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return services.UpdateUserRoleResponse(c, &user, UpdateUserRoleRequest.NewRoleID)
}

func GetUsersByRole(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &GetUsersByRoleRequest); err != nil {
		return err
	}

	return services.UsersByRoleResponse(c, GetUsersByRoleRequest.RoleId)
}

func DeleteUser(c *fiber.Ctx) error {
	var userRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.ParseRequestBody(c, &userRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(userRequest.Username, userRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return services.DeleteUserResponse(c, &user)
}
