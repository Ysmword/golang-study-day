package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
}
