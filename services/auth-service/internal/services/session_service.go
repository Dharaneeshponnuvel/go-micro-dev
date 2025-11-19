package services

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type SessionService struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db: db}
}

func (s *SessionService) CreateSession(session *models.Session) error {
	return s.db.Create(session).Error
}

// Single device login (student only)
func (s *SessionService) InvalidateOldSessions(userID string) {
	s.db.Model(&models.Session{}).
		Where("user_id = ? AND is_active = true", userID).
		Updates(map[string]interface{}{"is_active": false})
}
