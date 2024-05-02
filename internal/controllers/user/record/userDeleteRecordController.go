package controllers

import (
	services "SEP/internal/services/user/record"
	"github.com/labstack/echo/v4"
)

func UserDeleteRecordController(c echo.Context) error {
	recordid := c.Param("recordid")
	paramsMap := make(map[string]string)
	paramsMap["recordId"] = recordid
	return services.DeleteRecordService(paramsMap, c)
}
