package main

import "fmt"

func defer_call() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

func main() {
	defer_call()
}
