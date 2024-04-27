package controllers

import (
	services "SEP/internal/services/user/account"
	"github.com/labstack/echo/v4"
)

func UserConfirmController(c echo.Context) error {
	activationCode := c.Param("activationCode")
	return services.ConfirmUserService(activationCode, c)
}
