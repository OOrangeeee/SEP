package main

import (
	"SEP/internal/configs"
	"SEP/internal/router"
	"SEP/internal/utils"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()
	jwtS := os.Getenv("JWT_SECRET")
	utils.InitLog()
	configs.InitViper()
	utils.InitDB()
	configs.InitMiddleware(e, jwtS)
	router.GetRouterConfig(e)
	router.PostRouterConfig(e)
	router.PutRouterConfig(e)
	router.DeleteRouterConfig(e)
	e.Logger.Fatal(e.Start(":714"))
}
