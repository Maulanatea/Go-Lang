package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index, column:deleted_at"`
}
