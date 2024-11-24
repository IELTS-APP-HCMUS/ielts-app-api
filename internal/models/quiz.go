package models

import (
	"ielts-app-api/common"
	"time"

	"gorm.io/datatypes"
)

type Quiz struct {
	ID               int            `json:"id" gorm:"id,primaryKey"`
	Type             int            `json:"type" gorm:"type"`
	Mode             int            `json:"mode" gorm:"mode"`
	Title            string         `json:"title" gorm:"title"`
	Status           string         `json:"status" gorm:"status"`
	TotalQuestions   *int           `json:"total_questions,omitempty" gorm:"-"`
	Sort             *int           `json:"sort" gorm:"sort"`
	Time             *int           `json:"time" gorm:"time"`
	IsTest           *bool          `json:"is_test" gorm:"is_test"`
	SimplifiedID     *int           `json:"simplified_id" gorm:"simplified_id"`
	LimitSubmit      *int           `json:"limit_submit" gorm:"limit_submit"`
	Thumbnail        *string        `json:"thumbnail" gorm:"thumbnail"`
	QuizCode         string         `json:"quiz_code" gorm:"quiz_code"`
	Description      *string        `json:"description" gorm:"description"`
	Content          *string        `json:"content" gorm:"content"`
	Parts            []*Part        `json:"parts" gorm:"many2many:quiz_part;"`
	TagSearches      []*TagSearch   `json:"tags" gorm:"many2many:quiz_tag_search;"`
	UserCreated      string         `json:"user_created" gorm:"user_created"`
	UserUpdated      string         `json:"user_updated" gorm:"user_updated"`
	DateCreated      *time.Time     `json:"date_created" gorm:"date_created"`
	DateUpdated      *time.Time     `json:"date_updated" gorm:"date_updated"`
	QuizPart         []QuizPart     `json:"quiz_part" gorm:"foreignKey:Quiz"`
	QuizType         int            `json:"quiz_type" gorm:"quiz_type"`
	MockTestID       *int           `json:"mock_test_id"`
	MockTestType     int            `json:"mock_test_type"`
	Listening        *string        `json:"listening"`
	Question         *string        `json:"question"`
	Samples          *string        `json:"samples"`
	VoteCount        int            `json:"vote_count"`
	Questions        []*Question    `json:"questions,omitempty" gorm:"-"`
	QuizPartMs       []QuizPartM    `json:"quiz_part_ms,omitempty" gorm:"foreignKey:QuizID"`
	TotalSubmitted   int            `json:"total_submitted"`
	IsSubmitted      *bool          `json:"is_submitted,omitempty" gorm:"-"`
	Meta             datatypes.JSON `json:"meta,omitempty"`
	ShortDescription *string        `json:"short_description,omitempty"`
}

func (r Quiz) TableName() string {
	return common.POSTGRES_TABLE_NAME_QUIZ
}
