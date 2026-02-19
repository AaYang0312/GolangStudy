package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1、copy()函数复制切片
	// arr1 := []int{1, 2, 3, 4}
	// arr2 := make([]int, 4, 4)
	// a := copy(arr2, arr1)
	// arr2[0] = 2
	// fmt.Printf("a=%d\n", a) // 4, 复制的元素个数
	// fmt.Printf("arr1=%v, len=%d, cap=%d\n", arr1, len(arr1), cap(arr1))
	// fmt.Printf("arr2=%v, len=%d, cap=%d\n", arr2, len(arr2), cap(arr2))

	// 2、删除切片元素，没有专门方法，需要使用切片自身特性
	// a := []int{30, 31, 32, 33, 34, 35}
	// // 要删除索引为2的元素
	// a = append(a[:2], a[3:]...)

	// 3、关于修改字符串特定字符
	// s1 := "hello world"
	// arr1 := []byte(s1)

	// s2 := "你好golang" // 有汉字用rune类型
	// arr2 := []rune(s2)
	// arr2[0] = '我'
	// fmt.Println(string(arr2))

	// 4、数组切片的排序
	// 练习，选择排序
	// 从小到大
	// var numSlice = []int{9, 8, 7, 6, 5, 4}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := i + 1; j < len(numSlice); j++ {
	// 		if numSlice[i] > numSlice[j] {
	// 			numSlice[j], numSlice[i] = numSlice[i], numSlice[j]
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	// 从大到小
	// var numSlice = []int{9, 8, 7, 6, 5, 4}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := i + 1; j < len(numSlice); j++ {
	// 		if numSlice[i] < numSlice[j] {
	// 			numSlice[i], numSlice[i] = numSlice[i], numSlice[j]
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	// 冒泡，每轮选出最大的放正确位置
	// var numSlice = []int{9, 8, 7, 6, 5, 4}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := 0; j < len(numSlice)-1; j++ {
	// 		if numSlice[j] > numSlice[j+1] {
	// 			numSlice[j], numSlice[j+1] = numSlice[j+1], numSlice[j]
	// 		}
	// 	}
	// }
	// fmt.Println(numSlice)

	// go语言sort()排序
	var numSlice = []int{9, 8, 7, 6, 5, 4}
	sort.Ints(numSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(numSlice))) // 降序 Float64Slice, StringSlice
	fmt.Println(numSlice)

}
