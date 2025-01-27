package services

import (
	"context"
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) GetPlanById(ctx context.Context, userId string) ([]*models.Plan, error) {
	plans, err := s.planRepo.List(ctx, models.QueryParams{}, func(tx *gorm.DB) {
		tx.Where("user_id", userId)
	})
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (s *Service) CreatePlan(ctx context.Context, userId string, body models.PlanRequest) (*models.Plan, error) {
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
