package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"udpated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Title     string         `json:"title" gorm:"type:varchar(255)"`
	Email     string         `json:"email" gorm:"type:varchar(255)"`
}
