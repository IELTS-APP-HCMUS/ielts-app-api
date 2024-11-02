package models

import (
	"ielts-app-api/common"
	"time"
)

type TargetRequest struct {
	TargetStudyDuration *int     `json:"target_study_duration"`
	TargetReading       *float32 `json:"target_reading"`
	TargetListening     *float32 `json:"target_listening"`
	TargetSpeaking      *float32 `json:"target_speaking"`
	TargetWriting       *float32 `json:"target_writing"`
	NextExamDate        *string  `json:"next_exam_date"`
}

type Target struct {
	ID                  string    `json:"id" gorm:"type:uuid;primaryKey"`
	TargetStudyDuration int       `json:"target_study_duration" gorm:"column:target_study_duration"`
	TargetReading       float32   `json:"target_reading" gorm:"column:target_reading"`
	TargetListening     float32   `json:"target_listening" gorm:"column:target_listening"`
	TargetSpeaking      float32   `json:"target_speaking" gorm:"column:target_speaking"`
	TargetWriting       float32   `json:"target_writing" gorm:"column:target_writing"`
	NextExamDate        time.Time `json:"next_exam_date" gorm:"column:next_exam_date"`
}

func (Target) TableName() string {
	// define common variable here
	return common.POSTGRES_TABLE_NAME_TARGETS
}
