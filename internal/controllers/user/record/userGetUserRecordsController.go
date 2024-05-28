package controllers

import (
	services "SEP/internal/services/user/record"

	"github.com/labstack/echo/v4"
)

func UserGetUserRecordsController(c echo.Context) error {
	return services.GetUserRecordsService(c)
}

func UserGetAUserRecordController(c echo.Context) error {
	recordId := c.Param("recordsid")
	paramsMap := make(map[string]string)
	paramsMap["recordId"] = recordId
	return services.GetAUserRecordService(paramsMap, c)
}

func UserGetUserRecordsByPatientNameController(c echo.Context) error {
	patientName := c.QueryParam("patient-name")
	paramsMap := make(map[string]string)
	paramsMap["patientName"] = patientName
	return services.GetUserRecordsByPatientNameService(paramsMap, c)
}
