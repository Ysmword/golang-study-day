下面这段代码输出什么？

```go
func main() {
	m := map[int]string{0:"zero",1:"one"}
	for k,v := range m {
		fmt.Println(k,v)
	}
}
```

答：
乱序打印，k v值