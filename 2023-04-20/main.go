package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice) // 1,2,0,0,0
	change(slice[0:2]...)
	fmt.Println(slice) // 1,2,3,0,0,
}
