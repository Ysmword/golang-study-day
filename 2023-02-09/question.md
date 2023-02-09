今天给两道类似的题目，注意，有半数以上的人可能会做错！

题一：

```go
package main

func main() {
    var a int8 = -1
    var b int8 = -128 / a

    println(b)
}
```

题二：
```go
package main

func main() {
    const a int8 = -1
    var b int8 = -128 / a

    println(b)
}
```
它们分别输出什么？请写出你的答案，能解释原因最好。



答：
题一：-128
题二：编译错误

题目1：
因为var b intt8 = -128 / a不是常量表达式，因此untype常量-128隐式转化为int8类型(即和a的类型一致),所以-128/a的结果是int8类型，值是128，
超出了int8范围。因为结果不是常量，允许溢出,128的二进制表示10000000,正好是-128的补码

题目2：
对于 var b int8 = -128 / a，因为 a 是 int8 类型常量，所以 -128 / a 是常量表达式，在编译器计算，结果必然也是常量。因为 a 的类型是 int8，因此 -128 也会隐式转为 int8 类型，128 这个结果超过了 int8 的范围，但常量不允许溢出，因此编译报错。

得去恶补下计算机知识了，补码相关的都忘了。
详细可以看下：[补码相关]{https://www.cnblogs.com/zhangziqiu/archive/2011/03/30/ComputerCode.html}


