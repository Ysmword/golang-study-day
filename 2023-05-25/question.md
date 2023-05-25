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

答：不能编译

原因：值引用是不能被修改的。要么是map的Student改为指针，要么是重新赋值

