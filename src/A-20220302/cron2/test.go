package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds()) // 精确到秒

	var a int
	// 定时任务
	spec := "*/1 * * * * ?" // cron表达式，每秒一次
	c.AddFunc(spec, func() {
		if a == 10 {
			c.Stop()
		}
		fmt.Printf("[%d]11111\n", a)
		a++

	})

	c.Start()
	select {} // 阻塞主线程停止
}
