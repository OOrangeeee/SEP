package configs

import (
	securityCSRFControllers "SEP/internal/controllers/security/CSRF"
	useAccountControllers "SEP/internal/controllers/user/account"
	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	e.GET("/users/account/:id", nil)
	e.GET("/users/records-all/:userid", nil)
	e.GET("/users/records/:recordsid", nil)
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activationCode", useAccountControllers.UserConfirmController)
}

func PostRouterConfig(e *echo.Echo) {
	// 注册
	e.POST("/users/account", useAccountControllers.UserRegisterController)
	e.POST("/users/login", nil)
	e.POST("/detection", nil)
	e.POST("/segmentation", nil)
	e.POST("/track", nil)
}

func PutRouterConfig(e *echo.Echo) {
	e.PUT("/users/account", nil)
}

func DeleteRouterConfig(e *echo.Echo) {
	e.DELETE("/users/records", nil)
}
