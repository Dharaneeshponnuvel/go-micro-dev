package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BatchYearHandler struct {
	service *services.BatchYearService
}

func NewBatchYearHandler(db *gorm.DB) *BatchYearHandler {
	repo := repositories.NewBatchYearRepository(db)
	service := services.NewBatchYearService(repo)
	return &BatchYearHandler{service}
}

func (h *BatchYearHandler) CreateBatchYear(c *fiber.Ctx) error {
	var by models.BatchYear
	if err := c.BodyParser(&by); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.service.Create(&by); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create batch year"})
	}
	return c.Status(201).JSON(by)
}

func (h *BatchYearHandler) GetByInstitution(c *fiber.Ctx) error {
	instID := c.Params("institutionID")
	years, err := h.service.GetInstitutionBatchYears(instID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "No batch years found"})
	}
	return c.JSON(years)
}
