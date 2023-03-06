package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	fmt.Println(incr(&p))
	fmt.Println(p)
}
