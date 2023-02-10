package main

import "log"

func init() {
	log.Println("this is a test")
}

func init() {
	log.Println("this is a test1")
}

func init() {
	log.Println("this is a test2")
}

// 在一个包里面，可以有多个init函数，执行顺序没有明确定义
func main() {
	log.Println("this is main")
}
