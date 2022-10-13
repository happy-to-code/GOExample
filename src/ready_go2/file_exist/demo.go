package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	var filePath = "E:/20.06.16Project/GOExample/src/ready_go2/extend_d2/main.go"
	exists, err := PathExists(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)

	if exists {
		filenameWithSuffix := path.Base(filePath)
		fmt.Println("filenameWithSuffix:", filenameWithSuffix)

		split := strings.Split(filenameWithSuffix, ".")
		newName := split[0] + "_back"
		newName = newName + "." + split[1]

		name := strings.ReplaceAll(filePath, filenameWithSuffix, newName)
		fmt.Println("name:", name)

		// // 获取文件后缀
		// fileSuffix := path.Ext(filenameWithSuffix)
		// fmt.Println("fileSuffix:", fileSuffix)
		// 	重命名文件
		err := os.Rename(filePath, name)
		if err != nil {
			panic(err)
		}

	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
