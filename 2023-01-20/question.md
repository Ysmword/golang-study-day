下面这段代码能否通过编译，如果可以，输出什么？

```go
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

答：不能，因为append第二参数不能直接接收slice，需要使用...操作符