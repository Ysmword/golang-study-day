以下代码能否编译？

```go
package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {

	list = make(map[string]Student)

	student := Student{"Aceld"}

	list["student"] = student
	list["student"].Name = "LDB"

	fmt.Println(list["student"])
}
```

答：map[string]Student的value是一个 Student结构值，所以当list["student"] = student，是一个值拷贝的过程。而list["student"]
则是一个值引用。那么值引用的特点是`只读`。所以对list["student"].Name="LDB"的修改是不允许的
