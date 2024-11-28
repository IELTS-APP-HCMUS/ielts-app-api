package repositories

import (
	"context"
	"fmt"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"strings"

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
