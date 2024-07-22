package utils

import (
	"SEP/internal/models/dataModels"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var dataBaseUserName string
var dataBasePassword string
var dataBaseIp string
var dataBasePort string
var dataBaseName string

func InitDB() {
	dataBaseName = viper.GetString("database.dataBaseName")
	dataBaseUserName = viper.GetString("database.dataBaseUserName")
	dataBasePassword = viper.GetString("database.dataBasePassword")
	dataBaseIp = viper.GetString("database.dataBaseIp")
	dataBasePort = viper.GetString("database.dataBasePort")
	var err error
	DB, err = gorm.Open(postgres.Open("postgres://"+dataBaseUserName+":"+dataBasePassword+"@"+dataBaseIp+":"+dataBasePort+"/"+dataBaseName), &gorm.Config{})
	cnt := 0
	for err != nil && cnt < 50 {
		println("数据库连接失败，正在重试")
		DB, err = gorm.Open(postgres.Open("postgres://"+dataBaseUserName+":"+dataBasePassword+"@"+dataBaseIp+":"+dataBasePort+"/"+dataBaseName), &gorm.Config{})
		cnt++
		time.Sleep(time.Second)
	}
	DB.Model(&dataModels.User{})
	err = DB.AutoMigrate(&dataModels.User{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建用户表失败",
		}).Panic("创建用户表失败")
	}
	DB.Model(&dataModels.UserEmail{})
	err = DB.AutoMigrate(&dataModels.UserEmail{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建用户邮箱表失败",
		}).Panic("创建用户邮箱表失败")
	}
	DB.Model(&dataModels.Record{})
	err = DB.AutoMigrate(&dataModels.Record{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建记录表失败",
		}).Panic("创建记录表失败")
	}
}
