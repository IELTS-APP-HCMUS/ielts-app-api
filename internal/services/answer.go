package services

import (
	"context"
	"errors"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"ielts-app-api/internal/repositories"
	"sort"
	"time"

	"gorm.io/gorm"
)

func (s *Service) GetAnswer(ctx context.Context, userID string, answerID int) (*models.Answer, error) {
	// Get detail answer
	conds := []repositories.Clause{
		func(tx *gorm.DB) {
			tx.Where("id", answerID)
		},
		// func(tx *gorm.DB) {
		// 	ps := []common.Preload{
		// 		{
		// 			Model:    "QuizDetail",
		// 			Selected: []string{"id", "title"},
		// 		},
		// 	}

		// 	for _, p := range ps {
		// 		common.ApplyPreload(tx, p)
		// 	}
		// },
	}

	answer, err := s.answerRepo.GetDetailByConditions(ctx, conds...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	// get student info
	student, err := s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Select("id, first_name, last_name, fullname, avatar").Where("id = ?", answer.UserCreated)
	})
	if err != nil {
		return nil, err
	}
	answer.Student = student

	return answer, nil
}

func (s *Service) GetAnswerStatistic(ctx context.Context, studentID string, request *models.AnswerStatisticsQuery) (interface{}, error) {
	filters := []repositories.Clause{}
	var (
		statisticsByQuiz            models.AnswerStatistics
		statisticsByPassageOrQsType models.SuccessCounts
		err                         error
	)

	if request.StartedAt.IsZero() && request.EndedAt.IsZero() {
		request.EndedAt = time.Now()
		request.StartedAt = request.EndedAt.AddDate(0, 0, -365)
	} else if !request.StartedAt.IsZero() {
		request.EndedAt = time.Now()
	} else if !request.EndedAt.IsZero() {
		request.StartedAt = request.EndedAt.AddDate(0, 0, -365)
	}
	if request.Type == nil {
		return nil, nil
	}

	var mappingType = map[int]string{
		common.AnswerStatisticByPassage:    "passage",
		common.AnswerStatisticQuestionType: "question_type",
		common.AnswerStatisticByQuiz:       "answer_id",
	}

	filterSuccessQuizLog := func(tx *gorm.DB) *gorm.DB {
		return tx.Select(
			mappingType[*request.Type],
			"sum(total) as total",
			"sum(success) as success",
			"sum(failed) as failed",
			"sum(skipped) as skipped",
		).Where(mappingType[*request.Type] + " IS NOT NULL").Group(mappingType[*request.Type])
	}

	if *request.Type == common.AnswerStatisticByQuiz {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("user_created = ?", studentID)
		})
		if request.SkillId > 0 {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where("type = ?", request.SkillId)
			})
		}
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("date_created >= ? and date_created <= ?", request.StartedAt, request.EndedAt)
		})

		page, pageSize := common.GetPageAndPageSize(request.Page, request.PageSize)
		total, err := s.answerRepo.Count(ctx, models.QueryParams{}, filters...)
		if err != nil {
			return nil, err
		}
		resData := models.BaseListResponse{
			Total:    int(total),
			Page:     page,
			PageSize: pageSize,
			Items:    []*models.AnswerStatistic{},
		}

		if total == 0 {
			return &resData, nil
		}

		if len(request.Sort) == 0 {
			request.Sort = "date_created.desc"
		}

		statisticsByQuiz, err = s.answerRepo.Statistic.List(ctx, models.QueryParams{
			Limit:  pageSize,
			Offset: (page - 1) * pageSize,
			QuerySort: models.QuerySort{
				Origin: request.Sort,
			},
		}, filters...)

		if err != nil {
			return nil, err
		}
		resData.Items = statisticsByQuiz.Parse()
		return &resData, nil
	} else if *request.Type == common.AnswerStatisticByPassage || *request.Type == common.AnswerStatisticQuestionType {
		if request.SkillId > 0 {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where("skill = ?", request.SkillId)
			})
		}
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("user_id = ? AND create_at >= ? and create_at <= ?", studentID, request.StartedAt.Unix(), request.EndedAt.Unix())
		})
		filters = append(filters, func(tx *gorm.DB) {
			filterSuccessQuizLog(tx)
		})
		if *request.Type == common.AnswerStatisticByPassage {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where(mappingType[*request.Type] + " != 0") // passage != 0
			})
		} else {
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where(mappingType[*request.Type] + " != ''") // question_type != ''
			})
		}
		statisticsByPassageOrQsType, err = s.successQuizLogRepo.Statistic.List(ctx, models.QueryParams{
			QuerySort: models.QuerySort{
				Origin: request.Sort,
			},
		}, filters...)
		if err != nil {
			return nil, err
		}
		statisticsByPassageOrQsType = statisticsByPassageOrQsType.Parse()

		// default sort: correct percent asc
		if len(request.Sort) == 0 {
			sort.Slice(statisticsByPassageOrQsType, func(i, j int) bool {
				return statisticsByPassageOrQsType[i].CorrectPercent < statisticsByPassageOrQsType[j].CorrectPercent
			})
		}
		return &models.BaseListResponse{
			Total:    len(statisticsByPassageOrQsType),
			Page:     1,
			PageSize: len(statisticsByPassageOrQsType),
			Items:    statisticsByPassageOrQsType,
		}, nil
	}
	return nil, nil
}
