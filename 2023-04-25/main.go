package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Println(i, j)
		}
		fmt.Println("i", i)
	}
}
