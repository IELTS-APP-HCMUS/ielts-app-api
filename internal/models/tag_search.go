package models

import (
	"ielts-app-api/common"
)

type TagSearch struct {
	ID           int                  `json:"id" gorm:"id,primaryKey"`
	Title        string               `json:"title" gorm:"title"`
	Code         string               `json:"code" gorm:"code"`
	Quizzes      []*Quiz              `json:"-" gorm:"many2many:quiz_tag_search;"`
	TagPositions []*TagSearchPosition `json:"tag_positions" gorm:"many2many:tag_position_tag_search;joinForeignKey:tag_search_id;joinReferences:tag_position_id"`
}

func (TagSearch) TableName() string {
	return common.POSTGRES_TABLE_NAME_TAG_SEARCH
}

type TagSearchPosition struct {
	ID       int          `json:"id" gorm:"id,primaryKey"`
	Position string       `json:"position" gorm:"position"`
	Tags     []*TagSearch `json:"tags" gorm:"many2many:tag_position_tag_search;joinForeignKey:tag_position_id;joinReferences:tag_search_id"`
}

func (TagSearchPosition) TableName() string {
	return common.POSTGRES_TABLE_NAME_TAG_POSITION
}
