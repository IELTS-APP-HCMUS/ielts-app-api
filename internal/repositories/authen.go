package repositories

import (
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

type OTPRepository struct {
	BaseRepository[models.OTP]
}

func NewOTPRepository(db *gorm.DB) *OTPRepository {
	return &OTPRepository{
		BaseRepository: NewBaseRepository[models.OTP](db),
	}
}
