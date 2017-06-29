package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// Unix时间戳
	secs := now.Unix()
	// UnixNano时间戳
	nanos := now.UnixNano()
	fmt.Println(now)

	// UnixMillis需要手动转换
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 转为对应时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
