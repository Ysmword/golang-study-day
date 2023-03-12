下面这段代码输出什么？

```go
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
```

答：
13 和 23

接口。 一种类型实现多个接口,结构体work分别实现了接口A、B,所以变量a,b调用各自方法ShowA和ShowB