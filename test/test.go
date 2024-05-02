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
	result, err := featureTool.Track("E:\\university\\contest\\life\\out1\\P1_0.mp4")
	if err != nil {
		println(err.Error())
		panic(err)
	}
	println(result)
}
