package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", handlers.GetUsersByRole)
	app.Post("/login", handlers.Login)

	adminRoutes := app.Group("/admin/user", middlewares.ValidateAdminToken)
	{
		adminRoutes.Post("/register", handlers.Register)
		adminRoutes.Put("/username", handlers.UpdateUsername)
		adminRoutes.Put("/password", handlers.UpdatePassword)
		adminRoutes.Put("/role", handlers.UpdateUserRole)
		adminRoutes.Delete("/delete", handlers.DeleteUser)
	}
}
