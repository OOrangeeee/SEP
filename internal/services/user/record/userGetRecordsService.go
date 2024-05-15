package services

import (
	"SEP/internal/mappers"
	"SEP/internal/models/infoModels"
	"SEP/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	var recordsInfos []*infoModels.Record
	for _, record := range records {
		recordInfo := infoModels.Record{
			ID:          record.ID,
			URL:         record.URL,
			Type:        record.Type,
			Time:        record.Time,
			PatientName: record.PatientName,
		}
		recordsInfos = append(recordsInfos, &recordInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户记录成功",
		"records":         recordsInfos,
	})
}

func GetAUserRecordService(paramsMap map[string]string, c echo.Context) error {
	recordMapper := mappers.RecordMapper{}
	recordId := paramsMap["recordId"]
	if recordId == "" {
		utils.Log.WithField("error_message", "缺少记录id").Error("缺少记录id")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "缺少记录id",
		})
	}
	recordIdInt, err := strconv.ParseUint(recordId, 10, 32)
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
			"error_message": "获取用户记录失败",
		}).Error("获取用户记录失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户记录失败",
		})
	}
	recordInfo := infoModels.Record{
		ID:          record.ID,
		URL:         record.URL,
		Type:        record.Type,
		Time:        record.Time,
		PatientName: record.PatientName,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户记录成功",
		"record":          recordInfo,
	})
}

func GetUserRecordsByPatientNameService(paramsMap map[string]string, c echo.Context) error {
	recordMapper := mappers.RecordMapper{}
	patientName := paramsMap["patientName"]
	if patientName == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "患者姓名为空",
		}).Error("患者姓名为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "患者姓名为空",
		})
	}
	records, err := recordMapper.GetRecordsByPatientName(patientName)
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
	var recordsInfos []*infoModels.Record
	for _, record := range records {
		recordInfo := infoModels.Record{
			ID:          record.ID,
			URL:         record.URL,
			Type:        record.Type,
			Time:        record.Time,
			PatientName: record.PatientName,
		}
		recordsInfos = append(recordsInfos, &recordInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户记录成功",
		"records":         recordsInfos,
	})
}
