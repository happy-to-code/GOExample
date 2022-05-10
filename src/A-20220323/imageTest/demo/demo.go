package main

import (
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	qrcodeFile, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo\\qrcode.png")
	if err != nil {
		panic(err)
	}
	qrcodeImg, _, err := image.Decode(qrcodeFile)

	logoFile, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo\\logo.png")
	if err != nil {
		panic(err)
	}
	logoImg, _, err := image.Decode(logoFile)
	logoImg = resize.Resize(uint(100), uint(100), logoImg, resize.Lanczos3)

	outQrcodeImg := image.NewRGBA(qrcodeImg.Bounds())
	// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	draw.Draw(outQrcodeImg, qrcodeImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)

	offset := image.Pt((outQrcodeImg.Bounds().Max.X-logoImg.Bounds().Max.X)/2, (outQrcodeImg.Bounds().Max.Y-logoImg.Bounds().Max.Y)/2)
	draw.Draw(outQrcodeImg, qrcodeImg.Bounds().Add(offset), logoImg, image.Pt(0, 0), draw.Over)

	fileOut, _ := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo\\image.png")
	err = png.Encode(fileOut, outQrcodeImg)
	if err != nil {
		return
	}
}
