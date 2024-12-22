package models

import (
	"ielts-app-api/common"
)

type VocabRequest struct {
	Value           string `json:"value"`
	WordClass       string `json:"word_class,omitempty"`
	Meaning         string `json:"meaning,omitempty"`
	IPA             string `json:"ipa,omitempty"`
	Explanation     string `json:"explanation,omitempty"`
	Example         string `json:"example,omitempty"`
	IsLearnedStatus bool   `json:"status,omitempty"`
}

type Vocab struct {
	ID              int    `json:"id" gorm:"id,primaryKey"`
	Value           string `json:"value"`
	WordClass       string `json:"word_class,omitempty"`
	Meaning         string `json:"meaning,omitempty"`
	IPA             string `json:"ipa,omitempty" gorm:"column:ipa"`
	Explanation     string `json:"explanation,omitempty"`
	Example         string `json:"example,omitempty"`
	IsLearnedStatus bool   `json:"status,omitempty"`
	UserId          string `json:"user_id,omitempty" gorm:"column:user_id"`
	// CreatedAt       int    `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt       int    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Vocab) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_VOCAB
}
