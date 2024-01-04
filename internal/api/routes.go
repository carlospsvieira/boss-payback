package api

import (
	"boss-payback/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Related to User
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Put("/update-username", handlers.UpdateUsername)
	app.Put("/update-password", handlers.UpdatePassword)
	app.Delete("/delete-user", handlers.DeleteUser)

	// Related to Role
	app.Post("/new-role", handlers.CreateRole)
}
