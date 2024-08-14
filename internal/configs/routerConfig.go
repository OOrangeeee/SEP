package configs

import (
	featureControllers "SEP/internal/controllers/feature"
	securityCSRFControllers "SEP/internal/controllers/security/CSRF"
	useAccountControllers "SEP/internal/controllers/user/account"
	recordControllers "SEP/internal/controllers/user/record"

	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	// 获取用户信息
	e.GET("/users/account", useAccountControllers.UserGetInfoController)
	// 获取用户记录
	e.GET("/users/records-all", recordControllers.UserGetUserRecordsController)
	// 获取用户记录
	e.GET("/users/records/:recordsid", recordControllers.UserGetAUserRecordController)
	// 获取CSRF Token
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activationCode", useAccountControllers.UserConfirmController)
	// 获取用户记录
	e.GET("/users/records/patient", recordControllers.UserGetUserRecordsByPatientNameController)
}

func PostRouterConfig(e *echo.Echo) {
	// 注册
	e.POST("/users/account", useAccountControllers.UserRegisterController)
	// 登录
	e.POST("/users/login", useAccountControllers.UserLoginController)
	// 检测
	e.POST("/detection", featureControllers.DetectController)
	// 分割
	e.POST("/segmentation", featureControllers.SegmentController)
	// 跟踪
	e.POST("/track", featureControllers.TrackController)
}

func PutRouterConfig(e *echo.Echo) {
	// 修改昵称
	e.PUT("/users/account/nickname", useAccountControllers.UserUpdateNicknameController)
	// 修改密码
	e.PUT("/users/account/password", useAccountControllers.UserUpdatePasswordController)
}

func DeleteRouterConfig(e *echo.Echo) {
	// 删除记录
	e.DELETE("/users/records/:recordid", recordControllers.UserDeleteRecordController)
}
