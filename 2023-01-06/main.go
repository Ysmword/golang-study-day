package main

import "fmt"

func main() {
	test := "test"
	switch test {
	case "test":
		fmt.Println("test")
		fallthrough // 会继续执行紧跟下一个case
	case "test1":
		fmt.Println("test1")
	case "test2":
		fmt.Println("test2")
	default:
		fmt.Println("default")
	}
}
