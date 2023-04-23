下面代码里的 counter 的输出值？
```go
func main() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
```
A. 2

B. 3

C. 2 或 3

答：C
如果一开始A没有被获取，后面会被删除，m的长度只能为2