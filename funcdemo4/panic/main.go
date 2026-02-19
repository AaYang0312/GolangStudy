package main

import (
	"errors"
	"fmt"
)

// panic 可在任何地方引发，但 recover 只有在 defer 调用的函数中有效
func fn1() {
	fmt.Println("fn1")
}
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("fn2抛出异常")
}
func fn3(x, y int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(x / y) // 当 y = 0 抛出异常：runtime error: integer divide by zero
}

// 模拟读取文件
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件名错误")
}
func myFn() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("main1.go")
	if err != nil {
		panic(err)
	}
}
func main() {
	fn1()
	fn3(10, 0)
	fn2()
	myFn()
	fmt.Println("结束")
}
