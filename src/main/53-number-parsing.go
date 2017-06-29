package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// base:0 		->	自动推断字符串表示的数字的进制
	// biSize:64	->	整数是64位存储
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// 基础的10进制整数转换函数，常用
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("fad")
	fmt.Println(e)
}
