package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	imgPath := "E:\\20.06.16Project\\GOExample\\src\\A-20220323\\imageTest\\readFileByUrl2\\"
	imgUrl := "http://qiniu.yueda.vip/0000.jpg"

	fileName := path.Base(imgUrl)

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)

}
