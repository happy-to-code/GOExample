package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
)

func asyncLog(reader io.ReadCloser) (string, error) {
	// cache := "" // 缓存不足一行的日志信息
	// buf := make([]byte, 1024)
	//
	// var errStrList []string
	// for {
	// 	num, err := reader.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		return []string{}, err
	// 	}
	// 	if num > 0 {
	// 		b := buf[:num]
	// 		s := strings.Split(string(b), "\n")
	// 		line := strings.Join(s[:len(s)-1], "\n") // 取出整行的日志
	// 		cache = s[len(s)-1]
	// 		fmt.Printf(">>>>>>>>>%s%s\n", cache, line)
	// 		errStrList = append(errStrList, fmt.Sprintf("%s%s\n", cache, line))
	// 	}
	// }
	// fmt.Println("----------------->", errStrList)
	// return errStrList, nil

	bytes, err := ioutil.ReadAll(reader)
	fmt.Println("-------------------->", string(bytes))
	return string(bytes), err
}

func execute() error {
	cmd := exec.Command("sh", "-c", "E:/20.06.16Project/GOExample/src/AA-20220526/shell_call/scripts/curl.sh")

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
		return err
	}

	return nil
}

func main() {
	execute()
}
