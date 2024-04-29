package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
	"time"
)

var Log *logrus.Logger

func InitLog() {
	Log = logrus.New()
	Log.Formatter = &logrus.JSONFormatter{}
	Log.SetReportCaller(true)
	logFileLocation := filepath.Join("./logs", time.Now().Format("2006-01-02")+".log")
	Log.Out = &lumberjack.Logger{
		Filename:   logFileLocation,
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     30,
	}
}
