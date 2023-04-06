package main

import "fmt"

type myInt int

func (i myInt) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i myInt
	i.PrintInt()
}
