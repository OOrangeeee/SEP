package utils

import (
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type ImageTool struct {
}

func (it *ImageTool) ChangeColorsAndOverlay(firstImagePath, secondImagePath string) error {
	firstFile, err := os.Open(firstImagePath)
	if err != nil {
		return err
	}
	defer func(firstFile *os.File) {
		err := firstFile.Close()
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "文件关闭失败",
			})
		}
	}(firstFile)
	firstImg, _, err := image.Decode(firstFile)
	if err != nil {
		return err
	}
	firstBounds := firstImg.Bounds()
	secondFile, err := os.Open(secondImagePath)
	if err != nil {
		return err
	}
	defer func(secondFile *os.File) {
		err := secondFile.Close()
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "文件关闭失败",
			})
		}
	}(secondFile)
	secondImg, _, err := image.Decode(secondFile)
	if err != nil {
		return err
	}
	secondImgResized := resize.Resize(uint(firstBounds.Dx()), uint(firstBounds.Dy()), secondImg, resize.NearestNeighbor)
	processedImg := image.NewRGBA(firstBounds)
	for y := firstBounds.Min.Y; y < firstBounds.Max.Y; y++ {
		for x := firstBounds.Min.X; x < firstBounds.Max.X; x++ {
			pixel := secondImgResized.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			if originalColor.R == 0 && originalColor.G == 0 && originalColor.B == 0 {
				processedImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			} else if originalColor.R == 255 && originalColor.G == 255 && originalColor.B == 255 { // 白色变为绿色
				processedImg.Set(x, y, color.RGBA{0, 255, 0, 255})
			} else {
				processedImg.Set(x, y, originalColor)
			}
		}
	}
	overlay := image.NewRGBA(firstBounds)
	draw.Draw(overlay, firstBounds, firstImg, image.Point{}, draw.Src)
	draw.Draw(overlay, firstBounds, processedImg, image.Point{}, draw.Over)
	outFile, err := os.Create(secondImagePath)
	if err != nil {
		return err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "文件关闭失败",
			})
		}
	}(outFile)
	return png.Encode(outFile, overlay)
}
