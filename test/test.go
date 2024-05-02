package main

import (
	"SEP/internal/configs"
	"SEP/internal/utils"
)

func main() {
	utils.InitLog()
	configs.InitViper()
	featureTool := utils.FeatureTool{}
	result, err := featureTool.Detect("E:\\university\\contest\\life\\Polyp_detection\\images\\1.png")
	if err != nil {
		panic(err)
	}
	println(result)
	result, err = featureTool.Segment("E:\\university\\contest\\life\\Polyp_detection\\images\\3.png")
	if err != nil {
		panic(err)
	}
	println(result)
}
