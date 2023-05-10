package main

import "fmt"

func main() {
	a := new([]int)
	fmt.Println(a)

	b := new(chan int)
	fmt.Println(b)

	c := new(map[string]string)
	fmt.Println(c)

	d := make([]int, 0)
	d = append(d, 1)
	fmt.Println(d)

	e := make(chan int, 1)
	e <- 1
	fmt.Println(e)

	f := make(map[string]string)
	f["a"] = "a"
	fmt.Println(f)
}
