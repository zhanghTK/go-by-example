package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// 打印结构体实例
	// {1 2}
	fmt.Printf("%v\n", p)

	// 打印结构体实例
	// {x:1 y:2}
	fmt.Printf("%+v\n", p)

	// 打印结构体实例
	// main.point{x:1, y:2}
	fmt.Printf("%#v\n", p)

	// 打印值的类型
	// main.point
	fmt.Printf("%T\n", p)

	// 打印布尔值
	// true
	fmt.Printf("%t\n", true)

	// 标准的十进制
	// 123
	fmt.Printf("%d\n", 123)

	// 二进制表示形式
	// 1110
	fmt.Printf("%b\n", 14)

	// 整数的对应字符
	// !
	fmt.Printf("%c\n", 33)

	// 十六进制编码
	// 1c8
	fmt.Printf("%x\n", 456)

	// 对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化
	// 78.900000
	fmt.Printf("%f\n", 78.9)

	// 科学计数法
	// 1.234000e+08
	fmt.Printf("%e\n", 123400000.0)
	// 科学计数法
	// 1.234000E+08
	fmt.Printf("%E\n", 123400000.0)

	// 基本的字符串输出
	// "string"
	fmt.Printf("%s\n", "\"string\"")

	// 带有引号的输出
	// "\"string\""
	fmt.Printf("%q\n", "\"string\"")

	// base-16 编码的字符串，每个字节使用 2 个字符表示
	// 6865782074686973
	fmt.Printf("%x\n", "hex this")

	// 输出指针
	// 0xc04203c1d0
	fmt.Printf("%p\n", &p)

	// 指定宽度
	// |    12|   345|
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// 指定宽度.精度
	// |  1.20|  3.45|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// 左对齐
	// |1.20  |3.45  |
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 字符串对齐
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// 字符串左对齐
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf 则格式化并返回一个字符串而不带任何输出
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
