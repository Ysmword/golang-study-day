package main

import (
	"fmt"
	"time"
)

func suc() {
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5 // 这个也不能放在前面,会发生死锁的
	time.Sleep(1 * time.Second)
}

func main() {

	// 错误写法
	// ch1 := make(chan int)
	// ch1 <- 5
	// go fmt.Println(<-ch1) // <-ch1会先取值，取值的动作是在main的goroutine中

	suc() // 正确写法
}
