package main

import (
	"SEP/internal/configs"
	"SEP/internal/utils"
)

func main() {
	utils.InitLog()
	configs.InitViper()
	uploadTool := utils.UploadTool{}
	ans, err := uploadTool.UploadImage("E:/go/workplace/src/SEP/testing.png")
	if err != nil {
		println(err.Error())
	}
	println(ans)
}
