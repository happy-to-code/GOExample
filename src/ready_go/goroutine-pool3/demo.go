package main

import (
	"fmt"
	"time"
)

// panic: too many concurrent operations on a single file or socket (max 1048575)
func main() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	tick1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick1.C:
			for i := 0; i < 100000; i++ {
				go func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
					fmt.Println("xxx")
					time.Sleep(10 * time.Second)
				}()
			}
		}
	}
	time.Sleep(2 * time.Hour)
}
