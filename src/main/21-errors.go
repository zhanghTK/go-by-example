package main

import (
	"errors"
	"fmt"
)

// 错误通常最后一个返回值是error类型
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New构建一个使用给定错误信息的基本error值
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// 自定义错误类型
type argError struct {
	arg  int
	prob string
}

// 自定义错误类型实现error接口
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// &argError方法集合
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 使用自定义错误类型中的数据
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
