package repositories

import (
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
	BaseRepository[models.Plan]
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{
		db:             db,
		BaseRepository: NewBaseRepository[models.Plan](db),
	}
}
