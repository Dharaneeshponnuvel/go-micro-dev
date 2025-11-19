package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type ContentManagerService struct {
	db *gorm.DB
}

func NewContentManagerService(db *gorm.DB) *ContentManagerService {
	return &ContentManagerService{db}
}

func (s *ContentManagerService) Create(cm *models.ContentManager) error {
	return s.db.Create(cm).Error
}

func (s *ContentManagerService) AssignBatch(assign *models.ContentManagerBatch) error {
	return s.db.Create(assign).Error
}
