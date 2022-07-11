package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// fileName := "E:\\20.06.16Project\\GOExample\\src\\AA-20220526\\createFile\\11\\test.txt"
	fileName := "E:\\20.06.16Project\\GOExample\\src\\AA-20220526\\createFile\\11\\"
	PathExists(fileName)

	err := ioutil.WriteFile(fileName+"test.txt", []byte("hello world"), 0600)
	if err != nil {
		panic(err)
	}
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {

		} else {
			return true, nil
		}
	}
	return false, err
}
