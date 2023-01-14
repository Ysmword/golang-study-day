package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	// val 为固定值
	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
