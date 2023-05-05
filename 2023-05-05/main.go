package main

import "fmt"

func sum(x, y int) (total int, err error) {
	total = 4
	return x + y, nil
}

func main() {
	fmt.Println(sum(1, 2))
}
