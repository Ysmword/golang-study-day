package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 反射
	a := map[string]int{"1": 1, "2": 2}
	b := map[string]int{"1": 1, "2": 2}
	if reflect.DeepEqual(a, b) {
		fmt.Println("两个map相等")
	} else {
		fmt.Println("两个map不相等")
	}

	// 遍历判断
}
