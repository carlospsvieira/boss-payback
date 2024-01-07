package api

import (
	"boss-payback/internal/api/routes"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	routes.UserRoutes(app)
	routes.RoleRoutes(app)
	routes.ExpenseRoutes(app)
}
