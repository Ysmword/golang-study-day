下面这段代码正确的输出是什么？
```go
func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	f()
	fmt.Println("M")
}
```
A. F M D

B. D F M

C. F D M

答：C

解析：被调用函数里的defer语句在返回之前就会被执行。
