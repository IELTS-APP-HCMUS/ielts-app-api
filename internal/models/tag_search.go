package models

import (
	"ielts-app-api/common"
)

type TagSearch struct {
	ID      int     `json:"id" gorm:"id,primaryKey"`
	Title   string  `json:"title" gorm:"title"`
	Code    string  `json:"code" gorm:"code"`
	Quizzes []*Quiz `json:"-" gorm:"many2many:quiz_tag_search;"`
}

func (TagSearch) TableName() string {
	return common.POSTGRES_TABLE_NAME_TAG_SEARCH
}

type TagSearchPosition struct {
	ID       int    `json:"id" gorm:"id,primaryKey"`
	Position string `json:"position" gorm:"position"`
}

func (TagSearchPosition) TableName() string {
	return common.POSTGRES_TABLE_NAME_TAG_POSITION
}
