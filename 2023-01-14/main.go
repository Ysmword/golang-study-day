package main

import "fmt"

func main() {
	m := make(map[string]int)

	// fmt.Println(&m["test"])   // 无法编译通过
	fmt.Println(m["test"]) // 打印0
}
