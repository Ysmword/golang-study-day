下面这段代码输出什么？
```go
func main() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
```
A. compilation error

B. equal

C. not equal

答：A,数组类型不一样,不能比较,数组类型组成部分包括数据类型和数组长度


