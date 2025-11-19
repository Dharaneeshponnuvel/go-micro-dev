package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories"
)

type SessionService struct {
	repo *repositories.SessionRepository
}

func NewSessionService(repo *repositories.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

// Create new session
func (s *SessionService) CreateSession(session *models.Session) error {
	return s.repo.Create(session)
}

// Invalidate all previous sessions (single device login)
func (s *SessionService) InvalidateAllSessions(userID string) {
	s.repo.InvalidateAll(userID)
}

// Deactivate only current session (logout)
func (s *SessionService) DeactivateSession(userID, token string) {
	s.repo.DeactivateSession(userID, token)
}
