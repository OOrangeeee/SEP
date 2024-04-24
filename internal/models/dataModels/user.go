package dataModels

type User struct {
	UserID       int    `gorm:"primaryKey;autoIncrement"`
	UserName     string `gorm:"unique"`
	UserPassword string
	UserEmail    string `gorm:"unique"`
	UserNickName string
}
