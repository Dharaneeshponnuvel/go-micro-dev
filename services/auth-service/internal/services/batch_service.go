package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type BatchService struct {
	db *gorm.DB
}

func NewBatchService(db *gorm.DB) *BatchService {
	return &BatchService{db}
}

func (s *BatchService) Create(batch *models.Batch) error {
	return s.db.Create(batch).Error
}

func (s *BatchService) GetByBatchYear(batchYearID string) ([]models.Batch, error) {
	var batches []models.Batch
	err := s.db.Where("batch_year_id = ?", batchYearID).Find(&batches).Error
	return batches, err
}
