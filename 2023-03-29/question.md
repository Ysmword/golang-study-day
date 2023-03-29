以下代码输出什么？

```go
package main

import (
	"fmt"
)

func main() {
	var ans float64 = 15 + 25 + 5.2
	fmt.Println(ans)
}
```
A：不能编译

B：45

C：45.2

D：45.0

答：C

解析：15 + 25 + 5.2 是常量表达式，因为这个表达式的操作数都是无类型的常量，因为其中有 5.2，它的默认类型是浮点型，所以这个常量表达式的结果虽然是无类型的，但默认类型是浮点型。

答案详解：https://polarisxu.studygolang.com/posts/go/action/weekly-question-94/

