package dataModels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName           string `gorm:"unique"`
	UserPassword       string
	UserEmail          string `gorm:"unique"`
	UserNickName       string
	UserIsActive       bool
	UserActivationCode string `gorm:"unique"`
}
