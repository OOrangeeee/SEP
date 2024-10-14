package main

import (
	"SEP/internal/configs"
	"SEP/internal/router"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	utils.InitLog()
	configs.InitViper()
	utils.InitDB()
	configs.InitMiddleware(e)
	router.GetRouterConfig(e)
	router.PostRouterConfig(e)
	router.PutRouterConfig(e)
	router.DeleteRouterConfig(e)
	e.Logger.Fatal(e.Start(":714"))
}
