package handlers

import (
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleHandler struct {
	service *services.RoleService
}

func NewRoleHandler(db *gorm.DB) *RoleHandler {
	repo := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repo)
	return &RoleHandler{service}
}

func (h *RoleHandler) GetAllRoles(c *fiber.Ctx) error {
	roles, err := h.service.GetRoles()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch roles"})
	}
	return c.JSON(roles)
}
