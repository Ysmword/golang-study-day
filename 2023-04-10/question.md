下面这段代码输出什么？

```go
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
}
```

答：South
根据iota的用法推断South的值是2;另外，如果定义了String()方法,当使用fmt.Printf()、fmt.Print()
和fmt.Println()会自动使用String()方法,实现字符串打印

