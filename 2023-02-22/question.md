下面这段代码输出什么？
```go
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}
```
A.3

B.4

C.compilation error


答：操作符[i,j]。基于数组(切片)可以使用操作符[i,j]创建新的切片,从索引i,到索引j结束,截取已有数组的任意部分,返回新的切片,新切片的值包含原数组（切片）的 i 索引的值，但是不包含 j 索引的值。


截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。

所以例子中，切片 t 为 [4]，长度和容量都是 1。