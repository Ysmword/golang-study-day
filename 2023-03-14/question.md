下面代码中 A B 两处应该怎么修改才能顺利编译？

```go
func main() {
	var m map[string]int        //A
	m["a"] = 1
	if v := m["b"]; v != nil {  //B
		fmt.Println(v)
	}
}
```
答：
A:  m:= make(map[string]int)
B:  v!=0


