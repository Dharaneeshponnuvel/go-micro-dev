package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BatchHandler struct {
	service *services.BatchService
}

func NewBatchHandler(db *gorm.DB) *BatchHandler {
	return &BatchHandler{services.NewBatchService(db)}
}

func (h *BatchHandler) CreateBatch(c *fiber.Ctx) error {
	var batch models.Batch
	if err := c.BodyParser(&batch); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.service.Create(&batch)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create batch"})
	}

	return c.JSON(batch)
}

func (h *BatchHandler) GetByBatchYear(c *fiber.Ctx) error {
	batchYearID := c.Params("batchYearID")
	batches, err := h.service.GetByBatchYear(batchYearID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "No batches found"})
	}
	return c.JSON(batches)
}
