package models

import (
	"ielts-app-api/common"
)

type Role struct {
	ID       string `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	PublicID string `json:"public_id" gorm:"column:public_id"`
}

func (Role) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_ROLES
}
