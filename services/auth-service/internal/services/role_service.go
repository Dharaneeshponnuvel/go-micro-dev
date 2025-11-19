package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

// 📌 Get All Roles (for Register dropdown)
func (s *RoleService) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := s.db.Find(&roles).Error
	return roles, err
}

// 📌 Find Role by Name (used in Register & Login)
func (s *RoleService) FindByName(name string) (*models.Role, error) {
	var role models.Role
	err := s.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

// 📌 Create Role (if needed in future Admin Panel)
func (s *RoleService) CreateRole(role *models.Role) error {
	return s.db.FirstOrCreate(role, models.Role{Name: role.Name}).Error
}

// 📌 Delete Role
func (s *RoleService) DeleteRole(id string) error {
	return s.db.Delete(&models.Role{}, "id = ?", id).Error
}
