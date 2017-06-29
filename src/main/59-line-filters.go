package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用带缓冲的scanner读取控制台一行
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
