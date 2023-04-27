关于类型转化，下面选项正确的是？
```go
A.
type MyInt int
var i int = 1
var j MyInt = i

B.
type MyInt int
var i int = 1
var j MyInt = (MyInt)i

C.
type MyInt int
var i int = 1
var j MyInt = MyInt(i)

D.
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)
```

答：C

A: int类型不能转成MyInt类型
B: go不支持这种类型转换
D: i不是interface类型，不支持这种断言