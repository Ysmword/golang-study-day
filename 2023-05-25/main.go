package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {
	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	// list["student"].Name = "LDB"  不能这样赋值

	fmt.Println(list["student"])
}
