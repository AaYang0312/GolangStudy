package main

import "fmt"

// func fn1() {
// 	fmt.Println("fn1")
// }
// func fn2() {
// 	fn1()
// 	fmt.Println("fn2")
// }

// // 函数的递归调用
// func main() {
// 	fn2()
// }

func fn1(n int) {
	fmt.Println(n)

	if n > 0 {
		n--
		fn1(n)
	}
}
func fn2(n int) int {
	if n > 0 {
		return n + fn2(n-1)
	} else {
		return 0
	}
}
func fn3(n int) int {
	if n > 0 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}
func main() {
	// 递归输出1-所输入数的所有数字
	fn1(10)
	// 递归实现1-100的和
	n := fn2(100)
	fmt.Println(n)
	// 递归实现5的阶乘
	m := fn3(5)
	fmt.Println(m)
}
