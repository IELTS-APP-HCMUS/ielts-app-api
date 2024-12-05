package models

import "ielts-app-api/common"

type QuizTagSearch struct {
	ID          int `json:"id" gorm:"id,primaryKey"`
	QuizID      int `gorm:"index"`
	TagSearchID int `gorm:"index"`
}

func (QuizTagSearch) TableName() string {
	return common.POSTGRES_TABLE_NAME_QUIZ_TAG_SEARCH
}
