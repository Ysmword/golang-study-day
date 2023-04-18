下面这段代码输出什么？

```go
func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
```

答：7

记住一点：defer在return之前执行


return是非原子性的,在底层是分两步执行的:
第一步：返回值 赋值
defer,如果有defer函数,在第二步之前会执行
第二步：正真的RET返回


defer于返回值提前定义的关系：

如果没有提前定义:
1、首先默认创建了一个临时零值变量(假设为s)作为返回值,然后将i赋值给s,此时s的值是0。
2、然后执行defer里面的语句

如果提前定义:
1、因为返回值已经提前定义了，不会产生临时零值变量，返回值就是提前定义的变量，后续所有的操作也都是基于已经定义的变量，任何对于返回值变量的修改都会影响到返回值本身。



