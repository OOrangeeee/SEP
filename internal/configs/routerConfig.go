package configs

import (
	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	e.GET("/users/account/:id", nil)
	e.GET("/users/records-all/:userid", nil)
	e.GET("/users/records/:userid/:recordsid", nil)
}

func PostRouterConfig(e *echo.Echo) {
	e.POST("/users/account", nil)
	e.POST("/users/login", nil)
	e.POST("/detection", nil)
	e.POST("/segmentation", nil)
	e.POST("/track", nil)
}

func PutRouterConfig(e *echo.Echo) {
	e.PUT("/users/account/activation/token:", nil)
	e.PUT("/users/account", nil)
}

func DeleteRouterConfig(e *echo.Echo) {
	e.DELETE("/users/records", nil)
}
