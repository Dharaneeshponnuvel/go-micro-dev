package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type InstitutionHandler struct {
	service *services.InstitutionService
}

func NewInstitutionHandler(db *gorm.DB) *InstitutionHandler {
	repo := repositories.NewInstitutionRepository(db)
	service := services.NewInstitutionService(repo)
	return &InstitutionHandler{service}
}

func (h *InstitutionHandler) CreateInstitution(c *fiber.Ctx) error {
	var inst models.Institution
	if err := c.BodyParser(&inst); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	if err := h.service.CreateInstitution(&inst); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save institution"})
	}

	return c.Status(201).JSON(inst)
}

func (h *InstitutionHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	inst, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Institution not found"})
	}
	return c.JSON(inst)
}
