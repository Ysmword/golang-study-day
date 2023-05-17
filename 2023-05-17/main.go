package main

import "fmt"

const cl = 100

var bl = 123

func main() {
	fmt.Println(&bl, bl)
	fmt.Println(&cl, cl)
}
