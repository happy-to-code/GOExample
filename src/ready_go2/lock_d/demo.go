package main

import (
	"fmt"
	"sync"
)

// 因为mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型。
// 同时它们也没有其它的同步事件可以参考，这两个事件不可排序也就是可以并发的。
// 因为可能是并发的事件，所以main函数中的mu.Unlock()很有可能先发生，
// 而这个时刻mu互斥对象还处于未加锁的状态，从而会导致运行时异常。
func main1() {
	var mu sync.Mutex

	go func() {
		fmt.Println("你好, 世界")
		mu.Lock()
	}()

	mu.Unlock()
}

/*
修复的方式是在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，
解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，此时后台线程和主线程再没有其它的同步事件参考，
它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
*/
func main() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock()
}
