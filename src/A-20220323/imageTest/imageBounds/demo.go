package main

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	f1, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\imageBounds\\203.png")
	if err != nil {
		panic(err)
	}
	f3, err := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\imageBounds\\3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	// 解码
	backGroundFileImg, err := png.Decode(f1)
	if err != nil {
		panic(err)
	}

	// ==========================================
	// Transparent
	transparentImg := image.NewRGBA(image.Rect(0, 0, backGroundFileImg.Bounds().Dx()+100, backGroundFileImg.Bounds().Dy()+100))
	draw.Draw(transparentImg, transparentImg.Bounds(), image.Transparent, image.ZP, draw.Src)
	// for x := 0; x < transparentImg.Bounds().Dx(); x++ {
	// 	for y := 0; y < transparentImg.Bounds().Dy(); y++ {
	// 		transparentImg.Set(x, y, color.Transparent)
	// 	}
	// }

	// ==========================================

	// 新建一个蒙版
	// jpg := image.NewRGBA(backGroundFileImg.Bounds())
	// fmt.Printf(" jpg.Bounds():%+v\n", transparentImg.Bounds())
	// draw.Draw(jpg, jpg.Bounds(), backGroundFileImg, backGroundFileImg.Bounds().Min, draw.Over)
	//
	// fmt.Printf(" transparentImg.Bounds():%+v\n", transparentImg.Bounds())
	// fmt.Printf(" jpg.Bounds().Min:%+v\n", jpg.Bounds().Min)
	// draw.Draw(transparentImg, image.Rect(50, 50, backGroundFileImg.Bounds().Dx()-100, backGroundFileImg.Bounds().Dy()-100), jpg, jpg.Bounds().Min, draw.Over)

	// width := transparentImg.Bounds().Dx()
	// height := transparentImg.Bounds().Dy()
	// // for y := 0; y < height; y++ {
	// // 	for x := 0; x < width; x++ {
	// // 		jpg.Set(x, y, backGroundFileImg.At(x, y))
	// // 	}
	// // }
	//
	// // Rect(100, 100, 800, 800)
	// gc := draw2dimg.NewGraphicContext(jpg)
	// gc.SetStrokeColor(color.RGBA{0, 255, 0, 0xff})
	// gc.SetFillColor(color.RGBA{0, 0, 0, 0})
	// gc.SetLineWidth(5)
	// gc.BeginPath()
	// gc.MoveTo(0, 0)
	// gc.LineTo(float64(width), 0)
	// gc.LineTo(float64(width), float64(height))
	// gc.LineTo(0, float64(height))
	// gc.LineTo(0, 0)
	// gc.Close()
	// gc.FillStroke()
	// 写入新文件
	err = jpeg.Encode(f3, transparentImg, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")
}
