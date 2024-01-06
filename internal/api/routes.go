package api

import (
	"boss-payback/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Related to User
	app.Get("/users", handlers.GetUsersByRole)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Put("/username", handlers.UpdateUsername)
	app.Put("/password", handlers.UpdatePassword)
	app.Delete("/delete-user", handlers.DeleteUser)

	// Related to Role
	app.Get("/roles", handlers.GetRoles)
	app.Post("/new-role", handlers.CreateRole)
	app.Put("/role-name", handlers.UpdateRoleName)
	app.Put("/role-description", handlers.UpdateRoleDescription)
	app.Delete("/delete-role", handlers.DeleteRole)

	// Related to Expense
	app.Get("/expenses", handlers.GetExpenses)
	app.Post("/new-expense", handlers.CreateExpense)
	app.Put("/expense-amount", handlers.UpdateExpenseAmount)
	app.Put("/expense-description", handlers.UpdateExpenseDescription)
	app.Delete("/delete-expense", handlers.DeleteExpense)
}
