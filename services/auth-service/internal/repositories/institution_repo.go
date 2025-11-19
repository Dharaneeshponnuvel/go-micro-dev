package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type InstitutionRepository struct{ db *gorm.DB }

func NewInstitutionRepository(db *gorm.DB) *InstitutionRepository {
	return &InstitutionRepository{db}
}

func (r *InstitutionRepository) Create(inst *models.Institution) error {
	return r.db.Create(inst).Error
}

func (r *InstitutionRepository) FindByID(id string) (*models.Institution, error) {
	var inst models.Institution
	err := r.db.First(&inst, "id = ?", id).Error
	return &inst, err
}
