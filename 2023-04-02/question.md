下面选项正确的是？
```go
func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}
```
A. 1 2

B. compilation error


答： a和b是if中声明的变量,因此可以在else中打印出来
A
