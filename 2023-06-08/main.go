package main

import "fmt"

type person struct {
	name string
}

func main() {
	var m map[person]int
	p := person{"John"}
	fmt.Println(m[p])
}
