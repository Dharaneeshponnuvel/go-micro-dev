package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type InstitutionRepository struct {
	db *gorm.DB
}

func NewInstitutionRepository(db *gorm.DB) *InstitutionRepository {
	return &InstitutionRepository{db: db}
}

// ➕ Create Institution
func (r *InstitutionRepository) Create(inst *models.Institution) error {
	return r.db.Create(inst).Error
}

// 🔍 Find by ID
func (r *InstitutionRepository) FindByID(id string) (*models.Institution, error) {
	var inst models.Institution
	err := r.db.First(&inst, "id = ?", id).Error
	return &inst, err
}

// 🧾 Get All Institutions (future use – admin panel, dropdown, etc.)
func (r *InstitutionRepository) GetAll() ([]models.Institution, error) {
	var insts []models.Institution
	err := r.db.Find(&insts).Error
	return insts, err
}

// ✏️ Update Institution
func (r *InstitutionRepository) Update(inst *models.Institution) error {
	return r.db.Save(inst).Error
}

// ❌ Delete (Soft Delete with GORM)
func (r *InstitutionRepository) Delete(id string) error {
	return r.db.Delete(&models.Institution{}, "id = ?", id).Error
}
