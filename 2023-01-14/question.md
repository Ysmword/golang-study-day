以下代码是否能编译通过？

package main

import "fmt"

func main() {
	m := make(map[string]int)

	fmt.Println(&m["qcrao"])
}

答案：不能，不能对map的键值取址，考虑到map可自动扩容，map的数据元素的位置可能在这个过程中发生变化，因此go不允许获取map中的value的地址，这是在编译的时候就生效了的