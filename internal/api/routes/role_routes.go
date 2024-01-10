package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App) {

	adminRoutes := app.Group("/admin/role", middlewares.ValidateAdminToken)
	{
		adminRoutes.Get("/", handlers.GetRoles)
		adminRoutes.Post("/new", handlers.CreateRole)
		adminRoutes.Put("/name", handlers.UpdateRoleName)
		adminRoutes.Put("/description", handlers.UpdateRoleDescription)
		adminRoutes.Delete("/delete", handlers.DeleteRole)
	}
}
