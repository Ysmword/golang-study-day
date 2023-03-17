package main

import "fmt"

func main() {
	var x int32 = 32.0
	// var y int = x // 不能将int32类型的变量赋给int类型的变量
	var y int = int(x)
	var z rune = x
	fmt.Println(x, y, z)
}
