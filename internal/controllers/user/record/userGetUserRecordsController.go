package controllers

import (
	services "SEP/internal/services/user/record"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserGetUserRecordsController(c echo.Context) error {
	return services.GetUserRecordsService(c)
}

func UserGetAUserRecordController(c echo.Context) error {
	recordId := c.Param("recordsid")
	if recordId == "" {
		utils.Log.WithField("error_message", "缺少记录id").Error("缺少记录id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "缺少记录id",
		})
	}
	paramsMap := make(map[string]string)
	paramsMap["recordId"] = recordId
	return services.GetAUserRecordService(paramsMap, c)
}
