package models

import (
	"ielts-app-api/common"
	"time"
)

type MasterDataCategory struct {
	ID        int       `json:"id" gorm:"primaryKey,autoIncrement"` // Cột ID tự tăng (SERIAL trong PostgreSQL)
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type MasterData struct {
	ID         int       `json:"id" gorm:"primaryKey,autoIncrement"`
	CategoryID int       `json:"category_id" gorm:"column:category_id not null"` // Liên kết với master_data_categories
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (MasterDataCategory) TableName() string {
	return common.POSTGRES_TABLE_NAME_MASTER_DATA_CATEGORIES
}
func (MasterData) TableName() string {
	return common.POSTGRES_TABLE_NAME_MASTER_DATA
}
