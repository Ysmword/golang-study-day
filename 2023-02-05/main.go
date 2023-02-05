package main

import "fmt"

func main() {
	// var x = nil //
	var b interface{} = nil
	fmt.Println(b)
	// var c string = nil // 不能将nil复制给字符串
	var e error = nil
	fmt.Println(e)
}
