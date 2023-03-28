package main

import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}
}

func main() {
	p := f() // 或者 *f()
	fmt.Println(p.m)
}
