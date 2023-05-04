package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["a"] = "a"
	a["b"] = "b"
	a["c"] = "c"
	for key, value := range a {
		fmt.Println(key, value)
		delete(a, key)
	}
}
