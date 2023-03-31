return 之后的 defer 语句会执行吗，下面这段代码输出什么？

```go
var a bool = true
func main() {
	defer func(){
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func(){
		fmt.Println("3")
	}()
}
```
答：没有看到return,尴尬。。。
结果：2 1

原因：return 之后并没有将最后一个defer注册


defer关键字后面的函数或者方法想要执行必须先注册,return之后的defer是不能注册的,也就不能执行后面的函数方法