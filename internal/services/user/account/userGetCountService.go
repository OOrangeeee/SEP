package services

import (
	"SEP/internal/mappers"
	"SEP/internal/utils"
	"github.com/spf13/viper"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserInfo struct {
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserNickName string `json:"user_nick_name"`
}

func GetUserCount(c echo.Context) error {
	userMapper := mappers.UserMapper{}
	authToken := c.FormValue("token")
	if authToken != viper.GetString("config.token") {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": "权限不足",
		}).Error("权限不足")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "权限不足",
		})
	}
	users, err := userMapper.GetAllUsers()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户列表失败",
		}).Error("获取用户列表失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户列表失败",
		})
	}

	var userInfos []UserInfo
	for _, user := range users {
		userInfos = append(userInfos, UserInfo{
			UserName:     user.UserName,
			UserEmail:    user.UserEmail,
			UserNickName: user.UserNickName,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户列表成功",
		"count":           len(users),
		"userInfos":       userInfos,
	})
}
