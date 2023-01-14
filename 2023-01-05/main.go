package main

import "fmt"

type MyInt int

func main() {
	var i int = 1
	var j MyInt = MyInt(i)
	fmt.Println(j)
}
