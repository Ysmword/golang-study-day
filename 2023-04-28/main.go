package main

import "fmt"

func main() {
	a := 1
	switch a {
	case 1, 2:
		fmt.Println("1", a)
		fallthrough
	case 3, 4:
		fmt.Println("3", a)
	default:
		fmt.Println("test")
	}
}
