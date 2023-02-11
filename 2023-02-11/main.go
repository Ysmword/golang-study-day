package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello // 奇特，第一次见这种用法
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	h()
}
