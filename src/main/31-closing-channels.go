package main

import "fmt"

// 关闭通道则不能再向通道发送值了
// 一个非空的通道也是可以关闭的
// 关闭的非空通道中的值仍然可以获取到
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 关闭通道
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
