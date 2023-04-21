下面这段代码输出什么？

```go
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // 1 12 13 4 5
	fmt.Println("a = ", a) // 1 12 13 4 5
}
```
答：在循环过程a的值被修改了