下面代码输出什么？


```go
type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
}
```
A. 4

B. compilation error


答：B

类似于X=Y的赋值操作,必须知道X的地址,才能够将Y的值赋给X,但go中的map的value本身是不可寻地址。




