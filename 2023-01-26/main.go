package main

import "fmt"

func main() {
	list := new([]int)
	// list = append(list, 1)   // 不能，因为list是数组的指针
	fmt.Println(list)
}
