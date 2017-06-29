package main

import "fmt"

// 定义结构体person
// 包含一段name，age
type person struct {
	name string
	age  int
}

func main() {
	// 创建结构体
	fmt.Println(person{"Bob", 20})

	// 指定字段名
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段默认初始化零值
	fmt.Println(person{name: "Fred"})

	// 结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 使用.访问结构体的字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 使用.访问结构体指针的字段
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
