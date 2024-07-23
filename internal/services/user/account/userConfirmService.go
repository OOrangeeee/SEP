package services

import (
	"SEP/internal/mappers"
	"SEP/internal/utils"
	"github.com/spf13/viper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ConfirmUserService(ActivationCode string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	users, err := userMapper.GetUsersByUserActivationCode(ActivationCode)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "数据库查询失败",
		}).Error("数据库查询失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "数据库查询失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "激活码不存在").Error("激活码不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "激活码不存在",
		})
	}
	user := users[0]
	if user.UserIsActive {
		utils.Log.WithField("error_message", "用户已激活").Error("用户已激活")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户已激活",
		})
	}
	user.UserIsActive = true
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "数据库更新失败",
		}).Error("数据库更新失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "数据库更新失败",
		})
	}
	// 返回一个html页面
	html := viper.GetString("Confirm.return")
	html = strings.Replace(html, "{用户名}", user.UserName, -1)
	return c.HTML(http.StatusOK, html)
}
