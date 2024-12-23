package models

import (
	"ielts-app-api/common"
	"time"
)

type VocabRequest struct {
	Value           string `json:"value"`
	WordClass       string `json:"word_class,omitempty"`
	Meaning         string `json:"meaning,omitempty"`
	IPA             string `json:"ipa,omitempty"`
	Explanation     string `json:"explanation,omitempty"`
	Example         string `json:"example,omitempty"`
	IsLearnedStatus bool   `json:"is_learned_status,omitempty"`
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

func (Vocab) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_VOCAB
}
