对 add() 函数调用正确的是？
```go
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
```
A. add(1, 2)

B. add(1, 3, 7)

C. add([]int{1, 2})

D. add([]int{1, 3, 7}...)

答：ABD
add是可变函数，参数可以接收任意长度的int类型的数字
