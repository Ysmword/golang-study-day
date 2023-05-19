下面这段代码能否通过编译，如果可以，输出什么？


```go
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

答：s1 = append(s1,s2),append函数的第二个参数不能传递数组或者切片，需要使用...操作符