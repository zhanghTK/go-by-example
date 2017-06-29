package main

import "fmt"

// http://www.chongchonggou.com/g_648526226.html
// http://shanks.leanote.com/post/Untitled-55ca439338f41148cd000759-17
// http://blog.jobbole.com/107613/

// 接口中的方法，对于值类型接收器，不会包含指针类型接收器的方法
// 指针类型的接收器包含值类型接收器的方法
type Meower interface {
	Meow()
}

func Greet(meower Meower) {
	meower.Meow()
}

type Cat struct {
	Name  string
	Color string
}

func (c *Cat) Meow() {
	//func (c Cat) Meow() {
	fmt.Println("Name:", c.Name, "Color:", c.Color)
}

func main() {
	c := Cat{"a", "b"}
	Greet(c)
}
