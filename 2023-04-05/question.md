在 Go语言爱好者周刊第 104 期有一道题目，以下代码输出什么：
```go
package main

func main() {
  var x *struct {
    s [][32]byte
  }
  
  println(len(x.s[99]))
}
```
A：运行时 panic；

B：32；

C：编译错误；

D：0


答：B，简单来说不是求值，及时为nil调用，也不会发生panic
求值：取地址中的值



