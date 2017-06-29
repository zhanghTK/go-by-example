package main

import (
	"fmt"
)

// 接收任意数目的int作为参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 使用slice作为变参，数组不支持
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
