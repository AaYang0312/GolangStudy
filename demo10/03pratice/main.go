package main

import "fmt"

func main() {
	// 1、交换a,b的值，不借助中间变量
	// var a = 10
	// var b = 35

	// a += b
	// b = a - b
	// a -= b
	// print(a)
	// print(b)

	// 2、100天是多少星期零几天
	var day = 100
	var week = day / 7
	var left = day % 7
	fmt.Printf("100天为%d星期零%d天", week, left)

	/* 3、位运算符号(二进制)
	&: 与运算，全为1等于1
	|: 或运算，有1就为1
	^: 异或运算，不一样为1
	<<: 左移 5 101 5 << 2 = 10100
	>>: 右移 5 101 5 >> 2 = 1
	*/
}
