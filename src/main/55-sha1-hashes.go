package main

import (
	"crypto/sha1"
	"fmt"
)

// MD5：crypto/md5下md5.New()
func main() {
	s := "sha1 this string"
	// 产生一个散列值
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	// 16进制输出散列值
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
