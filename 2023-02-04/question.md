下面这段代码能否编译通过？如果可以，输出什么？

```go
const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
```


答：能被编译，并且打印结果为：0 2 zz zz 5
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可以理解为const语句块中的索引)

延申：
```go
const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

const(
    c = iota  // c为0
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
```
