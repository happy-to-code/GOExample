package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	outFile, err := os.Create("gopher2.png")
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		panic(err)
	}
	err = b.Flush()
	if err != nil {
		panic(err)
	}
}
