package main

import (
	"flag"
	"fmt"
)

// 未声明的标示默认都在最后的tail切片
// 声明的参数必须在未声明的参数之前
// -h可以获得帮助
// 提供未声明的参数也会获得帮助
func main() {
	// 默认值为foo的字符串标志word，描述为"a string"
	// 注意返回的是指针
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 解析
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
