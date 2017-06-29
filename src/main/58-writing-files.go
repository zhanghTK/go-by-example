package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 项文件写入数据
	// 如果文件不存在以perm权限创建
	// 如果文件存在清空再写
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./dat1", d1, 0644)
	check1(err)

	f, err := os.Create("./dat2")
	check1(err)

	defer f.Close()

	// 写入切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check1(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 写入字符串
	n3, err := f.WriteString("writes\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// 缓冲区的信息写入磁盘
	f.Sync()

	// 带缓冲区的写入器
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// 写入
	w.Flush()
}
