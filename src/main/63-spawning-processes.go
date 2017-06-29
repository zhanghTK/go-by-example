package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("cmd")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> cmd")
	fmt.Println(string(dateOut))
}
