以下代码能通过编译吗？为什么？

```go
package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "love"
	fmt.Println(peo.Speak(think))
}
```

答：上述代码报错的地方在 var peo People = Student{} 这条语句， Student{} 已经重写了父类 People{} 中的 Speak(string) string 方法，那么只需要用父类指针指向子类对象即可。