f1()、f2()、f3() 函数分别返回什么？

```go
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}


func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}


func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)  // r 是函数参数，不会影响外层r的值
	return 1
}
```


答：1 5 1

关键点：先执行defer在执行return