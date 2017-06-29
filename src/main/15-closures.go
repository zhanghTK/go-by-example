package main

import "fmt"

// intSeq函数返回一个新函数
// 新函数使用闭包的方式隐藏变量i
func inSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	// 闭包变量与闭包函数一一对应
	nextInt := inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 返回新的闭包函数和闭包变量
	newInts := inSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
}
