package services

import (
	"ielts-app-api/internal/repositories"
)

type Service struct {
	UserRepo *repositories.UserRepository
}

func NewService(repos ...interface{}) *Service {
	service := &Service{}
	for _, repo := range repos {
		switch repo.(type) {
		case *repositories.UserRepository:
			service.UserRepo = repo.(*repositories.UserRepository)
		default:
			panic("Unknown repository type provided")
		}
	}
	return service
}
