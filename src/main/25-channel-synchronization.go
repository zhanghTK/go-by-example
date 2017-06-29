package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	// 使用通道同步协程
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
