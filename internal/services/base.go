package services

import (
	"ielts-app-api/internal/repositories"
)

type Service struct {
	userRepo              *repositories.UserRepository
	targetRepo            *repositories.TargetRepository
	quizRepo              *repositories.QuizRepo
	quizSkillRepo         *repositories.QuizSkillRepo
	TargetRepo            *repositories.TargetRepository
	OTPRepo               *repositories.OTPRepository
	OTPAttemptRepo        *repositories.OTPAttemptRepository
	tagSearchRepo         *repositories.TagSearchRepository
	tagSearchPositionRepo *repositories.TagSearchPositionRepo
	answerRepo            *repositories.AnswerRepo
	successQuizLogRepo    *repositories.SuccessQuizLogRepo
	vocabBankRepo         *repositories.VocabBankRepository
	planRepo              *repositories.PlanRepository
	masterDateRepo        *repositories.MasterDataRepository
	vocabRepo             *repositories.VocabRepository
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
		case *repositories.OTPRepository:
			service.OTPRepo = repo.(*repositories.OTPRepository)
		case *repositories.OTPAttemptRepository:
			service.OTPAttemptRepo = repo.(*repositories.OTPAttemptRepository)
		case *repositories.TagSearchRepository:
			service.tagSearchRepo = repo.(*repositories.TagSearchRepository)
		case *repositories.TagSearchPositionRepo:
			service.tagSearchPositionRepo = repo.(*repositories.TagSearchPositionRepo)
		case *repositories.AnswerRepo:
			service.answerRepo = repo.(*repositories.AnswerRepo)
		case *repositories.SuccessQuizLogRepo:
			service.successQuizLogRepo = repo.(*repositories.SuccessQuizLogRepo)
		case *repositories.VocabBankRepository:
			service.vocabBankRepo = repo.(*repositories.VocabBankRepository)
		case *repositories.PlanRepository:
			service.planRepo = repo.(*repositories.PlanRepository)
		case *repositories.MasterDataRepository:
			service.masterDateRepo = repo.(*repositories.MasterDataRepository)
		case *repositories.VocabRepository:
			service.vocabRepo = repo.(*repositories.VocabRepository)
		default:
			panic("Unknown repository type provided")
		}
	}
	return service
}
