package main

import "fmt"

func f() (int, int) {
	return 0, 1
}

func main() {
	var x int
	// x, _ := f()
	x, _ = f()
	// x, y = f()
	x, y := f()
	fmt.Println(x, y)
}
