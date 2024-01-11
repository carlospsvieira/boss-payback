package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoutes(app *fiber.App) {
	app.Get("/expenses", handlers.GetExpenses)
	app.Get("/expenses/user", handlers.GetExpensesByUser)

	expenseRoutes := app.Group("/expense", middlewares.ValidateGeneralToken)
	{
		expenseRoutes.Post("/new", handlers.CreateExpense)
		expenseRoutes.Put("/amount", handlers.UpdateExpenseAmount)
		expenseRoutes.Put("/description", handlers.UpdateExpenseDescription)
	}

	adminRoutes := app.Group("/admin/expense", middlewares.ValidateAdminToken)
	{
		adminRoutes.Delete("/delete", handlers.DeleteExpense)
	}
}
