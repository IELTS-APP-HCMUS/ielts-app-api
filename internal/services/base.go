package services

import (
	"ielts-app-api/internal/repositories"
)

type Service struct {
	UserRepo       *repositories.UserRepository
	TargetRepo     *repositories.TargetRepository
	OTPRepo        *repositories.OTPRepository
	OTPAttemptRepo *repositories.OTPAttemptRepository
}

func NewService(repos ...interface{}) *Service {
	service := &Service{}
	for _, repo := range repos {
		switch repo.(type) {
		case *repositories.UserRepository:
			service.UserRepo = repo.(*repositories.UserRepository)
		case *repositories.TargetRepository:
			service.TargetRepo = repo.(*repositories.TargetRepository)
		case *repositories.OTPRepository:
			service.OTPRepo = repo.(*repositories.OTPRepository)
		case *repositories.OTPAttemptRepository:
			service.OTPAttemptRepo = repo.(*repositories.OTPAttemptRepository)
		default:
			panic("Unknown repository type provided")
		}
	}
	return service
}
