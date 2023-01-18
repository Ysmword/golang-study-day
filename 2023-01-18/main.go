package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	Name string
}

func main() {
	s := new(student)
	fmt.Println(s) // 打印&{}
	fmt.Println(unsafe.Sizeof(s))

	slice := make([]string, 2)
	fmt.Println(slice)
	fmt.Println(unsafe.Sizeof(slice))

	m := make(map[string]string)
	fmt.Println(m)

	ch := make(chan bool)
	fmt.Println(ch)
}
