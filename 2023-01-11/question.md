下面两段代码输出什么。
```go
// 1.
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}

// 2.
func main() {
	s := make([]int,0)
	s = append(s,1,2,3,4)
	fmt.Println(s)
}
```