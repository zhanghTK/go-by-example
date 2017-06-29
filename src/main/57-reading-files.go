package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 直接将文件读取
	dat, err := ioutil.ReadFile("./dat2")
	check(err)
	fmt.Println(string(dat))

	// 打开文件
	f, err := os.Open("./dat2")
	check(err)

	// 读取5个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 从第六个字节开始读
	o2, err := f.Seek(6, 0)
	check(err)
	// 读两个字节
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// 从第六个字节开始读
	o3, err := f.Seek(6, 0)
	check(err)
	// 读两个字节
	b3 := make([]byte, 2)
	// 至少读取两个字节
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 从文件头开始读
	_, err = f.Seek(0, 0)
	check(err)

	// 使用带缓冲的读取
	r4 := bufio.NewReader(f)
	// 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 关闭文件，建议defer
	f.Close()
}
