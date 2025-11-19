package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type RefreshTokenService struct {
	db *gorm.DB
}

func NewRefreshTokenService(db *gorm.DB) *RefreshTokenService {
	return &RefreshTokenService{db: db}
}

func (s *RefreshTokenService) SaveToken(rt *models.RefreshToken) error {
	return s.db.Create(rt).Error
}

func (s *RefreshTokenService) MarkUsed(id string) error {
	return s.db.Model(&models.RefreshToken{}).
		Where("id = ?", id).Update("is_used", true).Error
}
