package services

import (
	"SEP/internal/configs"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var configLock sync.RWMutex

func ChangeConfigService(c echo.Context) error {
	configPath := "./configs/config.yaml"
	// 检验权限
	authToken := c.FormValue("token")
	if authToken != viper.GetString("config.token") {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": "更新权限不足",
		}).Error("更新权限不足")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "更新权限不足",
		})
	}
	// 获取上传的文件
	file, err := c.FormFile("config")
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "获取上传文件失败",
		}).Error("获取上传文件失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "获取上传文件失败",
		})
	}

	// 检查文件名是否为 config.yaml
	if filepath.Base(file.Filename) != "config.yaml" {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": "上传的文件不是 config.yaml",
		}).Error("上传的文件不是 config.yaml")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "上传的文件必须是 config.yaml",
		})
	}

	// 获取写锁
	configLock.Lock()
	defer configLock.Unlock()

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "打开上传文件失败",
		}).Error("打开上传文件失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "打开上传文件失败",
		})
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			utils.Log.WithFields(map[string]interface{}{
				"error":         err,
				"error_message": "关闭上传文件失败",
			}).Error("关闭上传文件失败")
		}
	}(src)

	// 创建临时文件
	tempFile, err := os.CreateTemp(filepath.Dir(configPath), "config_temp_*.yaml")
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "创建临时文件失败",
		}).Error("创建临时文件失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "创建临时文件失败",
		})
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			utils.Log.WithFields(map[string]interface{}{
				"error":         err,
				"error_message": "删除临时文件失败",
			}).Error("删除临时文件失败")
		}
	}(tempFile.Name()) // 确保临时文件被删除

	// 复制文件内容到临时文件
	if _, err = io.Copy(tempFile, src); err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "复制文件内容失败",
		}).Error("复制文件内容失败")
		err = tempFile.Close()
		if err != nil {
			utils.Log.WithFields(map[string]interface{}{
				"error":         err,
				"error_message": "关闭中间文件失败",
			}).Error("关闭中间文件失败")
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "复制文件内容失败",
		})
	}
	err = tempFile.Close()
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "关闭中间文件失败",
		}).Error("关闭中间文件失败")
	}

	// 替换原文件
	if err := os.Rename(tempFile.Name(), configPath); err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error":         err,
			"error_message": "替换配置文件失败",
		}).Error("替换配置文件失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "替换配置文件失败",
		})
	}

	configs.InitViper()
	utils.Log.Info("配置文件更新成功")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "配置文件更新成功",
	})
}
