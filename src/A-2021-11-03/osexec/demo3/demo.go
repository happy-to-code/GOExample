package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// 第三种：执行命令，并区分stdout 和 stderr
// 上面的写法，无法实现区分标准输出和标准错误，只要换成下面种写法，就可以实现。
func main() {
	cmd := exec.Command("ls", "-l", "/var/log/*.log")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
