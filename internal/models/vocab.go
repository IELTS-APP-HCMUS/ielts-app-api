package models

import (
	"ielts-app-api/common"
	"time"

	"gorm.io/datatypes"
)

type VocabRequest struct {
	Key      string `json:"key,omitempty"`
	Word     string `json:"word"`
	WordType string `json:"type,omitempty"`
	Meaning  string `json:"meaning,omitempty"`
	IPA      string `json:"ipa,omitempty"`
	Note     string `json:"note,omitempty"`
	Example  string `json:"example,omitempty"`
	Status   string `json:"status,omitempty"`
}

type UserVocabBank struct {
	ID              int       `json:"id" gorm:"id,primaryKey;autoIncrement"`
	Key             string    `json:"key,omitempty"`
	Value           string    `json:"value,omitempty"`
	WordClass       string    `json:"word_class,omitempty"`
	Meaning         string    `json:"meaning,omitempty"`
	IPA             string    `json:"ipa,omitempty" gorm:"column:ipa"`
	Explanation     string    `json:"explanation,omitempty"`
	Example         string    `json:"example,omitempty"`
	IsLearnedStatus bool      `json:"is_learned_status,omitempty"`
	UserId          string    `json:"user_id,omitempty" gorm:"column:user_id"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type VocabQuery struct {
	Key string `form:"key"`
}

func (UserVocabBank) TableName() string {
	return common.POSTGRES_TABLE_NAME_USER_VOCAB_BANK
}

type LookUpVocabRequest struct {
	QuizId        int    `form:"quiz_id" binding:"required"`
	SentenceIndex int    `form:"sentence_index" binding:"required"`
	WordIndex     int    `form:"vocab_index" binding:"required"`
	Word          string `form:"word"`
}

type Vocab struct {
	ID          int            `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	VocabID     string         `json:"-" gorm:"column:vocab_id;not null"`
	Value       string         `json:"-" gorm:"column:value;not null"`
	WordDisplay string         `json:"word_display" gorm:"column:word_display;not null"`
	WordClass   string         `json:"word_class" gorm:"column:word_class;not null"`
	Meaning     string         `json:"meaning" gorm:"column:meaning;not null"`
	IPA         string         `json:"ipa" gorm:"column:ipa;not null"`
	Explanation string         `json:"explanation" gorm:"column:explanation;not null"`
	Collocation string         `json:"collocation" gorm:"column:collocation;not null"`
	Example     datatypes.JSON `json:"example" gorm:"column:example;not null"`
}

// TableName overrides the default table name for GORM
func (Vocab) TableName() string {
	return common.POSTGRES_TABLE_NAME_VOCAB_AI
}
