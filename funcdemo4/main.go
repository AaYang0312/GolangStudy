package main

import (
	"fmt"
)

// func f1() {
// 	fmt.Println("开始")

//		defer func() {
//			defer fmt.Println("aaaa")
//			fmt.Println("bbbb")
//		}()
//		fmt.Println("结束")
//	}
func f2() int {
	var a int
	defer func() {
		a++
	}()
	// fmt.Println("结束")
	return a // 0 返回时为0，返回后执行defer+1
}
func f3() (a int) {
	defer func() {
		a++
	}()
	// fmt.Println("结束")
	return a // 1 返回时为1，
}
func f4() (x int) {
	x = 5
	defer func(x int) {
		x++
		fmt.Println("执行x++")
	}(x)
	return x
}
func add(s string, a, b int) int {
	ret := a + b
	fmt.Println(s, a, b, ret)
	return ret
}
func main() {
	// defer语句会把后面跟随的语句延迟处理
	// 执行完其他后逆序执行defer
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)

	// //f1()
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4() )

	//defer语句在执行函数时，函数注册时需要确定所有参数的值
	x := 2
	y := 3
	defer add("AA", x, add("A", x, y))
	x = 10
	defer add("BB", x, add("B", x, y))
	y = 20
}

/*
	注册顺序:
	1. defer add("AA", x, add("A", x, y))
	2. defer add("BB", x, add("B", x, y))
	注册时会确认add中的add，所以先执行:
	1. add("A", x, y)
	2. add("B", x, y)
	完成后逆序执行defer语句:
	1. defer add("BB", x, add("B", x, y))
	2. defer add("AA", x, add("A", x, y))
*/
