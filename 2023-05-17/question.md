下面代码有什么问题？

package main

const cl = 100

var bl = 123

func main()  {
    println(&bl,bl)
    println(&cl,cl)
}


答：获取常量的地址，是非法行为

