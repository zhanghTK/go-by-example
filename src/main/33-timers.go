package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器
	timer1 := time.NewTimer(time.Second * 2)

	// 定时器使用通道阻塞
	//<-timer1.C
	t := <-timer1.C
	fmt.Println("Timer 1 expired", t)

	// 定时器失效前取消定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
