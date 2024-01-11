package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

/*
This system does not support initially more nor less than 3 roles (ID 1, 2 and 3).
You should either leave as they are or update the existing ones,
so there is no conflict with any default ID.
In case you need more or fewer than the existing roles,
make sure you check the middlewares that validate auth for each role.
*/
func RoleRoutes(app *fiber.App) {
	app.Get("/roles", handlers.GetRoles)

	adminRoutes := app.Group("/admin/role", middlewares.ValidateAdminToken)
	{
		adminRoutes.Post("role/new", handlers.CreateRole)
		adminRoutes.Put("/name", handlers.UpdateRoleName)
		adminRoutes.Put("/description", handlers.UpdateRoleDescription)
		adminRoutes.Delete("/delete", handlers.DeleteRole)
	}
}
