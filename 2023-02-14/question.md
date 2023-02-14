以下代码有什么问题？

```go
package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
```

答：输出结果不唯一，代码存在风险，所有go语句未必都能执行到
这是使用WaitGroup经常放下的错误！多次运行就会发现输出结果都会不同，甚至会发生报错的问题。这是因为go执行太快了，
导致wg.Add(1)还没有执行main函数就执行完毕了。wg.Add的位置放错了。真确的写法：
```go
package main

import (
	"sync"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
        wg.Add(1)
		go func(i int) {
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
```