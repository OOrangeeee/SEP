package dataModels

import (
	"gorm.io/gorm"
	"time"
)

type UserEmail struct {
	gorm.Model
	Email         string    `gorm:"unique"`
	EmailLastSent time.Time `gorm:"default:null"`
}
