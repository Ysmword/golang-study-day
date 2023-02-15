package main

import "fmt"

func test() {
	var ch chan int
	<-ch // 这么写会发生死锁
}

func test1() {
	ch := make(chan int)
	<-ch
}

func test2() {
	ch := make(chan int, 2)
	ch <- 2
	num, ok := <-ch
	if !ok {
		fmt.Println("chan 关闭")
		return
	}
	fmt.Println("num:", num)
}

func main() {
	test2()
}
