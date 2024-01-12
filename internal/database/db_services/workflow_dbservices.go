package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkflowInDB(c *fiber.Ctx, workflow *models.Workflow) error {
	if err := database.Instance.Db.Create(workflow).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetWorkflowsInDB(c *fiber.Ctx, workflows *[]models.Workflow) error {
	if err := database.Instance.Db.Find(workflows).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetWorkflowByApproverInDB(c *fiber.Ctx, workflows *[]models.Workflow, approverId uint) error {
	if err := database.Instance.Db.Where("approver_id = ?", approverId).Find(workflows).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetWorkflowByExpenseInDB(c *fiber.Ctx, workflows *[]models.Workflow, expenseId uint) error {
	if err := database.Instance.Db.Where("expense_id = ?", expenseId).Find(workflows).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateWorkflowStatusInDB(c *fiber.Ctx, id uint, updatedStatus string) error {
	var workflow models.Workflow
	if err := database.Instance.Db.Model(&workflow).Where("id = ?", id).Update("status", updatedStatus).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func UpdateWorkflowCommentsInDB(c *fiber.Ctx, id uint, updatedComments string) error {
	var workflow models.Workflow
	if err := database.Instance.Db.Model(&workflow).Where("id = ?", id).Update("comments", updatedComments).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteWorkflowInDB(c *fiber.Ctx, workflow *models.Workflow) error {
	if err := database.Instance.Db.Model(workflow).Where("id = ?", workflow.ID).Unscoped().Delete(&workflow).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
