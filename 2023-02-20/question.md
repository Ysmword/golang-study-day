这份面试经验总结中（其实谈不上总结，只是面试题的记录，并没有总结分析答案），有一道 Go 相关的题，也是一个老生常谈的问题：以下代码有什么问题，怎么解决？

total, sum := 0, 0
for i := 1; i <= 10; i++ {
    sum += i
    go func() {
        total += i
    }()
}
fmt.Printf("total:%d sum %d", total, sum)


字节面试Go的经验
[面试 Go 的经验]{https://zhuanlan.zhihu.com/p/132813717}

total的值不固定,发生了数据竞争;可能main运行完成, 创建的协程还没有运行完成

