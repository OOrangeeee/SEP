package utils

import (
	"github.com/labstack/echo/v4"
)

type CSRFTool struct {
}

func (cT *CSRFTool) GetCSRFToken(c echo.Context) string {
	csrfToken, ok := c.Get("csrf").(string)
	if !ok {
		Log.WithField("error_message", "无法获取CSRF Token").Error("无法获取CSRF Token")
		return ""
	}
	return csrfToken
}

func (cT *CSRFTool) SetCSRFToken(c echo.Context) bool {
	csrfToken := cT.GetCSRFToken(c)
	if csrfToken == "" {
		Log.WithField("error_message", "无法获取CSRF Token").Error("无法获取CSRF Token")
		return false
	}
	c.Response().Header().Set("X-CSRF-Token", csrfToken)
	return true
}
