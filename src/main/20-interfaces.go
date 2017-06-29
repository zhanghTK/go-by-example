package main

import (
	"fmt"
	"math"
)

// 定义一个几何体基本接口
type geometry interface {
	area() float64
	perim() float64
}

// 矩形的定义
type rect1 struct {
	width, height float64
}

// 圆形的定义
type circle struct {
	radius float64
}

// 矩形对几何体接口的实现
func (r rect1) area() float64 {
	return r.width * r.height
}
func (r rect1) perim() float64 {
	return 2 * (r.width + r.height)
}

// 圆形对几何体接口的实现
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 使用接口类型作为参数
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect1{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
