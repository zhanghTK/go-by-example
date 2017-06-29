package main

import (
	"fmt"
)

func main() {
	// 遍历slice，数组同理
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历同时获得索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s - %s\n", k, v)
	}

	// 遍历字符串unicode编码
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
