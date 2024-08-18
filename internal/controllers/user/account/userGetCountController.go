package controllers

import (
	services "SEP/internal/services/user/account"
	"github.com/labstack/echo/v4"
)

func UserGetCountController(c echo.Context) error {
	return services.GetUserCount(c)
}
