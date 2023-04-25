关于循环语句，下面说法正确的有（）

A. 循环语句既支持 for 关键字，也支持 while 和 do-while；

B. 关键字 for 的基本使用方法与 C/C++ 中没有任何差异；

C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；

D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；

答：C,D

分析：
A: 不支持while和do-while
B: for 还支持range关键字的循环
C: 支持：
```go
func main() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Println(i, j)
		}
		fmt.Println("i", i)
	}
}
```
D: 平行赋值,a[i], a[j] = a[j], a[i]。逗号赋值:a[i]=a[j],a[j]=a[i] 不行


