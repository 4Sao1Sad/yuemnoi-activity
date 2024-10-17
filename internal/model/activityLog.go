package model

import (
	"gorm.io/gorm"
)

type ActivityLog struct {
	gorm.Model
	LogDetail  string			`gorm:"type:text;not null"`
	UserId		 uint   	  `gorm:"not null"`
	Timestamp  string 		`gorm:"not null"`
}
