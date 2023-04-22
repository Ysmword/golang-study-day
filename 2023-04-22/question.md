下面这段代码输出结果正确吗？
```go
type Foo struct {
	bar string
}
func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}
```
输出：
{A} {B} {C}
&{A} &{B} &{C}


答：不对，因为value只被创建一次，每次都是去同一个地址
