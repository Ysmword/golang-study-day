package main

import "fmt"

func main() {
	var a int = 3
	var i interface{} = a
	fmt.Println(i)
}
