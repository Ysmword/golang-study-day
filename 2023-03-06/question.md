下面代码输出什么？

```go
func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p :=1
	incr(&p)
	fmt.Println(p)
}
```

A. 1

B. 2

C. 3


答：B

首先go不存在对指针进行运算,然后incr()函数里的p是*int类型的指针,指向main函数里的变量p的地址。第二行代码是将该地址的值执行一个自增操作。incr()返回自增或的结果