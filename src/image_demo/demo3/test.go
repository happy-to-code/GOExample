package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"os"
)

func main() {
	f, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\image_demo\\222.png")
	if err != nil {
		panic(err)
	}
	gopherImg, _, err := image.Decode(f) // 打开图片

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	for x := 0; x < img.Bounds().Dx(); x++ { // 将背景图涂黑
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, color.Black)
		}
	}
	draw.Draw(img, img.Bounds(), gopherImg, image.Pt(0, 0), draw.Over) // 将gopherImg绘制到背景图上
}
