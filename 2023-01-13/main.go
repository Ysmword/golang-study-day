package main

import "fmt"

func sum(x, y int) (total int, err error) {
	return x + y, nil
}

func main() {
	fmt.Println(sum(1, 2))
}
