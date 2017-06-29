package main

import (
	"fmt"
	"time"
)

// 关于协程调度
// http://blog.csdn.net/wangkai_123456/article/details/72780433
// http://blog.csdn.net/heiyeshuwu/article/details/51178268
// http://blog.amalcao.me/blog/2014/05/09/erlang-and-go-de-bing-fa-diao-du-qian-xi/
func main() {
	go func() {
		for true {
			fmt.Println("1")
		}
	}()

	go func() {
		for true {
			fmt.Println("2")
		}
	}()

	time.Sleep(time.Second * 100)
}
