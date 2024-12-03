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

const (
	defaultLimit    = 20
	defaultPage     = 1
	defaultPageSize = 10
	maxLimit        = 200
)

// GetPageAndPageSize validates and returns page size and limit
func GetPageAndPageSize(page, pageSize int) (int, int) {
	if page == 0 {
		page = defaultPage
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxLimit {
		pageSize = maxLimit
	}
	return page, pageSize
}

func (s *Service) GetAnswer(ctx context.Context, userID string, answerID int) (*models.Answer, error) {
	// Get detail answer
	conds := []repositories.Clause{
		func(tx *gorm.DB) {
			tx.Where("id", answerID)
		},
		func(tx *gorm.DB) {
			ps := []common.Preload{
				{
					Model:    "QuizDetail",
					Selected: []string{"id", "title"},
				},
			}

			for _, p := range ps {
				common.ApplyPreload(tx, p)
			}
		},
	}

	answer, err := s.answerRepo.GetDetailByConditions(ctx, conds...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	// isTeacher := false
	// if answer.UserCreated != userID {
	// 	isPermission, err := s.checkTeacherPermissionOnAnswer(ctx, userID, answer, answerID)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if !isPermission {
	// 		return nil, common.ErrActionNotAllowed
	// 	}
	// 	isTeacher = true
	// }

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

	quizTypes := []int{common.QuizTypeTest, common.QuizTypeMockTest}
	if request.QuizTypes != nil && len(*request.QuizTypes) > 0 {
		quizTypes = *request.QuizTypes
	}
	if len(quizTypes) == 1 {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz_type = ?", quizTypes[0])
		})
	} else {
		filters = append(filters, func(tx *gorm.DB) {
			tx.Where("quiz_type IN ?", quizTypes)
		})
	}
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

		if request.MockTestIDs != nil && len(*request.MockTestIDs) > 0 {
			quizzes, errQ := s.quizRepo.List(ctx, models.QueryParams{}, []repositories.Clause{
				func(tx *gorm.DB) {
					tx.Select("id").Where("mock_test_id IN ?", *request.MockTestIDs)
				},
			}...)
			if errQ != nil {
				return nil, errQ
			}
			quizIDs := []int{}
			for _, q := range quizzes {
				quizIDs = append(quizIDs, q.ID)
			}
			filters = append(filters, func(tx *gorm.DB) {
				tx.Where("quiz IN ?", quizIDs)
			})
		}
		page, pageSize := GetPageAndPageSize(request.Page, request.PageSize)
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

		// default sort: new -> old
		// if request.SkillId != common.QuizSkillTypeWritingSelfPractice2 {
		// 	filters = append(filters, func(tx *gorm.DB) {
		// 		tx.Preload("SuccessQuizLog", func(tx *gorm.DB) *gorm.DB {
		// 			return tx.Select(
		// 				mappingType[*request.Type],
		// 				"sum(total) as total",
		// 				"sum(success) as success",
		// 				"sum(failed) as failed",
		// 				"sum(skipped) as skipped",
		// 			).Where(mappingType[*request.Type] + " IS NOT NULL AND question_type !=''").Group(mappingType[*request.Type])
		// 		}).Preload("QuizDetail")
		// 	})
		// }

		if len(request.Sort) == 0 {
			request.Sort = "date_created.desc"
		}

		// var validTagSearchIDs []int

		// if request.SkillId == common.QuizSkillTypeWritingSelfPractice2 {
		// 	validTagSearchIDs, err = s.tagSearchRepo.FetchValidTagSearchIDs()
		// 	if err != nil {
		// 		return err, nil
		// 	}
		// 	filters = append(filters, func(tx *gorm.DB) {
		// 		tx.Preload("QuizDetail.TagSearches", func(tx *gorm.DB) *gorm.DB {
		// 			return tx.Where("tag_search.id IN ?", validTagSearchIDs).
		// 				Select("tag_search.id, tag_search.title")
		// 		}).Preload("QuizDetail")
		// 	})
		// }

		statisticsByQuiz, err = s.answerRepo.Statistic.List(ctx, models.QueryParams{
			Limit:  pageSize,
			Offset: (page - 1) * pageSize,
			QuerySort: models.QuerySort{
				Origin: request.Sort,
			},
		}, filters...)

		// for i, quiz := range statisticsByQuiz {
		// 	// if quiz.Type == common.QuizSkillTypeWritingSelfPractice2 {
		// 	// 	for _, tag := range quiz.QuizDetail.TagSearches {
		// 	// 		if common.Contains(validTagSearchIDs, tag.ID) {
		// 	// 			statisticsByQuiz[i].QuizTopicFormat = tag.Title
		// 	// 			break
		// 	// 		}
		// 	// 	}
		// 	// }
		// 	// pass

		// }

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
