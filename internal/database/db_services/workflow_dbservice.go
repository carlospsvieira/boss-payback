package db_services

import (
	"boss-payback/internal/database"
	"boss-payback/internal/database/models"
	"boss-payback/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkflowInDB(c *fiber.Ctx, workflow *models.Workflow) error {
	if err := database.DB.Db.Create(workflow).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func GetWorkflowsInDB(c *fiber.Ctx, workflows []models.Workflow) error {
	if err := database.DB.Db.Find(&workflows).Error; err != nil {
		return utils.HandleErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
