下面这段代码能否通过编译，不能的话原因是什么；如果通过，输出什么。

```go
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```

答：不能，因为这个是数组的指针
