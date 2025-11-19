package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{services.NewStudentService(db)}
}

func (h *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var st models.Student
	if err := c.BodyParser(&st); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	if err := h.service.CreateStudent(&st); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to create student"})
	}
	return c.Status(201).JSON(st)
}

func (h *StudentHandler) GetByBatch(c *fiber.Ctx) error {
	batchID := c.Params("batchID")
	students, err := h.service.GetByBatch(batchID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}
	return c.JSON(students)
}
