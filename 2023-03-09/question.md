一次两道题，因为相关的。

1、以下代码输出什么？
```go
package main

import (
    "fmt"
)

func main() {
    a := []int{2: 1}
    fmt.Println(a)
}
```
A：编译错误；B：[2 1]；C：[0 0 1]；D：[0 1]

答：C


2、以下代码输出什么？
```go
package main

func main() {
	var x = []int{4: 44, 55, 66, 1: 77, 88}
	println(len(x), x[2])
}
```
A：5 66；B：5 88；C：7 88；D：以上都不对

答案：C


详细解析：https://polarisxu.studygolang.com/posts/go/action/weekly-question-88/