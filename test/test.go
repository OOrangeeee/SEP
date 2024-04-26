package main

import (
	"SEP/internal/configs"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	configs.InitLog()
	configs.InitViper()
	configs.InitDB()
	configs.GetRouterConfig(e)
	configs.PostRouterConfig(e)
	configs.PutRouterConfig(e)
	configs.DeleteRouterConfig(e)
	configs.InitMiddleware(e)
	configs.InitMiddleware(e)
	e.GET("/api/csrf-token", func(c echo.Context) error {
		// 从上下文中获取 CSRF 令牌
		csrfToken := c.Get("csrf").(string)

		// 将 CSRF 令牌作为响应头返回
		c.Response().Header().Set("X-CSRF-Token", csrfToken)
		return c.NoContent(http.StatusOK)
	})
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":714"))
}
