package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 这个测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 编译一个表达式
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 测试匹配
	fmt.Println(r.MatchString("peach"))

	// 查找匹配字符串，返回字符串
	fmt.Println(r.FindString("peach punch"))

	// 查找第一次匹配的字符串，返回开始结束索引
	fmt.Println(r.FindStringIndex("peach punch"))

	// Submatch 返回完全匹配和局部匹配的字符串
	// p([a-z]+)ch 和 `([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Submatch 返回完全匹配和局部匹配的字符串开始结束索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 匹配[]byte
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式常量
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// 替换匹配部分为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// 传递指定内容到一个给定的函数
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
