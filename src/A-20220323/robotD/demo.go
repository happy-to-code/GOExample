package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 0; i < 30; i++ {

		exec.Command(`cmd`, `/c`, `start`, `https://www.jszzb.gov.cn/col22/81608.html`).Start()
	}
	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	// robotgo.MoveMouse(90, 50)
	// robotgo.MoveMouseSmooth(90, 50)

	// 按下鼠标左键
	// 第1个参数：left(左键) / center(中键，即：滚轮) / right(右键)
	// for i := 0; i < 5000; i++ {
	// 	time.Sleep(time.Millisecond * 300)
	// 	robotgo.MouseClick(`left`, false)
	// }

	fmt.Println("===========================")
}
