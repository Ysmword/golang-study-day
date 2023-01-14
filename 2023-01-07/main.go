package main

import "fmt"

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

type Student struct {
	Name string
}

func (s Student) PrintlnName() {
	fmt.Println(s.Name)
}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)

	s := &Student{Name: "name"}
	s.PrintlnName()
}
