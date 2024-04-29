package controllers

import (
	services "SEP/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func UserGetInfoController(c echo.Context) error {
	return services.GetUserInfoService(c)
}
