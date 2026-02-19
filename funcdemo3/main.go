package main

import "fmt"

/*

全局变量:
1. 常驻内存
2. 污染全局

局部变量:
1. 不常驻内存
2. 不污染全局

闭包:
1. 可以让一个变量常驻内存
2. 可以让一个变量不污染全局

*/
/*

闭包是指有权访问另一个函数作用域中变量的函数
常见的创建方法是在一个函数中创建另一个函数来访问函数中的局部变量

*/
func adder1() func() int {
	var i = 10
	return func() int {
		i += 1
		return i
	}
}
func adder2() func(y int) int {
	var i = 20
	return func(y int) int {
		i += y
		return i
	}
}
func main() {
	var a = adder1() //表示执行方法
	fmt.Println(a()) // 11
	fmt.Println(a()) // 12
	fmt.Println(a()) // 13

	var b = adder2()  //表示执行方法
	fmt.Println(b(1)) // 21
	fmt.Println(b(1)) // 22
	fmt.Println(b(1)) // 23
	fmt.Println(a())  // 14 变量i常驻内存，但是不污染全局
}
