以下代码输出什么？

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}
```

A：{"Time": "2020-12-20T00:00:00Z", "N": 5}；

B："2020-12-20T00:00:00Z"；

C：{"N": 5}；

D：<nil>

答：B

[参考答案](https://polarisxu.studygolang.com/posts/go/action/weekly-question-embed-time/)
简单来说：json.Marshal函数优先调用MarshalJSON,然后是MarshalText,如果都没有才会走正常类型编码逻辑。