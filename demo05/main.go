package main

import (
	"fmt"
)

func main() {
	// var age int8 = 20 //定义必须是int
	//int8取值范围: -128 - +127
	//(负零10000000规定为 - 128)

	// fmt.Printf("age=%v, 类型=%T", age, age)

	// var num uint8 = 130

	// fmt.Printf("num=%v, 类型=%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //查看占用的内存字节数

	// int 类型转换
	// var a1 int32 = 10
	// var a2 int64 = 21

	// fmt.Println(int64(a1) + a2)
	// fmt.Println(a1 + int32(a2)) //高位向低位转换可能截断

	// num := 30
	// fmt.Printf("num=%v, type=%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //表示64位计算机，int占用8字节

	num := 12

	fmt.Printf("num=%v\n", num) // %v表示原样输出
	fmt.Printf("num=%d\n", num) // %d表示十进制输出
	fmt.Printf("num=%b\n", num) // %b表示二进制输出
	fmt.Printf("num=%o\n", num) // %o表示八进制输出
	fmt.Printf("num=%x\n", num) // %x表示十六进制输出

}
