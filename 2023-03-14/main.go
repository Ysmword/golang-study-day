package main

import "fmt"

func main() {
	// var m map[string]int
	m := make(map[string]int)
	m["a"] = 1
	// if v := m["b"]; v != nil {
	// 	fmt.Println(v)
	// }

	if v := m["b"]; v != 0 {
		fmt.Println(v)
	}
}
