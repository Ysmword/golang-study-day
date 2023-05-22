package main

import "fmt"

type people struct {
	name string
}

func main() {
	p := &people{}
	fmt.Println(p.name)
	// fmt.Println((&p).name)  不行,&是取址符
	fmt.Println((*p).name)
}
