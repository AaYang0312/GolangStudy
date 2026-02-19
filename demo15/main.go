package main

import "fmt"

func main() {
	// 引用类型
	// 基础数据类型 和 数组是值类型

	// var a = 1
	// b := a
	// a = 10
	// fmt.Println(a, b)

	// var arr1 = [...]int{1, 2, 3, 4}
	// var arr2 = arr1
	// arr1[0] = 11
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// // 切片类型是引用类型

	var arrs1 = []int{}
	arrs1 = append(arrs1, 1, 2)
	var arrs2 = arrs1
	arrs1[0] = 11
	fmt.Println(arrs1)
	fmt.Println(arrs2)
	fmt.Printf("%T %T\n", arrs1, arrs2)

	// 定义多维数组

	var arrm = [3][2]string{
		{"Beijing", "Shanghai"},
		{"Toronto", "Murben"},
		{"New York", "Washinton"},
	}
	fmt.Println(len(arrm))
	fmt.Printf("%T", arrm)
}
