package models

import "time"

type User struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Role        string
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime"`
}

func (User) TableName() string {
	// define common variable here
	return "public.users"
}
