package services

import (
	"SEP/internal/mappers"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func DeleteUser(c echo.Context) error {
	authToken := c.FormValue("token")
	if authToken != viper.GetString("config.token") {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": "权限不足",
		}).Error("权限不足")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "权限不足",
		})
	}
	username := c.FormValue("username")
	userMapper := mappers.UserMapper{}
	userEmailMapper := mappers.UserEmailMapper{}
	recordMapper := mappers.RecordMapper{}
	users, err := userMapper.GetUsersByUserName(username)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": err.Error(),
		}).Error("查找用户出错")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "查找用户出错",
		})
	}
	if len(users) == 0 {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": "用户不存在",
		}).Error("用户不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	user := users[0]
	userEmailTemp := user.UserEmail
	userEmails, err := userEmailMapper.GetUserEmailsByUserEmail(userEmailTemp)
	userEmail := userEmails[0]
	err = userEmailMapper.DeleteUnscopedUserEmail(userEmail)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": err.Error(),
		}).Error("删除用户邮箱失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "删除用户邮箱失败",
		})
	}
	err = recordMapper.DeleteRecordsByUserId(user.ID)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": err.Error(),
		}).Error("删除用户记录失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "删除用户记录失败",
		})
	}
	err = userMapper.DeleteUnscopedUser(user)
	if err != nil {
		utils.Log.WithFields(map[string]interface{}{
			"error_message": err.Error(),
		}).Error("删除用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "删除用户失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":            user,
		"success_message": "删除成功",
	})
}
