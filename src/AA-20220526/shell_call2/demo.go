package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error) {

	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("sh", "-c", s)

	// 读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), err
}
func main() {
	s := "E:/20.06.16Project/GOExample/src/AA-20220526/shell_call/scripts/curl.sh"
	shell, err := exec_shell(s)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("back:", shell)
}
