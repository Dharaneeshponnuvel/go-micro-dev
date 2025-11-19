package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type ContentManagerRepository struct{ db *gorm.DB }

func NewContentManagerRepository(db *gorm.DB) *ContentManagerRepository {
	return &ContentManagerRepository{db}
}

func (r *ContentManagerRepository) Create(cm *models.ContentManager) error {
	return r.db.Create(cm).Error
}

func (r *ContentManagerRepository) AssignBatch(assign *models.ContentManagerBatch) error {
	return r.db.Create(assign).Error
}
