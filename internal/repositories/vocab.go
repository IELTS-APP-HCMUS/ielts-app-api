package repositories

import (
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

type VocabBankRepository struct {
	db *gorm.DB
	BaseRepository[models.UserVocabBank]
}

func NewVocabBankRepository(db *gorm.DB) *VocabBankRepository {
	return &VocabBankRepository{
		db:             db,
		BaseRepository: NewBaseRepository[models.UserVocabBank](db),
	}
}

type VocabRepository struct {
	db *gorm.DB
	BaseRepository[models.Vocab]
}

func NewVocabRepository(db *gorm.DB) *VocabRepository {
	return &VocabRepository{
		db:             db,
		BaseRepository: NewBaseRepository[models.Vocab](db),
	}
}
