package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoutes(app *fiber.App) {
	app.Get("/workflows", handlers.GetWorkflows)
	app.Get("/workflow/approver", handlers.GetWorkflowByApprover)
	app.Get("/workflow/expense", handlers.GetWorkflowByExpense)

	workflowRoutes := app.Group("/workflow", middlewares.ValidateToken)
	{
		workflowRoutes.Post("/new", handlers.CreateWorkflow)
		workflowRoutes.Put("/status", handlers.UpdateWorkflowStatus)
		workflowRoutes.Put("/comments", handlers.UpdateWorkflowComments)
		workflowRoutes.Delete("/delete", handlers.DeleteWorkflow)
	}
}
