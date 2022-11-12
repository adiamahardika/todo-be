package entity

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255)"`
	Email string `json:"email" gorm:"type:varchar(255)"`
}
