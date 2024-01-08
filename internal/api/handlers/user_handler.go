package handlers

import (
	"boss-payback/internal/api/services"
	"boss-payback/internal/database/db_services"
	"boss-payback/internal/database/models"
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

	if err := db_services.CreateUserInDB(c, &user); err != nil {
		return err
	}

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

	return services.LoginUserResponse(c, &user)
}

func UpdateUsername(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateUsernameRequest); err != nil {
		return err
	}

	if UpdateUsernameRequest.UpdatedUsername == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "New username cannot be empty")
	}

	user, err := helpers.FindUser(UpdateUsernameRequest.Username, UpdateUsernameRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := db_services.UpdateUsernameInDB(c, &user, UpdateUsernameRequest.UpdatedUsername); err != nil {
		return err
	}

	return services.UpdateUsernameResponse(c, UpdateUsernameRequest.UpdatedUsername)
}

func UpdatePassword(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdatePasswordRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(UpdatePasswordRequest.Username, UpdatePasswordRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !helpers.ValidatePassword(UpdatePasswordRequest.UpdatedPassword) {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "New password does not meet criteria")
	}

	hashedPassword, err := helpers.HashPassword(UpdatePasswordRequest.UpdatedPassword)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, "Failed to hash password")
	}

	if err := db_services.UpdatePasswordInDB(c, &user, hashedPassword); err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).SendString("Password updated!")
}

func UpdateUserRole(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &UpdateUserRoleRequest); err != nil {
		return err
	}

	user, err := helpers.FindUser(UpdateUserRoleRequest.Username, UpdateUserRoleRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := db_services.UpdateUserRoleInDB(c, &user, UpdateUserRoleRequest.RoleID); err != nil {
		return err
	}

	return services.UpdateUserRoleResponse(c, UpdateUserRoleRequest.RoleID)
}

func GetUsersByRole(c *fiber.Ctx) error {
	if err := utils.ParseRequestBody(c, &GetUsersByRoleRequest); err != nil {
		return err
	}

	users, err := db_services.UsersByRoleInDB(c, GetUsersByRoleRequest.RoleID)
	if err != nil {
		return err
	}

	return services.UsersByRoleResponse(c, GetUsersByRoleRequest.RoleID, users)
}

func DeleteUser(c *fiber.Ctx) error {
	var request UserRequest

	if err := utils.ParseRequestBody(c, &request); err != nil {
		return err
	}

	user, err := helpers.FindUser(request.Username, request.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := db_services.DeleteUserInDB(c, &user); err != nil {
		return err
	}

	return services.DeleteUserResponse(c, &user)
}
