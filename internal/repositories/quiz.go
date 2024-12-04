package repositories

import (
	"context"
	"errors"
	"fmt"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QuizRepo struct {
	db *gorm.DB
	BaseRepository[models.Quiz]
}

type QuizSkillRepo struct {
	db *gorm.DB
	BaseRepository[models.QuizSkill]
}

func NewQuizRepository(db *gorm.DB) *QuizRepo {
	baseRepo := NewBaseRepository[models.Quiz](db)
	return &QuizRepo{
		db:             db,
		BaseRepository: baseRepo,
	}
}

func NewQuizSkillRepository(db *gorm.DB) *QuizSkillRepo {
	baseRepo := NewBaseRepository[models.QuizSkill](db)
	return &QuizSkillRepo{
		db:             db,
		BaseRepository: baseRepo,
	}
}

func (r *QuizRepo) GetQuizIDsInCludeTagIDs(ctx context.Context, tagIDs []int) ([]int, error) {
	var qIDs []int

	tagIDsFmt := strings.ReplaceAll(fmt.Sprintf("%+v", tagIDs), " ", ", ")

	tx := r.db.Table(common.POSTGRES_TABLE_NAME_QUIZ_TAG_SEARCH).
		Select("quiz_id").
		Group("quiz_id").
		Where("quiz_id IS NOT NULL").
		Having(fmt.Sprintf("ARRAY%+v <@ ARRAY_AGG(tag_search_id)", tagIDsFmt))
	err := tx.Pluck("quiz_id", &qIDs).Error
	if err != nil {
		return nil, err
	}
	return qIDs, nil
}

func (r *QuizRepo) GetQuizDetailByConditions(ctx context.Context, quizId int) (*models.Quiz, error) {
	var results []struct {
		QuizID                 int
		QuizTitle              string
		QuizType               int
		PartID                 int
		QuizIDPart             int // Match with "QuizID" in the models.Part struct
		PartPassage            int
		QuestionID             int
		QuestionPartID         int
		QuestionType           string
		QuestionMultiple       datatypes.JSON
		QuestionGapFillInBlank *string
		// QuestionSingleChoiceRadio datatypes.JSON
		// QuestionSelection         datatypes.JSON
		// QuestionMultipleChoice    datatypes.JSON
		// QuestionSelectionOption   datatypes.JSON
	}

	// Perform a single query with JOINs
	err := r.db.Raw(`
		SELECT 
			q.*,
			p.id AS part_id, p.quiz AS quiz_id_part, p.passage AS part_passage,
			qu.id AS question_id, qu.part AS question_part_id, qu.question_type AS question_type, qu.multiple_choice AS question_multiple, qu.gap_fill_in_blank AS question_gap_fill_in_blank
		FROM 
			public.quiz AS q
		LEFT JOIN 
			public.part AS p ON q.id = p.quiz
		LEFT JOIN 
			public.question AS qu ON p.id = qu.part
		WHERE 
			q.id = ?
	`, quizId).Scan(&results).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrQuizNotFound
		}
		return nil, err
	}

	// Map the results into nested Quiz structure
	quiz := &models.Quiz{
		ID:    quizId,
		Title: results[0].QuizTitle,
		Type:  results[0].QuizType,
		Parts: []*models.Part{},
	}

	partMap := make(map[int]*models.Part) // To avoid duplicate parts
	for _, result := range results {
		// Check if the part already exists
		if _, exists := partMap[result.PartID]; !exists && result.PartID != 0 {
			part := &models.Part{
				ID:        result.PartID,
				Quiz:      result.QuizIDPart, // Match the alias in the query
				Passage:   result.PartPassage,
				Questions: []*models.Question{}, // Initialize Questions slice
			}
			quiz.Parts = append(quiz.Parts, part)
			partMap[result.PartID] = part
		}

		// Add questions to the appropriate part
		if result.QuestionID != 0 {
			question := &models.Question{
				ID:             result.QuestionID,
				Part:           &result.QuestionPartID,
				QuestionType:   result.QuestionType,
				MultipleChoice: result.QuestionMultiple,
				GapFillInBlank: result.QuestionGapFillInBlank,
			}
			if part, exists := partMap[result.PartID]; exists {
				part.Questions = append(part.Questions, question)
			}
		}
	}

	return quiz, nil
}
