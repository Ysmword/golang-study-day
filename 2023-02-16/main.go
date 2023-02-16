package main

import "fmt"

type Person struct {
	name string
}

func main() {
	var m map[Person]int
	p := Person{name: "mike"}
	fmt.Println(m[p])

	m = make(map[Person]int)
	m[p] = 9
	fmt.Println(m[p])
}
