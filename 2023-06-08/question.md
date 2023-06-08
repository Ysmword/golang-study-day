下面这段代码输出什么？
```go
type person struct {  
    name string
}

func main() {  
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
```
A.0

B.1

C.Compilation error


答： A; 获取不需要实例化
