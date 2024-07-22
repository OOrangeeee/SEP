package mappers

import (
	"SEP/internal/models/dataModels"
	"SEP/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type UserEmailMapper struct {
}

func (uem *UserEmailMapper) AddNewUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Create(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) DeleteUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Delete(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) UpdateUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Save(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) GetAllUserEmails() ([]*dataModels.UserEmail, error) {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails)
	return userEmails, result.Error
}

func (uem *UserEmailMapper) GetUserEmailsByUserEmail(userEmail string) ([]*dataModels.UserEmail, error) {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails, "Email=?", userEmail)
	return userEmails, result.Error
}

func (uem *UserEmailMapper) IsUserEmailSendInTimeRange(email string) bool {
	timeRange := viper.GetInt("email.emailOfRegister.timeRange")
	beforeTime := time.Now().Add(-time.Duration(timeRange) * time.Minute)
	var userEmail *dataModels.UserEmail
	result := utils.DB.First(&userEmail, "Email=?", email)
	if result.Error != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         result.Error,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
	}
	if userEmail.EmailLastSent.After(beforeTime) {
		return true
	}
	return false
}

func (uem *UserEmailMapper) IsExistUserEmail(email string) bool {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails, "Email=?", email)
	if result.Error != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         result.Error,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
	}
	if len(userEmails) == 0 {
		return false
	}
	return true
}
