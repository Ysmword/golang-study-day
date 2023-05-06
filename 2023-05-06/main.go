package main

import "fmt"

func main() {
	m := make(map[string]int)
	// fmt.Println(&m["q"])  // 无法获取地址
	fmt.Println(m)
}
