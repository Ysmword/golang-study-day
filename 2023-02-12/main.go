package main

import "fmt"

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	fmt.Println(i)
	// switch i.(type) { // i 不是interface类型,不能编译通过
	// case int:
	// 	println("int")
	// case string:
	// 	println("string")
	// case interface{}:
	// 	println("interface")
	// default:
	// 	println("unknown")
	// }
}
