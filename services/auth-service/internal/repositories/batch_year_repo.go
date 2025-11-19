package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type BatchYearRepository struct {
	db *gorm.DB
}

func NewBatchYearRepository(db *gorm.DB) *BatchYearRepository {
	return &BatchYearRepository{db}
}

// ➕ Create Batch Year
func (r *BatchYearRepository) Create(batchYear *models.BatchYear) error {
	return r.db.Create(batchYear).Error
}

// 📌 Find all batch years for each institution
func (r *BatchYearRepository) GetByInstitution(instID string) ([]models.BatchYear, error) {
	var years []models.BatchYear
	err := r.db.Where("institution_id = ?", instID).Find(&years).Error
	return years, err
}

// 🔍 Find One
func (r *BatchYearRepository) GetByID(id string) (*models.BatchYear, error) {
	var year models.BatchYear
	err := r.db.First(&year, "id = ?", id).Error
	return &year, err
}
