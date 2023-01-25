package main

const cl = 100

var bl = 123

func main() {
	println(&bl, bl)
	// println(&cl, cl)   常量不能取址
}
