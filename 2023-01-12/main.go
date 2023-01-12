package main

import "fmt"

func main() {
	testData := make(map[string]int)
	testData["1"] = 1
	testData["2"] = 2
	testData["3"] = 3
	testData["4"] = 4

	for key, val := range testData {
		fmt.Println(key, val)
		delete(testData, key)
	}
}
