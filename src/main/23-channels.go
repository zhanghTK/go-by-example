package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道：make(chan val-type)
	// 通道的发送和接收默认是阻塞的(无缓冲)
	// 直到发送方和接收方都准备完毕
	message := make(chan string)

	// 发送一个值到通道
	go func() { message <- "ping" }()

	// 从通道中接收一个值
	msg := <-message
	fmt.Println(msg)
}
