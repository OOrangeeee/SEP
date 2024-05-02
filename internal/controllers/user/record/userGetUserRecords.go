package controllers

import (
	services "SEP/internal/services/user/record"
	"github.com/labstack/echo/v4"
)

func UserGetUserRecordsController(c echo.Context) error {
	return services.GetUserRecordsService(c)
}
