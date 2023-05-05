下面这段代码有什么缺陷：

```go
func sum(x, y int)(total int, error) {
	return x+y, nil
}
```

答：
返回值的声明要么全声明，要么全不声明

