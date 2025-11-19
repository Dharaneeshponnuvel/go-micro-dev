package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type BatchYearService struct {
	db *gorm.DB
}

func NewBatchYearService(db *gorm.DB) *BatchYearService {
	return &BatchYearService{db}
}

func (s *BatchYearService) Create(batchYear *models.BatchYear) error {
	return s.db.Create(batchYear).Error
}

func (s *BatchYearService) GetInstitutionBatchYears(instID string) ([]models.BatchYear, error) {
	var years []models.BatchYear
	err := s.db.Where("institution_id = ?", instID).Find(&years).Error
	return years, err
}
