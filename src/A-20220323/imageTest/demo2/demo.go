package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo2\\dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo2\\20.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	jpg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) // 截取图片的一部分
	jpeg.Encode(file, jpg, nil)

}
