package controllers

import (
	services "SEP/internal/services/feature"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func DetectController(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文件上传失败",
		}).Error("文件上传失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文件上传失败",
		})
	}
	patientName := c.FormValue("patient-name")
	path, err := saveFileAndGetPath(file)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文件保存失败",
		}).Error("文件保存失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "文件保存失败",
		})
	}
	mapParams := make(map[string]string)
	mapParams["source"] = path
	mapParams["patientName"] = patientName
	return services.DetectService(mapParams, c)
}

func SegmentController(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文件上传失败",
		}).Error("文件上传失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文件上传失败",
		})
	}
	patientName := c.FormValue("patient-name")
	path, err := saveFileAndGetPath(file)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文件保存失败",
		}).Error("文件保存失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "文件保存失败",
		})
	}
	mapParams := make(map[string]string)
	mapParams["source"] = path
	mapParams["patientName"] = patientName
	return services.SegmentService(mapParams, c)
}

func saveFileAndGetPath(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "文件关闭失败",
			}).Error("文件关闭失败")
		}
	}(src)
	dstPath := "./uploads/" + file.Filename
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "目标文件关闭失败",
			}).Error("目标文件关闭失败")
		}
	}(dst)
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return dstPath, nil
}
