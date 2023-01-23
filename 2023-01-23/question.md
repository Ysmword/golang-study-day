Go101 作者发布了一道有意思的题目，这道题主要考察细节点，Go101 这本书也是以抠细节著称。看看这道题，以下程序输出什么？（单选）
```go
package main

const s = "Go101.org"

// len(s) == 9
// 1 << 9 == 512
// 512 / 128 == 4

var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func main() {
	println(a, b)
}
// A、0 0；B、0 4；C：4 0；D：4 4
```
这里涉及到两个知识点，在 Go 语言规范中都有明确的说明，但确实很细节。你答对了吗？原因知晓吗？

答案：C

知识点：
1、<<表示位运算中向左移动操作2
2、b中的运算因为s[:]数组，使得整个表达式不是一个常量位移表达式，按照规定会将1转化为byte,因此 1 << len(s[:]) 的结果也是 byte 类型，而 byte 类型最大只能表示 255，很显然 512 溢出了，结果为 0
更加详细的知识点可以看：https://polarisxu.studygolang.com/posts/go/action/interview-len-shift/



