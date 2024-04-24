package main

import (
	"SEP/internal/configs"
	"SEP/internal/models/dataModels"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	configs.InitViper()
	println(viper.GetString("database.dataBaseUserName"))
	println(viper.GetString("database.dataBasePassword"))
	println(viper.GetString("database.dataBaseIp"))
	println(viper.GetString("database.dataBasePort"))
	println(viper.GetString("database.dataBaseName"))

	configs.InitDB()

	err := configs.DB.AutoMigrate(&dataModels.User{})
	if err != nil {
		configs.Log.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Failed to create table")
	}

	user := dataModels.User{
		UserName:     "test1",
		UserPassword: "test1p",
		UserEmail:    "test1@example.com",
		UserNickName: "test1Nick",
	}

	err = configs.DB.Create(&user).Error
	if err != nil {
		configs.Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to create user")
	}

	println("User created")

}
