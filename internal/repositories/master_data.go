package repositories

import (
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

type MasterDataRepository struct {
	db *gorm.DB
	BaseRepository[models.MasterData]
}

func NewMasterDataRepository(db *gorm.DB) *MasterDataRepository {
	return &MasterDataRepository{
		db:             db,
		BaseRepository: NewBaseRepository[models.MasterData](db),
	}
}
