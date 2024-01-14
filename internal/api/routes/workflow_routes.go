package routes

import (
	"boss-payback/internal/api/handlers"
	"boss-payback/internal/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func WorkflowRoutes(app *fiber.App) {
	app.Get("/workflows", handlers.GetWorkflows)
	app.Get("workflow/approver/:id", handlers.GetWorkflowByApprover)
	app.Get("workflow/expense/:id", handlers.GetWorkflowByExpense)

	adminAndApproverRoutes := app.Group("/workflow", middlewares.ValidateAdminAndApproverToken)
	{
		adminAndApproverRoutes.Post("/new", handlers.CreateWorkflow)
		adminAndApproverRoutes.Put("/:id/status", handlers.UpdateWorkflowStatus)
		adminAndApproverRoutes.Put("/:id/comments", handlers.UpdateWorkflowComments)
	}

	adminRoutes := app.Group("/admin/workflow", middlewares.ValidateAdminAndApproverToken)
	{
		adminRoutes.Delete("/:id/delete", handlers.DeleteWorkflow)
	}

}
