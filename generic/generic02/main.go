package main

import "fmt"

// 定义一个getData的泛型方法
func getData[T any](value T) T {
	return value
}
func main() {
	str := getData[string]("this is str")
	fmt.Println(str)

	num := getData[int](123)
	fmt.Println(num)

	str2 := getData('2')
	fmt.Printf("%c", str2)
}
