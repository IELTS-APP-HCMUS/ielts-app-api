package models

import (
	"ielts-app-api/common"
	"time"
)

type PlanRequest struct {
	Activity string `json:"activity"`
	Time     string `json:"time"`
}

type Plan struct {
	ID        int       `json:"id" gorm:"id,primaryKey;autoIncrement"`
	Activity  string    `json:"activity,omitempty"`
	Time      string    `json:"time,omitempty"`
	UserId    string    `json:"user_id,omitempty" gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Plan) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_PLAN
}
