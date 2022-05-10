package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "http://qiniu.yueda.vip/0000.jpg"
	data := ReadTxtData(url)
	for k, v := range data {
		fmt.Println(k, "  ", v)
	}
}

func ReadTxtData(filePath string) map[string]interface{} {
	resp, err := http.Get(filePath)
	defer resp.Body.Close()
	reader := bufio.NewReaderSize(resp.Body, 1024*32)

	if err != nil {
		return nil
	}
	hashMapData := map[string]interface{}{}
	for {
		b, errR := reader.ReadBytes('\n') // 按照行读取，遇到\n结束读取
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		lineData := strings.TrimSuffix(strings.TrimSuffix(string(b), "\n"), "\r")
		if len(lineData) > 0 {
			hashMapData[lineData] = "1"
		}
	}

	return hashMapData
}
