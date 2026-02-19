package main

import "fmt"

func main() {
	// 1、关系运算符：==, !=, >, >=, <, <=
	// 2、逻辑运算：&&, ||, !
	var a int = 10
	if a < 15 || test() {
		print("执行")
	}
}

func test() bool {
	fmt.Println("test...")
	return true
}
