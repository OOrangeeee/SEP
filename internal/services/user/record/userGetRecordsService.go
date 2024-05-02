package services

import (
	"SEP/internal/mappers"
	"SEP/internal/models/infoModels"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetUserRecordsService(c echo.Context) error {
	recordMapper := mappers.RecordMapper{}
	userId := c.Get("userId").(uint)
	records, err := recordMapper.GetRecordsByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户记录失败",
		}).Error("获取用户记录失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户记录失败",
		})
	}
	// 根据每个record生成对应的infoModels.Record，放入一个切片中
	var recordsInfo []*infoModels.Record
	for _, record := range records {
		recordInfo := infoModels.Record{
			URL:         record.URL,
			Type:        record.Type,
			Time:        record.Time,
			PatientName: record.PatientName,
		}
		recordsInfo = append(recordsInfo, &recordInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户记录成功",
		"records":         recordsInfo,
	})
}
