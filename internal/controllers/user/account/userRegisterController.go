package controllers

import (
	services "SEP/internal/services/user/account"
	"github.com/labstack/echo/v4"
)

func UserRegisterController(c echo.Context) error {
	paramMap := make(map[string]string)
	paramMap["userName"] = c.FormValue("user-name")
	paramMap["userPassword"] = c.FormValue("user-password")
	paramMap["userEmail"] = c.FormValue("user-email")
	paramMap["userNickName"] = c.FormValue("user-nickname")
	return services.RegisterUserService(paramMap, c)
}
