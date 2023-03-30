下面的代码有几处语法问题，各是什么？

```go
package main
import (
    "fmt"
)
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}
```

答：nil不能赋给字符串类型，判断字符串是否为空，也不能使用nil来判断，而是使用空字符串判断