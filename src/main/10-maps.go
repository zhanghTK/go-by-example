package main

import "fmt"

func main() {
	// 使用make创建空map，make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len", len(m))

	// 移除一个键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 第二个返回值表明这个键是否在map中
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 声明且初始化一个map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
