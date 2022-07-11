package main

import (
	"bytes"
	"fmt"
	"path"

	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://horifon.oss-cn-shanghai.aliyuncs.com/20220521/50eaa48b14634f149af80e70fc076341.JPG"
	ext := path.Ext(url)
	fmt.Println(ext)
	base := path.Base(url)
	fmt.Println(base)
	// _, err := getDataByUrl(url)
	// if err != nil {
	// 	panic(err)
	// }
}

func getDataByUrl(url string) (img image.Image, err error) {
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("[%s]通过url获取数据失败,err:%s", url, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// 读取获取的[]byte数据
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("读取数据失败,err:%s", err.Error())
		return
	}

	if !strings.HasSuffix(url, ".jpg") &&
		!strings.HasSuffix(url, ".jpeg") &&
		!strings.HasSuffix(url, ".JPG") &&
		!strings.HasSuffix(url, ".png") {
		err = fmt.Errorf("[%s]不支持的图片类型,暂只支持.jpg、.png文件类型", url)
		return
	}

	// []byte 转 io.Reader
	reader := bytes.NewReader(data)
	if strings.HasSuffix(url, ".jpg") || strings.HasSuffix(url, ".jpeg") {
		// 此处jgeg.decode 有坑 会报 invalid JPEG format: missing SOI marker 错误
		// 所以当报错时再用 png.decode 试试
		img, err = jpeg.Decode(reader)
		if err != nil {
			fmt.Printf("jpeg.Decode err:%s", err.Error())
			reader2 := bytes.NewReader(data)
			img, err = png.Decode(reader2)
			if err != nil {
				err = fmt.Errorf("===>png.Decode err:%s", err.Error())
				return
			}
		}
	}

	if strings.HasSuffix(url, ".png") {
		img, err = png.Decode(reader)
		if err != nil {
			err = fmt.Errorf("png.Decode err:%s", err.Error())
			return
		}
	}

	return
}
