package model

import (
	"time"

	"gorm.io/gorm"
)

type ActivityLog struct {
	gorm.Model
	LogDetail  string			`gorm:"type:text;not null"`
	UserId		 string   	`gorm:"not null"`
	Timestamp  time.Time 	`gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
