package services

import (
	"SEP/internal/mappers"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func DeleteRecordService(paramsMap map[string]string, c echo.Context) error {
	recordMapper := mappers.RecordMapper{}
	recordId := paramsMap["recordId"]
	if recordId == "" {
		utils.Log.WithField("error_message", "缺少记录id").Error("缺少记录id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "缺少记录id",
		})
	}
	recordIdInt, err := strconv.ParseUint(recordId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "记录id转换失败",
		}).Error("记录id转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "记录id转换失败",
		})
	}
	record, err := recordMapper.GetRecordById(uint(recordIdInt))
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取记录失败",
		}).Error("获取记录失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "获取记录失败",
		})
	}
	err = recordMapper.DeleteRecord(record)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "删除记录失败",
		}).Error("删除记录失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "删除记录失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "删除记录成功",
	})
}
