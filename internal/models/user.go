package models

import (
	"ielts-app-api/common"
	"time"
)

type User struct {
	ID          string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email       string    `json:"email" gorm:"uniqueIndex"`
	FirstName   *string   `json:"first_name" gorm:"column:first_name"`
	LastName    *string   `json:"last_name" gorm:"column:last_name"`
	Password    string    `json:"-" gorm:"password"` // hide password
	RoleID      string    `json:"role_id" gorm:"column:role"`
	Status      string    `json:"status" gorm:"status"`
	IsActive    bool      `json:"is_active" gorm:"is_active"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number"`
	DateCreated time.Time `json:"date_created" gorm:"column:date_created;autoCreateTime"`
}

func (User) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_USERS
}
