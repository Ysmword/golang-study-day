package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	fmt.Println(len(a), cap(a))

	b := s[:2]
	fmt.Println(len(b), cap(b))

	c := s[1:2:cap(s)]
	fmt.Println(len(c), cap(c))
}
