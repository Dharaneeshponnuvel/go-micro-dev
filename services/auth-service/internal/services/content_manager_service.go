package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories"
)

type ContentManagerService struct {
	repo *repositories.ContentManagerRepository
}

func NewContentManagerService(repo *repositories.ContentManagerRepository) *ContentManagerService {
	return &ContentManagerService{repo: repo}
}

func (s *ContentManagerService) Create(cm *models.ContentManager) error {
	return s.repo.Create(cm)
}

func (s *ContentManagerService) AssignBatch(assign *models.ContentManagerBatch) error {
	return s.repo.AssignBatch(assign)
}

func (s *ContentManagerService) GetManagerDetails(userID string) (*models.ContentManager, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ContentManagerService) GetManagerBatches(managerID string) ([]models.ContentManagerBatch, error) {
	return s.repo.GetAssignedBatches(managerID)
}
