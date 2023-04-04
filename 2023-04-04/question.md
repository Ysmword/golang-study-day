下面这段代码输出什么？

```go
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
```


答案：
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4


程序执行到main()函数三行代码的时候,会先执行calc()函数的b参数,即：calc("10",a,b),输出10 1 2 3,得到值3
因为defer定义的函数是延迟函数,故calc("1",1,3)会被延迟执行。

程序执行到第五行时,同样先执行calc("20",a,b)输出：20 0 2 2得到值2,同样将calc("2",0,2)延迟执行；
程序执行到末尾的时,按照栈先进后出的方式依次执行: calc("2",0,2),calc("1",1,3),则一次输出2 0 2 2,1 1 3 4



