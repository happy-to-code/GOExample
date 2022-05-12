package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {

	fBg, err := os.Open("bkground.jpg")
	defer fBg.Close()
	bg, _, err := image.Decode(fBg)

	fSrc, err := os.Open("arrow1.jpg")
	defer fSrc.Close()
	src, _, err := image.Decode(fSrc)

	fMaskImg, err := os.Open("mask.jpg")
	defer fMaskImg.Close()
	maskImg, _, err := image.Decode(fMaskImg)

	bounds := src.Bounds() // you have defined that both src and mask are same size, and maskImg is a grayscale of the src image. So we'll use that common size.
	mask := image.NewAlpha(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			// get one of r, g, b on the mask image ...
			r, _, _, _ := maskImg.At(x, y).RGBA()
			// ... and set it as the alpha value on the mask.
			mask.SetAlpha(x, y, color.Alpha{uint8(255 - r)}) // Assuming that white is your transparency, subtract it from 255
		}
	}

	m := image.NewRGBA(bounds)
	draw.Draw(m, m.Bounds(), bg, image.ZP, draw.Src)

	draw.DrawMask(m, bounds, src, image.ZP, mask, image.ZP, draw.Over)

	toimg, _ := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\imageBounds3\\new.jpeg")
	defer toimg.Close()

	err = jpeg.Encode(toimg, m, nil)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
