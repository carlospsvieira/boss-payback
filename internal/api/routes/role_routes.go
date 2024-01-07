package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {
	app.Get("/roles", handlers.GetRoles)

	roleRoutes := app.Group("/role", middlewares.ValidateToken)
	{
		roleRoutes.Post("/new", handlers.CreateRole)
		roleRoutes.Put("/name", handlers.UpdateRoleName)
		roleRoutes.Put("/description", handlers.UpdateRoleDescription)
		roleRoutes.Delete("/delete", handlers.DeleteRole)
	}
}
