package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 加锁以及等待协程运行完成
func suc() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		wg.Add(1)
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			defer wg.Done()
			total += i
		}(i)
	}
	wg.Wait()
	fmt.Printf("suc total: %d, sum: %d\n", total, sum)
}

// 相比与加锁，使用原子操作会更好
func suc1() {
	var total int64
	sum := 0
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		sum += i
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	wg.Wait()
	fmt.Printf("suc1 total: %d, sum: %d\n", total, sum)
}

func main() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total: %d, sum: %d\n", total, sum)

	suc()

	suc1()
}
