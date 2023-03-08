以下代码输出什么？
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    go fmt.Println(<-ch1)
    ch1 <- 5
    time.Sleep(1 * time.Second)
}
```
A: 5

B: 不能编译

C: 运行时死锁

答：C

因此这道题目 go fmt.Println(<-ch1) 语句中的 <-ch1 是在 main goroutine 中求值的。这相当于一个无缓冲的 chan，发送和接收操作都在一个 goroutine 中（main goroutine）进行，因此造成死锁。

go 语句后面的函数调用，其参数会先求值，这和普通的函数调用求值一样。