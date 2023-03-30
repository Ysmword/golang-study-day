package main

import "fmt"

func main() {
	var x string = ""
	if x == "" {
		x = "default"
	}
	fmt.Println(x)
}
