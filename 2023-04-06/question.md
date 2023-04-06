下面这段代码输出什么？为什么？
```go
func (i int) PrintInt ()  {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}
```
A. 1

B. compilation error


答：基于类型创建的方法必须定义在同一个包内,上面的代码基于int类型创建了PrintInt方法,
由于int类型和方法PrintInt定义在不同一个的包内,所以会编译失败

解决办法：

```go
type Myint int

func (i Myint) PrintInt ()  {
	fmt.Println(i)
}

func main() {
	var i Myint = 1
	i.PrintInt()
}
```