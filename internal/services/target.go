package services

import (
	"context"
	"errors"
	"fmt"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"time"

	"gorm.io/gorm"
)

func (s *Service) GetTargetById(ctx context.Context, id string) (*models.Target, error) {
	fmt.Println(id)
	target, err := s.TargetRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (s *Service) CreateTarget(ctx context.Context, userId string, req models.TargetRequest) (*models.Target, error) {
	_, err := s.TargetRepo.GetByID(ctx, userId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		err = common.ErrTargetAlreadyExists
		return nil, err
	}

	layout := "2006-01-02 15:04:05"
	parsedNextExamDate, err := time.Parse(layout, *req.NextExamDate)
	if err != nil {
		return nil, err
	}

	newTarget := models.Target{
		ID:                  userId,
		TargetStudyDuration: *req.TargetStudyDuration,
		TargetReading:       *req.TargetReading,
		TargetListening:     *req.TargetListening,
		TargetSpeaking:      *req.TargetSpeaking,
		TargetWriting:       *req.TargetWriting,
		NextExamDate:        parsedNextExamDate,
	}
	createdTarget, err := s.TargetRepo.Create(ctx, &newTarget)
	if err != nil {
		return nil, err
	}
	return createdTarget, nil
}

func (s *Service) UpdateTarget(ctx context.Context, userId string, req models.TargetRequest) (*models.Target, error) {
	updateTarget, err := s.TargetRepo.GetByID(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	layout := "2006-01-02 15:04:05"
	parsedNextExamDate, err := time.Parse(layout, *req.NextExamDate)
	if err != nil {
		return nil, err
	}

	if req.TargetStudyDuration != nil {
		updateTarget.TargetStudyDuration = *req.TargetStudyDuration
	}
	if req.TargetReading != nil {
		updateTarget.TargetReading = *req.TargetReading
	}
	if req.TargetListening != nil {
		updateTarget.TargetListening = *req.TargetListening
	}
	if req.TargetSpeaking != nil {
		updateTarget.TargetSpeaking = *req.TargetSpeaking
	}
	if req.TargetWriting != nil {
		updateTarget.TargetWriting = *req.TargetWriting
	}
	if req.NextExamDate != nil {
		updateTarget.NextExamDate = parsedNextExamDate
	}
	updatedTarget, err := s.TargetRepo.Update(ctx, userId, updateTarget)
	if err != nil {
		return nil, err
	}
	return updatedTarget, nil
}
