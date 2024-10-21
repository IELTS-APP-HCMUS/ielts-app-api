package services

import (
	"context"
	"ielts-app-api/internal/models"
)

func (s *Service) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, err := s.UserRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
