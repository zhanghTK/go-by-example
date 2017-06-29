package main

import "fmt"

func main() {
	// 创建通道时指定缓存大小
	// 允许在没有接收方的情况下缓存限定数量的值
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
