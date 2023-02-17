下面这段代码输出什么？
```go
func hello(num ...int) {  
    num[0] = 18
}

func main() {  
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```
A.18

B.5

C.Compilation error

答：参数为数组的引用
