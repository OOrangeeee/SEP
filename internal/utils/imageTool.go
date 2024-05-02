package utils

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type ImageTool struct {
}

func (it *ImageTool) ChangeColorsAndOverlay(firstImagePath, secondImagePath string) error {
	// 打开第一张图片
	firstFile, err := os.Open(firstImagePath)
	if err != nil {
		return err
	}
	defer firstFile.Close()

	// 解码第一张图片
	firstImg, _, err := image.Decode(firstFile)
	if err != nil {
		return err
	}
	firstBounds := firstImg.Bounds()

	// 打开第二张图片
	secondFile, err := os.Open(secondImagePath)
	if err != nil {
		return err
	}
	defer secondFile.Close()

	// 解码第二张图片
	secondImg, _, err := image.Decode(secondFile)
	if err != nil {
		return err
	}

	// 使用 nfnt/resize 调整第二张图片的大小以匹配第一张图片
	secondImgResized := resize.Resize(uint(firstBounds.Dx()), uint(firstBounds.Dy()), secondImg, resize.NearestNeighbor)

	// 创建一个新的同样大小的图像
	processedImg := image.NewRGBA(firstBounds)

	// 处理调整大小后的第二张图片的像素
	for y := firstBounds.Min.Y; y < firstBounds.Max.Y; y++ {
		for x := firstBounds.Min.X; x < firstBounds.Max.X; x++ {
			pixel := secondImgResized.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			// 黑色变为透明
			if originalColor.R == 0 && originalColor.G == 0 && originalColor.B == 0 {
				processedImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			} else if originalColor.R == 255 && originalColor.G == 255 && originalColor.B == 255 { // 白色变为绿色
				processedImg.Set(x, y, color.RGBA{0, 255, 0, 255})
			} else {
				processedImg.Set(x, y, originalColor)
			}
		}
	}

	// 创建一个新的画布并将第一张图片作为背景
	overlay := image.NewRGBA(firstBounds)
	draw.Draw(overlay, firstBounds, firstImg, image.Point{}, draw.Src)
	// 将处理后的第二张图片覆盖到第一张图片上
	draw.Draw(overlay, firstBounds, processedImg, image.Point{}, draw.Over)

	// 保存到第二张图片的位置，覆盖原文件
	outFile, err := os.Create(secondImagePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, overlay)
}
