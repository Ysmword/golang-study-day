package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println("1", person.age) //28

	// 2.
	defer func(p *Person) {
		fmt.Println("2", p.age) //28
	}(person)

	// 3.
	defer func() {
		fmt.Println("3", person.age) // 29
	}()

	person = &Person{29}
}
