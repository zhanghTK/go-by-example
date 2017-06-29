package main

import "fmt"

type rect struct {
	width, height int
}

// 指针类型接收器方法
// 指针类型接收器优点：
//  1. 避免方法在调用时产生一个拷贝
//  2. 方法能改改变接受的数据
func (r *rect) area() int {
	return r.width * r.height
}

// 值类型接收器方法
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
