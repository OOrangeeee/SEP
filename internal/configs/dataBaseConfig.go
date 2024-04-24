package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Failed to connect to database")
	}
}
