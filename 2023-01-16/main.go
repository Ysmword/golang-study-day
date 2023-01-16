package main

import (
	"fmt"
	"unsafe"
)

// 信号传输
func useChan() {
	a := make(chan struct{}, 1)
	go func() {
		a <- struct{}{}
	}()

	<-a
	fmt.Println("test")
}

// 集合
type set struct {
	data map[string]struct{}
	len  int
}

// 只声明方法的结构体
type empty struct{}

func (e empty) Name() {
	fmt.Println("empty")
}

func main() {
	e := empty{}
	fmt.Println(unsafe.Sizeof(e))
	e.Name()
	useChan()
}
