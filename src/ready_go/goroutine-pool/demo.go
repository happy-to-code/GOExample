package main

import (
	"fmt"
	"math/rand"
)

type Job struct { // 任务
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct { // 结果
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func main() {
	// 需要2个管道
	// 任务管道
	jobChan := make(chan *Job, 128)
	// 结果管道
	resultChan := make(chan *Result, 128)

	// 工作池（goroutine池）
	createPool(64, jobChan, resultChan) // 工作池中有64个Goroutine在工作

	// 负责打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道，进行打印
		for result := range resultChan { // 从通道resultChan接收值
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan) // 给函数传入参数，立即执行

	// 主协程
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 函数createPool：创建工作池
// 工作池里的Goroutine负责计算从jobChan取数字，然后计算各位之和，再输出到resultChan
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 开 num 个协程，做计算工作
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum // 读取随机数
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				// 运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
