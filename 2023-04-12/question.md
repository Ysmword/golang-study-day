以下代码的输出结果：
```go
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case foo <- <-bar:
		default:
			println("default")
		}
	}()
	wg.Wait()
}
```
A：default

B：panic

答：B,发生死锁

case 如果是channel,可以发送数据，也可以接受数据。

发送死锁的原因：case会被先运算,然后在从select中选值。然而bar无值,就从里面获取获取值,因此发生了死锁




