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

答：不能，
继承与多态的特点在golang对多态的特点体现在语法上并不明显。
我们知道发生多态的几个要素：
1、有interface接口，并且有接口定义方法。
2、有子类去重写interface接口。
3、有父类指针只向子类的具体对象。
那么，满足上述3个条件，就可以产生多态效果，就是父类指针可以调用子类的具体方法。

所以上述代码报错的地方在 var peo People = Stduent{} 这条语句， Student{} 已经重写了父类 People{} 中的 Speak(string) string 方法，那么只需要用父类指针指向子类对象即可。（Go 中不叫父类，这里是为了好理解）
所以应该改成 var peo People = &Student{} 即可编译通过。（People 为 interface 类型，就是指针类型）；
还有因为peo是非指针实例，不拥有指针类型的方法，去掉Speak中的指针也可以






