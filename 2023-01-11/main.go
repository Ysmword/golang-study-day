package main

import "fmt"

func test() {
	s := make([]int, 5) // 5，表示创建一个长度为5的切片
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func test1() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

func main() {
	test()
	test1()
}
