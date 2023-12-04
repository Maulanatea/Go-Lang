package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	Id         uint           `json:"id" gorm:"primarykey"`
	Image      string         `json:"image"`
	CategoryId uint           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
