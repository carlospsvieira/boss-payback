package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", handlers.GetUsersByRole)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	userRoutes := app.Group("/user", middlewares.ValidateToken)
	{
		userRoutes.Put("/username", handlers.UpdateUsername)
		userRoutes.Put("/password", handlers.UpdatePassword)
		userRoutes.Put("/role", handlers.UpdateUserRole)
		userRoutes.Delete("/delete", handlers.DeleteUser)
	}
}
