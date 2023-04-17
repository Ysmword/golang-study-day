下面这段代码输出什么？为什么？

```go
func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}
```

答：随机打印值，可能一直打印2 3

解析：
for range 使用变量声明(:=)的形式迭代变量，需要注意的是，变量i、v在每次循环体中都会被重用,而不是重新声明。
各个goruntine中输出的i,v值都是for range循环结束后的i,v最终值,而不是各个goruntine启动时的i,v值。可以理解为闭包
引用,使用的是上下文环境的值。