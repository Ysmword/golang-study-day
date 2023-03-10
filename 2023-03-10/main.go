package main

import "fmt"

func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("s1 yes is nil")
	}

	if s2 == nil {
		fmt.Println("s2 yes is nil")
	}
}
