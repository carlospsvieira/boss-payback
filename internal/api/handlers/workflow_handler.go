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
	var workflows []models.Workflow

	approverId, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.GetWorkflowByApproverInDB(c, &workflows, approverId)

	return services.GetWorkflowByApproverResponse(c, &workflows, approverId)
}

func GetWorkflowByExpense(c *fiber.Ctx) error {
	var workflows []models.Workflow

	expenseId, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.GetWorkflowByExpenseInDB(c, &workflows, expenseId)

	return services.GetWorkflowByExpenseResponse(c, &workflows, expenseId)
}

func UpdateWorkflowStatus(c *fiber.Ctx) error {
	var request UpdateWorkflowStatusRequest

	utils.ParseRequestBody(c, request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.UpdateWorkflowStatusInDB(c, id, request.Status)

	return services.UpdateWorkflowStatusResponse(c, id, request.Status)
}

func UpdateWorkflowComments(c *fiber.Ctx) error {
	var request UpdateWorkflowCommentsRequest

	utils.ParseRequestBody(c, request)

	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.UpdateWorkflowCommentsInDB(c, id, request.Comments)

	return services.UpdateWorkflowCommentsResponse(c, id, request.Comments)
}

func DeleteWorkflow(c *fiber.Ctx) error {
	id, err := utils.ParseUint(c.Params("id"))
	if err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db_services.DeleteWorkflowInDB(c, id)

	return services.DeleteWorkflowResponse(c, id)
}
