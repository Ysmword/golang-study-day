下面这段代码输出什么?

```go
func hello(i int) {  
    fmt.Println(i)
}
func main() {  
    i := 5
    defer hello(i)
    i = i + 10
}
```

答：打印5

解析：hello函数的参数在执行defer语句时候就会保存一份副本,在实际调用hello()函数时用