package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoutes(app *fiber.App) {

	workflowRoutes := app.Group("/workflow", middlewares.ValidateAdminAndApproverToken)
	{
		workflowRoutes.Get("/workflows", handlers.GetWorkflows)
		workflowRoutes.Get("/approver", handlers.GetWorkflowByApprover)
		workflowRoutes.Get("/expense", handlers.GetWorkflowByExpense)
		workflowRoutes.Post("/new", handlers.CreateWorkflow)
		workflowRoutes.Put("/status", handlers.UpdateWorkflowStatus)
		workflowRoutes.Put("/comments", handlers.UpdateWorkflowComments)
		workflowRoutes.Delete("/delete", handlers.DeleteWorkflow)
	}
}
