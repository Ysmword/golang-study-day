有下面 3 行代码：

// 32 位机器
1）var x int32 = 32.0
2）var y int = x
3）var z rune = x
它们是否能编译通过？为什么？

如果面试时问这道题，你需要想想面试官想考察你什么。

答：
1）正确,字面量是无类型(untyped)的，32.0是无类型的浮点数字面量，因此他可以赋值给任意数字相关类型变量（常量）
2）错误,Go语言不会做隐式类型转换, int 和 int32是不同类型, 因此上题中2）不能编译
3）正确,rune是int32别名，因此题目中3)也能编译通过

[go类型](https://polarisxu.studygolang.com/posts/go/action/type-and-alias/)