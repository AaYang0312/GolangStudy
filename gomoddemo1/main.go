package main

import (
	"fmt"
	_ "gomoddemo1/calc" // 把 C 设为后面的缩写，写_为匿名，匿名引用仍会执行包内的 init() 函数
)

func main() {
	// sum := C.Add(1, 2)
	// fmt.Println(sum)
	// //fmt.Println(calc.aaa) // 首字母小写无法访问
	// fmt.Println(C.AAA)

	fmt.Println("执行完毕")
}
