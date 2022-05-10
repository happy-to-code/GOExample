package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
)

func main() {

}

// 创建二维码
func createAvatar(content string, qrCodeSize int) (image.Image, error) {
	// fmt.Println("content:", content)
	if len(content) == 0 {
		return nil, nil
	}
	if qrCodeSize <= 0 {
		qrCodeSize = 250
	}

	bgImg, err := createQrCode(content, qrCodeSize)
	return bgImg, err
}

func createQrCode(content string, qrCodeSize int) (img image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, qrcode.Highest)

	if err != nil {
		return nil, fmt.Errorf("创建二维码失败")
	}
	qrCode.DisableBorder = true
	img = qrCode.Image(qrCodeSize)
	return img, nil
}
