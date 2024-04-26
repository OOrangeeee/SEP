package main

import (
	"SEP/internal/configs"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.InitLog()
	configs.InitViper()
	configs.InitDB()
	configs.InitMiddleware(e)
	configs.GetRouterConfig(e)
	configs.PostRouterConfig(e)
	configs.PutRouterConfig(e)
	configs.DeleteRouterConfig(e)
}
