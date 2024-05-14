package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCSRFTokenService(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "CSRF Token 获取成功",
	})
}
