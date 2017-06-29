package main

import (
	"fmt"
	"time"
)

// 协程——轻量级的线程

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {

	// 协程中函数调用
	go f("goroutine")

	// 协程中调用匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 函数调用
	f("direct")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
