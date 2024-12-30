package models

import (
	"ielts-app-api/common"
	"time"
)

type VocabRequest struct {
	Word     string `json:"word"`
	WordType string `json:"type,omitempty"`
	Meaning  string `json:"meaning,omitempty"`
	IPA      string `json:"ipa,omitempty"`
	Note     string `json:"note,omitempty"`
	Example  string `json:"example,omitempty"`
	Status   string `json:"status,omitempty"`
}

type Vocab struct {
	ID              int       `json:"id" gorm:"id,primaryKey;autoIncrement"`
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
	Value string `form:"value"`
}

func (Vocab) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_VOCAB
}
