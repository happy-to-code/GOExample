package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// PathExists 判断目录是否存在
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
func init() {
	var _dir string
	sysType := runtime.GOOS
	if sysType == "windows" {
		// windows系统
		_dir = "conf1"
	} else {
		_dir = "./conf1"
	}

	exist, err := PathExists(_dir)
	if err != nil {
		log.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		log.Printf("has dir![%v]\n", _dir)
	} else {
		log.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
		} else {
			log.Printf("mkdir success!\n")
		}
	}

}

func writeResult(fileName, data string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理
		fmt.Println("########", err)
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)
		data = data + "\n"

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(data), n)
	}

	defer f.Close()
	return nil
}

func main() {
	timeStr := time.Unix(time.Now().Unix(), 0).Format("20060102150405")
	var fileName = "E:\\20.06.16Project\\GOExample\\src\\A-2021-08-11\\readFile\\conf\\" + timeStr + "-hasFix.txt"
	f, err2 := os.Create(fileName)
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	fi, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-2021-08-11\\readFile\\conf\\fixData.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		sLine := string(a)
		fmt.Println(sLine)
		split := strings.Split(sLine, ",")
		fmt.Println("--->", split)

		writeResult(fileName, sLine)
	}
}
