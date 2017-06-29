package main

import (
	"fmt"
	"time"
)

// 使用协程、通道、打点器限制处理速率
func main() {
	fmt.Println("====================永久限速====================")
	// 限制接收请求
	requests := getRequest()

	// NewTicker的封装，无法停止，不能被回收
	// 相当于速率限制的管理器
	limiter := time.Tick(time.Second * 2)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("====================临时限速====================")

	// 临时进行速率限制
	// 限速器通道有初始值，初始通道不阻塞
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		// 打点器通道的值给限速器通道
		for t := range time.Tick(time.Second * 2) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := getRequest()
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
func getRequest() chan int {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	return requests
}
