package main

import "fmt"

func main() {
	// strA := 'abc' + '123'  // 非法
	strB := "abc" + "123"
	fmt.Println(strB)

	// strC := '123' + "abc" // 非法

	fmt.Println(fmt.Sprintf("abc%d", 123))

}
