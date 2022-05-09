package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func main() {
	f, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\image_demo\\222.png")
	if err != nil {
		panic(err)
	}
	img, formatName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(formatName)
	fmt.Println(img.Bounds())
	fmt.Println(img.ColorModel())
}
