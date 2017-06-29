package main

import "fmt"

// 使用通道接收string类型的通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings：使用通道发送数据
// pongs：使用通道接收数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
