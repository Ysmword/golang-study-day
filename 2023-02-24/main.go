package main

import "fmt"

func main() {
	s := make([]int, 2)
	fmt.Println("s", len(s))
	fmt.Println("s", cap(s))

	a := [2]int{1, 2}
	fmt.Println("a", len(a))
	fmt.Println("a", cap(a))

	m := make(map[int]int)
	m[1] = 1
	fmt.Println("m", len(m))
	// fmt.Println(cap(m)) 不能使用

	c := make(chan int, 2)
	c <- 2
	c <- 3
	fmt.Println("c", len(c)) // 实际长度跟make里面的参数无关,跟chan有多少个元素有关
	fmt.Println("c", cap(c))
}
