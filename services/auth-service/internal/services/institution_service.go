package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type InstitutionService struct {
	db *gorm.DB
}

func NewInstitutionService(db *gorm.DB) *InstitutionService {
	return &InstitutionService{db}
}

func (s *InstitutionService) CreateInstitution(inst *models.Institution) error {
	return s.db.Create(inst).Error
}

func (s *InstitutionService) GetByID(id string) (*models.Institution, error) {
	var inst models.Institution
	err := s.db.First(&inst, "id = ?", id).Error
	return &inst, err
}
