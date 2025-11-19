package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type StudentService struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{db}
}

func (s *StudentService) CreateStudent(su *models.Student) error {
	return s.db.Create(su).Error
}

func (s *StudentService) GetByBatch(batchID string) ([]models.Student, error) {
	var students []models.Student
	err := s.db.Where("batch_id = ?", batchID).
		Preload("User").
		Find(&students).Error
	return students, err
}
