下面这段代码能否通过编译，不能的话原因是什么；如果通过，输出什么。

```go
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```

答案：不能，new函数返回指针类型的数组，而append函数接收非指针类型的数组


