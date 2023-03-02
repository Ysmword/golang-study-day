这道题看起来很迷惑，目测很多人可能会答错。以下代码输出什么？（这是来自 《Go爱好者周刊》的一道题目，正确率才 12%）
```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
)

func main() {
    t := struct {
        time.Time
        N int
    }{
        time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
        5,
    }

    m, _ := json.Marshal(t)
    fmt.Printf("%s", m)
}
```
A：{"Time": "2020-12-20T00:00:00Z", "N": 5 }；

B："2020-12-20T00:00:00Z"；

C：{"N": 5}；

D：< nil >

答案：B

time.Time为内嵌，可以理解为继承,t的方法集包含了time.Time的方法集。

json.Marshal的执行顺序为:
如果值实现了 json.Marshaler 接口并且不是 nil 指针，则 Marshal 函数会调用其 MarshalJSON 方法以生成 JSON。如果不存在 MarshalJSON 方法，但该值实现 encoding.TextMarshaler 接口，则 Marshal 函数调用其 MarshalText 方法并将结果编码为 JSON 字符串。



