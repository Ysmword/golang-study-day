下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？

```go
type S struct {
	m string
}

func f() *S {
	return __  //A
}

func main() {
	p := __    //B
	fmt.Println(p.m) //print "foo"
}
```