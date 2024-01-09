package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoutes(app *fiber.App) {
	app.Get("/workflows", handlers.GetWorkflows)
	app.Get("/workflow/user", handlers.GetWorkflowByApprover)

	workflowRoutes := app.Group("/workflow", middlewares.ValidateToken)
	{
		workflowRoutes.Post("/workflow/new", handlers.CreateWorkflow)
		// workflowRoutes.Put("/description", handlers.UpdateWorkflowDescription)
		// workflowRoutes.Delete("/delete", handlers.DeleteWorkflow)
	}
}
