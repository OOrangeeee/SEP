package main

import (
	"SEP/internal/configs"
	"SEP/internal/utils"
)

func main() {
	utils.InitLog()
	configs.InitViper()
	featureTool := utils.FeatureTool{}
	result, err := featureTool.Detect("E:\\go\\workplace\\src\\SEP\\ai\\Polyp_detection\\images\\1.png")
	if err != nil {
		println(err)
	}
	println(result)
}
