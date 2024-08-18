package controllers

import (
	services "SEP/internal/services/security/config"
	"github.com/labstack/echo/v4"
)

func ChangeConfig(c echo.Context) error {
	return services.ChangeConfigService(c)
}
