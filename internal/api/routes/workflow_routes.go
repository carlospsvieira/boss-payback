package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoutes(app *fiber.App) {
	app.Get("/expenses/user", handlers.GetExpensesByUser)

	expenseRoutes := app.Group("/expense", middlewares.ValidateToken)
	{
		expenseRoutes.Post("/new", handlers.CreateExpense)
		expenseRoutes.Put("/amount", handlers.UpdateExpenseAmount)
		expenseRoutes.Put("/description", handlers.UpdateExpenseDescription)
		expenseRoutes.Delete("/delete", handlers.DeleteExpense)
	}
}
