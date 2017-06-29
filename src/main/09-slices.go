package main

import "fmt"

func main() {
	// slice的类型仅由包含的元素决定
	// 使用make创建slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len", len(s))

	// append返回一个新的slice，包含一个或多个新值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片操作，左闭右开
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 左空，右闭
	l = s[:5]
	fmt.Println("sl2:", l)

	// 左闭，右空
	l = s[2:]
	fmt.Println("sl3:", l)

	// 声明同时初始化一个slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// slice组成多维数组结构
	// 内部slice长度可以不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
