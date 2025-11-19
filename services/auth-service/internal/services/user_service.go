package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *UserService) FindByID(id string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Role").Where("id = ?", id).First(&user).Error
	return &user, err
}
