package services

import (
	"boss-payback/internal/database/models"

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

func GetWorkflowsResponse(c *fiber.Ctx, workflows []models.Workflow) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    workflows,
		"message": "Successfully fetched all workflows",
	})
}
