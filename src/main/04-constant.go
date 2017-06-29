package main

import (
	"fmt"
	"math"
)

// const用于声明一个常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// const语句可以出现在任何var语句出现的地方
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常数是没有确定的类型的
	// 知道它们被给定了一个类型
	// 比如说一次显示的类型转化
	fmt.Println(int64(d))

	// 当上下文需要时，一个数可以被给定一个类型
	// 例如变量赋值或函数调用
	// 这里的math.Sin函数需要一个float64的参数
	fmt.Println(math.Sin(n))
}
