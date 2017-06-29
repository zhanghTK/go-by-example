package main

import (
	"fmt"
	"os"
	"time"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	filePath := "./" + time.Unix(time.Now().Unix(), 0).Format("20060102-0304")
	f := createFile(filePath)
	defer closeFile(f)
	writeFile(f)
}
