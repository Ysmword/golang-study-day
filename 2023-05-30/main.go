package main

import "fmt"

type People interface {
	Show()
}

type Student struct {
}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAA")
	} else {
		fmt.Println("BBBBB")
	}
}
