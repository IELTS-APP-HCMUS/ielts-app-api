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

func (s *Service) CreateTarget(ctx context.Context, userId string, target models.TargetRequest) (*models.Target, error) {
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
	parsedNextExamDate, err := time.Parse(layout, target.NextExamDate)
	if err != nil {
		return nil, err
	}

	newTarget := models.Target{
		ID:                  userId,
		TargetStudyDuration: target.TargetStudyDuration,
		TargetReading:       target.TargetReading,
		TargetListening:     target.TargetListening,
		TargetSpeaking:      target.TargetSpeaking,
		TargetWriting:       target.TargetWriting,
		NextExamDate:        parsedNextExamDate,
	}
	createdTarget, err := s.TargetRepo.Create(ctx, &newTarget)
	if err != nil {
		return nil, err
	}
	return createdTarget, nil
}

func (s *Service) UpdateTarget(ctx context.Context, userId string, target models.TargetRequest) (*models.Target, error) {
	_, err := s.TargetRepo.GetByID(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = common.ErrTargetNotFound
			return nil, err
		}
		return nil, err
	}

	layout := "2006-01-02 15:04:05"
	parsedNextExamDate, err := time.Parse(layout, target.NextExamDate)
	if err != nil {
		return nil, err
	}

	updateTarget := models.Target{}
	if target.TargetStudyDuration != 0 {
		updateTarget.TargetStudyDuration = target.TargetStudyDuration
	}
	if target.TargetReading != 0 {
		updateTarget.TargetReading = target.TargetReading
	}
	if target.TargetListening != 0 {
		updateTarget.TargetListening = target.TargetListening
	}
	if target.TargetSpeaking != 0 {
		updateTarget.TargetSpeaking = target.TargetSpeaking
	}
	if target.TargetWriting != 0 {
		updateTarget.TargetWriting = target.TargetWriting
	}
	if target.NextExamDate != "" {
		updateTarget.NextExamDate = parsedNextExamDate
	}
	updatedTarget, err := s.TargetRepo.Update(ctx, userId, &updateTarget)
	if err != nil {
		return nil, err
	}
	return updatedTarget, nil
}
