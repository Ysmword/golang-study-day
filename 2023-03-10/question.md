下面代码下划线处可以填入哪个选项？

```go
func main() {
	var s1 []int
	var s2 = []int{}
	if __ == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}
```

- A. s1
- B. s2
- C. s1、s2 都可以


答：C


nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合
