package main

import "fmt"

const a = 10

var b = 100

func init() {
	fmt.Println("init", a)
	fmt.Println("init", b)
}

func main() {
	fmt.Println("main")
}
