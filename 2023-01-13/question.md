问题：

下面这段代码有什么缺陷：
func sum(x, y int)(total int, error) {
	return x+y, nil
}


答案：返回值有一个命名了，其他的也需要命名
