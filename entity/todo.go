package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint           `json:"id" gorm:"primarykey"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"udpated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	ActivityGroupId string         `json:"activity_group_id" form:"activity_group_id" gorm:"type:varchar(255)"`
	Title           string         `json:"title" gorm:"type:varchar(255)"`
	IsActive        bool           `json:"is_active,default=true" gorm:"type:boolean;default:true"`
	Priority        string         `json:"priority" form:"priority" gorm:"type:varchar(255);default:very-high"`
}
