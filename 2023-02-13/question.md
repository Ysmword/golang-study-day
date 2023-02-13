执行下面的代码会发生什么？

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
```

答：过早关闭channel,又因为休眠期间，有协程会向关闭的channel发送数据，从而触发panic

详细：
1、给一个nil channel 发送数据，造成永久阻塞
2、从一个nil channel 接收数据，造成永久阻塞
3、给一个已关闭的channel发送数据，引起panic
4、从一个已关闭的channel接收数据，如果缓存区为空，则返回一个零值
5、无缓存的channel是同步的，而有缓存的channel是非同步的