package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ExpenseRoutes(app *fiber.App) {
	app.Get("/expenses", handlers.GetExpenses)

	expenseRoutes := app.Group("/expense", middlewares.ValidateToken)
	{
		expenseRoutes.Post("/new-expense", handlers.CreateExpense)
		expenseRoutes.Put("/expense-amount", handlers.UpdateExpenseAmount)
		expenseRoutes.Put("/expense-description", handlers.UpdateExpenseDescription)
		expenseRoutes.Delete("/delete-expense", handlers.DeleteExpense)
	}
}
