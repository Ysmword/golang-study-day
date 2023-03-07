package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 3, 7))
	fmt.Println(add([]int{1, 3, 7}...))
}
