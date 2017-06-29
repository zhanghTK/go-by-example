package main

import "fmt"

// if语句可以不使用圆括号，但花括号是必须的
// 不支持三目
func main() {
	// 基本例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 不要else只使用if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在条件语句前可以有一个赋值语句
	// 任何这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is ngative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
