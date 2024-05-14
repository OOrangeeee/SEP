package services

import (
	"SEP/internal/mappers"
	"SEP/internal/models/infoModels"
	"SEP/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetUserInfoService(c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userId := c.Get("userId").(uint)
	isAdmin := c.Get("isAdmin").(bool)
	users, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户信息失败",
		}).Error("获取用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户信息失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	user := users[0]
	userInfo := infoModels.User{
		UserName:     user.UserName,
		UserEmail:    user.UserEmail,
		UserNickName: user.UserNickName,
		UserIsAdmin:  isAdmin,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户信息成功",
		"userInfo":        userInfo,
	})
}
