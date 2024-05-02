package configs

import (
	featureControllers "SEP/internal/controllers/feature"
	securityCSRFControllers "SEP/internal/controllers/security/CSRF"
	useAccountControllers "SEP/internal/controllers/user/account"
	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	// 获取用户信息
	e.GET("/users/account", useAccountControllers.UserGetInfoController)
	e.GET("/users/records-all/:userid", nil)
	e.GET("/users/records/:recordsid", nil)
	// 获取CSRF Token
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activationCode", useAccountControllers.UserConfirmController)
}

func PostRouterConfig(e *echo.Echo) {
	// 注册
	e.POST("/users/account", useAccountControllers.UserRegisterController)
	// 登录
	e.POST("/users/login", useAccountControllers.UserLoginController)
	// 检测
	e.POST("/detection", featureControllers.DetectController)
	e.POST("/segmentation", nil)
	e.POST("/track", nil)
}

func PutRouterConfig(e *echo.Echo) {
	// 修改昵称
	e.PUT("/users/account/nickname", useAccountControllers.UserUpdateNicknameController)
	// 修改密码
	e.PUT("/users/account/password", useAccountControllers.UserUpdatePasswordController)
}

func DeleteRouterConfig(e *echo.Echo) {
	e.DELETE("/users/records", nil)
}
