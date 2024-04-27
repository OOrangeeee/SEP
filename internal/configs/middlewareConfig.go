package configs

import (
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitMiddleware(e *echo.Echo) {

	//recover
	e.Use(middleware.Recover())

	//logger
	logFileLocation := filepath.Join("./logs", "httpLog.log")
	logFile, err := os.OpenFile(logFileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "打开httpLog文件失败",
		}).Panic("打开httpLog文件失败")
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "time:${time_rfc3339}\nmethod:${method}\nuri:${uri}\nstatus:${status}\nerror:${error}\nhost:${host}\npath:${path}\n\n\n",
	}))

	//csrf
	//配置文件区分开发和生产环境
	cookieSecure := viper.GetBool("CSRF.cookieSecure")
	csrfConfig := middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookiePath:     "/",
		CookieSecure:   cookieSecure,
		CookieHTTPOnly: true,
		ContextKey:     "csrf",
		CookieMaxAge:   60,
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))

}
