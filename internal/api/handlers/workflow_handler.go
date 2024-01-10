package handlers

import (
	"boss-payback/internal/api/services"
	"boss-payback/internal/database/db_services"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkflow(c *fiber.Ctx) error {
	var workflow models.Workflow

	utils.ParseRequestBody(c, &workflow)

	db_services.CreateWorkflowInDB(c, &workflow)

	return services.CreateWorkflowResponse(c, &workflow)
}

func GetWorkflows(c *fiber.Ctx) error {
	var workflows []models.Workflow

	db_services.GetWorkflowsInDB(c, &workflows)

	return services.GetWorkflowsResponse(c, &workflows)
}

func GetWorkflowByApprover(c *fiber.Ctx) error {
	var request GetWorkflowByApproverRequest
	var workflows []models.Workflow
	utils.ParseRequestBody(c, &request)

	db_services.GetWorkflowByApproverInDB(c, &workflows, request.ApproverID)

	return services.GetWorkflowByApproverResponse(c, &workflows, request.ApproverID)
}

func GetWorkflowByExpense(c *fiber.Ctx) error {
	var request GetWorkflowByExpenseRequest
	var workflows []models.Workflow

	utils.ParseRequestBody(c, &request)

	db_services.GetWorkflowByExpenseInDB(c, &workflows, request.ExpenseID)

	return services.GetWorkflowByExpenseResponse(c, &workflows, request.ExpenseID)
}

func UpdateWorkflowStatus(c *fiber.Ctx) error {
	var request UpdateWorkflowStatusRequest

	utils.ParseRequestBody(c, request)

	db_services.UpdateWorkflowStatusInDB(c, request.ID, request.Status)

	return services.UpdateWorkflowStatusResponse(c, request.ID, request.Status)
}

func UpdateWorkflowComments(c *fiber.Ctx) error {
	var request UpdateWorkflowCommentsRequest

	utils.ParseRequestBody(c, request)

	db_services.UpdateWorkflowCommentsInDB(c, request.ID, request.Comments)

	return services.UpdateWorkflowCommentsResponse(c, request.ID, request.Comments)
}

func DeleteWorkflow(c *fiber.Ctx) error {
	var workflow models.Workflow

	utils.ParseRequestBody(c, &workflow)

	db_services.DeleteWorkflowInDB(c, &workflow)

	return services.DeleteWorkflowResponse(c, &workflow)
}
