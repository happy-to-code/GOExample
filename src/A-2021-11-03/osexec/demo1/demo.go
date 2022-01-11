package main

import (
	"log"
	"os/exec"
)

// 直接调用 Cmd 对象的 Run 函数，返回的只有成功和失败，获取不到任何输出的结果。

func main() {
	cmd := exec.Command("ls", "-l", "/var/log/")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
