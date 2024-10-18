package models

import (
	"ielts-app-api/common"
	"time"
)

type User struct {
	ID          string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email       string  `gorm:"uniqueIndex"`
	FirstName   *string `gorm:"column:first_name"`
	LastName    *string `gorm:"column:last_name"`
	Password    string
	Role        string
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
}

func (User) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_USERS
}
