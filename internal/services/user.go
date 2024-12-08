package services

import (
	"context"
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) GetUserProfileById(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Preload("Role").Where("id = ?", id)
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
