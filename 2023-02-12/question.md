下面这段代码能否编译通过？如果可以，输出什么？

```go
func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
```

答：不能编译通过，因为i是int类型。只有接口类型才能使用类型断言
