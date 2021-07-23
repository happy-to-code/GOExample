package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	var body = `
		{
			"keysss":0000
		}
`
	// 当前时间字符串
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02-15-04-05")
	var fileName string
	sysType := runtime.GOOS
	if sysType == "windows" {
		fileName = "datafile/" + timeStr + "-" + "resHash"
	} else {
		fileName = "./datafile/" + timeStr + "-" + "resHash"
	}
	dstFile, err := os.Create(fileName)
	if err != nil {
		// return "", resHash, fmt.Errorf("创建文件出错:%v\n", err.Error())
		fmt.Printf("创建文件出错:%v\n", err.Error())
	}
	defer dstFile.Close()
	dstFile.Write([]byte(body))
}
