package services

import (
	"boss-payback/internal/database/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkflowResponse(c *fiber.Ctx, workflow *models.Workflow) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"approver": workflow.ApproverID,
			"expense":  workflow.ExpenseID,
			"status":   workflow.Status,
			"comment":  workflow.Comments,
		},
		"message": "New workflow created!",
	})
}

func GetWorkflowsResponse(c *fiber.Ctx, workflows *[]models.Workflow) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    workflows,
		"message": "Successfully fetched all workflows",
	})
}

func GetWorkflowByApproverResponse(c *fiber.Ctx, workflows *[]models.Workflow, approverId uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    workflows,
		"message": fmt.Sprintf("Successfully fetched workflows from Approver %d", approverId),
	})
}

func GetWorkflowByExpenseResponse(c *fiber.Ctx, workflows *[]models.Workflow, expenseId uint) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    workflows,
		"message": fmt.Sprintf("Successfully fetched workflows by Expense %d", expenseId),
	})
}

func UpdateWorkflowStatusResponse(c *fiber.Ctx, id uint, updatedStatus string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"status": updatedStatus,
		},
		"message": fmt.Sprintf("Workflow id %d status was updated.", id),
	})
}

func UpdateWorkflowCommentsResponse(c *fiber.Ctx, id uint, updatedComments string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"status": updatedComments,
		},
		"message": fmt.Sprintf("Workflow id %d comments was updated.", id),
	})
}

func DeleteWorkflowResponse(c *fiber.Ctx, workflow *models.Workflow) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("Workflow with id %d was deleted!", workflow.ID),
	})
}
