package models

import (
	"ielts-app-api/common"
	"time"
)

type SignupRequest struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	IdToken  *string `json:"id_token,omitempty"`
}

type OTP struct {
	Email     string    `gorm:"primaryKey;size:255;not null" json:"email"`
	OTP       string    `gorm:"size:6;not null" json:"otp"`
	Expiry    time.Time `gorm:"not null" json:"expiry"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (OTP) TableName() string {
	return common.POSTGRES_TABLE_NAME_OTPS
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required"`
}
