package services

import (
	"context"
	"ielts-app-api/internal/models"
)

// func (s *Service) GetMasterData(ctx context.Context) ([]*models.MasterData, error) {

// }

func (s *Service) CreatePla1n(ctx context.Context, userId string, body models.PlanRequest) (*models.Plan, error) {
	plan := &models.Plan{
		Activity: body.Activity,
		Time:     body.Time,
		UserId:   userId,
	}

	plan, err := s.planRepo.Create(ctx, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
