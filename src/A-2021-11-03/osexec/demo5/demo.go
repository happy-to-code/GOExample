package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// 使用 os 库的 Setenv 函数来设置的环境变量，是作用于整个进程的生命周期的。

func main() {
	os.Setenv("NAME", "wangbm")
	cmd := exec.Command("echo", os.ExpandEnv("$NAME"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}
