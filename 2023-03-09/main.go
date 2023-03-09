package main

import "fmt"

func slice1() {
	a := []int{2: 1}
	fmt.Println(a)
}

func slice2() {
	a := []int{4: 44, 55, 66, 1: 77, 88}
	// 键值对中key表示索引，value表示值,比如4:44表示在索引为4的时候，其值为44
	// a: 0 77 88 0 44 55 66
	fmt.Println(len(a), a[2], a)
}

func main() {
	slice1()
	slice2()
}
