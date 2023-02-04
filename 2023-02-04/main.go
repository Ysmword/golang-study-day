package main

import "fmt"

const (
	x = iota
	_ // 1
	y // 2
	z = "zz"
	k        // zz
	p = iota // 5
)

const (
	c = iota // 遇到新的const会被重置值为0
)

func main() {
	fmt.Println(x, y, z, k, p, c)
}
