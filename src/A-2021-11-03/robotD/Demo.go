package main

import (
	"github.com/go-vgo/robotgo"
	"os/exec"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		cmd := exec.Command("cmd", "/c", "start https://so.csdn.net/so/search?q=yida%26yueda&t=&u=")
		// cmd := exec.Command("cmd", "/c", "start https://www.baidu.com")
		cmd.Run()
		time.Sleep(time.Second * 3)
		// robotgo.MoveClick(838, 137, "left", false)
		robotgo.MoveClick(367, 248, "left", false)
		time.Sleep(time.Second * 2)
		robotgo.MoveClick(503, 878, "left", true)
	}

}
