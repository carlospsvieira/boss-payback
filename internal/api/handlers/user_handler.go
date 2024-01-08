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

	utils.ParseRequestBody(c, &user)

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

	db_services.CreateUserInDB(c, &user)

	return services.CreateUserResponse(c, &user)
}

func Login(c *fiber.Ctx) error {
	var request UserRequest

	utils.ParseRequestBody(c, &request)

	user, err := helpers.FindUser(request.Username, request.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return services.LoginUserResponse(c, &user)
}

func UpdateUsername(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateUsernameRequest)

	if UpdateUsernameRequest.UpdatedUsername == "" {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, "New username cannot be empty")
	}

	user, err := helpers.FindUser(UpdateUsernameRequest.Username, UpdateUsernameRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	db_services.UpdateUsernameInDB(c, &user, UpdateUsernameRequest.UpdatedUsername)

	return services.UpdateUsernameResponse(c, UpdateUsernameRequest.UpdatedUsername)
}

func UpdatePassword(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdatePasswordRequest)

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

	db_services.UpdatePasswordInDB(c, &user, hashedPassword)

	return c.Status(fiber.StatusNoContent).SendString("Password updated!")
}

func UpdateUserRole(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &UpdateUserRoleRequest)

	user, err := helpers.FindUser(UpdateUserRoleRequest.Username, UpdateUserRoleRequest.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	db_services.UpdateUserRoleInDB(c, &user, UpdateUserRoleRequest.RoleID)

	return services.UpdateUserRoleResponse(c, UpdateUserRoleRequest.RoleID)
}

func GetUsersByRole(c *fiber.Ctx) error {
	utils.ParseRequestBody(c, &GetUsersByRoleRequest)

	users, err := db_services.UsersByRoleInDB(c, GetUsersByRoleRequest.RoleID)
	if err != nil {
		return err
	}

	return services.UsersByRoleResponse(c, GetUsersByRoleRequest.RoleID, users)
}

func DeleteUser(c *fiber.Ctx) error {
	var request UserRequest

	utils.ParseRequestBody(c, &request)

	user, err := helpers.FindUser(request.Username, request.Password)
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	db_services.DeleteUserInDB(c, &user)

	return services.DeleteUserResponse(c, &user)
}
