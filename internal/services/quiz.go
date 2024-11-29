package services

import (
	"context"
	"errors"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"ielts-app-api/internal/repositories"

	"gorm.io/gorm"
)

func (s *Service) GetQuizzes(ctx context.Context, userID string, request *models.ListQuizzesParamsUri) (*models.BaseListResponse, error) {
	var (
		filters = []repositories.Clause{}
		quizIDs = []int{}
		err     error
	)
	page, pageSize := common.GetPageAndPageSize(request.Page, request.PageSize)

	resData := models.BaseListResponse{
		Total:    0,
		Page:     page,
		PageSize: pageSize,
		Items:    []*models.Quiz{},
	}

	if request.TagPassage != nil ||
		request.TagSection != nil ||
		request.TagQuestionType != nil ||
		request.TagTask != nil ||
		request.TagTopic != nil ||
		request.TagBookType != nil {

		tagIDs := []int{}
		if request.TagSection != nil {
			tagIDs = append(tagIDs, *request.TagSection)
		}
		if request.TagPassage != nil {
			tagIDs = append(tagIDs, *request.TagPassage)
		}
		if request.TagQuestionType != nil {
			tagIDs = append(tagIDs, *request.TagQuestionType)
		}

		if request.TagTask != nil {
			tagIDs = append(tagIDs, *request.TagTask)
		}

		if request.TagTopic != nil {
			tagIDs = append(tagIDs, *request.TagTopic)
		}

		if request.TagBookType != nil {
			tagIDs = append(tagIDs, *request.TagBookType)
		}

		// get quizIDs have matched tags
		quizIDs, err = s.quizRepo.GetQuizIDsInCludeTagIDs(ctx, tagIDs)
		if err != nil {
			return nil, err
		}

		if len(quizIDs) == 0 {
			return &resData, nil
		}

		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz.id IN ?", quizIDs)
		})
	}

	if request.IsTest != nil {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz.is_test = ?", *request.IsTest)
		})
	}

	if request.Type != nil {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz.type = ?", *request.Type)
		})
	}

	if request.Status != nil && len(*request.Status) > 0 {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz.status = ?", *request.Status)
		})
	}

	if request.Search != nil && len(*request.Search) > 0 {
		var quesQuizIDs []int

		if len(quesQuizIDs) > 0 {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where("quiz.title ILIKE ? OR id IN (?)", "%"+*request.Search+"%", quesQuizIDs)
			})
		} else {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where("quiz.title ILIKE ?", "%"+*request.Search+"%")
			})
		}
	}

	total, err := s.quizRepo.Count(ctx, models.QueryParams{}, filters...)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return &resData, nil
	}

	resData.Total = int(total)

	// Preload tagSearches
	filters = append(filters, func(tx *gorm.DB) {
		tx.Preload("TagSearches")
	})

	records, err := s.quizRepo.List(
		ctx,
		models.QueryParams{
			Limit:  pageSize,
			Offset: (page - 1) * pageSize,
			QuerySort: models.QuerySort{
				Origin: request.Sort,
			},
		},
		filters...,
	)

	if err != nil {
		return nil, err
	}

	quizIDs = []int{}
	for _, record := range records {
		quizIDs = append(quizIDs, record.ID)
	}

	resData.Items = records
	return &resData, nil
}

func (s *Service) GetQuiz(ctx context.Context, req *models.QuizParamsUri, userID string) (*models.Quiz, error) {
	var (
		quiz *models.Quiz
		err  error
	)
	filters := []repositories.Clause{}
	filters = append(filters, func(tx *gorm.DB) {
		tx.Preload("Parts", func(db *gorm.DB) *gorm.DB {
			return db.Joins("INNER JOIN quiz_part ON quiz_part.quiz_id = ? AND quiz_part.part_id = part.id", req.QuizID).Order("quiz_part.sort")
		}).Preload("Parts.Questions", func(db *gorm.DB) *gorm.DB {
			return db.Order("question.sort")
		}).Where("id", req.QuizID)
	})

	// Preload vocabs: Milestone 2
	quiz, err = s.quizRepo.GetDetailByConditions(
		ctx,
		filters...,
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	return quiz, nil
}
