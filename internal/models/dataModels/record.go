package dataModels

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model
	UserId      int
	URL         string
	Type        string
	Time        time.Time
	PatientName string
}
