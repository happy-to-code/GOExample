package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo4\\simhei.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 80, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

var text = []string{
	"地支：苏州相城",
	"姓名：张三",
	"电话：18890908888",
}

// 使用从服务器直接读取文件的方式
func main() {
	// 获取router路由对象
	r := gin.New()

	// 创建请求
	r.GET("/hello", createPic)

	r.Run(":9000") // 参数为空 默认监听8080端口

}

func createPic(c *gin.Context) {
	// detFile, err := os.Create("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo6\\dst.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer detFile.Close()

	// toPastePic, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo5\\qrcode.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer toPastePic.Close()
	// qrcodeImg, _ := png.Decode(toPastePic) // 对要粘贴的图片解码

	// ===========================================================================================
	imgUrl := "http://qiniu.yueda.vip/0000.jpg"

	// 获取远端图片
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// []byte 转 io.Reader
	reader := bytes.NewReader(data)
	qrcodeImg, _ := jpeg.Decode(reader)
	// ===========================================================================================

	// 重新调整二维码图片尺寸
	qrcodeImg = resize.Resize(314, 314, qrcodeImg, resize.Lanczos3)

	backGroundFile, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\demo6\\kong.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer backGroundFile.Close()
	backGroundFileImg, _ := jpeg.Decode(backGroundFile)

	// jpg := image.NewRGBA(image.Rect(0, 0, 300, 300))
	jpg := image.NewRGBA(backGroundFileImg.Bounds())

	draw.Draw(jpg, jpg.Bounds(), backGroundFileImg, backGroundFileImg.Bounds().Min, draw.Over)

	// 向画布中写入文字
	fontRender(jpg) // 首先将一个图片信息存入jpg

	// 粘贴二维码
	draw.Draw(jpg, qrcodeImg.Bounds().Add(image.Pt(60, 150)), qrcodeImg, qrcodeImg.Bounds().Min, draw.Src)  // 截取图片的一部分
	draw.Draw(jpg, qrcodeImg.Bounds().Add(image.Pt(435, 150)), qrcodeImg, qrcodeImg.Bounds().Min, draw.Src) // 截取图片的一部分
	// draw.Draw(jpg, qrcodeImg.Bounds().Add(image.Pt(60, 610)), qrcodeImg, qrcodeImg.Bounds().Min, draw.Src)  // 截取图片的一部分
	// draw.Draw(jpg, qrcodeImg.Bounds().Add(image.Pt(435, 610)), qrcodeImg, qrcodeImg.Bounds().Min, draw.Src) // 截取图片的一部分

	png.Encode(c.Writer, jpg)
	c.Writer.Header().Set("Conetnt-Type", "image/jpg")
	c.Writer.Flush()

	// jpeg.Encode(detFile, jpg, nil)

}

func fontRender(jpg *image.RGBA) {
	flag.Parse()
	fontBytes, err := ioutil.ReadFile(*fontfile) // 读取字体文件
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes) // 解析字体文件
	if err != nil {
		log.Println(err)
		return
	}

	black := image.Black
	// white := image.White

	// draw.Draw(jpg, jpg.Bounds(), white, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)       // 设置字体分辨率
	c.SetFont(f)         // 设置字体
	c.SetFontSize(*size) // 设置字体大小
	c.SetClip(jpg.Bounds())
	c.SetDst(jpg)
	c.SetSrc(black)

	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	fmt.Println("-------------------->", 10+int(c.PointToFixed(*size)>>6))
	// Draw the text.
	pt := freetype.Pt(900, 150+int(c.PointToFixed(*size)>>6))
	for _, s := range text {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}

}
