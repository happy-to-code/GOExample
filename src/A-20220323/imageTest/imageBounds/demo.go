package main

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	f1, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\imageBounds\\203.png")
	if err != nil {
		panic(err)
	}
	f3, err := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\imageBounds\\4.png")
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
	// ImageTypeToRGBA64(&backGroundFileImg)
	// transparentImg = OpacityAdjust(transparentImg, 0)
	// draw.Draw(transparentImg, transparentImg.Bounds(), &image.Uniform{C: color.RGBA{R: 0, G: 0, B: 0, A: 255}}, image.ZP, draw.Over)
	// for x := 0; x < transparentImg.Bounds().Dx(); x++ {
	// 	for y := 0; y < transparentImg.Bounds().Dy(); y++ {
	// 		transparentImg.Set(x, y, image.NewUniform(color.Alpha{0}))
	// 	}
	// }

	// ==========================================

	// 新建一个蒙版
	// jpg := image.NewRGBA(backGroundFileImg.Bounds())
	// fmt.Printf(" jpg.Bounds():%+v\n", transparentImg.Bounds())
	// draw.Draw(jpg, jpg.Bounds(), backGroundFileImg, backGroundFileImg.Bounds().Min, draw.Over)

	// fmt.Printf(" transparentImg.Bounds():%+v\n", transparentImg.Bounds())
	fmt.Println("--->", transparentImg.Bounds().Dx())
	fmt.Println("--->", transparentImg.Bounds().Dy())
	draw.Draw(transparentImg,
		image.Rect(50, 50, transparentImg.Bounds().Dx(), transparentImg.Bounds().Dy()),
		backGroundFileImg,
		image.Point{},
		draw.Over)

	width := transparentImg.Bounds().Dx()
	height := transparentImg.Bounds().Dy()

	gc := draw2dimg.NewGraphicContext(transparentImg)
	gc.SetStrokeColor(color.RGBA{0, 255, 0, 0xff})
	gc.SetFillColor(color.RGBA{0, 0, 0, 0})
	gc.SetLineWidth(5)
	gc.BeginPath()
	gc.MoveTo(0, 0)
	gc.LineTo(float64(width), 0)
	gc.LineTo(float64(width), float64(height))
	gc.LineTo(0, float64(height))
	gc.LineTo(0, 0)
	gc.Close()
	gc.FillStroke()
	// 写入新文件
	err = png.Encode(f3, transparentImg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")
}

// OpacityAdjust 将输入图像m的透明度变为原来的倍数。若原来为完成全不透明，则percentage = 0.5将变为半透明
func OpacityAdjust(m *image.RGBA, percentage float64) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			opacity := uint16(float64(a) * percentage)
			// 颜色模型转换，至关重要！
			v := newRgba.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity})
			// Alpha = 0: Full transparent
			rr, gg, bb, aa := v.RGBA()
			newRgba.SetRGBA64(i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
		}
	}
	return newRgba
}
func ImageTypeToRGBA64(m *image.Image) *image.RGBA64 {
	bounds := (*m).Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA64(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := (*m).At(i, j)
			r, g, b, a := colorRgb.RGBA()
			nR := uint16(r)
			nG := uint16(g)
			nB := uint16(b)
			alpha := uint16(a)
			newRgba.SetRGBA64(i, j, color.RGBA64{R: nR, G: nG, B: nB, A: alpha})
		}
	}
	return newRgba
}
