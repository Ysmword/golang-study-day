package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t)

	fmt.Println(a[1:2])
}
