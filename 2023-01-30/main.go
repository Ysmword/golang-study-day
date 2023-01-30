package main

import "fmt"

type student struct {
	name string
}

func main() {
	p := &student{}
	fmt.Println(p.name)
	// fmt.Println((&p).name)  p已经是指针了，在&表示取的是指针的地址了
	fmt.Println((*p).name)
	// fmt.Println(p->name)  不允许这样访问
}
