package services

import (
	"ielts-app-api/internal/repositories"
)

type Service struct {
	userRepo      *repositories.UserRepository
	targetRepo    *repositories.TargetRepository
	quizRepo      *repositories.QuizRepo
	quizSkillRepo *repositories.QuizSkillRepo
}

func NewService(repos ...interface{}) *Service {
	service := &Service{}
	for _, repo := range repos {
		switch repo.(type) {
		case *repositories.UserRepository:
			service.userRepo = repo.(*repositories.UserRepository)
		case *repositories.TargetRepository:
			service.targetRepo = repo.(*repositories.TargetRepository)
		case *repositories.QuizRepo:
			service.quizRepo = repo.(*repositories.QuizRepo)
		case *repositories.QuizSkillRepo:
			service.quizSkillRepo = repo.(*repositories.QuizSkillRepo)
		default:
			panic("Unknown repository type provided")
		}
	}
	return service
}
