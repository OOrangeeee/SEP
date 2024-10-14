package configs

import (
	"SEP/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

func InitMiddleware(e *echo.Echo, jwtSecret string) {

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

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://base.polypeye.cn", "https://new.polypeye.cn"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Csrf-Token", "Origin", "Accept"},
		ExposeHeaders:    []string{"X-Csrf-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//csrf
	csrfConfig := middleware.CSRFConfig{
		Skipper: func(c echo.Context) bool {
			if (c.Path() == "/config" && c.Request().Method == "PUT") || (c.Path() == "/users" && c.Request().Method == "DELETE") {
				return true
			}
			return false
		},
		TokenLookup:    "header:X-Csrf-Token",
		ContextKey:     "csrf",
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		TokenLength:    32,
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			csrfToken := c.Get("csrf")
			if csrfToken != nil {
				if token, ok := csrfToken.(string); ok && token != "" {
					c.Response().Header().Add("X-Csrf-Token", token)
				}
			}
			return next(c)
		}
	})

	//JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			if (c.Path() == "/users" && c.Request().Method == "GET") || (c.Path() == "/users" && c.Request().Method == "DELETE") || (c.Path() == "/config" && c.Request().Method == "PUT") || (c.Path() == "/csrf-token" && c.Request().Method == "GET") || (c.Path() == "/users/account/activation/:activationCode" && c.Request().Method == "GET") || (c.Path() == "/users/account" && c.Request().Method == "POST") || (c.Path() == "/users/login" && c.Request().Method == "POST") {
				return true
			}
			utils.Log.WithFields(logrus.Fields{
				"jwtSecret": jwtSecret,
			}).Info("JWT中间件启用")
			return false
		},
		SigningKey:  jwtSecret,
		TokenLookup: "header:Authorization:Bearer ",
		ErrorHandler: func(c echo.Context, err error) error {
			utils.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"jwtS":  jwtSecret,
				"jwt":   c.Request().Header.Get("Authorization"),
			}).Error("JWT validation failed")
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired JWT")
		},

		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims, ok := user.Claims.(jwt.MapClaims)
			if !ok {
				c.Logger().Error("无法断言JWT claims为MapClaims类型")
				return
			}
			userIdTmp, ok := claims["UserId"].(float64)
			userId := uint(userIdTmp)
			if !ok {
				c.Logger().Error("用户ID claims类型断言错误")
				return
			}
			isAdmin, ok := claims["IsAdmin"].(bool)
			if !ok {
				c.Logger().Error("管理员状态claims类型断言错误")
				return
			}
			utils.Log.WithFields(logrus.Fields{
				"userId":          userId,
				"isAdmin":         isAdmin,
				"success_message": "用户登录成功",
			}).Info("用户登录成功")

			c.Set("userId", userId)
			c.Set("isAdmin", isAdmin)
		},
	}))
}
