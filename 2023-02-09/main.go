package main

func question1() {
	var a int8 = -1
	var b int8 = -128 / a

	println(b)
}

func question2() {
	const a int8 = -1
	// var b int8 = -128 / a // 超范围了
	// println(b)
}

func main() {
	question1()
	question2()
}
