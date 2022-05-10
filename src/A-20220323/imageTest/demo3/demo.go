package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

// 一张图片贴到另一张图片上面
func main() {
	detFile, err := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo3\\dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer detFile.Close()

	toPasteFile, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo3\\20.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer toPasteFile.Close()
	toPasteFileImg, _ := jpeg.Decode(toPasteFile)

	backGroundFile, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo3\\10.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer backGroundFile.Close()
	backGroundFileImg, _ := jpeg.Decode(backGroundFile)

	// jpg := image.NewRGBA(image.Rect(0, 0, 300, 300))
	jpg := image.NewRGBA(backGroundFileImg.Bounds())

	// dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op
	draw.Draw(jpg, jpg.Bounds(), backGroundFileImg, backGroundFileImg.Bounds().Min, draw.Over)               // 首先将一个图片信息存入jpg
	draw.Draw(jpg, jpg.Bounds(), toPasteFileImg, toPasteFileImg.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over) // 将另外一张图片信息存入jpg

	// draw.DrawMask(jpg, jpg.Bounds(), toPasteFileImg, toPasteFileImg.Bounds().Min, backGroundFileImg, backGroundFileImg.Bounds().Min, draw.Src) // 利用这种方法不能够将两个图片直接合成？目前尚不知道原因。

	jpeg.Encode(detFile, jpg, nil)

}
