package services

import (
	"SEP/internal/mappers"
	"SEP/internal/models/dataModels"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func DetectService(paramsMap map[string]string, c echo.Context) error {
	csrfTool := utils.CSRFTool{}
	recordMapper := mappers.RecordMapper{}
	userId := c.Get("userId").(uint)
	source := paramsMap["source"]
	patientName := paramsMap["patientName"]
	if source == "" {
		utils.Log.WithField("error_message", "参数错误").Error("参数错误")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "参数错误",
		})
	}
	featureTool := utils.FeatureTool{}
	result, err := featureTool.Detect(source)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "检测失败",
		}).Error("检测失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "检测失败",
		})
	}
	newRecord := dataModels.Record{
		UserId:      userId,
		URL:         result,
		Type:        "detect",
		Time:        time.Now(),
		PatientName: patientName,
	}
	err = recordMapper.AddRecord(&newRecord)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "记录保存失败",
		}).Error("记录保存失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "记录保存失败",
		})
	}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success_message": "检测成功，请前往记录查看结果",
	})
}

func SegmentService(paramsMap map[string]string, c echo.Context) error {
	csrfTool := utils.CSRFTool{}
	recordMapper := mappers.RecordMapper{}
	userId := c.Get("userId").(uint)
	source := paramsMap["source"]
	patientName := paramsMap["patientName"]
	if source == "" {
		utils.Log.WithField("error_message", "参数错误").Error("参数错误")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "参数错误",
		})
	}
	featureTool := utils.FeatureTool{}
	result, err := featureTool.Segment(source)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "分割失败",
		}).Error("分割失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "分割失败",
		})
	}
	newRecord := dataModels.Record{
		UserId:      userId,
		URL:         result,
		Type:        "segment",
		Time:        time.Now(),
		PatientName: patientName,
	}
	err = recordMapper.AddRecord(&newRecord)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "记录保存失败",
		}).Error("记录保存失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "记录保存失败",
		})
	}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success_message": "分割成功，请前往记录查看结果",
	})
}
