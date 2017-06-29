package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义打点器
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	// 停止打点器
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
