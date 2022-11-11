package entity

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        string `json:"is_active"`
	Priority        string `json:"priority" default:"very-high"`
}
