package entity

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ActivityGroupId string `json:"activity_group_id" gorm:"type:varchar(255)"`
	Title           string `json:"title" gorm:"type:varchar(255)"`
	IsActive        string `json:"is_active" gorm:"type:varchar(255)"`
	Priority        string `json:"priority" gorm:"type:varchar(255);default:very-high"`
}
