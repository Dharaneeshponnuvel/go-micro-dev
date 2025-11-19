package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ContentManagerHandler struct {
	service *services.ContentManagerService
}

func NewContentManagerHandler(db *gorm.DB) *ContentManagerHandler {
	return &ContentManagerHandler{services.NewContentManagerService(db)}
}

func (h *ContentManagerHandler) AssignBatch(c *fiber.Ctx) error {
	var assign models.ContentManagerBatch
	if err := c.BodyParser(&assign); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := h.service.AssignBatch(&assign)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to assign batch"})
	}
	return c.JSON(fiber.Map{"success": true})
}
