下面这段代码能否通过编译？如果通过，输出什么？

```go
package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int =0
	var i1 MyInt1 = i  // 不能将int类型值赋值给MyInt1类型
	var i2 MyInt2 = i  // 不能将MyInt1类型赋值给int类型
	fmt.Println(i1,i2)
}
```

答案：不能编译通过
type MyInt1 int 表示基于int创建新类型MyInt1
type MyInt2 = int 表示给int创建一个别名（MyInt2）



