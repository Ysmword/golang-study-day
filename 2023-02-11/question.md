下面这段代码输出什么以及原因？

```go
func hello() []string {  
    return nil
}

func main() {  
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}
```
A. nil

B. not nil

C. compilation error


答：B，是将函数 hello 赋值给变量 h，而不是函数的返回值（即不是进行函数调用），所以输出 not nil。