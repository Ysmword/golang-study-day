关于字符串连接，下面语法正确的是？

A. str := 'abc' + '123'

B. str := "abc" + "123"

C. str := '123' + "abc"

D. fmt.Sprintf("abc%d", 123)

答：BD

A: 单引号表示单个字符，类型是rune
C: 就是在A的基础上，两个类型完全不一样没法直接相加